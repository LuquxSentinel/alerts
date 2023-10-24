package api

import (
	"context"
	"github.com/luqus/s/tokens"
	"time"
	
	"github.com/gofiber/fiber/v2"
	"github.com/luqus/s/types"
)

func (api *APIServer) registerUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	registerInput := new(types.RegisterInput)
	if err := c.BodyParser(registerInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid register input")
	}

	//check if phone number or email exists
	count, err := api.authStorage.CheckIfEmailExists(ctx, registerInput.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("registration failed")
	}

	if count > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON("email already in use")
	}

	count, err = api.authStorage.CheckIfPhoneNumberExists(ctx, registerInput.PhoneNumber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("registration failed")
	}

	if count > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON("phone number already in use")
	}

	user := new(types.User)
	user.Password = registerInput.Password
	err = user.HashPassword()
	if err !=nil {
		return c.Status(fiber.StatusInternalServerError).JSON("registration failed")
	}
	user.SetUID()
	user.Email = registerInput.Email
	user.FirstName = registerInput.FirstName
	user.LastName = registerInput.LastName
	user.PhoneNumber = registerInput.PhoneNumber
	user.CreatedAt = time.Now().UTC()

	// Insert user into database
	err = api.authStorage.CreateUser(ctx, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("registration failed")
	}


	return c.Status(fiber.StatusOK).JSON("user successfully created")
}

func (api *APIServer) login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	loginInput := new(types.LoginInput)
	if err := c.BodyParser(loginInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid login input")
	}

	//fetch  user from database by email
	user, err := api.authStorage.GetUserByEmail(ctx, loginInput.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("wrong email or password")
	}

	// validate password
	err = user.VerifyPassword(loginInput.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("wrong email or password")
	}

	// generate jwt
	authorization, err := tokens.GenerateJwt(user.UID, user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("token generation error")
	}

	// set authorization header
	c.Set("authorization", authorization)

	return c.Status(fiber.StatusOK).JSON(user.FilterResponse())
}
package main

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/luqus/s/storage"
	"github.com/luqus/s/types"
)

type APIServer struct {
	listenAddr   string
	router       *fiber.App
	authStorage  storage.AuthenticationStorage
	alertStorage storage.AlertStorage
}

func NewAPIServer(listenAddr string, authstore storage.AuthenticationStorage, alertStorage storage.AlertStorage) *APIServer {
	return &APIServer{
		listenAddr:   listenAddr,
		router:       fiber.New(),
		authStorage:  authstore,
		alertStorage: alertStorage,
	}
}

func (api *APIServer) Run() error {

	return api.router.Listen(api.listenAddr)
}

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
	user.HashPassword()
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
		return c.Status(fiber.StatusBadRequest).JSON("invalif login input")
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

	// TODO: generate jwt
	authorization := ""

	// set authorization header
	c.Set("authorization", authorization)

	return c.Status(fiber.StatusOK).JSON(user.FilterResponse())
}

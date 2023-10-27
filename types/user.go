package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	UID         string             `bson:"uid"`
	Email       string
	Password    string
	PhoneNumber string
	FirstName   string
	LastName    string
	CreatedAt   time.Time
}

func (u *User) HashPassword() error {
	b, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}

	u.Password = string(b)
	return nil
}

func (u *User) VerifyPassword(plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword))

}

func (u *User) SetUID() {
	u.ID = primitive.NewObjectID()
	u.UID = u.ID.Hex()
}

type ResponseUser struct {
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	CreatedAt   time.Time `json:"created_at"`
}

type RegisterInput struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Password    string `json:"password"`
}

func (u *User) FilterResponse() *ResponseUser {
	return &ResponseUser{
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		CreatedAt:   u.CreatedAt,
	}
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

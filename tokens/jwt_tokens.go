package tokens

// import (
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

// type Claims struct {
// 	UID   string
// 	Email string
// 	jwt.RegisteredClaims
// }

// func GenerateJwt(uid string, email string) (string, error) {
// 	claims := &Claims{
// 		UID:   uid,
// 		Email: email,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiredDate: jwt.NewNumericDate(time.Now().UTC().UnixNano()),
// 		},
// 	}

// 	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(""))
// 	if err != nil {
// 		return "", err
// 	}

// 	return signedToken, nil
// }

// func ValidateJwt(signedToken string) error {
// 	claims := new(Claims)

// }

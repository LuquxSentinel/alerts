package tokens

 import (
	 "errors"
	 "os"
	 "time"

 	"github.com/golang-jwt/jwt/v5"
 )

 type Claims struct {
 	UID   string
 	Email string
 	jwt.RegisteredClaims
 }

 func GenerateJwt(uid string, email string) (string, error) {
 	claims := &Claims{
 		UID:   uid,
 		Email: email,
 		RegisteredClaims: jwt.RegisteredClaims{
 			ExpiresAt: jwt.NewNumericDate(time.Now().Local().UTC()),
 		},
 	}

 	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(""))
 	if err != nil {
 		return "", err
 	}

 	return signedToken, nil
 }

 func ValidateJwt(signedToken string) (string, error) {

	 key := os.Getenv("SECRET_KEY")
	 claims := new(Claims)

	 token, err := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		 return []byte(key), nil
	 })
	 if err != nil {
		return "",err
	 }

	 if !token.Valid {
		return "", errors.New("invalid authorization header")
	 }

	 return claims.UID, nil

 }

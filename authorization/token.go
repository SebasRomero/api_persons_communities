package authorization

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sebasromero/api/model"
)

func GenerateToken(data *model.Login) (string, error) {
	claim := model.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "Sebasco",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(t string) (model.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &model.Claim{}, verifyFunction)
	if err != nil {
		return model.Claim{}, err
	}
	if !token.Valid {
		return model.Claim{}, errors.New("Invalid token")
	}

	claim, ok := token.Claims.(*model.Claim)
	if !ok {
		return model.Claim{}, errors.New("Could not get the claim")
	}

	return *claim, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}

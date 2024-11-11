package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/escoutdoor/bookit/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthTokenClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

func (uc *usecase) ParseToken(jwtToken string) (uuid.UUID, error) {
	const op = "AuthUseCase.ParseToken"

	token, err := jwt.ParseWithClaims(jwtToken, &AuthTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s: unexpected signing method: %v", op, t.Header["alg"])
		}

		return []byte(uc.signKey), nil
	})
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			return uuid.Nil, fmt.Errorf("it doesn't look like a token")
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return uuid.Nil, fmt.Errorf("invalid token signature")
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			return uuid.Nil, fmt.Errorf("token is either expired or not active yet")
		default:
			return uuid.Nil, fmt.Errorf("invalid token")
		}
	}
	claims, ok := token.Claims.(*AuthTokenClaims)
	if !ok {
		return uuid.Nil, err
	}
	return claims.UserID, nil
}

func (uc *usecase) genToken(user *model.User) (string, error) {
	const op = "AuthUseCase.genToken"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &AuthTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: user.ID,
	})

	tokenStr, err := token.SignedString([]byte(uc.signKey))
	if err != nil {
		return "", fmt.Errorf("%s: signed string %s", op, err)
	}

	return tokenStr, nil
}

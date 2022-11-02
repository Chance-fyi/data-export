package jwt

import (
	"context"
	"data-export/pkg/g"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

type config struct {
	RedisKey   string
	Secret     string
	Ttl        int
	RefreshTtl int
}

var Config config

type Claims struct {
	RefreshExpiresAt time.Time
	jwt.RegisteredClaims
}

func Login(id int) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, Claims{
		RefreshExpiresAt: time.Now().Add(time.Second * time.Duration(Config.RefreshTtl)),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(Config.Ttl))),
			ID:        strconv.Itoa(id),
		},
	}).SignedString([]byte(Config.Secret))

	if err == nil {
		err = g.Redis().HSet(context.Background(), Config.RedisKey, id, token).Err()
	}

	return token, err
}

func Parse(token string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.Secret), nil
	})
	if err == nil {
		if claims, ok := t.Claims.(*Claims); ok && t.Valid {
			result, err := g.Redis().HGet(context.Background(), Config.RedisKey, claims.ID).Result()
			if err != nil || token != result {
				return nil, jwt.ErrInvalidKey
			}
			return claims, nil
		}
	}
	return nil, err
}

func RefreshToken(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.Secret), nil
	})
	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		return "", err
	}
	if claims, ok := t.Claims.(*Claims); ok {
		result, err := g.Redis().HGet(context.Background(), Config.RedisKey, claims.ID).Result()
		if err != nil || token != result {
			return "", jwt.ErrInvalidKey
		}
		if time.Now().Before(claims.RefreshExpiresAt) {
			id, err := strconv.Atoi(claims.ID)
			if err != nil {
				return "", jwt.ErrInvalidKey
			}
			return Login(id)
		}
	}

	return "", jwt.ErrInvalidKey
}

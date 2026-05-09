package services

import (
	"AureliaReadsBackend/domain/entities"
	"AureliaReadsBackend/domain/services"
	"AureliaReadsBackend/domain/utils"
	"errors"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

type jwtServiceImpl struct {
	secret string
}

func NewJwtService(secret string) services.JwtService {
	return &jwtServiceImpl{secret: secret}
}

type myClaims struct {
	TypeToken utils.TokenType `json:"type_token"`
	UserID    entities.UserID `json:"user_id"`
	jwt.RegisteredClaims
}

func (i jwtServiceImpl) createToken(typeToken utils.TokenType, uid entities.UserID) (string, error) {
	claims := myClaims{
		TypeToken: typeToken,
		UserID:    uid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(i.secret))
}

func (i jwtServiceImpl) NewJwt(uid entities.UserID) (*entities.JwtTokens, error) {
	access, accessErr := i.createToken(utils.AccessToken, uid)
	refresh, refreshErr := i.createToken(utils.RefreshToken, uid)

	if accessErr != nil || refreshErr != nil {
		return nil, accessErr
	}

	return &entities.JwtTokens{
		Access:  access,
		Refresh: refresh,
	}, nil
}

func (i jwtServiceImpl) VerifyToken(token string, typeToken utils.TokenType) (*entities.JwtPayload, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&myClaims{},
		func(token *jwt.Token) (any, error) {
			return []byte(i.secret), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*myClaims)
	if !ok || !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.TypeToken != typeToken {
		return nil, errors.New("invalid token type")
	}

	return &entities.JwtPayload{
		UserID: claims.UserID,
	}, nil
}

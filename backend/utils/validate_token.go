package utils

import (
	"clientapp/config"
	"errors"

	"github.com/golang-jwt/jwt"
)

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {

	// Parseamos el token con los claims
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Verificamos que el metodo de la firma sea el esperado
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("metodo de firma incorrecto")
		}

		// Devolvemos la clave secreta para validar la firma
		return []byte(config.AppVars.JwtKey), nil
	})

	// Si hubo un error en el parsin o el token no es valido, lo retornamos
	if err != nil {
		return nil, err
	}

	// Verificamos si los claims son validos y los devolvemos en caso de serlo
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token no válido")
}

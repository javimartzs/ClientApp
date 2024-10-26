package middlewares

import (
	"clientapp/utils"
	"net/http"
	"strings"
)

// AuthMiddleware es un middleware que protege rutas que requieren autenticacion
// Verifica que el usuario tenga un JWT valido
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Obtenemos el token jwt desde el encabezado Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Falta el encabezado Authorization", http.StatusUnauthorized)
			return
		}

		// Extraemos el token JWT quitando el prefijo "Bearer "
		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(token)

		// Si el token es invalido o ha expirado devolvemos un error
		if err != nil {
			http.Error(w, "Token invalido o expirado", http.StatusUnauthorized)
			return
		}

		// Guardamos los datos del usuario en los encabezados de la solicitud
		r.Header.Set("user_id", claims["user_id"].(string))
		r.Header.Set("role", claims["role"].(string))

		// Continuaamos ocn el siguiente handler si el token es valido
		next.ServeHTTP(w, r)
	}
}

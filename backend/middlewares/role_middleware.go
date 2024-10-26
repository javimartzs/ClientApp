package middlewares

import "net/http"

// ManagerOnlyMiddleware verifica si el usuario tiene rol de "Manager"
//
//	Solo los usuarios con este rol pueden acceder a rutas protegidas
func ManagerOnlyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Obtenemos el rol del usuario desde los encabezados
		role := r.Header.Get("role")

		// Si el rol no es "manager", devolvemos un error de autorizacion
		if role != "manager" {
			http.Error(w, "Acceso no autorizado. Se requiere otro rol", http.StatusForbidden)
			return
		}

		// Si el rol es "manager", permitimos continuar con el siguiente handler
		next.ServeHTTP(w, r)
	}
}

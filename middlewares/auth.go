package middlewares

import (
	"net/http"
	"strings"

	"github.com/mhdiiilham/gorm/db"
	h "github.com/mhdiiilham/gorm/helpers"
	m "github.com/mhdiiilham/gorm/models"
)

// IsAuthenticated ...
func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Token format: Authorization = Bearer <token>
		bearerToken := r.Header.Get("Authorization")

		if bearerToken == "" {
			m.RespondError(w, http.StatusUnauthorized, "Not Authorized!")
			return
		}
		
		splitToken := strings.Split(bearerToken, " ")

		if len(splitToken) <= 1 {
			m.RespondError(w, http.StatusUnauthorized, "Not Authorized!")
			return
		}
		e, err := h.ExtractedJWT(splitToken[1])
		if err != nil {
			m.RespondError(w, http.StatusUnauthorized, "Not Authorized!")
			return
		}
		check := db.Connection().Where("email = ?", e)

		if check.Error != nil {
			m.RespondError(w, http.StatusUnauthorized, "Not Authorized!")
			return
		}

		next.ServeHTTP(w, r)
	})
}

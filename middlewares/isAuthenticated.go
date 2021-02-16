package middlewares

import (
	"goauth/helpers"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			helpers.Respond(w, helpers.BaseResponseBody{
				Data:    nil,
				Status:  http.StatusForbidden,
				Message: "Invalid Authorization",
			})
			return
		}

		token := strings.Split(tokenString, " ")
		if len(token) == 1 || (token[0] != "Bearer" && token[0] != "Basic") {
			helpers.Respond(w, helpers.BaseResponseBody{
				Data:    nil,
				Status:  http.StatusForbidden,
				Message: "Invalid Authorization",
			})
			return
		}
		var pass bool
		if token[0] == "Bearer" {
			pass = handleAuth(r, token[1])
		} else if token[0] == "Basic" {
			pass = handleBasicAuth(r)
		}

		if !pass {
			helpers.Respond(w, helpers.BaseResponseBody{
				Data:    nil,
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func handleBasicAuth(r *http.Request) bool {
	USERNAME_BASIC_AUTH := os.Getenv("USERNAME_BASIC_AUTH")
	PASSWORD_BASIC_AUTH := os.Getenv("PASSWORD_BASIC_AUTH")
	user, pass, _ := r.BasicAuth()
	if user == USERNAME_BASIC_AUTH && pass == PASSWORD_BASIC_AUTH {
		return true
	}
	return false
}

func handleAuth(r *http.Request, token string) bool {

	claims, err := helpers.VerifyToken(token)
	if err != nil {
		return false
	}
	id := claims.(jwt.MapClaims)["id"].(string)
	username := claims.(jwt.MapClaims)["username"].(string)
	r.Header.Set("user_id", id)
	r.Header.Set("username", username)
	return true
}

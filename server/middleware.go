package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-squads/reuni-server/response"
	"github.com/go-squads/reuni-server/services"
	"github.com/gorilla/mux"

	"github.com/go-squads/reuni-server/authenticator"
)

func withAuthenticator(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.TrimLeft(token, "Bearer")
		token = strings.TrimLeft(token, " ")
		obj, res := authenticator.VerifyUserJWToken(token)
		if res {
			log.Println("User", obj, "access from", r.RemoteAddr)
			next.ServeHTTP(w, r)
		} else {
			log.Println("Access from", r.RemoteAddr, "not authorized")
			response.ResponseHelper(w, http.StatusUnauthorized, response.ContentText, "")
			return
		}

	}
}

func validateAgentTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		serviceName := mux.Vars(r)["service_name"]
		res, err := services.ValidateTokenProcess(serviceName, token)
		if err != nil {
			response.ResponseHelper(w, http.StatusInternalServerError, response.ContentText, "")
			return
		}
		if res {
			next.ServeHTTP(w, r)
		} else {
			response.ResponseHelper(w, http.StatusForbidden, response.ContentText, "")
		}
	}
}
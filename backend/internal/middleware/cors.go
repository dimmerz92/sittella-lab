package middleware

import (
	"net/http"
	"os"
	"strings"
)

type CORSValidator struct {
	handler        http.Handler
	allowedDomains []string
}

func (v *CORSValidator) validDomain(origin string) bool {
	for _, domain := range v.allowedDomains {
		if origin == domain {
			return true
		}
	}
	return false
}

// wraps the mux
func (v *CORSValidator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if v.validDomain(origin) {
		// set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// handle OPTIONS method
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// handle actual request method
		v.handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

// instantiates validator
func NewValidator(handler http.Handler) *CORSValidator {
	allowedDomains := strings.Split(os.Getenv("ALLOWED"), ",")
	return &CORSValidator{handler, allowedDomains}
}

package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/imdinnesh/safepass/pkg/auth"
)

type RouteAuthType string

const (
	AuthPublic    RouteAuthType = "public"
	AuthProtected RouteAuthType = "protected"
	AuthOTP       RouteAuthType = "otp"
)

type AuthMiddleware struct {
	JWT *auth.JWTValidator
	OTP *auth.OTPStore
}

func NewAuthMiddleware(jwt *auth.JWTValidator, otp *auth.OTPStore) *AuthMiddleware {
	return &AuthMiddleware{JWT: jwt, OTP: otp}
}

func (a *AuthMiddleware) Wrap(next http.Handler, mode RouteAuthType) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if mode == AuthPublic {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "missing auth token", http.StatusUnauthorized)
			return
		}

		claims, err := a.JWT.Validate(authHeader)
		if err != nil {
			http.Error(w, "invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims["user_id"])
		ctx = context.WithValue(ctx, RoleKey, claims["role"])

		if mode == AuthOTP {
			sessionID := r.Header.Get("X-Session-ID")
			if sessionID == "" {
				http.Error(w, "missing session ID", http.StatusUnauthorized)
				return
			}
			ok, err := a.OTP.ValidateSession(ctx, sessionID)
			if !ok || err != nil {
				http.Error(w, "invalid session", http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

package middleware

import (
	"net/http"

	"github.com/unrolled/secure"
)

func Helmet(next http.Handler) http.Handler {
	helmet := secure.New(secure.Options{
		STSSeconds:            31536000,
		STSIncludeSubdomains:  true,
		STSPreload:            true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "script-src $NONCE",
	})

	return helmet.Handler(next)
}

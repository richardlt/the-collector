package users

import "fmt"

var config *c

type c struct{ jwtSecret, secret, appURI, redirectURI string }

// Init .
func Init(jwtSecret, secret, appURI string) {
	config = &c{
		jwtSecret:   jwtSecret,
		secret:      secret,
		appURI:      appURI,
		redirectURI: fmt.Sprintf("%s/api/auth/callback", appURI),
	}
}

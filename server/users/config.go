package users

import "fmt"

var config *c

type c struct{ jwtSecret, appURI, redirectURI string }

// Init .
func Init(jwtSecret, appURI string) {
	config = &c{
		jwtSecret:   jwtSecret,
		appURI:      appURI,
		redirectURI: fmt.Sprintf("%s/api/auth/callback", appURI),
	}
}

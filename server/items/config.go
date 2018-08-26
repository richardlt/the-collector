package items

var config *c

type c struct{ jwtSecret string }

// Init .
func Init(jwtSecret string) { config = &c{jwtSecret: jwtSecret} }

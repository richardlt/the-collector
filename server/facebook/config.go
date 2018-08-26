package facebook

var config *c

type c struct{ appID, appSecret string }

// Init .
func Init(appID, appSecret string) {
	config = &c{appID: appID, appSecret: appSecret}
}

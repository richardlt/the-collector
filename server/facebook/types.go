package facebook

// Error .
type Error struct {
	Type    string `json:"type"`
	Code    int    `json:"code"`
	TraceID string `json:"fbtrace_id"`
	Message string `json:"message"`
}

// User .
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Error *Error `json:"error"`
}

// AccessToken .
type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Error       *Error `json:"error"`
}

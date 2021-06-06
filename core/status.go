package core

const (
	Error            = 400
	Unauthorized     = 401
	UnauthorizedFail = 402
)

var statusText = map[int]string{
	Error:            "Error",
	Unauthorized:     "Unauthorized",
	UnauthorizedFail: "UnauthorizedFail",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}

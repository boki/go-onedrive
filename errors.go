package onedrive

import "errors"

// Errors
var (
	ErrFileTooLarge = errors.New("file is too large for simple upload")
)

// The Error type defines the basic structure of errors that are returned from
// the OneDrive API.
// See: http://onedrive.github.io/misc/errors.htm
type Error struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	InnerError *Error `json:"innererror"`
}

func (e Error) Error() string {
	return e.Message
}

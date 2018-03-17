package onedrive

import "errors"

// Errors
var (
	ErrFileTooLarge = errors.New("file is too large for simple upload")
)

// The Error type defines the basic structure of errors that are returned from
// the OneDrive API.
// See: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/concepts/errors
type Error struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	InnerError *Error `json:"innererror"`
}

func (e *Error) Error() string {
	return e.Message
}

// IsError checks if the error chain contains the expected OneDrive error code.
func (e *Error) IsError(code string) bool {
	for i := e; i != nil; i = i.InnerError {
		if i.Code == code {
			return true
		}
	}
	return false
}

package reloadly

//ErrorResponse is the default error response for twilio API's
type Error struct {
	TimeStamp string        `json:"timeStamp"`
	Message   string        `json:"message"`
	Path      string        `json:"path"`
	ErrorCode string        `json:"errorCode"`
	InfoLink  interface{}   `json:"infoLink"`
	Details   []interface{} `json:"details"`
}

//Implements Golang's error interface.
func (e *Error) Error() string{
	return e.Message
}
//GetErrorCode Gets the Error Code
func (e *Error) GetErrorCode() string{
	return e.ErrorCode
}
//GetErrorInfo Fetches InfoLink about the Error
func (e *Error) GetErrorInfo() interface{}{
	return e.InfoLink
}
//GetErrorDetails Returns Details about the Error!
func (e *Error) GetErrorDetails() []interface{}{
	return e.Details
}
//GetErrorDetails Returns the Errors Timestamp!
func (e *Error) GetErrorTimeStamp() string{
	return e.TimeStamp
}
//GetErrorPath Returns the Errors Path
func (e *Error) GetErrorPath() string{
	return e.Path
}
package Err

//ErrorResponse is the default error response for Reloadlys API's
type ErrorResponse struct {
	TimeStamp string        `json:"timeStamp"`
	Message   string        `json:"message"`
	Path      string        `json:"path"`
	ErrorCode string        `json:"errorCode"`
	InfoLink  interface{}   `json:"infoLink"`
	Details   []interface{} `json:"details"`
}

//Implements Golang's error interface.
func (e *ErrorResponse) Error() string{
	return e.Message
}
//GetErrorCode Gets the Error Code
func (e *ErrorResponse) GetErrorCode() string{
	return e.ErrorCode
}
//GetErrorInfo Fetches InfoLink about the Error
func (e *ErrorResponse) GetErrorInfo() interface{}{
	return e.InfoLink
}
//GetErrorDetails Returns Details about the Error!
func (e *ErrorResponse) GetErrorDetails() []interface{}{
	return e.Details
}
//GetErrorDetails Returns the Errors Timestamp!
func (e *ErrorResponse) GetErrorTimeStamp() string{
	return e.TimeStamp
}
//GetErrorPath Returns the Errors Path
func (e *ErrorResponse) GetErrorPath() string{
	return e.Path
}
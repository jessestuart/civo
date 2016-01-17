package api

import "fmt"

type HTTPError struct {
	Message    string
	URL        string
	StatusCode int
}

func (e *HTTPError) Error() string {
	return e.Message
}

func HTTPErrorNew(message, url string, statusCode int) *HTTPError {
	message = fmt.Sprintf("%s: %v (%s)", message, statusCode, url)
	return &HTTPError{Message: message, URL: url, StatusCode: statusCode}
}

func HTTPErrorNewf(message, url string, statusCode int, a ...interface{}) *HTTPError {
	return &HTTPError{Message: fmt.Sprintf(message, a...), URL: url, StatusCode: statusCode}
}

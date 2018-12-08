// Package jsend implements the JSend specification (https://labs.omniti.com/labs/jsend).
package jsend

import (
	"encoding/json"
	"net/http"
)

const (
	StatusSuccess = "success"
	StatusError   = "error"
	StatusFail    = "fail"
)

type jsendResponse struct {
	rw         http.ResponseWriter
	data       map[string]interface{}
	message    string
	statusCode int
	code       int
}

// Write writes a JSON encoded message that satisfies the JSend specification.
// Besides receiving an http.ResponseWriter, it also receives functional options to set the various JSend fields.
func Write(w http.ResponseWriter, options ...func(*jsendResponse)) (int, error) {
	resp := &jsendResponse{rw: w}

	// Call functional options.
	for _, option := range options {
		option(resp)
	}

	b, err := buildResponse(resp)

	if err != nil {
		return 0, err
	}

	resp.rw.Header().Set("Content-Type", "application/json")
	resp.rw.WriteHeader(resp.statusCode)
	return resp.rw.Write(b)
}

// Data is a functional option that sets the data field.
func Data(d map[string]interface{}) func(*jsendResponse) {
	return func(r *jsendResponse) {
		r.data = d
	}
}

// Message is a functional option that sets the message field.
func Message(m string) func(*jsendResponse) {
	return func(r *jsendResponse) {
		r.message = m
	}
}

// StatusCode is a functional option that sets the status code to be used in the HTTP response.
func StatusCode(c int) func(*jsendResponse) {
	return func(r *jsendResponse) {
		r.statusCode = c
	}
}

// Code is a functional option that sets the code field
func Code(c int) func(*jsendResponse) {
	return func(r *jsendResponse) {
		r.code = c
	}
}

// buildResponse builds and returns the JSON encoding for a valid JSend output.
func buildResponse(r *jsendResponse) ([]byte, error) {
	// Default to status code of 200.
	if r.statusCode == 0 {
		r.statusCode = 200
	}

	// Determine the status message.
	status := StatusSuccess
	switch {
	case r.statusCode >= 500:
		status = StatusError
		break
	case r.statusCode >= 400 && r.statusCode < 500:
		status = StatusFail
		break
	}

	// We initialise the response with a status field - all responses require it.
	resp := map[string]interface{}{
		"status": status,
	}

	// Required fields depend on status.
	switch {
	case status == StatusSuccess || status == StatusFail:
		// Required.
		if len(r.data) > 0 {
			resp["data"] = r.data
		} else {
			resp["data"] = nil
		}

		// Optional.
		if r.message != "" {
			resp["message"] = r.message
		}
		break
	case status == StatusError:
		// Optional.
		if len(r.data) > 0 {
			resp["data"] = r.data
		}

		// Required.
		if r.message == "" {
			resp["message"] = "Undefined error"
		} else {
			resp["message"] = r.message
		}

		// Optional.
		if r.code > 0 {
			resp["code"] = r.code
		}
	}

	j, err := json.Marshal(resp)
	if err != nil {
		return j, err
	}

	return j, nil
}

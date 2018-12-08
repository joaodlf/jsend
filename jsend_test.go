package jsend

import (
	"net/http/httptest"
	"testing"
)

var testCases = []struct {
	jsendResponse
	expected string
}{
	{jsendResponse{message: "test1"}, `{"data":null,"status":"success"}`},
	{jsendResponse{data: map[string]interface{}{"test": 2}}, `{"data":{"test":2},"status":"success"}`},
	{jsendResponse{statusCode: 400, message: "test3", data: map[string]interface{}{"test": 3}}, `{"data":{"test":3},"status":"fail"}`},
	{jsendResponse{statusCode: 500, code: 1, message: "test4", data: map[string]interface{}{"test": 4}}, `{"code":1,"data":{"test":4},"message":"test4","status":"error"}`},
	{jsendResponse{statusCode: 500, code: 1}, `{"code":1,"message":"Undefined error","status":"error"}`},
}

func TestSend(t *testing.T) {
	for _, tc := range testCases {
		rr := httptest.NewRecorder()

		_, err := Write(rr,
			Message(tc.message),
			StatusCode(tc.statusCode),
			Code(tc.code),
			Data(tc.data))

		if err != nil {
			t.Error(err)
		}

		body := rr.Body.String()

		if rr.Body.String() != tc.expected {
			t.Errorf("handler returned unexpected body: got %v want %v", body, tc.expected)
		}
	}
}

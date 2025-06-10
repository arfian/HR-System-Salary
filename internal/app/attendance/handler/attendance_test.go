package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mocks "hr-system-salary/internal/app/attendance/service/mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_AddAttendanceEmployee(t *testing.T) {
	type depsMock struct {
		attendenceSvc *mocks.IAttendanceService
	}

	type resp struct {
		message    string
		success    bool
		request_id string
	}

	tests := []struct {
		message    string
		request_id string
		success    bool
		mockCall   func(m depsMock)
	}{
		{
			message:  "fail: invalid token",
			success:  true,
			mockCall: func(m depsMock) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.message, func(t *testing.T) {
			attendenceSvc := mocks.NewAttendanceServiceItf(t)
			m := depsMock{
				attendenceSvc: attendenceSvc,
			}

			tt.mockCall(m)

			h := &handler{
				// attendanceService: attendenceSvc,
			}

			req := httptest.NewRequest("POST", fmt.Sprintf("http://localhost:8089/v1/api/attendance/employee"), nil)
			w := httptest.NewRecorder()

			resultBody := ginHandleFunc(req, w, h.AddAttendanceEmployee, "/v1/api/attendance/employee", http.MethodGet)
			if tt.success {
				assert.Contains(t, resultBody["success"], tt.success)
			}
		})
	}
}

func ginHandleFunc(req *http.Request, w *httptest.ResponseRecorder, handler gin.HandlerFunc, url string, method string) map[string]interface{} {
	r := gin.Default()
	r.Handle(method, url, handler)

	r.ServeHTTP(w, req)
	readBody, _ := io.ReadAll(w.Body)
	var resultBody map[string]interface{}
	json.Unmarshal(readBody, &resultBody)

	return resultBody
}

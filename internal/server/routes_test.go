package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelloWorldHandler(t *testing.T) {
	s := &Server{}
	r := gin.New()
	r.GET("/", s.HelloWorldHandler)
	// Create a test HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	// Serve the HTTP request
	r.ServeHTTP(rr, req)
	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	// Check the response body
	expected := "{\"message\":\"Hello World\"}"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestTripAdvisorHandlerWithValidURL(t *testing.T) {
	s := &Server{}
	r := gin.New()
	r.POST("/tripadvisor", s.TripAdvisorHandler)

	// Tạo request body với URL hợp lệ
	requestBody := `{"url": "https://www.tripadvisor.com/Hotel_Review-g298085-d27540572-Reviews-V_hotel_Da_Nang_Beach-Da_Nang.html"}`
	req, err := http.NewRequest("POST", "/tripadvisor", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Tạo ResponseRecorder để ghi lại phản hồi
	rr := httptest.NewRecorder()

	// Gửi request đến handler
	r.ServeHTTP(rr, req)

	// Kiểm tra status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Kiểm tra nội dung phản hồi không rỗng
	if rr.Body.String() == "" {
		t.Errorf("Handler returned empty body, expected non-empty data")
	}
}

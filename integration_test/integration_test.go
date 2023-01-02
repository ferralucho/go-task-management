package integration_test

import (
	. "github.com/Eun/go-hit"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

const (
	// Attempts connection
	host       = "localhost:8082"
	healthPath = "http://" + host + "/healthz"
	attempts   = 2

	// HTTP REST
	basePath = "http://" + host + "/v1"
)

func TestMain(m *testing.M) {
	err := healthCheck(attempts)
	if err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", host, err)
	}

	log.Printf("Integration tests: host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck(attempts int) error {
	var err error

	for attempts > 0 {
		err = Do(Get(healthPath), Expect().Status().Equal(http.StatusOK))
		if err == nil {
			return nil
		}

		log.Printf("Integration tests: url %s is not available, attempts left: %d", healthPath, attempts)
		time.Sleep(time.Second)
		attempts--
	}

	return err
}

// HTTP POST: /v1/management/task
func TestHTTPDoCreateTask(t *testing.T) {
	body := `{
    "type": "task",
    "title": "Clean the rocket",
    "category": "Maintenance"
}`
	Test(t,
		Description("DoTask Success"),
		Post(basePath+"/management/task"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
	)

	body = `{
	}`
	Test(t,
		Description("DoTask Fail"),
		Post(basePath+"management/task"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".error").Equal("invalid request body"),
	)
}

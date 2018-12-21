package kind_test

import (
	"fmt"
	"github.com/kris-nova/kind-test/assertion"
	"github.com/kris-nova/logger"
	"testing"
	"time"
)


const (	APISleepSeconds   = 5
	APISocketAttempts = 40
	APIEndpoint       = "localhost"
)

var (
	APIPort = ""
)

func TestApiListen(t *testing.T) {
	success := false
	for i := 0; i < APISocketAttempts; i++ {
		_, err := assertion.AssertTcpSocketAcceptsConnection(fmt.Sprintf("%s:%s", APIEndpoint, APIPort), "opening a new socket connection against the Kubernetes API")
		if err != nil {
			logger.Info("Attempting to open a socket to the Kubernetes API: %v...\n", err)
			time.Sleep(time.Duration(APISleepSeconds) * time.Second)
			continue
		}
		success = true
	}
	if !success {
		t.Fatalf("Unable to connect to Kubernetes API")
	}
}


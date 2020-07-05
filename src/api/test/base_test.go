package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/skhut/go-testing/src/api/app"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	fmt.Println("about to start the application")
	go app.StartApp()
	fmt.Println("application started. about to start test cases.....")
	os.Exit(m.Run())
}

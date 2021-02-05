package servicev1

import (
	"log"
	"os"
	"testing"
)

func BeforeTest() {}

func TestMain(m *testing.M) {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags)
	BeforeTest()
	defer AfterTest()
	os.Exit(m.Run())
}

func AfterTest() {}

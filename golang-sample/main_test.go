package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BeforeTest() {}

func TestMain(m *testing.M) {
	log.SetFlags(log.LstdFlags)
	log.SetOutput(os.Stdout)
	BeforeTest()
	defer AfterTest()
	os.Exit(m.Run())
}

func AfterTest() {}

func Test_All(t *testing.T) {
	assert.True(t, true)
}

func Benchmark_All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Println("main")
	}
}

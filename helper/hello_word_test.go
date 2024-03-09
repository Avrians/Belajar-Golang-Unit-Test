package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Budi")
	if result != "Hello Budi" {
		// error
		panic("Result is not Hello Budi")
	}
}

func TestHelloWorldBudi(t *testing.T) {
	result := HelloWorld("Budi")
	if result != "Hello Budi" {
		// error
		panic("Result is not Hello Budi")
	}
}
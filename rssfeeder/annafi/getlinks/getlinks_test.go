package main

import (
    "testing"
    "os"
    "fmt"
)
func TestMain(m *testing.M) {
    setup()
    code := m.Run() 
    shutdown()
    os.Exit(code)
}

func setup() {
	
	fmt.Println("Start")
	
}
func shutdown() {
	fmt.Println("Shutdown")
	
}

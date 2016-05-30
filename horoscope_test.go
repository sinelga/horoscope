package main

import (
    "testing"
    "os"
)
func TestMain(m *testing.M) {
//    setup()
    code := m.Run() 
//    shutdown()
    os.Exit(code)
}



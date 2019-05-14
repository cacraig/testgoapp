package main

import (
    "testgoapp/app"
    "fmt"
)

func main() {
    port := ":3000"
    fmt.Println("Hello world!")
    app := &app.App{}
    app.Run(port)
}
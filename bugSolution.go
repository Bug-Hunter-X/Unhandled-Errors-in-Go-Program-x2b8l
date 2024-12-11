package main

import (

    "fmt"

    "os"

)

func main() {

    file, err := os.Open("my_file.txt")

    if err != nil {

        fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)

        os.Exit(1)

    }

    defer file.Close()

    // Process the file here...

    fmt.Println("File processed successfully")

}
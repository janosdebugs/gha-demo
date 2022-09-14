package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func main() {
    fmt.Printf("The workspace directory is located at %s with the following contents:\n", os.Getenv("GITHUB_WORKSPACE"))

    files, err := ioutil.ReadDir(os.Getenv("GITHUB_WORKSPACE"))
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
        fmt.Println(f.Name())
    }
}

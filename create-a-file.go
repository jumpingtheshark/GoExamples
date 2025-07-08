https://gosamples.dev/create-file/#:~:text=To%20create%20a%20new%20empty,removed%20without%20deleting%20the%20file.

package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    f, err := os.Create("testFile.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    fmt.Println(f.Name())
}

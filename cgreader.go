package cgreader

import (
    "fmt"
    //"io"
    //"io/ioutil"
)

func RunProgram(input string) {
    fmt.Println("test")
    var hello string
    fmt.Scanf("%s\n", &hello)
    fmt.Printf("%s", "Word: " + hello)

    /*file, err := ioutil.ReadFile(input)
    if err == nil {
        fmt.Printf("%s", file)
    }*/
}
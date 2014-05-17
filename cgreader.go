package cgreader

import (
    "fmt"
    //"io"
    "io/ioutil"
)

func RunProgram(input string) {
    file, err := ioutil.ReadFile(input)
    if err == nil {
        fmt.Printf("%s", file)
    }
}
Codingame Reader
=================

Small Go package to simulate the Codingame programs offline on your computer.

# How to use
The implementation is quite straightforward and simple. Your offline will be almost identical as your online codingame code, with the difference being that the input comes via _cgreader_ rather than via _stdin_.

## Template: manual input program

  ```
  package main

  import (
      "github.com/glendc/cgreader"                      // cgreader package
  )

  func main() {
      cgreader.RunAndValidateProgramManual(
          "<INPUT TEXT FILE>",                          // program input
          "<OUTPUT TEXT FILE>",                         // expected output
          true,                                         // show output?
          func(ch <-chan string) string {               // program main
              return "<YOUR FINAL OUTPUT HERE>"         // program output
          })
  }

  ```
  
## Template: flow input program

  ```
  package main

  import (
	  "github.com/glendc/cgreader"                      // cgreader package
  )

  type Program struct {
  	// Variables needed in your program logic
  }

  func (p *Program) Update(input string) {
  	// Called as long as receiving input
  }

  func (p *Program) GetOutput() string {
	  return "<FINAL OUTPUT HERE>"
  }

  func main() {
	  cgreader.RunAndValidateFlowProgram(
		  "<INPUT FILE>",                               // program input
		  "<TEST FILE>",								// expected output
		  true,											// show output?
		  &Program{})									// program
}

  ```

# Example(s):

## Ascii Art

You can find the source code of the example [here](https://github.com/GlenDC/CodingGame/blob/master/go/ascii_art.go).

It will output the following output:

  ```
  ### 
  #   
  ##  
  #   
  ### 
  
  Program is correct!
  ```

# Feedback

Any feedback is welcome, and can be given along the bug reports as an issue report.

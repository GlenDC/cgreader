Codingame Reader
=================

Small Go package to simulate the Codingame programs offline on your computer.

# Index

1. [Quick Guide](#quick-guide)
  1. [Types of programs](#types-of-programs)
  1. [Manual Program](#manual-program)
    1. [Run a manual program](#run-a-manual-program)
      1. Template
      1. Example
    1. [Run and validate a manual program](#run-and-validate-a-manual-program)
      1. Template
      1. Example
    1. [Run and self-validate a manual program](#run-and-self-validate-a-manual-program)
      1. Template
      1. Example
  1. [Flow Program](#flow-program)
    1. [Run a flow program](#run-a-flow-program)
      1. Template
      1. Example
    1. [Run and validate a flow program](#run-and-validate-a-flow-program)
      1. Template
      1. Example
    1. [Run and self-validate a flow program](#run-and-self-validate-a-flow-program)
      1. Template
      1. Example
  1. [Target Program](#target-program)
    1. [Run a target program](#run-a-target-program)
      1. Template
      1. Example
    1. [Run and validate a target program](#run-and-validate-a-target-program)
      1. Template
      1. Example
    1. [Run and self-validate a target program](#run-and-self-validate-a-target-program)
      1. Template
      1. Example
1. [Codingame Offline Code Convertor](#Codingame Offline Code Convertor)
  1. [How to use](#how-to-use)
1. Feedback

# Quick Guide
The implementation is quite straightforward and simple. Your offline will be almost identical as your online codingame code, with the difference being that the input comes via _cgreader_ rather than via _stdin_.

## Types of programs

_TODO: write an awesome chapter_

## Manual Program

### Run a manual program

#### Template

_TODO: write this template_

#### Example

_TODO: write this example_

### Run and validate a manual program

#### Template

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

#### Example

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

### Run and self-validate a manual program

#### Template

_TODO: write this template_

#### Example

_TODO: write this example_

## Flow Program

### Run a flow program

#### Template

_TODO: write this template_

#### Example

_TODO: write this example_

### Run and validate a flow program

#### Template

_TODO: write this template_

#### Example

_TODO: write this example_

### Run and self-validate a flow program

#### Template

_TODO: write this template_

#### Example

_TODO: write this example_

## Manual Program

### Run a target program

#### Template

_TODO: write this template_

#### Example

_TODO: write this example_

### Run and validate a target program

#### Template

_TODO: write this template_

#### Example

_TODO: write this example_

### Run and self-validate a target program

#### Template

_TODO: write this template_

#### Example

_TODO: write this example_

# Codingame Offline Code Convertor

You might want to test out your code on [the official Codingame website](http://www.codingame.com). This is possible thanks to the Codingame Offline Code Convertor or _cgocc_ in short. It' a command line utility that:

1. removes all the cgreader specific code
1. adds all the necacary code in order to work in the online environment
1. copy the parsed code to your clipboard

## How to use

Using the cgocc utility is so easy that it can be summarized in 3 steps:

1. Type and enter in your terminal: ``cgocc [go program]``
  1. Make sure your code has been formatted with the go fmt utility
  2. The cgocc utility only works for programs that use the cgreader pkg
2. Go to the correct challenge to be found on the training page on [the official Codingame website](http://www.codingame.com).
3. Paste your code & run it!

# Feedback

Any feedback is welcome, and can be given along the bug reports as an issue report or by [mailing me](mailto:contact@glendc.com).

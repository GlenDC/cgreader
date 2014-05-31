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
1. [Codingame Offline Code Convertor](#codingame-offline-code-convertor)
  1. [How to use](#how-to-use)
1. Feedback

# Quick Guide
The implementation is quite straightforward and simple. Your offline will be almost identical as your online codingame code, with the difference being that the input comes via _cgreader_ rather than via _stdin_.

You can find the [descriptions](https://github.com/glendc/CodingGame/desc), [input](https://github.com/glendc/CodingGame/desc) and optionally the [output](https://github.com/glendc/CodingGame/output) text files all [here](https://github.com/glendc/CodingGame) or on [the official Codingame website](http://www.codingame.com). They are also the authors of all these files.

## Types of programs

[Codingame](http://www.codingame.com) has a lot challenges. These challenges can be deviced in types of programs based on how they receive input and what the goal of the challenge is.

1. [Manual Program](#manual-program): This is the most simple program and just requires you to write a simple _main_ function that takes a _string channel_ as its input. This channel will give you the input line by line and it's up to you how to interpret the received input. The output of the program has to be returned at the end of this method.
1. [Flow Program](#flow-program): This program is quite similar to a manual program. The biggest difference is that you'll have to define a struct which has two methods. Each frame the _Update_ method will be called, receiving a line of input. At the end of the program the output will be asked via the _GetOutput_ method.
1. [Target Program](#target-program): Some challenges are based on win and lose conditions. These are the most complex program and require extra work from the user in order to do these challenges offline, as you'll have to write the logic of the challenge, on top of your usual challenge code. So how does a target program works?
  1. You'll write a struct based on the _TargetProgram_ interface
  1. The initial input will be parsed and have to be manually interpred by you via the _InitialInput_ method.
  1. The program runs and calls each frame the _Update_ method, which will return your output for that frame
    1. This output can also be traced if wanted.
  1. Each frame your output will be used and update the game state via the _DoMove_ method
  1. The program exits if the _LoseConditionCheck_- or/and _WinConditionCheck_ method returns true

All programs, except target programs can be either:

1. run the program
1. run and validate the program based on a test text file
1. run and validate the program based on a validation lambda

With all three options you can also echo your final output if wanted.

Suggestions to improve a type of program, or to define a new type of program are welcome and can be filed as an issue or a pull request.

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
<<<<<<< HEAD
=======
  
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
>>>>>>> 892d2c107be0e463249d274242743016ff4299e6

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

Any feedback is welcome, and can be given along the bug reports as an issue report or by [mailing me](mailto:contact@glendc.com). Pull requests are also welcome on the condition that your commits are clean and bring additional values to the cgreader pkg.

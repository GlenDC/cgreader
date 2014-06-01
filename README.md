Codingame Reader
=================

Small Go package to simulate the Codingame programs offline on your computer.

# Index

1. [Quick Guide](#quick-guide)
  1. [Types of programs](#types-of-programs)
  1. [Manual Program](#manual-program)
    1. [Run a manual program](#run-a-manual-program)
    1. [Run and validate a manual program](#run-and-validate-a-manual-program)
    1. [Run and self-validate a manual program](#run-and-self-validate-a-manual-program)
  1. [Flow Program](#flow-program)
    1. [Run a flow program](#run-a-flow-program)
    1. [Run and validate a flow program](#run-and-validate-a-flow-program)
    1. [Run and self-validate a flow program](#run-and-self-validate-a-flow-program)
  1. [Target Program](#target-program)
  1. [How to use](#how-to-use)
1. Feedback

# Quick Guide
The implementation is quite straightforward and simple. Your offline will be almost identical as your online codingame code, with the difference being that the input comes via _cgreader_ rather than via _stdin_.

You can find the [descriptions](https://github.com/GlenDC/Codingame/tree/master/descriptions), [input](https://github.com/GlenDC/Codingame/tree/master/input) and optionally the [output](https://github.com/GlenDC/Codingame/tree/master/output) text files all [here](https://github.com/glendc/Codingame) or on [the official Codingame website](http://www.codingame.com).

## Types of programs

[Codingame](http://www.codingame.com) has a lot challenges. These challenges can be devided in types of programs based on how they receive input and what the goal of the challenge is.

1. [Manual Program](#manual-program): This is the most simple program and just requires you to write a simple _main_ function that takes a _string channel_ as its input. This channel will give you the input line by line and it's up to you how to interpret the received input. The output of the program has to be returned at the end of this method.
1. [Flow Program](#flow-program): This program is quite similar to a manual program. The biggest difference is that you'll have to define a struct which has two methods. Each frame the _Update_ method will be called, receiving a line of input. At the end of the program the output will be asked via the _GetOutput_ method.
1. [Target Program](#target-program): Some challenges are based on win and lose conditions. These are the most complex program and require extra work from the user in order to do these challenges offline, as you'll have to write the logic of the challenge, on top of your usual challenge code. So how does a target program works?
  1. You'll write a struct based on the _TargetProgram_ interface
  1. The initial input will be parsed and have to be manually interpred by you via the _InitialInput_ method.
  1. The program runs and calls each frame the _Update_ method, using the input given via the _GetInput_ method. _Update_ will return your output for that frame
    1. This output can also be traced if wanted.
  1. Each frame your output will be used and update the game state via the _SetOutput_ method
  1. The program exits if the _LoseConditionCheck_- or/and _WinConditionCheck_ method returns true

All programs, except target programs can either:

1. run the program
1. run and validate the program based on a test text file
1. run and validate the program based on a validation lambda

With all three options you can also echo your final output if wanted.

You can find template for bigger (target) programs [here](https://github.com/GlenDC/Codingame/tree/master/templates/go), to allow you to just concentrate on the challenge(s) and not distract you too much with the program logic itself. Feel free to add templates yourself or improve existing ones via a pull request on that repository.

Suggestions to improve a type of program, or to define a new type of program are welcome and can be filed as an issue or a pull request.

## Manual Program

### Run a manual program

#### Template

    package main

    import (
      "github.com/glendc/cgreader"                      // cgreader package
    )

    func main() {
      cgreader.RunManualProgram(
          "<INPUT TEXT FILE>",                          // program input
          func(ch <-chan string) string {               // program main
              return "<YOUR FINAL OUTPUT HERE>"         // program output
      })
    }

#### Example

    package main

    import (
      "fmt"
      "github.com/glendc/cgreader"
      "strings"
    )

    func main() {
      cgreader.RunManualProgram(
        "../../input/ascii_1.txt",
        func(ch <-chan string) string {
          var width, height int
          var text string

          fmt.Sscanln(<-ch, &width)
          fmt.Sscanln(<-ch, &height)
          fmt.Sscanln(<-ch, &text)

          text = strings.ToUpper(text)

          ascii := make([]string, height)
          for i := 0; i < height; i++ {
            ascii[i] = <-ch
          }

          output := make([]string, height)
          for _, char := range text {
            character := int(char) - 65
            if character < 0 || character > 26 {
              character = 26
            }
            for i := range output {
              position := character * width
              output[i] += ascii[i][position : position+width]
            }
          }

          var program_output string

          for _, line := range output {
            program_output += line + "\n"
          }

          return program_output
        })
    }

##### Output:

    ### 
    #   
    ##  
    #   
    ### 

### Run and validate a manual program

#### Template

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

#### Example

    package main

    import (
      "fmt"
      "github.com/glendc/cgreader"
      "strings"
    )

    func main() {
      cgreader.RunAndValidateManualProgram(
        "../input/ascii_1.txt",
        "../output/ascii_1.txt",
        true,
        func(ch <-chan string) string {
          var width, height int
          var text string

          fmt.Sscanln(<-ch, &width)
          fmt.Sscanln(<-ch, &height)
          fmt.Sscanln(<-ch, &text)

          text = strings.ToUpper(text)

          ascii := make([]string, height)
          for i := 0; i < height; i++ {
            ascii[i] = <-ch
          }

          output := make([]string, height)
          for _, char := range text {
            character := int(char) - 65
            if character < 0 || character > 26 {
              character = 26
            }
            for i := range output {
              position := character * width
              output[i] += ascii[i][position : position+width]
            }
          }

          var program_output string

          for _, line := range output {
            program_output += line + "\n"
          }

          return program_output
        })
    }

##### Output:

    ### 
    #   
    ##  
    #   
    ### 

    Program is correct!

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

## Target Program

#### Template

    package main

    import (
      "github.com/glendc/cgreader"
    )

    type Program struct {
      // define a structure, and optionally define member variables
    }

    func (program *Program) ParseInitialData(ch <-chan string) {
      // parse the initial data, just like in a manual program
    }

    func (program *Program) GetInput() (ch chan string) {
      ch = make(chan string)
      go func() {
        // pass the challenge input into the channel
      }()
      return
    }

    func (program *Program) Update(ch <-chan string) string {
      // your solution logic will be defined here
      // return an output string, based on the input given via the channel
    }

    func (program *Program) SetOutput(output string) string {
      // define challenge logic using the output of the user
      // return optional trace message
    }

    func (program *Program) LoseConditionCheck() bool {
      // return true if user has lost
    }

    func (program *Program) WinConditionCheck() bool {
      // return true if user has won
    }

    func main() {
      cgreader.RunTargetProgram( // Execute your program offline
        "program_1.txt", // The challenge input file
        true,            // Trace output && state?
        &Program{})      // Create the program && pass it by reference
    }

#### Example

    package main

    import (
      "fmt"
      "github.com/glendc/cgreader"
      "strings"
    )

    type Vector struct {
      x, y int
    }

    type Ragnarok struct {
      thor, target, dimensions Vector
      energy                   int
    }

    func GetDirection(a, b string, x, y, v int) <-chan string {
      ch := make(chan string)
      go func() {
        difference := x - y
        switch {
        case difference < 0:
          ch <- a
        case difference > 0:
          ch <- b
        default:
          ch <- ""
        }
        close(ch)
      }()
      return ch
    }

    func (ragnarok *Ragnarok) ParseInitialData(ch <-chan string) {
      fmt.Sscanf(
        <-ch,
        "%d %d %d %d %d %d %d \n",
        &ragnarok.dimensions.x,
        &ragnarok.dimensions.y,
        &ragnarok.thor.x,
        &ragnarok.thor.y,
        &ragnarok.target.x,
        &ragnarok.target.y,
        &ragnarok.energy)
    }

    func (ragnarok *Ragnarok) GetInput() (ch chan string) {
      ch = make(chan string)
      go func() {
        ch <- fmt.Sprintf("%d", ragnarok.energy)
      }()
      return
    }

    func (ragnarok *Ragnarok) Update(ch <-chan string) string {
      channel_b := GetDirection("N", "S", ragnarok.target.y, ragnarok.thor.y, ragnarok.thor.y)
      channel_a := GetDirection("E", "W", ragnarok.thor.x, ragnarok.target.x, ragnarok.thor.x)

      result_b := <-channel_b
      result_a := <-channel_a

      return fmt.Sprint(result_b + result_a)
    }

    func (ragnarok *Ragnarok) SetOutput(output string) string {
      if strings.Contains(output, "N") {
        ragnarok.thor.y -= 1
      } else if strings.Contains(output, "S") {
        ragnarok.thor.y += 1
      }

      if strings.Contains(output, "E") {
        ragnarok.thor.x += 1
      } else if strings.Contains(output, "W") {
        ragnarok.thor.x -= 1
      }

      ragnarok.energy -= 1

      return fmt.Sprintf(
        "Target = (%d,%d)\nThor = (%d,%d)\nEnergy = %d",
        ragnarok.target.x,
        ragnarok.target.y,
        ragnarok.thor.x,
        ragnarok.thor.y,
        ragnarok.energy)
    }

    func (ragnarok *Ragnarok) LoseConditionCheck() bool {
      if ragnarok.energy <= 0 {
        return true
      }

      x, y := ragnarok.thor.x, ragnarok.thor.y
      dx, dy := ragnarok.dimensions.x, ragnarok.dimensions.y

      if x < 0 || x >= dx || y < 0 || y >= dy {
        return true
      }

      return false
    }

    func (ragnarok *Ragnarok) WinConditionCheck() bool {
      return ragnarok.target == ragnarok.thor
    }

    func main() {
      cgreader.RunTargetProgram("../../input/ragnarok_1.txt", true, &Ragnarok{})
    }


##### Output:

    E
    Target = (10,8)
    Thor = (8,8)
    Energy = 2

    E
    Target = (10,8)
    Thor = (9,8)
    Energy = 1

    E
    Target = (10,8)
    Thor = (10,8)
    Energy = 0
    
    Program is correct!

# Codingame Offline Code Convertor

You might want to test out your code on [the official Codingame website](http://www.codingame.com). This is possible thanks to the Codingame Offline Code Convertor or _cgocc_ in short. It' a command line utility that:

1. removes all the cgreader specific code
1. adds all the necacary code in order to work in the online environment
1. copy the parsed code to your clipboard

_Note: This utility hasn't been developed yet. [Mail me](mailto:contact@glendc.com) for more information about **cgocc**_.

#### How to use

Using the cgocc utility is so easy that it can be summarized in 3 steps:

1. Type and enter in your terminal: ``cgocc [go program]``
  1. Make sure your code has been formatted with the go fmt utility
  2. The cgocc utility only works for programs that use the cgreader pkg
2. Go to the correct challenge to be found on the training page on [the official Codingame website](http://www.codingame.com).
3. Paste your code & run it!

# Feedback

Any feedback is welcome, and can be given along the bug reports as an issue report or by [mailing me](mailto:contact@glendc.com). Pull requests are also welcome on the condition that your commits are clean and bring additional values to the cgreader pkg.

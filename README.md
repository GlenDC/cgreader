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
  1. [Target Program](#target-program)
    1. [Predefined Target Challenges](#predefined-target-challenges)
      1. [Ragnarok Example](#ragnarok-example)
      1. [List of Predefined Challenges](#list-of-predefined-challenges)
    1. [Template and Example](#template-and-example)
  1. [Challenge map in your terminal](#challenge-map-in-your-terminal)
1. [Feedback](#feedback)

# Quick Guide
The implementation is quite straightforward and simple. Your offline will be almost identical as your online codingame code, with the difference being that the input comes via _cgreader_ rather than via _stdin_.

You can find the [descriptions](https://github.com/GlenDC/Codingame/tree/master/descriptions), [input](https://github.com/GlenDC/Codingame/tree/master/input) and optionally the [output](https://github.com/GlenDC/Codingame/tree/master/output) text files all [here](https://github.com/glendc/Codingame) or on [the official Codingame website](http://www.codingame.com).

## Types of programs

[Codingame](http://www.codingame.com) has a lot challenges. These challenges can be devided in types of programs based on how they receive input and what the goal of the challenge is.

1. [Manual Program](#manual-program): This is the most simple program and just requires you to write a simple _main_ function that takes a _string channel_ as its input. This channel will give you the input line by line and it's up to you how to interpret the received input. The output of the program has to be returned at the end of this method.
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

**Target Programs shouldn't be used as a user, instead use the correct [Predefined Challenge Program](#list-of-predefined-challenges)!**

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

## Target Program

### Predefined Target Challenges

The target program challenge type was created to allow you to play more complex challenges, such as Ragnarok, offline. However with target programs you still need to program the Challenge and AI logic yourself, which isn't the goal of the Codingame challenges at all. [The Predefined Target Challenges](#list-of-predefined-challenges) allow you to start on the challenge instantly and it keeps your code base exactly the same as if it were an online submission. 

#### How to convert your offline PT solution code to use online?

There is no real reason why you would want to convert your offline PT challenge code, but let's say you want to do so. It's possible and easy, as your code base will remain quite similar.

Let's say we have the follow _psuedo_ offline [Ragnarok](https://raw.githubusercontent.com/GlenDC/Codingame/master/descriptions/ragnarok.md) PT solution code:

    package man
    
    import (
      "github.com/glendc/cgreader"
      // packages...
    )
    
    // definition of functions, types and variables...
    
    func Initialize(ch <-chan string) {
      // parse the initial input, no output expected...
    }
    
    func Update(ch <-chan string) string {
      // the code of your solution logic will be defined here...
      // return "output"
    }
    
    func main() {
      cgreader.RunRagnarok("ragnarok_1.txt", true, Initialize, Update)
    }
    
After your converted this code **manually**, you will end up with the following online version:

    package main
    
    import (
       // packages...
    )
    
    // definition of functions, types and variables...
    
    func main() {
      // parse the initial input, no output expected...
      
      for {
        // the code of your solution logic will be defined here...
        fmt.Print("output")
      }
    }
    
As you can see it's quite similar. On top of this you'll have to convert code that makes use of the channel input parameter, to use the standard input instead. (e.g. ``fmt.Sscanf(<-ch`` to ``fmt.Scanf(``)

#### Ragnarok Example

You can find my ragnarok solution code [here](https://github.com/GlenDC/Codingame/blob/master/solutions/go/ragnarok.go). Thanks to all [the predefined Ragnarok logic](https://github.com/GlenDC/cgreader/blob/master/ragnarok.go), I can simply run the program via the ``cgreader.RunRagnarok`` function, which makes use of the default way to run a Target Program defined [here](https://github.com/GlenDC/cgreader/blob/master/cgreader.go). All this results in a clean and good looking solution. 

#### List of Predefined Challenges

_You can find templates for all [the ready-to-use challenges](#ready-to-use-challenges) [here](https://github.com/GlenDC/Codingame/tree/master/templates/go), this will allow you to solve the challenge straight away!_

##### ready-to-use challenges:

* **Power of Thor**: [description](https://raw.githubusercontent.com/GlenDC/Codingame/master/descriptions/ragnarok.md), [template](https://github.com/GlenDC/Codingame/blob/master/templates/go/ragnarok.go) and [solution](https://github.com/GlenDC/Codingame/blob/master/solutions/go/ragnarok.go)
* **Thor Vs. Giants**: [description](https://raw.githubusercontent.com/GlenDC/Codingame/master/descriptions/ragnarok_giants.md), [template](https://github.com/GlenDC/Codingame/blob/master/templates/go/ragnarok_giants.go) and [solution](https://github.com/GlenDC/Codingame/blob/master/solutions/go/ragnarok_giants.go)

##### challenges to be developed:

* Skynet Final - Level 1
* Skynet - The Chasm
* Kirk's Quest - The descent
* Mars Lander - Level 1
* Indiana - Level 1
* Mars Lander - Level 2
* Skynet Finale - Level 2
* Skynet - The Bridge
* Kirk's Quest - The labyrinth
* Indiana - Level 2
* Indiana - Level 3
* Mars Lander - Level 3

_Contributions on the "reverse engineering" of these challenges are more than welcome!_

### Template and Example

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
    
# Challenge map in your terminal

For challenges like [ragnarok](https://raw.githubusercontent.com/GlenDC/Codingame/master/descriptions/ragnarok.md) you might want to have a map, like you would have in [the online Codingame version](http://www.codingame.com). For this you can use the ``cgreader.DrawMap(...)`` function. You can see a working ragnarok solution with map [here](https://github.com/GlenDC/Codingame/blob/master/solutions/go/ragnarok.go).

### Ragnarok Map Example

##### Map after first move:

    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  T  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  H  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  +  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .

##### Map after last move:

    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  H  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  +  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  +  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  +  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  +  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  +  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  +  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  +  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  +  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  +  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  +  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  +  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  +  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  +  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  +  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  +  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  +  .  .  .  .  .  .  .  .  
    .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 

# Feedback

Any feedback is welcome, and can be given along the bug reports as an issue report or by [mailing me](mailto:contact@glendc.com). Pull requests are also welcome on the condition that your commits are clean and bring additional values to the cgreader pkg.

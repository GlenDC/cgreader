Codingame Reader
=================

Small Go package to simulate the Codingame programs offline on your computer.

# Index

1. [Quick Guide](#quick-guide)
  1. [Types of programs](#types-of-programs)
  1. [Manual Program](#manual-program)
    1. [Run a manual program](#run-a-manual-program)
    1. [Run and validate a manual program](#run-and-validate-a-manual-program)
    1. [Run your program with multiple input files](#run-your-program-with-multiple-input-files)
  1. [Target Program](#target-program)
    1. [Predefined Target Challenges](#predefined-target-challenges)
      1. [Run with multiple input files](#run-with-multiple-input-files)
      1. [How to convert your offline PT solution code to use online?](#how-to-convert-your-offline-pt-solution-code-to-use-online)
      1. [Ragnarok Example](#ragnarok-example)
      1. [List of Predefined Challenges](#list-of-predefined-challenges)
    1. [Template and Example](#template-and-example)
  1. [Challenge map in your terminal](#challenge-map-in-your-terminal)
  1. [Configuration](#configuration)
    1. [Challenge timeout](#challenge-timeout)
    1. [Framerate in Target Challenges](#framerate-in-target-challenges)
    1. [Output Callback](#output-callback)
  1. [Support for other languages](#support-for-other-languages)
1. [Feedback](#feedback)

# Quick Guide
The implementation is quite straightforward and simple. Your offline will be almost identical as your online codingame code, with the difference being that the input/output goes via [channels](http://golang.org/doc/effective_go.html#channels) rather than via the _stdin_ and _stdout_.

You can find the [descriptions](https://github.com/GlenDC/Codingame/tree/master/descriptions), [input](https://github.com/GlenDC/Codingame/tree/master/input) and optionally the [output](https://github.com/GlenDC/Codingame/tree/master/output) text files all [here](https://github.com/glendc/Codingame) or on [the official Codingame website](http://www.codingame.com).

## Types of programs

[Codingame](http://www.codingame.com) has a lot challenges. These challenges can be devided in types of programs based on how they receive input and what the goal of the challenge is.

1. [Manual Program](#manual-program): This is the most simple program and just requires you to write a simple _main_ function that takes a _string channel_ as its input. This channel will give you the input line by line and it's up to you how to interpret the received input. The output of the program has to be returned via the output channel.
1. [Target Program](#target-program): Some challenges are based on win and lose conditions. These are the most complex program and require extra work from the user in order to do these challenges offline, as you'll have to write the logic of the challenge, on top of your usual challenge code. Because of this there are the [predefined challenge programs](#list-of-predefined-challenges), that do all this hard work for you. But anyway... How does a target program work?
  1. You'll write a struct based on the _TargetProgram_ interface
  1. The initial input will be parsed and have to be manually interpred by you via the _InitialInput_ method.
  1. The program runs and calls each frame the _Update_ method, using the input given via the _GetInput_ method. _Update_ will return your output for that frame via the output channel.
    1. This output can also be traced if wanted.
  1. Each frame your output will be used and update the game state via the _SetOutput_ method
  1. The program exits if the _LoseConditionCheck_- or/and _WinConditionCheck_ method returns true

Manual programs can:

1. run the program _or_
1. run and validate the program based on a test text file

With both options you can also echo your output if wanted.

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
          "<INPUT TEXT FILE>",                          // program input source
          func(input <-chan string, output chan string) {            
              // your solution here
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
        func(input <-chan string, output chan string) {
          var width, height int
          var text string

          fmt.Sscanln(<-input, &width)
          fmt.Sscanln(<-input, &height)
          fmt.Sscanln(<-input, &text)

          text = strings.ToUpper(text)

          ascii := make([]string, height)
          for i := 0; i < height; i++ {
            ascii[i] = <-input
          }

          lines := make([]string, height)
          for _, char := range text {
            character := int(char) - 65
            if character < 0 || character > 26 {
              character = 26
            }
            for i := range lines {
              position := character * width
              lines[i] += ascii[i][position : position+width]
            }
          }

          for _, line := range lines {
            output <- line
          }
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
            "<INPUT TEXT FILE>",                          // program input file
            "<OUTPUT TEXT FILE>",                         // expected output file
            true,                                         // show output?
            func(input <-chan string, output chan string) {               // program main
                // your solution here
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
        func(input <-chan string, output chan string) {
          var width, height int
          var text string

          fmt.Sscanln(<-input, &width)
          fmt.Sscanln(<-input, &height)
          fmt.Sscanln(<-input, &text)

          text = strings.ToUpper(text)

          ascii := make([]string, height)
          for i := 0; i < height; i++ {
            ascii[i] = <-input
          }

          lines := make([]string, height)
          for _, char := range text {
            character := int(char) - 65
            if character < 0 || character > 26 {
              character = 26
            }
            for i := range lines {
              position := character * width
              lines[i] += ascii[i][position : position+width]
            }
          }

          for _, line := range lines {
            output <- line;
          }
        })
    }

##### Output:

    ### 
    #   
    ##  
    #   
    ### 

    Program is correct!
    
### Run your program with multiple input files

Running each input file one by one is quite tedious. Therefore it is possible to run multiple input files at once, both for validated and normal manual programs. The definition of these functions is similar, except for the file parameter(s).

##### Examples

###### Multiple manual programs

    // run multiple manual programs at once
    // the output is seperated by a newline character
    func RunManualPrograms(input []string, main ProgramMain)

###### Multiple validated manual programs

    // run and validated multiple manual programs at once
    // the output is seperated by a newline character
    func RunAndValidateManualPrograms(input, test []string, echo bool, main ProgramMain)

## Target Program

### Predefined Target Challenges

The target program challenge type was created to allow you to play more complex challenges, such as Ragnarok, offline. However with target programs you still need to program the Challenge and AI logic yourself, which isn't the goal of the Codingame challenges at all. [The Predefined Target Challenges](#list-of-predefined-challenges) allow you to start on the challenge instantly and it keeps your code base exactly the same as if it were an online submission. 
    
#### Run with multiple input files

Running each input file one by one is quite tedious. Therefore it is possible to run multiple input files at once. The definition of this function is similar to the normal one, except for the input file parameter.

##### Examples

    // run len(in) amount of times the target program
    // with each time a different input file.
    func RunTargetPrograms(input []string, trace bool, program TargetProgram)
    
    // each predefined target challenge allows you to do the same in a similar fashion
    // example for the ragnarok challenge:
    func RunRagnarokPrograms(input []string, trace bool, initialize UserInitializeFunction, update UserUpdateFunction)

#### How to convert your offline PT solution code to use online?

There is no real reason why you would want to convert your offline PT challenge code, but let's say you want to do so. It's possible and easy, as your code base will remain quite similar.

Let's say we have the follow _psuedo_ offline [Ragnarok](https://raw.githubusercontent.com/GlenDC/Codingame/master/descriptions/ragnarok.md) PT solution code:

    package man
    
    import (
      "github.com/glendc/cgreader"
      // packages...
    )
    
    // definition of functions, types and variables...
    
    func Initialize(input <-chan string) {
      // parse the initial input, no output expected...
    }
    
    func Update(input <-chan string, output chan string) {
      // the code of your solution logic will be defined here...
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
    
As you can see it's quite similar. On top of this you'll have to convert code that makes use of the channel input parameter, to use the standard input instead. (e.g. ``fmt.Sscanf(<-input`` to ``fmt.Scanf(`` and ``output <- fmt.Sprintf`` to ``fmt.Printf``)

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

    func (program *Program) GetInput() (input chan string) {
      input = make(chan string)
      go func() {
        // pass the challenge input into the channel
      }()
      return
    }

    func (program *Program) Update(input <-chan string, output chan string) {
      // your solution logic will be defined here
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

Take a look at [the ragnarok predefined target challenge](https://github.com/GlenDC/cgreader/blob/master/ragnarok.go), which is a simple and clear example on how to implement a target program. It also shows that it requires much more coding and knowledge to develop one, than to actually solve it, which is the reason why a raw target program shouldn't be used by _end-users_.
    
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
    
# Configuration

## Challenge timeout

Challenges are considered and reported as invalid, when they take longer than _1 second_, the default timeout value. This value can be set to a custom value with the ``SetTimeout`` function, in case a challenge requires a different value.

An example:

    cgreader.SetTimeout(42.0) // the challenge timeout is now 42 seconds

## Framerate in Target Challenges

In challenges like [ragnarok giants](https://raw.githubusercontent.com/GlenDC/Codingame/master/descriptions/ragnarok_giants.md) you might want to have a slower framerate. By default it is not limited, and goes as fast as your code allows too. Using either the ``SetDelay`` or the ``SetFrameRate`` allows you to configure this and make your game run slower.

Some examples:

    cgreader.SetFrameRate(60) // == SetDelay(1000ms/60)
    
    cgreader.SetDelay(100) // Sleep each frame for 100ms == 0.1s

## Output Callback

All format in cgreader is formatted, and uses fmt.Printf as _stdout_. You can call the ``SetPrintfCallback`` function in case you want to redirect the output to somewhere else:

    // the type for the Printf callback, used in cgreader
    type PrintfCallback func(format string, a ...interface{})
    
    // set cgreader's Printf callback via this function
    func SetPrintfCallback(callback PrintfCallback)

# Support for other languages

Even though cgreader is written in go, it is not the only language supported. In case you want to find out what other languages are supported you can [click here](https://github.com/GlenDC/cgreader/wiki) to go to the wiki. There you'll also find detailed information about the different functions for each language, and how to use them.

# Feedback

Any feedback is welcome, and can be given along the bug reports as an issue report or by [mailing me](mailto:contact@glendc.com). Pull requests are also welcome on the condition that your commits are clean and bring additional values to the cgreader pkg.

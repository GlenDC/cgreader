Codingame Reader
=================

Codingame Reader allows you to code, run and validate your program on your computer, without any need for internet access. Not yet convinced with the sweet perks of being able to code in your own IDE with your fancy tools. What if I would tell you that you can test your program in one run for all the different scenarios and that you get a smiley in case you did well?! See, you like it, don't you?

Are you a developer that wants to contribute to the library (written in Go)? Please head over to [the manual](https://github.com/GlenDC/cgreader/wiki/manual) so that you can know everything you need to know.

**Disclaimer:** All these challenges are originally created by [Codingame](http://www.codingame.com), I merely offer the service to run them offline in your terminal or IDE. 

The implementation is quite straightforward and simple. Your offline will be almost identical as your online codingame code, with the difference being that the input/output goes via [channels](http://golang.org/doc/effective_go.html#channels) rather than via the _stdin_ and _stdout_. And that your logic will be encapsulated within functions, instead of the main function and an endless loop.

You can find the [descriptions](https://github.com/GlenDC/Codingame/tree/master/descriptions), [input](https://github.com/GlenDC/Codingame/tree/master/input) and optionally the [output](https://github.com/GlenDC/Codingame/tree/master/output) text files all [here](https://github.com/glendc/Codingame) or on [the official Codingame website](http://www.codingame.com).

# Index

1. [Learn how to use it within 5 minutes](#learn-how-to-use-it-within-5-minutes)
1. [List of interactive challenges](#list-of-interactive-challenges)
1. [Configuration](#configuration)
1. [Support for other languages](#support-for-other-languages)
1. [Feedback](#feedback)

# Learn how to use it within 5 minutes

First you'll have to get cgreader at your computer. If you're programming in Go it's just a mather of ``import github.com/glendc/cgreader``. In case you use another language, than please go to [the wiki](https://github.com/GlenDC/cgreader/wiki), to see how it has to be done for yours. Is your language not supported? File an issue, so that we can start talking about it.

The public interface is small, and pretty easy to use. It only relies on a couple of global functions defined by cgreader (which only use primitive types). This makes it not only easy to use, but also makes it consistent between he different supported languages. In theory you only ahve to use one function (that is, to run a challenge). However there are [a couple of functions](#configuration) to your disposal in case you want to tweak some stuff.

## Run a static challenge

Running a static challenge easy and straight forward, you can do it in the following way:

    // run your static program for one scenario
    func RunStaticProgram(input, output string, trace bool, main ProgramMain)

    // run your static program for multiple scenarios
    func RunStaticPrograms(input, output []string, trace bool, main ProgramMain)

If you think away the syntax (which might be different from your language), you just have to call the function with 2 paramters to define your input and output text file(s), a boolean value to define if you want to see your generated output and a void function that encapsulates your program logic.

## Run an interactive challenge

Running an interactive challenge is just as easy and straight forward as a static challenge, with slightly different parameters.

    // run your interactive program for one scenario
    func RunInteractiveProgram(type, input string, trace bool, init initFunc, update updateFunc)

    // run your interactive program for multiple scenarios
    func RunInteractivePrograms(type string, input []string, trace bool, init initFunc, update updateFunc)

If you think away the syntax (which might be different from your language), you just have to call the function with the first parameter that defines your program type (a string), the input file(s) for the program, a boolean value to define if you want to see your generated output and 2 void functions that encapsulates your program logic for both the initial input and the program loop.

## Run a sandboxed program

In case you're programming in a languaged parsed by this repository, such as **Brainfuck** you might want to run the program without using any input and validation. This is possible via the ``RunProgram`` function which still gives you the timout protection and runtime report.

    func RunProgram(main SandboxProgramFunction)

# List of interactive challenges

You can find the program type (id) for all the supported interactive challenges within square brackets just next to their official name.

## ready-to-use challenges:

* [**ragnarok**] Power of Thor: [description](https://raw.githubusercontent.com/GlenDC/Codingame/master/descriptions/ragnarok.md)
  * Solutions: [Go](https://github.com/GlenDC/Codingame/blob/master/solutions/go/ragnarok.go)
* [**ragnarok_giants**] Thor Vs. Giants: [description](https://raw.githubusercontent.com/GlenDC/Codingame/master/descriptions/ragnarok_giants.md)
  * Solutions: [Go](https://github.com/GlenDC/Codingame/blob/master/solutions/go/ragnarok_giants.go)
* [**kirk**] Kirk's Quest - The descent: [description](https://raw.githubusercontent.com/GlenDC/Codingame/master/descriptions/kirk.md)
  * Solutions: [Go](https://github.com/GlenDC/Codingame/blob/master/solutions/go/kirk.go)

## challenges to be developed:

* [**skynet\_final\_1**] Skynet Final - Level 1
* [**skynet**] Skynet - The Chasm
* [**mars\_lander\_1**] Mars Lander - Level 1
* [**shadow\_knight\_1**] Shadow of the Knight - 1
* [**indiana\_1**] Indiana - Level 1
* [**mars\_lander\_2**] Mars Lander - Level 2
* [**skynet\_final\_2**] Skynet Finale - Level 2
* [**shadow\_knight\_2**] Shadow of the Knight - 2
* [**skynet\_bridge**] Skynet - The Bridge
* [**kirk\_labyrinth**] Kirk's Quest - The labyrinth
* [**indiana\_2**] Indiana - Level 2
* [**indiana\_3**] Indiana - Level 3
* [**mars\_lander\_3**] Mars Lander - Level 3

_Contributions on the "reverse engineering" of these challenges are more than welcome!_

# Configuration

## Challenge timeout

Challenges are considered and reported as invalid, when they take longer than _1 second_, the default timeout value. This value can be set to a custom value with the ``SetTimeout`` function, in case a challenge requires a different value.

An example:

    SetTimeout(42.0) // the challenge timeout is now 42 seconds

## Framerate in Target Challenges

In challenges like [ragnarok giants](https://raw.githubusercontent.com/GlenDC/Codingame/master/descriptions/ragnarok_giants.md) you might want to have a slower framerate. By default it is not limited, and goes as fast as your code allows too. Using either the ``SetDelay`` or the ``SetFrameRate`` allows you to configure this and make your game run slower.

Some examples:

    SetFrameRate(60) // == SetDelay(1000ms/60)
    
    SetDelay(100) // Sleep each frame for 100ms == 0.1s

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

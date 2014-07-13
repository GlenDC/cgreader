package main

type CommandFunction func([]*Command)

type Command struct {
	run      CommandFunction
	children []*Command
}

func (command *Command) excecute() {
	command.run(command.children)
}

func (command *Command) add(cmd *Command) {
	command.children = append(command.children, cmd)
}

func CreateLinearGroup() *Command {
	return &Command{func(commands []*Command) {
		for _, child := range commands {
			child.excecute()
		}
	}, nil}
}

func CreateLoopGroup() *Command {
	return &Command{func(commands []*Command) {
		for programBuffer[programIndex] != 0 {
			for _, child := range commands {
				child.excecute()
			}
		}
	}, nil}
}

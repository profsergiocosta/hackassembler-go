package command

// Haskell: data Command = Arithmetic String| Pop String Int|  Push String Int

type Command interface {
	isCommand()
}

type ACommand struct {
	At string
}

func (_ ACommand) isCommand() {}

type CCommand struct {
	Dest string
	Comp string
	Jump string
}

func (_ CCommand) isCommand() {}

type LCommand struct {
	Label string
}

func (_ LCommand) isCommand() {}

type UndefinedCommand struct {
	Label string
}

func (_ UndefinedCommand) isCommand() {}

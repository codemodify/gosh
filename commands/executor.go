package commands

// Executor -
type Executor struct {
	commands []Command
}

// NewExecutor -
func NewExecutor(commands []Command) *Executor {
	return &Executor{
		commands: commands,
	}
}

// Execute -
func (thisRef *Executor) Execute(command string) error {
	for _, v := range thisRef.commands {
		if v.CanHandle(command) {
			return v.Execute(command)
		}
	}

	return nil
}

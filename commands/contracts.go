package commands

// Command -
type Command interface {
	CanHandle(command string) bool
	Execute(command string) error
}

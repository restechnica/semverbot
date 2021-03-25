package commands

// Commander interface to run commands.
type Commander interface {
	Output(name string, arg ...string) (output string, err error)
	Run(name string, arg ...string) (err error)
}

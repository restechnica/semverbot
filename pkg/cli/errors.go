package cli

func NewCommandError(err error) CommandError {
	return CommandError{Err: err}
}

type CommandError struct {
	Err error
}

func (e CommandError) Error() string {
	return e.Err.Error()
}

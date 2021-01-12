package args

type ArgLineError struct {
	msg string
}

func (e *ArgLineError) Error() string {
	return e.msg
}

func NewArgLineError(msg string) *ArgLineError {
	return &ArgLineError{msg: msg}
}

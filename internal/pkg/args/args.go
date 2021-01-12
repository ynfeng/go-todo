package args

import "os"

var usage = `
	Usage: todo <command> [args] [--options]
	command:
		list [-a] : list unfinished todo
              -a  : list all todo 
`

type Args struct {
}

func (args *Args) Cmd() (*string, error) {
	if len(os.Args) == 0 {
		return nil, NewArgLineError(usage)
	}
	return &os.Args[0], nil
}

func NewArgs(argsLine []string) *Args {
	os.Args = argsLine
	return &Args{}
}

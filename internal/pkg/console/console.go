package console

import (
	"fmt"
	"github.com/ynfeng/todo/internal/pkg/todo"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func PrintItems(items []todo.Item) {
	for i, item := range items {
		fmt.Fprintf(out, "%d. %s\n", i+1, item.Name())
	}
	fmt.Fprintf(out, "\n")
}

func Printf(format string, a ...interface{}) {
	fmt.Fprintf(out, format, a...)
}

func SetOut(writer io.Writer) {
	out = writer
}

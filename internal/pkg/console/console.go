package console

import (
	"fmt"
	"github.com/ynfeng/todo/internal/pkg/todo"
)

func PrintItems(items []todo.Item) {
	for i, item := range items {
		fmt.Printf("%d. %s\n", i+1, item.Name())
	}
	fmt.Println()
}

func Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

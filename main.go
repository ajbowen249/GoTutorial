package main

import (
	"fmt"
	"github.com/ajbowen249/GoTutorial/table"
)

func main() {
	tb := table.New()
	tb.CellAlign = table.CENTERLEFTBIAS
	tb.AddColumn("XValues")
	tb.AddRow("1111")
	tb.Output(func(s string) {
		fmt.Println(s)
	})
}

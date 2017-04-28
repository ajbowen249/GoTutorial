package main

import (
	"fmt"
	al "github.com/ajbowen249/GoTutorial/algorithms"
	ds "github.com/ajbowen249/GoTutorial/datastructures"
	"github.com/ajbowen249/GoTutorial/table"
	"math/rand"
	"time"
)

func main() {
	tb := table.New()
	tb.AddColumn("NumItems")
	tb.AddColumn("Selection Sort Time (ms)")
	tb.AddColumn("Binary Sort Time (ms)")

	for i := 0; i <= 18000; i += 1000 {
		selection, binary := sortTest(i)
		tb.AddRow(fmt.Sprintf("%v", i), fmt.Sprintf("%.3f", selection.Seconds()*1000), fmt.Sprintf("%.3f", binary.Seconds()*1000))
	}

	tb.Output(func(row string) { fmt.Println(row) })
}

func sortTest(length int) (time.Duration, time.Duration) {
	input := make([]int, length)

	rand.Seed(time.Now().UnixNano())
	for i := range input {
		input[i] = rand.Int()
	}

	selectSortStart := time.Now()

	al.SortInt(input)
	selectSortDuration := time.Since(selectSortStart)

	treeSortStart := time.Now()
	resultBuffer := make([]int, length)
	index := 0

	tree := new(ds.IntBinaryTree)

	for i := range input {
		tree.Add(input[i])
	}

	tree.VisitAscending(func(node *ds.IntBinaryNode) {
		resultBuffer[index] = node.Value
		index++
	})

	treeSortDuration := time.Since(treeSortStart)

	return selectSortDuration, treeSortDuration
}

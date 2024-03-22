package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/brianshannan/rosalind/utils"
	"github.com/brianshannan/rosalind/utils/datastructures/graph"
)

func main() {
	flag.Parse()
	data, err := utils.ReadFileFromArgStripTrailingNewline()
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	numVertices, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(err)
	}

	g := graph.NewGraph(numVertices)
	for _, line := range lines[1:] {
		strVals := strings.Split(line, " ")
		v1, v2 := utils.ParseInt(strVals[0])-1, utils.ParseInt(strVals[1])-1
		g.AddEdge(v1, v2)
		g.AddEdge(v2, v1)
	}
	fmt.Println(g.FindNumIslands() - 1)
}

package graph

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

var (
	tinyUG = `13
13
0 5
4 3
0 1
9 12
6 4
5 4
0 2
11 12
9 10
0 6
7 8
9 11
5 3
`

	tinyDG = `13
22
4 2
2 3
3 2
6 0
0 1
2 0
11 12
12 9
9 10
9 11
7 9
10 12
11 4
4 3
3 5
6 8
8 6
5 4
0 5
6 4
6 9
7 6`

	tinyDAG = `13
15
2 3
0 6
0 1
2 0
11 12
9 12
9 10
9 11
3 5
8 7
5 4
0 5
6 4
6 9
7 6`
)

func readGraph(in []byte) (int, int, [][]int) {
	var (
		vertices int
		edges    int
		pairs    = make([][]int, 0)
	)
	scanner := bufio.NewScanner(bytes.NewReader(in))
	line := 0
	for scanner.Scan() {
		if line == 0 {
			vertices, _ = strconv.Atoi(scanner.Text())
			line++
			continue
		} else if line == 1 {
			edges, _ = strconv.Atoi(scanner.Text())
			line++
			continue
		}

		parts := strings.Split(scanner.Text(), " ")
		v, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])
		pairs = append(pairs, []int{v, w})

		line++
	}

	return vertices, edges, pairs
}

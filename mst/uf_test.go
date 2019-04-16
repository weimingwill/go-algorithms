package mst

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

var tinyUF = `10
4 3
3 8
6 5
9 4
2 1
8 9
5 0
7 2
6 1
1 0
6 7`

func ExampleUnionFind() {
	scanner := bufio.NewScanner(bytes.NewReader([]byte(tinyUF)))

	var uf *UnionFind
	line := 0
	for scanner.Scan() {
		if line == 0 {
			n, _ := strconv.Atoi(scanner.Text())
			uf = NewUnionFind(n)
			line++
			continue
		}
		parts := strings.Split(scanner.Text(), " ")
		p, _ := strconv.Atoi(parts[0])
		q, _ := strconv.Atoi(parts[1])
		if uf.Connected(p, q) {
			continue
		}
		uf.Union(p, q)
		line++
	}

	fmt.Println(uf.Count())

	// Output:
	// 2
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, start int
	// fmt.Fscan을 사용하면 공백/개행 상관없이 입력을 매우 간단하게 받을 수 있습니다.
	fmt.Fscan(reader, &n, &m, &start)

	// 인접 행렬 대신 인접 리스트(slice의 slice)를 사용하여 메모리와 순회 시간을 절약합니다.
	graph := make([][]int, n+1)

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// 방문할 수 있는 정점이 여러 개인 경우에는 정점 번호가 작은 것을 먼저 방문하도록 정렬
	for i := 1; i <= n; i++ {
		sort.Ints(graph[i])
	}

	// DFS
	visited := make([]bool, n+1)
	dfs(graph, start, visited, writer)
	fmt.Fprintln(writer)

	// BFS
	visited = make([]bool, n+1)
	bfs(graph, start, visited, writer)
	fmt.Fprintln(writer)
}

func dfs(graph [][]int, curr int, visited []bool, writer *bufio.Writer) {
	visited[curr] = true
	// 매번 문자열을 연결(+)하지 않고 곧바로 출력 스트림에 씁니다.
	fmt.Fprintf(writer, "%d ", curr)

	for _, next := range graph[curr] {
		if !visited[next] {
			dfs(graph, next, visited, writer)
		}
	}
}

func bfs(graph [][]int, start int, visited []bool, writer *bufio.Writer) {
	// container/list 대신 일반 슬라이스를 큐로 사용하는 것이 더 빠르고 직관적입니다.
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:] // pop
		fmt.Fprintf(writer, "%d ", curr)

		for _, next := range graph[curr] {
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next) // push
			}
		}
	}
}

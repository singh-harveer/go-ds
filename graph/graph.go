package graph

/**
Notes-->
	1- All check for visited graph in undirected graph to avoid infinite cycle.
**/
func DepthFirstTraversal(source int, graph map[int][]int) []int {
	var stack []int
	stack = append(stack, source)

	var result []int
	for len(stack) > 0 {
		var current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current)
		for _, neighbor := range graph[current] {
			stack = append(stack, neighbor)
		}
	}

	return result
}

func BreathFirstTraversal(source int, graph map[int][]int) []int {
	var queue []int
	queue = append(queue, source)

	var result []int
	for len(queue) > 0 {
		var current = queue[0]
		queue = queue[1:]
		result = append(result, current)
		for _, neighbor := range graph[current] {
			queue = append(queue, neighbor)
		}
	}

	return result
}

func HasPath(source int, dest int, graph map[int][]int) bool {
	var queue []int
	queue = append(queue, source)

	for len(queue) > 0 {
		var current = queue[0]
		queue = queue[1:]
		if current == dest {
			return true
		}
		for _, neighbor := range graph[current] {
			queue = append(queue, neighbor)
		}
	}

	return false
}

type obj struct {
	node int
	dist int
}

func ShortestPath(edges [][]int, src int, dst int) int {
	var graph = buildGraph(edges)
	var visited = make(map[int]struct{})

	var queue []obj
	queue = append(queue, obj{src, 0})

	for len(queue) > 0 {
		var current = queue[0]
		queue = queue[1:]

		if current.node == dst {
			return current.dist
		}

		for _, heighbor := range graph[current.node] {
			if _, ok := visited[current.node]; !ok {
				visited[current.node] = struct{}{}
				queue = append(queue, obj{heighbor, current.dist + 1})
			}
		}
	}

	// if came here which means there is no path
	return -1
}

func buildGraph(edges [][]int) map[int][]int {
	var m = make(map[int][]int)

	for _, pair := range edges {
		var first, second = pair[0], pair[1]

		if _, ok := m[first]; !ok {
			m[first] = []int{}
		}

		if _, ok := m[second]; !ok {
			m[second] = []int{}
		}

		m[first] = append(m[first], second)
		m[second] = append(m[second], first)
	}

	return m
}

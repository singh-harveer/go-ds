package graph

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

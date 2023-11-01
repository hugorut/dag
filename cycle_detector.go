package dag

type CycleDetector struct {
	visited  map[interface{}]bool
	recStack map[interface{}]bool
	hasCycle bool
	dag      *DAG
}

func (cd *CycleDetector) Visit(v interface{}) {
	if cd.hasCycle {
		return
	}

	stack := []interface{}{v}

	for len(stack) > 0 {
		vertex := stack[len(stack)-1]

		if !cd.visited[vertex] {
			cd.visited[vertex] = true
			cd.recStack[vertex] = true
		}

		anyUnvisitedNeighbor := false
		for neighbor := range cd.dag.outboundEdge[vertex] {
			if !cd.visited[neighbor] {
				stack = append(stack, neighbor)
				anyUnvisitedNeighbor = true
				break
			} else if cd.recStack[neighbor] {
				cd.hasCycle = true
				return
			}
		}

		if !anyUnvisitedNeighbor {
			stack = stack[:len(stack)-1] // pop from stack
			cd.recStack[vertex] = false  // remove from recursion stack
		}
	}
}

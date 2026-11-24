package main

import "fmt"

// INF marks "no path" in the distance matrix.
const INF = int(^uint(0) >> 1 / 2) // half of MaxInt: addition can't overflow

// FloydWarshall returns all-pairs shortest distances. dist[i][j] on input is
// the direct edge weight (INF if absent, 0 on the diagonal); the returned
// matrix holds shortest path costs. Negative edges are fine; a negative
// value on the diagonal afterwards reveals a negative cycle.
func FloydWarshall(dist [][]int) [][]int {
	n := len(dist)
	d := make([][]int, n)
	for i := range d {
		d[i] = append([]int{}, dist[i]...)
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if d[i][k] == INF {
				continue
			}
			for j := 0; j < n; j++ {
				if d[k][j] != INF && d[i][k]+d[k][j] < d[i][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}
	return d
}

func main() {
	m := [][]int{
		{0, 3, INF, 7},
		{8, 0, 2, INF},
		{5, INF, 0, 1},
		{2, INF, INF, 0},
	}
	for _, row := range FloydWarshall(m) {
		fmt.Println(row)
	}
}

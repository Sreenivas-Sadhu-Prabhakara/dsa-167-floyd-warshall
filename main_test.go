package main

import (
	"reflect"
	"testing"
)

func TestFloydWarshall(t *testing.T) {
	m := [][]int{
		{0, 3, INF, 7},
		{8, 0, 2, INF},
		{5, INF, 0, 1},
		{2, INF, INF, 0},
	}
	want := [][]int{
		{0, 3, 5, 6},
		{5, 0, 2, 3},
		{3, 6, 0, 1},
		{2, 5, 7, 0},
	}
	if got := FloydWarshall(m); !reflect.DeepEqual(got, want) {
		t.Errorf("FloydWarshall = %v, want %v", got, want)
	}

	// Disconnected pair stays INF.
	m2 := [][]int{{0, INF}, {INF, 0}}
	got := FloydWarshall(m2)
	if got[0][1] != INF || got[1][0] != INF {
		t.Errorf("disconnected = %v", got)
	}

	// Negative edge without a cycle.
	m3 := [][]int{{0, -2, INF}, {INF, 0, 3}, {INF, INF, 0}}
	got = FloydWarshall(m3)
	if got[0][2] != 1 {
		t.Errorf("negative edge path = %v", got)
	}

	// Negative cycle shows as a negative diagonal.
	m4 := [][]int{{0, 1, INF}, {INF, 0, -3}, {1, INF, 0}}
	got = FloydWarshall(m4)
	neg := false
	for i := range got {
		if got[i][i] < 0 {
			neg = true
		}
	}
	if !neg {
		t.Error("negative cycle must surface on the diagonal")
	}
}

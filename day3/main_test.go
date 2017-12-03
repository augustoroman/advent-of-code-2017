package main

import "testing"

func TestRing(t *testing.T) {
	testcases := []struct{ n, ring int }{
		{1, 1},
		{2, 3},
		{5, 3},
		{7, 3},
		{9, 3},
		{10, 5},
		{24, 5},
		{25, 5},
		{26, 7},
		{35, 7},
		{49, 7},
		{50, 9},
	}
	for _, test := range testcases {
		r := ring(test.n)
		if r != test.ring {
			t.Errorf("Expected %d to be in ring %d, but got ring %d", test.n, test.ring, r)
		}
	}
}

func TestPos(t *testing.T) {
	testcases := []struct{ n, x, y int }{
		{1, 0, 0},
		{2, 1, 0},
		{5, -1, -1},
		{7, -1, 1},
		{9, 1, 1},
		{10, 2, 1},
		{24, 1, 2},
		{25, 2, 2},
		{26, 3, 2},
		{35, -1, -3},
		{49, 3, 3},
		{50, 4, 3},
	}
	for _, test := range testcases {
		x, y := pos(test.n)
		if x != test.x || y != test.y {
			t.Errorf("%3d: exp:(%2d,%2d)  got:(%2d,%2d)", test.n, test.x, test.y, x, y)
		}
	}
}

func TestDist(t *testing.T) {
	testcases := []struct{ n, dist int }{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}
	for _, test := range testcases {
		dist := dist(test.n)
		if dist != test.dist {
			t.Errorf("%d: exp %d, got %d", test.n, test.dist, dist)
		}
	}
}

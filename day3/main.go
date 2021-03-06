package main

import (
	"flag"
	"fmt"
	"math"
)

// 37  36  35  34  33  32  31
// 38  17  16  15  14  13  30
// 39  18   5   4   3  12  29
// 40  19   6   1   2  11  28
// 41  20   7   8   9  10  27
// 42  21  22  23  24  25  26
// 43  44  45  46  47  48  49

// Ring | Start # | End #
//  1   |   1     |   1  = 1^2
//  3   |   2     |   9  = 3^2
//  5   |  10     |   25 = 5^2
//  7   |  26     |   49 = 7^2

// Data from square 1 is carried 0 steps, since it's at the access port.
// Data from square 12 is carried 3 steps, such as: down, left, left.
// Data from square 23 is carried only 2 steps: up twice.
// Data from square 1024 must be carried 31 steps.

func ring(n int) int {
	r0 := int(math.Ceil(math.Sqrt(float64(n))))
	if r0%2 == 0 {
		return r0 + 1
	}
	return r0
}

type Ring int

func (r Ring) diameter() int { return int(r) }
func (r Ring) radius() int   { return int(r-1) / 2 }
func (r Ring) start() int    { return r.prev().end() + 1 }
func (r Ring) end() int      { return int(r * r) }
func (r Ring) prev() Ring {
	if r == 1 {
		return 0
	}
	return r - 2
}

func (r Ring) startPos() (x, y int) {
	return r.radius(), r.prev().radius()
}
func (r Ring) endPos() (x, y int) {
	return r.radius(), r.radius()
}
func (r Ring) walkDir(x, y int) (dx, dy, maxdist int) {
	rad := r.radius()
	if x == rad && y != -rad {
		return 0, -1, -(-rad - y)
	} else if x == -rad && y != rad {
		return 0, 1, rad - y
	} else if y == rad && x != rad {
		return 1, 0, rad - x
	} else if y == -rad && x != -rad {
		return -1, 0, -(-rad - x)
	}
	panic(fmt.Errorf("(%d,%d) not on ring %d [rad=%d]", x, y, r, rad))
}
func (r Ring) walk(sx, sy, n int) (x, y int) {
	// log.Println("walking ", n, " from ", sx, sy)
	for n > 0 {
		dx, dy, maxd := r.walkDir(sx, sy)
		// log.Println(n, "steps left: dx/dy: ", dx, dy, " up to ", maxd, " from: ", sx, sy)
		if maxd > n {
			maxd = n
		}
		sx += dx * maxd
		sy += dy * maxd
		n -= maxd
	}
	// log.Println(sx, sy, n)
	return sx, sy
}

func pos(n int) (x, y int) {
	// log.Println("------------------")
	r := Ring(ring(n))
	// log.Println("n:", n, " --> ring ", r, " starts at ", r.start())
	sx, sy := r.startPos()
	// log.Println(n, sx, sy, r)
	x, y = r.walk(sx, sy, n-r.start())
	return x, y
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func dist(n int) int {
	x, y := pos(n)
	return abs(x) + abs(y)
}

type walker struct {
	x, y int
	n    int
	r    Ring
}

func NewWalker() *walker { return &walker{0, 0, 1, 1} }

func (w *walker) Step() {
	w.n++
	r := Ring(ring(w.n))
	if r != w.r {
		w.x++
		w.r = r
		return
	} else {
		w.x, w.y = r.walk(w.x, w.y, 1)
	}
}

type xy struct{ x, y int }
type Store map[xy]int

func (s Store) update(x, y int) int {
	sum := 0
	for xx := x - 1; xx <= x+1; xx++ {
		for yy := y - 1; yy <= y+1; yy++ {
			sum += s[xy{xx, yy}]
		}
	}
	s[xy{x, y}] = sum
	return sum
}

func main() {
	n := flag.Int("n", 0, "val to check")
	flag.Parse()
	x, y := pos(*n)
	d := dist(*n)
	fmt.Printf("%d -> (%d,%d) -> %d\n", *n, x, y, d)

	w := NewWalker()
	s := Store{}
	s[xy{0, 0}] = 1
	for {
		w.Step()
		val := s.update(w.x, w.y)
		fmt.Printf("Wrote %d to (%d,%d)\n", val, w.x, w.y)
		if val > *n {
			break
		}
	}
}

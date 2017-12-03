package main

import (
	"flag"
	"fmt"
	"math"
)

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

func main() {
	n := flag.Int("n", 0, "val to check")
	flag.Parse()
	x, y := pos(*n)
	d := dist(*n)
	fmt.Printf("%d -> (%d,%d) -> %d\n", *n, x, y, d)
}

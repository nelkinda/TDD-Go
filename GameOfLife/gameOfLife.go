package GameOfLife

type Point struct {
	x, y int
}

func (p Point) getNeighbors() []Point {
	return []Point{
		{p.x - 1, p.y - 1},
		{p.x - 1, p.y},
		{p.x - 1, p.y + 1},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
		{p.x + 1, p.y - 1},
		{p.x + 1, p.y},
		{p.x + 1, p.y + 1},
	}
}

type Universe struct {
	points map[Point]bool
}

func EmptyUniverse() Universe {
	return Universe{points: make(map[Point]bool)}
}

func NewUniverse(points []Point) Universe {
	universe := Universe{}
	universe.points = make(map[Point]bool)
	for _, p := range points {
		universe.points[p] = true
	}
	return universe
}

func (u Universe) next() Universe {
	var points []Point
	points = u.cellSurvival(points)
	points = u.cellBirth(points)
	return NewUniverse(points)
}

func (u Universe) cellSurvival(points []Point) []Point {
	for p := range u.points {
		if u.survives(p) {
			points = append(points, p)
		}
	}
	return points
}

func (u Universe) survives(p Point) bool {
	count := u.countLiveNeighbors(p)
	return count == 2 || count == 3
}

func (u Universe) cellBirth(points []Point) []Point {
	for p := range u.points {
		for _, n := range p.getNeighbors() {
			if !(u.points[n]) {
				liveNeighbors := u.countLiveNeighbors(n)
				if liveNeighbors == 3 {
					points = append(points, n)
				}
			}
		}
	}
	return points
}

func (u Universe) countLiveNeighbors(p Point) int {
	liveNeighbors := 0
	for _, n := range p.getNeighbors() {
		if u.points[n] {
			liveNeighbors++
		}
	}
	return liveNeighbors
}

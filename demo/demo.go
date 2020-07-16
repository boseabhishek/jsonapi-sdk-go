package demo

type Point struct {
	X, Y, Z int
}

func NewPoint(x, y, z int) *Point {
	return &Point{
		X: x,
		Y: y,
		Z: z,
	}
}

func (p *Point) Equal() bool {
	return p.X == p.Y && p.X == p.Z && p.Y == p.Z
}

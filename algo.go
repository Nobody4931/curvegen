package main

type Point struct {
	X, Y float64
}

func NewPolynomialThroughPoints(points []Point) *Polynomial {
	if len(points) < 1 {
		panic("points must contain at least one value")
	}

	f := NewConstantPolynomial(points[0].Y)
	g := NewZeroPolynomial(points[0].X)

	for i := 1; i < len(points); i++ {
		x, y := points[i].X, points[i].Y
		f = f.Add(g.MulC((y - f.Eval(x)) / g.Eval(x)))
		g = g.Mul(NewZeroPolynomial(x))
	}

	return f.Trim()
}

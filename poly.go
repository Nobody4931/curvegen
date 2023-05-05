package main

type Polynomial struct {
	Degree int
	Coefficients []float64
}

func NewPolynomial(degree int) *Polynomial {
	return &Polynomial{
		Degree: degree,
		Coefficients: make([]float64, degree + 1),
	}
}


func (poly *Polynomial) Coefficient(coIdx int) float64 {
	if 0 > coIdx || coIdx > poly.Degree {
		return 0
	}
	return poly.Coefficients[coIdx]
}

func (poly *Polynomial) Eval(x float64) float64 {
	var result float64 = 0
	var exp float64 = 1
	for _, coeff := range poly.Coefficients {
		result += coeff * exp
		exp *= x
	}
	return result
}


func (poly *Polynomial) Trim() *Polynomial {
	degree := poly.Degree
	for degree >= 0 {
		if !fequal(poly.Coefficients[degree], 0) {
			break
		}
		degree--
	}
	result := NewPolynomial(degree)
	copy(result.Coefficients, poly.Coefficients)
	return result
}


func (poly *Polynomial) Neg() *Polynomial {
	result := NewPolynomial(poly.Degree)
	for coIdx := range result.Coefficients {
		result.Coefficients[coIdx] = -poly.Coefficient(coIdx)
	}
	return result
}

func (poly *Polynomial) Add(other *Polynomial) *Polynomial {
	result := NewPolynomial(max(poly.Degree, other.Degree))
	for coIdx := range result.Coefficients {
		result.Coefficients[coIdx] = poly.Coefficient(coIdx) + other.Coefficient(coIdx)
	}
	return result
}

func (poly *Polynomial) Sub(other *Polynomial) *Polynomial {
	return poly.Add(other.Neg())
}

func (poly *Polynomial) Mul(other *Polynomial) *Polynomial {
	result := NewPolynomial(poly.Degree + other.Degree)
	for coIdxPoly, coeffPoly := range poly.Coefficients {
		for coIdxOther, coeffOther := range other.Coefficients {
			result.Coefficients[coIdxPoly + coIdxOther] += coeffPoly * coeffOther
		}
	}
	return result
}

package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return 4 / (1 + math.Pow(x, 2))
}

func f2(x float64) float64 {
	return 1 / math.Pow(x*math.Pow(x, 2), 0.5)
}

func f3(x float64) float64 {
	return 1 / (1 - math.Log(x))
}

func rightRectangleMethod(a, b, h float64) float64 {
	sum := 0.0
	for x := a; x < b; x += h {
		sum += f(x)
	}
	return h * sum
}

func midRectangleMethod(a, b, h float64) float64 {
	sum := 0.0
	for x := a + h/2; x < b; x += h {
		sum += f(x)
	}
	return h * sum
}

func simpsonsMethod(a, b, h float64) float64 {
	n := int((b - a) / h)
	if n%2 != 0 {
		n++
	}
	sum := f(a) + f(b)

	for i := 2; i < n; i += 2 {
		x := a + float64(i)*h
		sum += 4 * f(x)
	}

	for i := 1; i < n; i += 2 {
		x := a + float64(i)*h
		sum += 2 * f(x)
	}

	return (h / 3) * sum
}

func rightRectangleMethodWithN(a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.0
	for i := 1; i <= n; i++ {
		x := a + float64(i)*h
		sum += f2(x)
	}
	return h * sum
}

func rombergIntegration(a float64, b, initialN int, tolerance float64) (float64, error) {
	var result, prevResult float64
	n := initialN
	delta := math.MaxFloat64

	for delta >= tolerance {
		prevResult = result
		result = rightRectangleMethodWithN(a, float64(b), n)
		if prevResult != 0 {
			delta = math.Abs(result - prevResult)
		}
		n *= 2
	}

	return result, nil
}

func simpsonsMethod3(a, b, h float64) float64 {
	n := int((b - a) / h)
	if n%2 != 0 {
		n++
	}
	sum := f3(a) + f3(b)

	for i := 2; i < n; i += 2 {
		x := a + float64(i)*h
		sum += 4 * f3(x)
	}

	for i := 1; i < n; i += 2 {
		x := a + float64(i)*h
		sum += 2 * f3(x)
	}

	return (h / 3) * sum
}

func main() {
	a := 0.0
	b := 1.0
	ширина_правих_прямокутників := 0.00017320508
	ширина_середніх_прямокутниіків := 0.01
	ширина_сімсона := 0.003248572349

	rightResult := rightRectangleMethod(a, b, ширина_правих_прямокутників)

	midResult := midRectangleMethod(a, b, ширина_середніх_прямокутниіків)

	simpson := simpsonsMethod(a, b, ширина_сімсона)

	fmt.Printf("Метод правих прямокутників: %.4f\n", rightResult)
	fmt.Printf("Метод середніх прямокутників: %.4f\n", midResult)
	fmt.Printf("Метод Сімпсона: %.4f\n", simpson)

	r, _ := rombergIntegration(0.2, 1, 1, 0.28)
	fmt.Printf("I2: %.4f\n", r)

	r = simpsonsMethod3(0, 1, 0.010143)
	fmt.Printf("Integral is: %.2f\n", r)
}

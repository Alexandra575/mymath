package mymath

import "math"

func Yn(n int, x float64) float64 {
   
    if n < 0 {
        return 0
    }
    
    result := 0.0
    for k := 0; k <= n; k++ {
        result += math.Pow(-1, float64(k)) * (math.Pow(x, float64(n-k)) / float64(factorial(k)))
    }
    return result
}

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}
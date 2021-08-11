package main

import (
    "math"
)

// 値を2乗する
func Square(x int) int {
    result := math.Pow(float64(x), 2.0)
    return int(result)
}

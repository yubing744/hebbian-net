package pkg

import "math"

// 欧几里得距离计算
func euclideanDistance(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		diff := a[i] - b[i]
		sum += diff * diff
	}
	return sum // 不需要开平方，比较距离大小即可
}

// Sigmoid激活函数
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

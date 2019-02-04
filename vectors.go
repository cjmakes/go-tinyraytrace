package main

import "math"

// Magnitude return the length of a vector
func (v *vec3) Magnitude() float64 {
	var len float64
	for _, i := range v {
		len += math.Pow(i, 2)
	}
	return math.Sqrt(len)
}

// Adds v2 to v
func (v *vec3) Add(v2 vec3) (ret vec3) {
	for i := range ret {
		ret[i] = v[i] + v2[i]
	}
	return
}

// Subtractesv2 from v
func (v *vec3) Sub(v2 vec3) (ret vec3) {
	for i := range ret {
		ret[i] = v[i] - v2[i]
	}
	return
}

// ScalarProduct computes v multiplied by scalar k
func (v *vec3) ScalarProduct(k float64) (ret vec3) {
	for i, v := range v {
		ret[i] = v * k
	}
	return
}

// DotProduct computes the dot product of v * v2
func (v *vec3) DotProduct(v2 vec3) float64 {
	dp := 0.0
	for i := range v {
		dp += v[i] * v2[i]
	}
	return dp
}

// Abs returns the absolute value
func (v *vec3) Abs() (ret vec3) {
	for i := range v {
		ret[i] = math.Abs(v[i])
	}
	return
}

// Norm returns a unit vector in the direction of v
func (v *vec3) Norm() (ret vec3) {
	return v.ScalarProduct(1 / v.Magnitude())
}

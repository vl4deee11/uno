package engine

import "errors"

func MinE(values []float32) (float32, error) {
	if len(values) == 0{
		return -1, errors.New("empty array")
	}

	min := values[0]
	for i := range values {
		if values[i] < min {
			min = values[i]
		}
	}
	return min, nil
}

func MaxE(values []float32) (float32, error) {
	if len(values) == 0{
		return -1, errors.New("empty array")
	}

	max := values[0]
	for i := range values {
		if values[i] > max {
			max = values[i]
		}
	}
	return max, nil
}
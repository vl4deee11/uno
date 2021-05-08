package engine

func maxFloat32(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}

func minFloat32(x, y float32) float32 {
	if x < y {
		return x
	}
	return y
}

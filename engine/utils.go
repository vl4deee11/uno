package engine

func Min(values []*Pair) *Pair {
	if len(values) == 0{
		return nil
	}

	min := values[0].E
	mP := values[0]
	for _, p := range values {
		if p.E < min {
			min = p.E
			mP = p
		}
	}
	return mP
}

func Max(values []*Pair) *Pair {
	if len(values) == 0{
		return nil
	}

	max := values[0].E
	mP := values[0]
	for _, p := range values {
		if p.E > max {
			max = p.E
			mP = p
		}
	}
	return mP
}

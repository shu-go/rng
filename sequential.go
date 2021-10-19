package rng

type Sequential interface {
	Next() Sequential
	Prev() Sequential

	Less(Sequential) bool
	Equal(Sequential) bool
}

func Max(a, b Sequential) Sequential {
	if a.Less(b) {
		return b
	}
	return a
}

func Min(a, b Sequential) Sequential {
	if a.Less(b) {
		return a
	}
	return b
}

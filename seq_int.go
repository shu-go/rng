package rng

type Int int

func (s Int) Next() Sequential {
	i := int(s)
	return Int(i + 1)
}

func (s Int) Prev() Sequential {
	i := int(s)
	return Int(i - 1)
}

func (s Int) Less(b Sequential) bool {
	if bb, ok := b.(Int); ok {
		return s < bb
	}
	return false
}

func (s Int) Equal(b Sequential) bool {
	if bb, ok := b.(Int); ok {
		return s == bb
	}
	return false
}

package rng

import (
	"fmt"
)

type Range struct {
	Start, End Sequential
}

var Invalid Range = Range{Int(1), Int(0)}

func NewRange(s, e Sequential) Range {
	return Range{
		Start: s,
		End:   e,
	}
}

func (r Range) IsValid() bool {
	return r.Start.Equal(r.End) || r.Start.Less(r.End)
}

func (r Range) ContainsSeq(s Sequential) bool {
	if !r.IsValid() {
		return false
	}

	return s.Less(r.Start) || r.End.Less(s)
}

func (r Range) ContainsRange(a Range) bool {
	if !r.IsValid() || !a.IsValid() {
		return false
	}

	if a.Start.Less(r.Start) {
		return false
	}

	if r.End.Less(a.End) {
		return false
	}

	return true
}

func (r Range) IsIntersecting(a Range) bool {
	if !r.IsValid() {
		return false
	}

	if !a.IsValid() {
		return false
	}

	// ハズレ
	if a.End.Less(r.Start) || r.End.Less(a.Start) {
		return false
	}

	return true
}

func (r Range) Equal(a Range) bool {
	return r.IsValid() == a.IsValid() && r.Start.Equal(a.Start) && r.End.Equal(a.End)
}

func (r Range) Add(a Range) (r1, r2 Range) {
	if !r.IsIntersecting(a) {
		if r.End.Next().Equal(a.Start) {
			return NewRange(r.Start, a.End), Invalid
		} else if a.End.Next().Equal(r.Start) {
			return NewRange(a.Start, r.End), Invalid
		}
		return r, a
	}

	var s, e Sequential
	if r.Start.Less(a.Start) {
		s = r.Start
	} else {
		s = a.Start
	}
	if r.End.Less(a.End) {
		e = a.End
	} else {
		e = r.End
	}

	return NewRange(s, e), Invalid
}

func (r Range) Minus(a Range) (r1, r2 Range, intersect bool) {
	if !r.IsValid() {
		return r, Invalid, false
	}

	if !a.IsValid() {
		return r, Invalid, false
	}

	// ハズレ
	if a.End.Less(r.Start) || r.End.Less(a.Start) {
		return r, Invalid, false
	}

	// a が r を包含
	if (a.Start.Equal(r.Start) || a.Start.Less(r.Start)) && (r.End.Equal(a.End) || r.End.Less(a.End)) {
		return Invalid, Invalid, true
	}

	// r が a を包含
	if r.Start.Less(a.Start) && a.End.Less(r.End) {
		return NewRange(r.Start, a.Start.Prev()), NewRange(a.End.Next(), r.End), true
	}

	// r;    s    e
	// a:  s     e
	if a.Start.Less(r.Start) {
		return NewRange(a.End.Next(), r.End), Invalid, true
	}

	if r.Start.Less(a.Start) {
		return NewRange(r.Start, a.Start.Prev()), Invalid, true
	}

	panic(fmt.Sprintf("%v minus %v", r, a))
}

func (r Range) String() string {
	if !r.IsValid() {
		return "[INVALID]"
	}
	return fmt.Sprintf("[%v, %v]", r.Start, r.End)
}

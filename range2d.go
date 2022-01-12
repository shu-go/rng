package rng

import (
	"sort"
)

type Range2D struct {
	R1, R2 Range
}

func NewRange2D(s1, e1, s2, e2 Sequential) Range2D {
	return Range2D{
		R1: NewRange(s1, e1),
		R2: NewRange(s2, e2),
	}
}

var Invalid2D Range2D = Range2D{
	R1: NewRange(Int(1), Int(0)),
	R2: NewRange(Int(1), Int(0)),
}

func (r Range2D) IsValid() bool {
	return r.R1.IsValid() && r.R2.IsValid()
}

func (r Range2D) IsIntersecting(a Range2D) bool {
	if !r.IsValid() {
		return false
	}

	if !a.IsValid() {
		return false
	}

	// ハズレ
	if r.R1.IsIntersecting(a.R1) || r.R2.IsIntersecting(a.R2) {
		return false
	}

	return true
}

func (r Range2D) Equal(a Range2D) bool {
	return r.R1.Equal(a.R1) && r.R2.Equal(a.R2)
}

func (r Range2D) Add(a Range2D) (r1, r2 Range2D) {
	/*
	 * +---++--+  +---+-+--+
	 * |   |/  /  |   / |  /
	 * +---++--+  +---+-+--+
	 */
	if r.R1.Equal(a.R1) && (r.R2.IsIntersecting(a.R2) || r.R2.End.Next().Equal(a.R2.Start) || r.R2.Start.Equal(a.R2.End.Next())) {
		r2 := r.R2
		if a.R2.Start.Less(r.R2.Start) {
			r2.Start = a.R2.Start
		}
		if r.R2.End.Less(a.R2.End) {
			r2.End = a.R2.End
		}
		return NewRange2D(r.R1.Start, r.R1.End, r2.Start, r2.End), Invalid2D
	}
	if r.R2.Equal(a.R2) && (r.R1.IsIntersecting(a.R1) || r.R1.End.Next().Equal(a.R1.Start) || r.R1.Start.Equal(a.R1.End.Next())) {
		r1 := r.R1
		if a.R1.Start.Less(r.R1.Start) {
			r1.Start = a.R1.Start
		}
		if r.R1.End.Less(a.R1.End) {
			r1.End = a.R1.End
		}
		return NewRange2D(r1.Start, r1.End, r.R2.Start, r.R2.End), Invalid2D
	}

	if !r.R1.IsIntersecting(a.R1) || !r.R2.IsIntersecting(a.R2) {
		return r, a
	}

	return r, a
}

/*
 * +--+---+ r
 * |      |
 * |  +------+
 * |  |//////| a
 * +--|//////|
 *    +------+
 *
 * split
 * +--+---+
 * |  |   |
 * +--+------+
 * |  |//////|
 * +--|//////|
 *    +------+
 *
 * minus
 * +--+---+ r1->
 * |  |   |
 * +--+---+
 * |  |
 * +--+
 *
 * r2
 * |
 * v
 *
 * join
 * +--+---+
 * |      |
 * +--+---+
 * |  |
 * +--+
 */
func (r Range2D) Minus(a Range2D, optJoin ...bool) []Range2D {
	if a.R1.ContainsRange(r.R1) && a.R2.ContainsRange(r.R2) {
		return nil
	}

	rr := make([]Range2D, 0)

	rr = append(rr, r)
	if !a.R1.ContainsRange(r.R1) {
		for i := 0; i < len(rr); i++ {
			tmp1, tmp2 := rr[i].SplitByR1(a.R1.Start)
			if tmp2.IsValid() {
				rr = append(rr, Invalid2D)
				copy(rr[i+2:], rr[i+1:])
				rr[i] = tmp1
				rr[i+1] = tmp2
				i++
			}
		}
		for i := 0; i < len(rr); i++ {
			tmp1, tmp2 := rr[i].SplitByR1(a.R1.End.Next())
			if tmp2.IsValid() {
				rr = append(rr, Invalid2D)
				copy(rr[i+2:], rr[i+1:])
				rr[i] = tmp1
				rr[i+1] = tmp2
				i++
			}
		}
	}
	if !a.R2.ContainsRange(r.R2) {
		for i := 0; i < len(rr); i++ {
			tmp1, tmp2 := rr[i].SplitByR2(a.R2.Start)
			if tmp2.IsValid() {
				rr = append(rr, Invalid2D)
				copy(rr[i+2:], rr[i+1:])
				rr[i] = tmp1
				rr[i+1] = tmp2
				i++
			}
		}
		for i := 0; i < len(rr); i++ {
			tmp1, tmp2 := rr[i].SplitByR2(a.R2.End.Next())
			if tmp2.IsValid() {
				rr = append(rr, Invalid2D)
				copy(rr[i+2:], rr[i+1:])
				rr[i] = tmp1
				rr[i+1] = tmp2
				i++
			}
		}
	}
	for i := len(rr) - 1; i >= 0; i-- {
		if a.R1.ContainsRange(rr[i].R1) && a.R2.ContainsRange(rr[i].R2) {
			rr = append(rr[:i], rr[i+1:]...)
		}
	}

	if len(optJoin) == 0 || optJoin[0] {
		joinByR1(&rr)
	}

	return rr
}

func (r Range2D) SplitByR1(p1 Sequential) (Range2D, Range2D) {
	if p1.Equal(r.R1.Start) ||
		p1.Less(r.R1.Start) ||
		r.R1.End.Less(p1) {
		//
		return r, Invalid2D
	}

	return NewRange2D(r.R1.Start, p1.Prev(), r.R2.Start, r.R2.End), NewRange2D(p1, r.R1.End, r.R2.Start, r.R2.End)
}

func (r Range2D) SplitByR2(p2 Sequential) (Range2D, Range2D) {
	if p2.Equal(r.R2.Start) ||
		p2.Less(r.R2.Start) ||
		r.R2.End.Less(p2) {
		//
		return r, Invalid2D
	}

	return NewRange2D(r.R1.Start, r.R1.End, r.R2.Start, p2.Prev()), NewRange2D(r.R1.Start, r.R1.End, p2, r.R2.End)
}

func joinByR1(rr *[]Range2D) {
	sortRange2DR1R2(*rr)

	for i := 0; i < len(*rr)-1; i++ {
		joined := false
		if !(*rr)[i].R1.Equal((*rr)[i+1].R1) {
			continue
		}
		if (*rr)[i].R2.End.Next().Equal((*rr)[i+1].R2.Start) {
			joined = true
			(*rr)[i].R2.End = (*rr)[i+1].R2.End
			*rr = append((*rr)[:i+1], (*rr)[i+1+1:]...)
		}
		if joined {
			i--
		}
	}

	for i := 0; i < len(*rr)-1; i++ {
		joined := false
		if !(*rr)[i].R2.Equal((*rr)[i+1].R2) {
			continue
		}
		if (*rr)[i].R1.End.Next().Equal((*rr)[i+1].R1.Start) {
			joined = true
			(*rr)[i].R1.End = (*rr)[i+1].R1.End
			*rr = append((*rr)[:i+1], (*rr)[i+1+1:]...)
		}
		if joined {
			i--
		}
	}
}

func sortRange2DR1R2(rr []Range2D) {
	sort.Slice(rr, func(i, j int) bool {
		if rr[i].R1.Start.Less(rr[j].R1.Start) {
			return true
		}
		if rr[i].R1.Start.Equal(rr[j].R1.Start) {
			if rr[i].R2.Start.Less(rr[j].R2.Start) {
				return true
			}
			if rr[i].R2.Start.Equal(rr[j].R2.Start) {
				if rr[i].R1.End.Less(rr[j].R1.End) {
					return true
				}
				if rr[i].R1.End.Equal(rr[j].R1.End) {
					if rr[i].R2.End.Less(rr[j].R2.End) {
						return true
					} else {
						return false
					}
				}
			}
		}
		return false
	})
}

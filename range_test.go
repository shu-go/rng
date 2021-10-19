package rng_test

import (
	"testing"

	"github.com/shu-go/gotwant"
	"github.com/shu-go/rng"
)

func TestIntRangeHougan(t *testing.T) {
	a, b := rng.NewRange(rng.Int(0), rng.Int(100)), rng.NewRange(rng.Int(25), rng.Int(30))

	t.Run("Minus", func(t *testing.T) {
		c, d, is := a.Minus(b)
		gotwant.Test(t, is, true)
		gotwant.Test(t, c.String(), "[0, 24]")
		gotwant.Test(t, d.String(), "[31, 100]")
	})

	t.Run("Add", func(t *testing.T) {
		c, d := a.Add(b)
		gotwant.Test(t, c.String(), "[0, 100]")
		gotwant.Test(t, d.String(), rng.Invalid.String())
	})
}

func TestIPv4RangeHougan(t *testing.T) {
	a, b := rng.NewRange(rng.NewIPv4("0.0.0.0"), rng.NewIPv4("100.0.0.0")), rng.NewRange(rng.NewIPv4("25.0.0.0"), rng.NewIPv4("30.0.0.0"))
	c, d, is := a.Minus(b)

	gotwant.Test(t, is, true)
	gotwant.Test(t, c.String(), "[[0 0 0 0], [24 255 255 255]]")
	gotwant.Test(t, d.String(), "[[30 0 0 1], [100 0 0 0]]")
}

func TestIntRangeHouganSare(t *testing.T) {
	a, b := rng.NewRange(rng.Int(0), rng.Int(100)), rng.NewRange(rng.Int(25), rng.Int(30))
	t.Run("Minus", func(t *testing.T) {
		c, d, is := b.Minus(a)
		gotwant.Test(t, is, true)
		gotwant.Test(t, c.String(), rng.Invalid.String())
		gotwant.Test(t, d.String(), rng.Invalid.String())
	})

	t.Run("Add", func(t *testing.T) {
		c, d := b.Add(a)
		gotwant.Test(t, c.String(), "[0, 100]")
		gotwant.Test(t, d.String(), rng.Invalid.String())
	})
}

func TestIntRangeKousa(t *testing.T) {
	a, b := rng.NewRange(rng.Int(0), rng.Int(100)), rng.NewRange(rng.Int(25), rng.Int(200))

	t.Run("Minus", func(t *testing.T) {
		c, d, is := a.Minus(b)
		gotwant.Test(t, is, true)
		gotwant.Test(t, c.String(), "[0, 24]")
		gotwant.Test(t, d.String(), rng.Invalid.String())

		c, d, is = b.Minus(a)
		gotwant.Test(t, is, true)
		gotwant.Test(t, c.String(), "[101, 200]")
		gotwant.Test(t, d.String(), rng.Invalid.String())
	})

	t.Run("Add", func(t *testing.T) {
		c, d := a.Add(b)
		gotwant.Test(t, c.String(), "[0, 200]")
		gotwant.Test(t, d.String(), rng.Invalid.String())

		c, d = b.Add(a)
		gotwant.Test(t, c.String(), "[0, 200]")
		gotwant.Test(t, d.String(), rng.Invalid.String())
	})
}

func TestIntRangeMiss(t *testing.T) {
	a, b := rng.NewRange(rng.Int(0), rng.Int(100)), rng.NewRange(rng.Int(125), rng.Int(200))

	t.Run("Minus", func(t *testing.T) {
		c, d, is := a.Minus(b)
		gotwant.Test(t, is, false)
		gotwant.Test(t, c.String(), "[0, 100]")
		gotwant.Test(t, d.String(), rng.Invalid.String())

		c, d, is = b.Minus(a)
		gotwant.Test(t, is, false)
		gotwant.Test(t, c.String(), "[125, 200]")
		gotwant.Test(t, d.String(), rng.Invalid.String())
	})

	t.Run("Add", func(t *testing.T) {
		c, d := a.Add(b)
		gotwant.Test(t, c.String(), "[0, 100]")
		gotwant.Test(t, d.String(), "[125, 200]")

		a, b = rng.NewRange(rng.Int(0), rng.Int(100)), rng.NewRange(rng.Int(101), rng.Int(200))
		c, d = a.Add(b)
		gotwant.Test(t, c.String(), "[0, 200]")
		gotwant.Test(t, d.String(), rng.Invalid.String())
	})
}

func TestIntRangeEqual(t *testing.T) {
	a, b := rng.NewRange(rng.Int(0), rng.Int(100)), rng.NewRange(rng.Int(0), rng.Int(100))
	c, d, is := a.Minus(b)

	gotwant.Test(t, is, true)
	gotwant.Test(t, c.String(), rng.Invalid.String())
	gotwant.Test(t, d.String(), rng.Invalid.String())
}

func TestContainsRange(t *testing.T) {
	gotwant.TestExpr(t, "equal point", rng.NewRange(rng.Int(0), rng.Int(0)).ContainsRange(rng.NewRange(rng.Int(0), rng.Int(0))))
	gotwant.TestExpr(t, "equal", rng.NewRange(rng.Int(0), rng.Int(1)).ContainsRange(rng.NewRange(rng.Int(0), rng.Int(1))))
	gotwant.TestExpr(t, "r1", rng.NewRange(rng.Int(-1), rng.Int(1)).ContainsRange(rng.NewRange(rng.Int(0), rng.Int(1))))
	gotwant.TestExpr(t, "r2", rng.NewRange(rng.Int(0), rng.Int(2)).ContainsRange(rng.NewRange(rng.Int(0), rng.Int(1))))
	gotwant.TestExpr(t, "r1r2", rng.NewRange(rng.Int(-1), rng.Int(2)).ContainsRange(rng.NewRange(rng.Int(0), rng.Int(1))))
	gotwant.TestExpr(t, "!r1", !rng.NewRange(rng.Int(0), rng.Int(1)).ContainsRange(rng.NewRange(rng.Int(-1), rng.Int(1))))
	gotwant.TestExpr(t, "!r2", !rng.NewRange(rng.Int(0), rng.Int(1)).ContainsRange(rng.NewRange(rng.Int(0), rng.Int(2))))
	gotwant.TestExpr(t, "!r1r2", !rng.NewRange(rng.Int(0), rng.Int(1)).ContainsRange(rng.NewRange(rng.Int(-1), rng.Int(2))))
}

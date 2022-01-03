package rng_test

import (
	"fmt"
	"testing"

	"github.com/shu-go/gotwant"
	"github.com/shu-go/rng"
)

func Benchmark2D(b *testing.B) {
	b.Run("Add", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r1 := rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(0), rng.Int(50))
			r2 := rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(100), rng.Int(150))
			r1.Add(r2)
		}
	})
	b.Run("Minus", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r1 := rng.NewRange2D(rng.Int(150), rng.Int(250), rng.Int(100), rng.Int(11100))
			r2 := rng.NewRange2D(rng.Int(100), rng.Int(200), rng.Int(1), rng.Int(255))
			r1.Minus(r2)
		}
	})
}

func Test2DAdd(t *testing.T) {
	t.Run("NonIntersecting", func(t *testing.T) {
		r1 := rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(0), rng.Int(50))
		r2 := rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(100), rng.Int(150))
		a, b := r1.Add(r2)
		gotwant.Test(t, a, r1)
		gotwant.Test(t, b, r2)

		r1 = rng.NewRange2D(rng.Int(0), rng.Int(50), rng.Int(0), rng.Int(50))
		r2 = rng.NewRange2D(rng.Int(150), rng.Int(200), rng.Int(100), rng.Int(150))
		a, b = r1.Add(r2)
		gotwant.Test(t, a, r1)
		gotwant.Test(t, b, r2)
	})

	t.Run("JoinBy1D", func(t *testing.T) {
		r1 := rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(0), rng.Int(50))
		r2 := rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(51), rng.Int(150))
		a, b := r1.Add(r2)
		gotwant.Test(t, a, rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(0), rng.Int(150)))
		gotwant.Test(t, b, rng.Invalid2D)

		r1 = rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(0), rng.Int(50))
		r2 = rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(25), rng.Int(150))
		a, b = r1.Add(r2)
		gotwant.Test(t, a, rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(0), rng.Int(150)))
		gotwant.Test(t, b, rng.Invalid2D)
	})

	t.Run("Contained", func(t *testing.T) {
		r1 := rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(0), rng.Int(50))
		r2 := rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(25), rng.Int(30))
		a, b := r1.Add(r2)
		gotwant.Test(t, a, r1)
		gotwant.Test(t, b, rng.Invalid2D)

		a, b = r2.Add(r1)
		gotwant.Test(t, a, r1)
		gotwant.Test(t, b, rng.Invalid2D)
	})

	t.Run("Intersecting", func(t *testing.T) {
		r1 := rng.NewRange2D(rng.Int(50), rng.Int(100), rng.Int(50), rng.Int(100))
		r2 := rng.NewRange2D(rng.Int(0), rng.Int(150), rng.Int(0), rng.Int(150))
		a, b := r1.Add(r2)
		gotwant.Test(t, a, r1)
		gotwant.Test(t, b, r2)
	})
}

func Test2DMinus(t *testing.T) {
	t.Run("NonIntersecting", func(t *testing.T) {
		r1 := rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(0), rng.Int(50))
		r2 := rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(100), rng.Int(150))
		rr := r1.Minus(r2)
		gotwant.Test(t, len(rr), 1)
		gotwant.Test(t, rr[0], r1)

		r1 = rng.NewRange2D(rng.Int(0), rng.Int(50), rng.Int(0), rng.Int(50))
		r2 = rng.NewRange2D(rng.Int(150), rng.Int(200), rng.Int(100), rng.Int(150))
		rr = r1.Minus(r2)
		gotwant.Test(t, len(rr), 1)
		gotwant.Test(t, rr[0], r1)
	})

	t.Run("R1 equal", func(t *testing.T) {
		r1 := rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(0), rng.Int(100))
		r2 := rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(51), rng.Int(150))
		rr := r1.Minus(r2)
		gotwant.Test(t, len(rr), 1)
		gotwant.Test(t, rr[0], rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(0), rng.Int(50)))

		r1 = rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(51), rng.Int(150))
		r2 = rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(0), rng.Int(100))
		rr = r1.Minus(r2)
		gotwant.Test(t, len(rr), 1)
		gotwant.Test(t, rr[0], rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(101), rng.Int(150)))
	})

	t.Run("R2 equal", func(t *testing.T) {
		r1 := rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(0), rng.Int(100))
		r2 := rng.NewRange2D(rng.Int(51), rng.Int(150), rng.Int(0), rng.Int(100))
		rr := r1.Minus(r2)
		gotwant.Test(t, len(rr), 1)
		gotwant.Test(t, rr[0], rng.NewRange2D(rng.Int(0), rng.Int(50), rng.Int(0), rng.Int(100)))

		r1 = rng.NewRange2D(rng.Int(51), rng.Int(150), rng.Int(0), rng.Int(100))
		r2 = rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(0), rng.Int(100))
		rr = r1.Minus(r2)
		gotwant.Test(t, len(rr), 1)
		gotwant.Test(t, rr[0], rng.NewRange2D(rng.Int(101), rng.Int(150), rng.Int(0), rng.Int(100)))
	})

	t.Run("ContainedIn", func(t *testing.T) {
		r1 := rng.NewRange2D(rng.Int(0), rng.Int(100), rng.Int(0), rng.Int(100))
		r2 := rng.NewRange2D(rng.Int(25), rng.Int(50), rng.Int(25), rng.Int(50))
		rr := r1.Minus(r2)
		gotwant.Test(t, len(rr), 4)
		gotwant.Test(t, rr[0], rng.NewRange2D(rng.Int(0), rng.Int(24), rng.Int(0), rng.Int(100)))
		gotwant.Test(t, rr[1], rng.NewRange2D(rng.Int(25), rng.Int(50), rng.Int(0), rng.Int(24)))
		gotwant.Test(t, rr[2], rng.NewRange2D(rng.Int(25), rng.Int(50), rng.Int(51), rng.Int(100)))
		gotwant.Test(t, rr[3], rng.NewRange2D(rng.Int(51), rng.Int(100), rng.Int(0), rng.Int(100)))
	})

	t.Run("rb", func(t *testing.T) {
		r1 := rng.NewRange2D(rng.Int(150), rng.Int(250), rng.Int(100), rng.Int(11100))
		r2 := rng.NewRange2D(rng.Int(100), rng.Int(200), rng.Int(1), rng.Int(255))
		rr := r1.Minus(r2)
		gotwant.Test(t, len(rr), 2)
		gotwant.Test(t, rr[0], rng.NewRange2D(rng.Int(150), rng.Int(200), rng.Int(256), rng.Int(11100)))
		gotwant.Test(t, rr[1], rng.NewRange2D(rng.Int(201), rng.Int(250), rng.Int(100), rng.Int(11100)))
	})

	t.Run("1D", func(t *testing.T) {
		r1 := rng.NewRange2D(rng.Int(0), rng.Int(65535), rng.IPv4{192, 168, 0, 0}, rng.IPv4{192, 168, 255, 255})
		r2 := rng.NewRange2D(rng.Int(100), rng.Int(100), rng.IPv4{192, 168, 1, 0}, rng.IPv4{192, 168, 1, 255})
		rr := r1.Minus(r2)
		fmt.Printf("%v\n", r1)
		fmt.Printf("%v\n", r2)
		gotwant.Test(t, len(rr), 4)
		gotwant.Test(t, rr[0], rng.NewRange2D(rng.Int(0), rng.Int(99), rng.IPv4{192, 168, 0, 0}, rng.IPv4{192, 168, 255, 255}))
		gotwant.Test(t, rr[1], rng.NewRange2D(rng.Int(100), rng.Int(100), rng.IPv4{192, 168, 0, 0}, rng.IPv4{192, 168, 0, 255}))
		gotwant.Test(t, rr[2], rng.NewRange2D(rng.Int(100), rng.Int(100), rng.IPv4{192, 168, 2, 0}, rng.IPv4{192, 168, 255, 255}))
		gotwant.Test(t, rr[3], rng.NewRange2D(rng.Int(101), rng.Int(65535), rng.IPv4{192, 168, 0, 0}, rng.IPv4{192, 168, 255, 255}))
	})
}

func TestIPv4(t *testing.T) {
	i := rng.IPv4{192, 168, 1, 0}
	ii := i.Prev()
	gotwant.Test(t, ii, rng.IPv4{192, 168, 0, 255})

	i = rng.IPv4{192, 168, 0, 255}
	ii = i.Next()
	gotwant.Test(t, ii, rng.IPv4{192, 168, 1, 0})
}

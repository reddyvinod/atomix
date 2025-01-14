package atomix

import (
	"math"
	"sync/atomic"
)

// Float64 is an atomic wrapper around float64.
type Float64 struct {
	v uint64
}

// NewFloat64 creates a float64.
func NewFloat64(f float64) *Float64 {
	return &Float64{math.Float64bits(f)}
}

// Load atomically the value.
func (f *Float64) Load() float64 {
	return math.Float64frombits(atomic.LoadUint64(&f.v))
}

// Store atomically the passed value.
func (f *Float64) Store(s float64) {
	atomic.StoreUint64(&f.v, math.Float64bits(s))
}

// Add atomically and return the new value.
func (f *Float64) Add(s float64) float64 {
	for {
		old := f.Load()
		new := old + s
		if f.CAS(old, new) {
			return new
		}
	}
}

// Sub atomically and return the new value.
func (f *Float64) Sub(s float64) float64 {
	return f.Add(-s)
}

// CAS is an atomic Compare-and-swap.
func (f *Float64) CAS(old, new float64) bool {
	return atomic.CompareAndSwapUint64(&f.v, math.Float64bits(old), math.Float64bits(new))
}

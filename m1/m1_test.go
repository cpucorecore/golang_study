package m1_test

import (
	"github.com/cpucorecore/golang_study/m2"
	"testing"
)

func Test_DepentMe(t *testing.T) {
	m2.Hello()
}

func Benchmark_DepentMe(b *testing.B) {
	m2.Hello()
}

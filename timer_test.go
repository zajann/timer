package timer

import (
	"testing"
	"time"
)

func BenchmarkTimer(b *testing.B) {
	Init(0)

	for i := 0; i < b.N; i++ {
		GetTime()
	}
}

func BenchmarkOriginTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now()
	}
}

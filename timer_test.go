package timer

import (
	"testing"
	"time"
)

func BenchmarkGetFormatStringWithTimer(b *testing.B) {
	Init(0)

	for i := 0; i < b.N; i++ {
		GetTimeFormat("2006-01-02 15:04:05")
	}
}

func BenchmarkGetFormatStringWithoutTimer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now().Format("2006-01-02 15:04:05")
	}
}

func BenchmarkGetMonthStringWithTimer(b *testing.B) {
	Init(0)

	for i := 0; i < b.N; i++ {
		GetMonthString()
	}
}

func BenchmarkGetMonthStringWithoutTimer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now().Month().String()
	}
}

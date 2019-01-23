package mytime

import (
	"fmt"
	"testing"
)

func TestStrToTime(t *testing.T) {
	a := StrToTime("2019-01-02T12:05:11")
	fmt.Println(a)
}

func BenchmarkStrToTimeWith100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrToTime("2019-01-02T12:05:11")
	}
}

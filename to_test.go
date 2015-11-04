package util

import (
	"testing"
)

func TestRoundFloat64(t *testing.T) {
	t.Log("test RoundFloat64 begin")
	data := []struct {
		input  float64
		prec   int
		output float64
	}{
		{10.12341, 4, 10.1234},
		{23.441, 2, 23.44},
		{23.447, 2, 23.45},
		{0.34516, 4, 0.3452},
	}
	for _, v := range data {
		tmp := RoundFloat64(v.input, v.prec)
		if tmp != v.output {
			t.Fatal("wish result:", v.output, "| real result:", tmp)
		} else {
			t.Log("wish result:", v.output, "| real result:", tmp)
		}
	}
	t.Log("test RoundFloat64 end")
}

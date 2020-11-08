package filter

import (
	"testing"
)

func TestSHash(t *testing.T) {
	type args struct {
		s string
	}
	a1 := args{
		s: "adsff",
	}
	a2 := args{
		s: "adsff1",
	}
	a3 := args{
		s: "adsff2",
	}
	a4 := args{
		s: "adsff3",
	}
	a5 := args{
		s: "adsff4",
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: a1,
		},
		{
			name: "2",
			args: a2,
		},
		{
			name: "3",
			args: a3,
		},
		{
			name: "4",
			args: a4,
		},
		{
			name: "5",
			args: a5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(shash(tt.args.s))
		})
	}
}

func TestBloom_Contains(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		b    *Bloom
		args args
		want bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Contains(tt.args.s); got != tt.want {
				t.Errorf("Bloom.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

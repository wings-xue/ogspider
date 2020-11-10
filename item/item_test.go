package item

import "testing"

func TestHasValue(t *testing.T) {
	type args struct {
		name  string
		model Field
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				name:  "BaseCSS",
				model: Field{Name: "asdf"},
			},
			want: true,
		},
		{
			args: args{
				name:  "Name",
				model: Field{Name: "asdfas"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasValue(tt.args.name, tt.args.model); got != tt.want {
				t.Errorf("HasValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

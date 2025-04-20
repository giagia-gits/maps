package math

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		x, y interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "x and y is int",
			args: args{
				x: int(1),
				y: int(2),
			},
			want: int(1) + int(2),
		},
		{
			name: "x and y is uint",
			args: args{
				x: uint(1),
				y: uint(2),
			},
			want: uint(1) + uint(2),
		},
		{
			name: "x and y is float",
			args: args{
				x: float64(1.1),
				y: float64(2.2),
			},
			want: float64(1.1) + float64(2.2),
		},
		{
			name: "x and y is others",
			args: args{
				x: "abc",
				y: "xyz",
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.args.x, tt.args.y)
			if result != tt.want {
				t.Errorf("expected %v, but got %v", tt.want, result)
			}
		})
	}
}

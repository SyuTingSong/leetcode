package q6

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"LEETCODEISHIRING",
			args{"LEETCODEISHIRING", 3},
			"LCIRETOESIIGEDHN",
		},
		{
			"LEETCODEISHIRING",
			args{"LEETCODEISHIRING", 4},
			"LDREOEIIECIHNTSG",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

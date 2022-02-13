package codepractice

import "testing"

func Test_getLenght(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestLength1",
			args: args{
				str: "Malli",
			},
			want: 5,
		},
		{
			name: "TestLength2",
			args: args{
				str: "Mallikarjun",
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLength(tt.args.str); got != tt.want {
				t.Errorf("getLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

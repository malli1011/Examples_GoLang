package codepractice

import (
	"fmt"
	"testing"
)

//TestCalcSquare simple test
//Test should start with Test and followed by method name under test
//Test file should have _test.go extension
func Test_calcSquareSimple(t *testing.T) {
	expectedVal := 100
	res, _ := calcSquare(10)

	if res != expectedVal {
		t.Error("Expect", expectedVal, "Got", res)
		t.Fail()
	}

}

//parameterised tests
//https://go.dev/blog/subtests
func Test_calcSquare(t *testing.T) {
	type args struct {
		i int
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Positive Value", args: args{i: 10}, want: 100, wantErr: false},
		{name: "Negative Value", args: args{i: -1}, want: 0, wantErr: true},
		{name: "Zero Value", args: args{i: 0}, want: 0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calcSquare(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcSquare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calcSquare() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//Example for the documentation
//https://go.dev/blog/examples
func Example_calcSquare() {
	fmt.Println(calcSquare(10))
	//Output:
	//100 <nil>
}

//go test -bench .
//go test -bench custom_error_ex.go

func Benchmark_calcSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = calcSquare(10)
	}
}

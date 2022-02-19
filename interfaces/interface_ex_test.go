package interfaces

import "testing"

func Test_printGreeting(t *testing.T) {
	type args struct {
		b bot
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "English bot test", args: args{englishBot{}}},
		{name: "Spanish bot test", args: args{spanishBot{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printGreeting(tt.args.b)
		})
	}
}

func Test_printGreeting2(t *testing.T) {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

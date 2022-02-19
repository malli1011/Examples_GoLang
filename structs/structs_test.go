package structs

import "testing"

func Test_demo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"Frist test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			demo()
		})
	}
}

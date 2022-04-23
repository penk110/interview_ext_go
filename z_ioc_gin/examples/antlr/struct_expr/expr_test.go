package examples

import "testing"

func Test_structExpr(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "structExpr",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			structExpr()
		})
	}
}

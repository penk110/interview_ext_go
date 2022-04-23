package func_expr

import "testing"

func Test_funcExpr(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "funcExpr",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			funcExpr()
		})
	}
}

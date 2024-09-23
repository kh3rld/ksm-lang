package eval

import (
	"reflect"
	"testing"

	"github.com/kh3rld/ksm-lang/parser"
)

func TestEvaluator_Eval(t *testing.T) {
	type args struct {
		node parser.Node
	}
	tests := []struct {
		name string
		e    *Evaluator
		args args
		want *Number
	}{
		{
			name: "Single number",
			e:    &Evaluator{},
			args: args{node: &parser.NumberExpr{Value: 42}},
			want: &Number{Value: 42},
		},
		{
			name: "Addition of two numbers",
			e:    &Evaluator{},
			args: args{
				node: &parser.BinaryExpr{
					Left:     &parser.NumberExpr{Value: 3},
					Right:    &parser.NumberExpr{Value: 5},
					Operator: "+",
				},
			},
			want: &Number{Value: 8},
		},
		{
			name: "Complex expression",
			e:    &Evaluator{},
			args: args{
				node: &parser.BinaryExpr{
					Left: &parser.BinaryExpr{
						Left:     &parser.NumberExpr{Value: 3},
						Right:    &parser.NumberExpr{Value: 2},
						Operator: "+",
					},
					Right:    &parser.NumberExpr{Value: 4},
					Operator: "-",
				},
			},
			want: &Number{Value: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Evaluator{}
			if got := e.Eval(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Evaluator.Eval() = %v, want %v", got, tt.want)
			}
		})
	}
}

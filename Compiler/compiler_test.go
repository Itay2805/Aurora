package compiler

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestImmidiateExpression(t *testing.T) {
	ast := BuildAST(
		`123
		"test"
		something`)

	expectedAST := &CompilerContext{
		Expressions: []Expr{
			IntImmidiateExpr{
				Value: 123,
			},
			StringImmidateExpr{
				Value: "test",
			},
			IdentImmidiateExpr{
				Identifier: "something",
			},
		},
	}

	if !reflect.DeepEqual(ast, expectedAST) {
		t.Errorf("Got:\n%+v\n Expected:\n%+v\n", ast, expectedAST)
	}
}

func TestFunctionCallsAndMemberAccesses(t *testing.T) {
	ast := BuildAST(
		`test.member
		method()
		test.method()
		test(123, test())
		`)

	expectedAST := &CompilerContext{
		Expressions: []Expr{
			MemberAccessExpr{
				AccessFrom: IdentImmidiateExpr{
					Identifier: "test",
				},
				MemberName: "member",
			},
			FunctionCallExpr{
				Function: IdentImmidiateExpr{
					Identifier: "method",
				},
				Parameters: []Expr{},
			},
			FunctionCallExpr{
				Function: MemberAccessExpr{
					AccessFrom: IdentImmidiateExpr{
						Identifier: "test",
					},
					MemberName: "method",
				},
				Parameters: []Expr{},
			},
			FunctionCallExpr{
				Function: IdentImmidiateExpr{
					Identifier: "test",
				},
				Parameters: []Expr{
					IntImmidiateExpr{
						Value: 123,
					},
					FunctionCallExpr{
						Function: IdentImmidiateExpr{
							Identifier: "test",
						},
						Parameters: []Expr{},
					},
				},
			},
		},
	}
	{
		b, err := json.Marshal(ast)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
	}
	{
		b, err := json.Marshal(expectedAST)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
	}
	if !reflect.DeepEqual(ast, expectedAST) {
		t.Errorf("Got:\n%+v\n Expected:\n%+v\n", ast, expectedAST)
	}
}

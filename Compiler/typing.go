package compiler

// GetExprType this is really improvised and uses alot
//  of assumptions but it will do for now I guess
func GetExprType(context *ProgramContext, expr Expr) *Type {
	switch v := expr.(type) {
	case IntImmidiateExpr:
		return &Type{
			GCPtr:    false,
			Name:     "i64",
			Optional: false,
			Ptr:      0,
			Ref:      false,
		}
	case StringImmidateExpr:
		return &Type{
			GCPtr:    false,
			Name:     "string",
			Optional: false,
			Ptr:      0,
			Ref:      false,
		}
	case IdentImmidiateExpr:
		return &context.SearchForVariableInCurrentOrGlobal(v.Identifier).Type
	case MemberAccessExpr:
		/// will need to make this much better, member access is going to be annoying...
		/// for now we assume that global must be
		ident, ok := v.AccessFrom.(IdentImmidiateExpr)
		if ok {
			va := context.SearchForVariableInModule(v.MemberName, ident.Identifier)
			if va.Export || va.Module == context.CurrentModule {
				return &va.Type
			}
		}
		return nil

	/// will need to change these, so it will return depends on the object
	/// since we want to have operator overloading in the future
	case LHSOperatorExpr:
		return GetExprType(context, v.On)
	case LROperatorExpr:
		return GetExprType(context, v.Left)

	case AssignExpr:
		return GetExprType(context, v.Target)
	}
	return nil
}

func TypeCompare(a Type, b Type) bool {
	if a.Name != b.Name {
		return false
	}
	if a.GCPtr != b.GCPtr {
		return false
	}
	if a.Ptr != b.Ptr {
		return false
	}
	if a.Ref != b.Ref {
		return false
	}
	return true
}

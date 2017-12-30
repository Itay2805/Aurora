package compiler

import (
	"log"
)

type ProgramContext struct {
	Modules       map[string]*Module
	CurrentModule *Module
}

func (ctx ProgramContext) SearchForVariableInCurrentOrGlobal(name string) *Variable {
	for _, v := range ctx.CurrentModule.Variables {
		if v.Name == name {
			return v
		}
	}
	if ctx.CurrentModule.Name != "Global" {
		return ctx.SearchForVariableInModule(name, "Global")
	}
	return nil
}

func (ctx ProgramContext) SearchForVariableInModule(name string, module string) *Variable {
	mod, exists := ctx.Modules[module]
	if !exists {
		log.Fatalln("The module " + module + " does not exists")
		return nil
	}
	for _, v := range mod.Variables {
		if v.Name == name {
			return v
		}
	}
	return nil
}

type Module struct {
	Name      string
	Variables map[string]*Variable
}

type Variable struct {
	Module *Module
	Name   string
	Type   Type
	Export bool
}

func SetupProgram(ctx *ASTContext) *ProgramContext {
	program := ProgramContext{
		Modules: make(map[string]*Module),
	}
	// add global module
	program.Modules["Global"] = &Module{
		Name:      "Global",
		Variables: make(map[string]*Variable),
	}
	program.CurrentModule = program.Modules["Global"]

	/// First we should search for all the variables
	/// we also need a way to make sure if we do
	///
	/// module B { var export b := A.a; }
	/// module A { var export a := 5; }
	///
	/// it will request A to be processed so it will
	/// find the type of a to set the type of b

	global := program.Modules["Global"]
	for _, block := range ctx.Blocks {
		v, ok := block.(VariableStmt)

		if ok {
			_, ok := global.Variables[v.Name]
			if ok {
				log.Println("variable " + v.Name + " already exists")
				return nil
			}
			xtype := v.Type
			if xtype.Name == "" {
				temp := GetExprType(&program, v.Value)
				if temp == nil {
					log.Println("Could not infer type of " + v.Name)
					return nil
				}
				xtype = *temp
			}
			if v.Value == nil && !v.Type.Optional {
				log.Println("tried to define un-optional variable without a value (" + v.Name + ")")
				return nil
			}
			if !TypeCompare(xtype, *GetExprType(&program, v.Value)) {
				if !(GetExprType(&program, v.Value).Name == "i64" && (xtype.Name == "i8" || xtype.Name == "i16" || xtype.Name == "i32")) {
					log.Println("expresion does not match the type of variable " + v.Name)
					return nil
				}
			}
			global.Variables[v.Name] = &Variable{
				Module: global,
				Name:   v.Name,
				Export: v.Export,
				Type:   xtype,
			}
		}
	}

	return &program
}

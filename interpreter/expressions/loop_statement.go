package expressions

import (
	"blom/ast"
	"blom/env"
	"blom/env/objects"
)

func InterpretForLoopStatement(interpreter Interpreter, environment *env.Environment[objects.Object], statement *ast.ForLoopStatement) {
	declaration := statement.Declaration
	condition := statement.Condition
	step := statement.Step
	body := statement.Body

	if declaration != nil {
		interpreter.InterpretStatement(declaration, environment)
	}

	for {
		conditionValue := interpreter.InterpretStatement(condition, environment)
		if conditionValue == nil {
			break
		}

		if conditionValue.(*objects.BooleanObject).Value {
			interpreter.InterpretBlock(body, environment)
			interpreter.InterpretStatement(step, environment)
		} else {
			break
		}
	}
}

func InterpretWhileLoopStatement(interpreter Interpreter, environment *env.Environment[objects.Object], statement *ast.WhileLoopStatement) {
	condition := statement.Condition
	body := statement.Body

	for {
		conditionValue := interpreter.InterpretStatement(condition, environment)
		if conditionValue == nil {
			break
		}

		if conditionValue.(*objects.BooleanObject).Value {
			interpreter.InterpretBlock(body, environment)
		} else {
			break
		}
	}
}

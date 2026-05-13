package main

import (
	"fmt"
)

// The idea here is that the evaluate function performs a huge switch statement on each possible
// expression, and computes the logic for each, recursing until it's able to return a literal
// Every function lives within the Interpreter struct

type Interpreter struct {}

func (i *Interpreter) evaluate(expression Expr) (any, error) {
	switch e := expression.(type) {
	case Literal:
		return e.Value, nil
	case Grouping:
		group, err := i.evaluate(e.Expression)
		if err != nil {
			return nil, err
		}
		return group, nil
	case Unary:
		right, err := i.evaluate(e.Right)
		if err != nil {
			return nil, err
		}

		switch e.Operator.Type {
		case BANG:
			return i.isTruthy(right), nil
		case MINUS:
			// this functions ensure that token following a minus is a number. if not, it returns an error
			right, err := i.toNumber(e.Operator, right)
			if err != nil {
				return nil, err
			}
			return -right, nil
		default:
			return nil, fmt.Errorf("unreachable: unknown expression type")
		}
	case Binary:
		left, err := i.evaluate(e.Left)
		if err != nil {
			return nil, err
		}
		right, err := i.evaluate(e.Right)
		if err != nil {
			return nil, err
		}
		// for binary operations, most operators only work if boths operands are the same type
		// usually numbers, so we have to check the types of left and right to ensure it is the
		// case. If we not, we return an error
		switch e.Operator.Type {
		case GREATER:
			l, r, err := i.toNumbers(e.Operator, left, right)
			if err != nil {
				return nil, err
			}
			return l > r, nil
		case GREATER_EQUAL:
			l, r, err := i.toNumbers(e.Operator, left, right)
			if err != nil {
				return nil, err
			}
			return l >= r, nil
		case LESS:
			l, r, err := i.toNumbers(e.Operator, left, right)
			if err != nil {
				return nil, err
			}
			return l < r, nil
		case LESS_EQUAL:
			l, r, err := i.toNumbers(e.Operator, left, right)
			if err != nil {
				return nil, err
			}
			return l <= r, nil
		case BANG_EQUAL:
			l, r, err := i.toNumbers(e.Operator, left, right)
			if err != nil {
				return nil, err
			}
			return !i.isEqual(l, r), nil
		case EQUAL_EQUAL:
			l, r, err := i.toNumbers(e.Operator, left, right)
			if err != nil {
				return nil, err
			}
			return i.isEqual(l, r), nil
		case MINUS:
			l, r, err := i.toNumbers(e.Operator, left, right)
			if err != nil {
				return nil, err
			}
			return l - r, nil
		case PLUS:
			switch l := left.(type) {
			case float64:
				if r, ok := right.(float64); ok {
					return l + r, nil
				}
			case string:
				if r, ok := right.(string); ok {
					return l + r, nil
				}
			}
			return nil, &RuntimeError{Token: e.Operator, Message: "Operands must be two numbers or two strings"}
		case SLASH:
			l, r, err := i.toNumbers(e.Operator, left, right)
			if err != nil {
				return nil, err
			}
			return l / r, nil
		case STAR:
			l, r, err := i.toNumbers(e.Operator, left, right)
			if err != nil {
				return nil, err
			}
			return l * r, nil
		default:
			return nil, &RuntimeError{Token: e.Operator, Message: "Operands must be numbers"}
		}
	default:
		return nil, &RuntimeError{Token: Token{}, Message: "Operands must be numbers"}
	}	
}

func (i *Interpreter) toNumbers(operator Token, left, right any) (float64, float64, error) {
	l, lok := left.(float64)
	r, rok := right.(float64)
	if !lok || !rok {
		return 0, 0, &RuntimeError{Token: operator, Message: "Operands must be numbers"}
	}
	return l, r, nil
}

func (i *Interpreter) toNumber(operator Token, operand any) (float64, error) {
	if num, ok := operand.(float64); ok {
		return num, nil
	}
	return 0, &RuntimeError{Token: operator, Message: "Operand must be a number"}
}

func (i *Interpreter) isTruthy(object any) bool {
	if object == nil {
		return false
	}
	if v, ok := object.(bool); ok {
		return v
	}
	return true

}

func (i *Interpreter) isEqual(left any, right any) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil {
		return false
	}
	return left == right
}

func (i *Interpreter) Interpret(expression Expr) error {
	value, err := i.evaluate(expression)
	if err != nil {
		return err
	}
	fmt.Println(i.stringify(value))
	return nil
}

func (i *Interpreter) stringify(object any) string {
	if object == nil {
		return "nil"
	}

	if num, ok := object.(float64); ok {
		text := fmt.Sprintf("%g", num)
		return text
	}
	return fmt.Sprintf("%v", object)
}
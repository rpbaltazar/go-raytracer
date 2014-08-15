package gomath

import "testing"

func Test_QuadraticEquationSolver(t *testing.T) {
  a := -2.5
  b := 3.29
  c := 3.22
  out := QuadraticSolution {NumSolutions: 2, S1: 1.9698551749335749 , S2: -0.6538551749335747}

  if res := QuadraticEquationSolver(a,b,c); res != out{
    t.Errorf("QuadraticEquationSolver(%v, %v, %v) => expected %v, got %v", a, b, c, out, res)
  }

  a = 4
  b = 4
  c = 1
  out = QuadraticSolution {NumSolutions: 1, S1: -0.5}

  if res := QuadraticEquationSolver(a,b,c); res != out{
    t.Errorf("QuadraticEquationSolver(%v, %v, %v) => expected %v, got %v", a, b, c, out, res)
  }

  a = 4
  b = 3
  c = 1
  out = QuadraticSolution {NumSolutions: 0}

  if res := QuadraticEquationSolver(a,b,c); res != out{
    t.Errorf("QuadraticEquationSolver(%v, %v, %v) => expected %v, got %v", a, b, c, out, res)
  }

}

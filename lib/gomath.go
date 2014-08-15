package gomath

import "math"

type QuadraticSolution struct{
  S1, S2 float64
  NumSolutions int
}

func QuadraticEquationSolver(a float64, b float64, c float64) QuadraticSolution{
  delta := math.Pow(b, 2) - 4*a*c

  if delta < 0{
    return QuadraticSolution{NumSolutions: 0}
  }else if delta == 0{
    return QuadraticSolution{NumSolutions: 1, S1: (-b) / (2*a)}
  }else {
    s1 := (-b - math.Sqrt(delta))/(2*a)
    s2 := (-b + math.Sqrt(delta))/(2*a)
    return QuadraticSolution{NumSolutions: 2, S1: s1, S2: s2}
  }
}

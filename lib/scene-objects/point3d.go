package scene_objects

import "math"

type Point3D struct {
  X, Y, Z float64
}

func (p *Point3D) Subtract(p1 Point3D) Point3D {
  return Point3D{ X: p.X-p1.X, Y: p.Y-p1.Y, Z: p.Z-p1.Z}
}

func (p *Point3D) Add(p1 Point3D) Point3D {
  return Point3D{ X: p.X+p1.X, Y: p.Y+p1.Y, Z: p.Z+p1.Z}
}

func (p *Point3D) Multiply(f1 float64) Point3D {
  return Point3D{ X: p.X*f1, Y: p.Y*f1, Z: p.Z*f1}
}

func (p *Point3D) ScalarProd(p1 Point3D) float64 {
  return p.X*p1.X + p.Y*p1.Y + p.Z*p1.Z
}

func (p *Point3D) Pow(f float64) Point3D {
  return Point3D{ X: math.Pow(p.X, f), Y: math.Pow(p.Y, f), Z: math.Pow(p.Z, f) }
}

func (p *Point3D) ComputeDistance(p1 Point3D) float64{
  a := math.Pow( p1.X - p.X, 2)
  b := math.Pow( p1.Y - p.Y, 2)
  c := math.Pow( p1.Z - p.Z, 2)

  return math.Sqrt (a+b+c)
}

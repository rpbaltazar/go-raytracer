package main

import "fmt"
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

func (p *Point3D) PointMultiply(p1 Point3D) Point3D {
  return Point3D{ X: p.X*p1.X, Y: p.Y*p1.Y, Z: p.Z*p1.Z}
}

func (p *Point3D) ScalarProd(p1 Point3D) float64 {
  return  p.X*p1.X + p.Y*p1.Y + p.Z*p1.Z
}

func (p *Point3D) Pow(f float64) Point3D {
  return Point3D{ X: math.Pow(p.X, f), Y: math.Pow(p.Y, f), Z: math.Pow(p.Z, f) }
}

type Ray struct {
  Origin, Direction Point3D
}

type Camera struct {
  Position Point3D
}

type Sphere struct {
  Center Point3D
  SphereRay float64
}

var imageWidth, imageHeight int = 800, 640

func main() {

  camera := Camera{ Point3D{X: float64(imageWidth)/2, Y: float64(imageHeight)/2, Z: -100 } }

  scene := loadScene()

  fmt.Println(camera)

  for px := 0; px < imageWidth; px++{
    for py :=0; py < imageHeight; py++{
      //Create ray with origin ox, oy and direction from px, py to ox, oy
      primRay := computeRay(px, py, camera)
      // fmt.Println(primRay)
      //Scene.isIntersectedBy(primRay)
      for idx := 0; idx < len(scene); idx ++{
        inter := intersects(primRay, scene[idx])
        fmt.Println(inter)
      }
    }
  }
}

func intersects(primRay Ray, obj Sphere) int{
  v := primRay.Direction
  o := primRay.Origin
  c := obj.Center
  r := obj.SphereRay

  oc := o.Subtract(c)

  A := v.ScalarProd(v)
  B := 2* oc.ScalarProd(v)
  C := oc.ScalarProd(oc) - math.Pow(r,2)

  delta := math.Pow(B, 2) - 4*A*C
  //two solutions - calculate distance and return the closest
  if delta > 0{
    return 2
  // one solution - return distance
  }else if delta == 0{
    return 1
  }else{
    return -1
  }
}

func loadScene() []Sphere{
  obj := Sphere{ Center: Point3D {X: 150, Y: 150, Z: 150}, SphereRay: 50 }
  return []Sphere{obj}
}

func computeRay(px, py int, cam Camera) Ray {
  rayDirX := float64(px) - cam.Position.X
  rayDirY := float64(py) - cam.Position.Y
  rayDirZ := -cam.Position.Z

  return Ray{ Origin: cam.Position, Direction: Point3D {rayDirX, rayDirY, rayDirZ} }
}

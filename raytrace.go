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

type Intersection struct {
  Object Sphere
  Distance float64
  HitPoint Point3D
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
      //Scene.isIntersectedBy(primRay)
      curr_intersection := new(Intersection)

      for idx := 0; idx < len(scene); idx ++{
        new_intersection := intersects(primRay, scene[idx])
        //intersects
        if (new_intersection.Distance > 0){
          fmt.Println("new intersection", new_intersection)
        }
        if new_intersection.Distance > 0 && new_intersection.Distance < curr_intersection.Distance{
          curr_intersection.Distance = new_intersection.Distance
          curr_intersection.Object = new_intersection.Object
        }
      }
      if (curr_intersection.Distance > 0){
        fmt.Println(curr_intersection)
      }
    }
  }
}

// TODO: Move equation solver to specific method (math lib maybe)
// solution[] = [sol1, sol2]

func intersects(primRay Ray, obj Sphere) Intersection{
  v := primRay.Direction
  o := primRay.Origin
  c := obj.Center
  r := obj.SphereRay

  oc := o.Subtract(c)

  A := v.ScalarProd(v)
  B := 2* oc.ScalarProd(v)
  C := oc.ScalarProd(oc) - math.Pow(r,2)

  delta := math.Pow(B, 2) - 4*A*C
  intersection := Intersection{ Distance: -1}
  //two solutions - calculate distance and return the closest
  if delta > 0{
    sol1 := (-B + delta)/(2*A)
    sol2 := (-B - delta)/(2*A)

    hitP1 := computeHitPoint(sol1, primRay)
    hitP2 := computeHitPoint(sol2, primRay)

    distance1 := primRay.Origin.ComputeDistance(hitP1)
    distance2 := primRay.Origin.ComputeDistance(hitP2)

    if distance1 > distance2{
      intersection.Distance = distance2
      intersection.HitPoint = hitP2
    } else{
      intersection.Distance = distance1
      intersection.HitPoint = hitP1
    }

    intersection.Object = obj
    return intersection
  // one solution - return distance
  }else if delta == 0{
    sol := (-B)/(2*A)
    hitP := computeHitPoint(sol, primRay)
    distance := primRay.Origin.ComputeDistance(hitP)

    intersection.HitPoint = hitP
    intersection.Distance = distance
    intersection.Object = obj

    return intersection
  }else{
    return intersection
  }
}

func computeHitPoint(t float64, r Ray) Point3D {
  hX := r.Origin.X + t*r.Direction.X
  hY := r.Origin.Y + t*r.Direction.Y
  hZ := r.Origin.Z + t*r.Direction.Z

  hitP := Point3D { X: hX, Y: hY, Z: hZ }
  return hitP
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

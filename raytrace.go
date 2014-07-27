package main

import "fmt"

type Point3D struct {
  X, Y, Z int

}
func (p *Point3D) Subtract(p1 Point3D) Point3D {
  return Point3D{ X: p.X-p1.X, Y: p.Y-p1.Y, Z: p.Z-p1.Z}
}

type Ray struct {
  Origin, Direction Point3D
}

type Camera struct {
  Position Point3D
}

type Sphere struct {
  Center Point3D
  SphereRay int
}

var imageWidth, imageHeight int = 800, 640

func main() {

  camera := Camera{ Point3D{X: imageWidth/2, Y: imageHeight/2, Z: -100 } }

  scene := loadScene()

  fmt.Println(camera)

  for px := 0; px < imageWidth; px++{
    for py :=0; py < imageHeight; py++{
      //Create ray with origin ox, oy and direction from px, py to ox, oy
      primRay := computeRay(px, py, camera)
      // fmt.Println(primRay)
      //Scene.isIntersectedBy(primRay)
      for idx := 0; idx < len(scene); idx ++{
        if intersects(primRay, scene[idx]){
          fmt.Println(primRay)
        }
      }
    }
  }
}

func intersects(primRay Ray, obj Sphere) bool{
  return true
}

func loadScene() []Sphere{
  obj := Sphere{ Center: Point3D {X: 150, Y: 150, Z: 150}, SphereRay: 50 }
  return []Sphere{obj}
}

func computeRay(px, py int, cam Camera) Ray {
  rayDirX := px - cam.Position.X
  rayDirY := py - cam.Position.Y
  rayDirZ := -cam.Position.Z

  return Ray{ Origin: cam.Position, Direction: Point3D {rayDirX, rayDirY, rayDirZ} }
}

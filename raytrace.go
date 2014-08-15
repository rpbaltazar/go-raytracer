package main

import (
  "fmt"
  "math"
  "go-raytracer/lib/scene-objects"
  "go-raytracer/lib"
)

var imageWidth, imageHeight int = 800, 640

func main() {

  camera := scene_objects.Camera{ scene_objects.Point3D{X: float64(imageWidth)/2, Y: float64(imageHeight)/2, Z: -100 } }

  scene := loadScene()

  fmt.Println(camera)

  for px := 0; px < imageWidth; px++{
    for py :=0; py < imageHeight; py++{
      //Create ray with origin ox, oy and direction from px, py to ox, oy
      primRay := computeRay(px, py, camera)
      //Scene.isIntersectedBy(primRay)
      curr_intersection := new(scene_objects.Intersection)
      curr_intersection.Distance = math.MaxFloat64

      for idx := 0; idx < len(scene); idx ++{
        new_intersection := intersects(primRay, scene[idx])
        //intersects
        fmt.Println(new_intersection.Distance, new_intersection.Distance < curr_intersection.Distance)
        if new_intersection.Distance > 0 && new_intersection.Distance < curr_intersection.Distance{
          curr_intersection.Distance = new_intersection.Distance
          curr_intersection.Object = new_intersection.Object
        }
      }
      if (curr_intersection.Distance > 0){
        fmt.Println("Ray intersects Object @", curr_intersection)
      }
    }
  }
}

func intersects(primRay scene_objects.Ray, obj scene_objects.Sphere) scene_objects.Intersection{
  v := primRay.Direction
  o := primRay.Origin
  c := obj.Center
  r := obj.SphereRay

  oc := o.Subtract(c)

  A := v.ScalarProd(v)
  B := 2* oc.ScalarProd(v)
  C := oc.ScalarProd(oc) - math.Pow(r,2)

  intersection := scene_objects.Intersection{Distance: -1}
  quadraticSolution := gomath.QuadraticEquationSolver(A, B, C)

  if quadraticSolution.NumSolutions == 0{
    return intersection
  } else if quadraticSolution.NumSolutions == 1 {
    hitP1 := primRay.PointInRay(quadraticSolution.S1)
    distance1:= primRay.Origin.ComputeDistance(hitP1)

    intersection.Distance = distance1
    intersection.HitPoint = hitP1
  } else{

    hitP1 := primRay.PointInRay(quadraticSolution.S1)
    distance1:= primRay.Origin.ComputeDistance(hitP1)

    hitP2 := primRay.PointInRay(quadraticSolution.S2)
    distance2:= primRay.Origin.ComputeDistance(hitP2)

    if distance1 > distance2{
      intersection.Distance = distance2
      intersection.HitPoint = hitP2
    } else{
      intersection.Distance = distance1
      intersection.HitPoint = hitP1
    }
  }

  return intersection
}

func loadScene() []scene_objects.Sphere{
  obj := scene_objects.Sphere{ Center: scene_objects.Point3D {X: 150, Y: 150, Z: 150}, SphereRay: 50 }
  return []scene_objects.Sphere{obj}
}

func computeRay(px, py int, cam scene_objects.Camera) scene_objects.Ray {
  rayDirX := float64(px) - cam.Position.X
  rayDirY := float64(py) - cam.Position.Y
  rayDirZ := -cam.Position.Z

  return scene_objects.Ray{ Origin: cam.Position, Direction: scene_objects.Point3D {rayDirX, rayDirY, rayDirZ} }
}

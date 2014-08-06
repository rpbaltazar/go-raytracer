package main

import (
  "fmt"
  "math"
  "go-raytracer/lib/scene-objects"
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

// TODO: Move equation solver to specific method (math lib maybe)
// solution[] = [sol1, sol2]

func intersects(primRay scene_objects.Ray, obj scene_objects.Sphere) scene_objects.Intersection{
  v := primRay.Direction
  o := primRay.Origin
  c := obj.Center
  r := obj.SphereRay

  oc := o.Subtract(c)

  A := v.ScalarProd(v)
  B := 2* oc.ScalarProd(v)
  C := oc.ScalarProd(oc) - math.Pow(r,2)

  delta := math.Pow(B, 2) - 4*A*C
  intersection := scene_objects.Intersection{ Distance: -1}
  //two solutions - calculate distance and return the closest
  if delta > 0{
    sol1 := (-B + delta)/(2*A)
    sol2 := (-B - delta)/(2*A)

    // hitP1 := computeHitPoint(sol1, primRay)
    hitP1 := primRay.Point(sol1)
    // hitP2 := computeHitPoint(sol2, primRay)
    hitP2 := primRay.Point(sol2)

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
    // hitP := computeHitPoint(sol, primRay)
    hitP := primRay.Point(sol)
    distance := primRay.Origin.ComputeDistance(hitP)

    intersection.HitPoint = hitP
    intersection.Distance = distance
    intersection.Object = obj

    return intersection
  }else{
    return intersection
  }
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

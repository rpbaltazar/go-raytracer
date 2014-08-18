package main

import (
  // "fmt"
  "math"
  "go-raytracer/lib/scene-objects"
  "go-raytracer/lib"
  "image"
  "image/color"
  "image/png"
  "os"
  "time"
)

const (
  ImageWidth = 800
  ImageHeight = 640
)

var (
  white color.Color = color.RGBA{255, 255, 255, 255}
  black color.Color = color.RGBA{0, 0, 0, 255}
)

func main() {

  camera := scene_objects.Camera{ scene_objects.Point3D{X: float64(ImageWidth)/2, Y: float64(ImageHeight)/2, Z: -100 } }
  light := scene_objects.Light { scene_objects.Point3D{X: 0, Y: 300, Z: 300} }
  render := image.NewRGBA(image.Rect(0,0,ImageWidth, ImageHeight))

  scene := loadScene()

  for px := 0; px < ImageWidth; px++{
    for py :=0; py < ImageHeight; py++{
      primRay := computeRay(px, py, camera)
      currIntersection := new(scene_objects.Intersection)
      currIntersection.Distance = -1

      for idx := 0; idx < len(scene); idx ++{
        newIntersection := intersects(primRay, scene[idx])
        if newIntersection.Distance > 0 && (newIntersection.Distance < currIntersection.Distance || currIntersection.Distance == -1){
          currIntersection.Distance = newIntersection.Distance
          currIntersection.Object = newIntersection.Object
          currIntersection.ObjectId = idx
        }
      }
      if (currIntersection.Distance > 0){
        shadowRay := scene_objects.Ray{Origin: currIntersection.HitPoint, Direction: light.Position.Subtract(currIntersection.HitPoint) }
        isInShadow := false

        for idx := 0; idx < len(scene); idx ++{
          if idx == currIntersection.ObjectId{
            continue
          }
          shadowIntersection:= intersects(shadowRay, scene[idx])
          if shadowIntersection.Distance > 0{
            isInShadow = true
            break
          }
        }
        //if hit point is in shadow, then pixel is black
        if isInShadow{
          render.Set(px, py, black)
        //else we set the color
        }else{
          //TODO: Determine the light brightness to determine the color.
          //Appel's algo?
          render.Set(px, py, currIntersection.Object.SphereColor)
        }
      }
    }
  }
  export(render)
}
//TODO: This method should be renamed to isIntersected and
// moved to the sphere file.
// All other objects should implement the same method with
// their own equations for intersection
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
    intersection.Object = obj

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
    intersection.Object = obj
  }

  return intersection
}

func loadScene() []scene_objects.Sphere{
  obj := scene_objects.Sphere{
    Center: scene_objects.Point3D {X: 150, Y: 150, Z: 150},
    SphereRay: 50,
    SphereColor: color.RGBA{0,0,255,255},
    RefractiveIndex: 1.5,
  }
  return []scene_objects.Sphere{obj}
}

func computeRay(px, py int, cam scene_objects.Camera) scene_objects.Ray {
  rayDirX := float64(px) - cam.Position.X
  rayDirY := float64(py) - cam.Position.Y
  rayDirZ := -cam.Position.Z

  normalizedDirection := scene_objects.Point3D {rayDirX, rayDirY, rayDirZ}
  normalizedDirection = normalizedDirection.Normalize()

  return scene_objects.Ray{ Origin: cam.Position, Direction: normalizedDirection }
}

func export(render image.Image) {
  w, _ := os.Create("render_"+time.Now().Format("20060102150405")+".png")
  defer w.Close()
  png.Encode(w, render) //Encode writes the Image render to w in PNG format.
}

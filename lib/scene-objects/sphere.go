package scene_objects

import "image/color"

type Sphere struct {
  Center Point3D
  SphereRay float64
  SphereColor color.RGBA

  //Properties of the object:
  //Should they be referenced in an Object Interface?
  RefractiveIndex float64
}

// func (s *Sphere) isIntersected(r Ray) Intersection {
//   return Intersection{ Object: s, Distance: 10.4}
// }

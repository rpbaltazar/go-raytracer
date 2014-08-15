package scene_objects

type Ray struct {
  Origin, Direction Point3D
}

func (r *Ray) PointInRay(t float64) Point3D {
  hX := r.Origin.X + t*r.Direction.X
  hY := r.Origin.Y + t*r.Direction.Y
  hZ := r.Origin.Z + t*r.Direction.Z

  hitP := Point3D { X: hX, Y: hY, Z: hZ }
  return hitP
}


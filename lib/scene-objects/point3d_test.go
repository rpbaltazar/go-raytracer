package scene_objects

import "testing"
import "math"

func Test_Point3D_Subtract_Point3D(t *testing.T){
  p1 := Point3D {X: 1, Y: 1, Z: 1}
  p2 := Point3D {X: 2, Y: 2, Z: 2}
  out:= Point3D {X: -1, Y: -1, Z: -1}

  if res:=p1.Subtract(p2); res != out{
    t.Errorf("(%v).Subtract(%v) => expected %v, got %v", p1, p2, out, res)
  }
}

func Test_Point3D_Add_Point3D(t *testing.T){
  p1 := Point3D {X: 1, Y: 1, Z: 1}
  p2 := Point3D {X: 2, Y: 2, Z: 2}
  out:= Point3D {X: 3, Y: 3, Z: 3}

  if res:=p1.Add(p2); res != out{
    t.Errorf("(%v).Add(%v) => expected %v, got %v", p1, p2, out, res)
  }
}

func Test_Point3D_Multiply_Float(t *testing.T){
  v1 := Point3D {X: 1, Y: 1, Z: 1}
  f := 2.5
  out:= Point3D {X: 2.5, Y: 2.5, Z: 2.5}

  if res:=v1.Multiply(f); res != out{
    t.Errorf("(%v).Multiply(%v) => expected %v, got %v", v1, f, out, res)
  }
}

func Test_Point3D_ScalarProd_Point3D(t *testing.T){
  v1 := Point3D {X: 1, Y: 1, Z: 1}
  v2 := Point3D {X: 5, Y: -3, Z: 4.5}
  out:= 6.5

  if res:=v1.ScalarProd(v2); res != out{
    t.Errorf("(%v).ScalarProd(%v) => expected %v, got %v", v1, v2, out, res)
  }
}

func Test_Point3D_Normalize(t *testing.T){
  v1 := Point3D {X: 1, Y: 1, Z: 1}
  norm := math.Sqrt(3)
  out:= Point3D {X: 1/norm, Y: 1/norm, Z: 1/norm}

  if res:=v1.Normalize(); res != out{
    t.Errorf("(%v).Normalize() => expected %v, got %v", v1, out, res)
  }

  v2 := Point3D {X: 5, Y: -3, Z: 4.5}
  norm = math.Sqrt(54.25)
  out= Point3D {X: 5/norm, Y: -3/norm, Z: 4.5/norm}

  if res:=v2.Normalize(); res != out{
    t.Errorf("(%v).Normalize() => expected %v, got %v", v2, out, res)
  }
}

func Test_Point3D_ComputeDistance_Point3D(t *testing.T){
  p1 := Point3D {X: 1, Y: 1, Z: 1}
  p2 := Point3D {X: 5, Y: -3, Z: 4.5}
  out:= math.Sqrt(44.25)

  if res:=p1.ComputeDistance(p2); res != out{
    t.Errorf("(%v).ComputeDistance(%v) => expected %v, got %v", p1, p2, out, res)
  }
}

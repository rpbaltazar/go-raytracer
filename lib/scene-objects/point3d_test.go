package scene_objects

import "testing"

func TestSubtract(t *testing.T){
  if testing.Short(){
    t.Skip("Skipping Test point subtract point in short suite")
  }

  p1 := Point3D {X: 1, Y: 1, Z: 1}
  p2 := Point3D {X: 2, Y: 2, Z: 2}
  out:= Point3D {X: -1, Y: -1, Z: -1}

  if res:=p1.Subtract(p2); res != out{
    t.Errorf("(%v).Subtract(%v) => expected %v, got %v", p1, p2, out, res)
  }

}

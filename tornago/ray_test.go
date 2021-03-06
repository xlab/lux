package tornago

import (
	"github.com/luxengine/lux/flops"
	"github.com/luxengine/lux/glm"
	"testing"
)

func TestRay_FromTo(t *testing.T) {
	tests := []struct {
		From, To glm.Vec3
		// expected values
		Origin, Direction, Destination glm.Vec3
		Len                            float32
	}{
		/*{
			From:      glm.Vec3{},
			To:        glm.Vec3{},
			Origin:    glm.Vec3{},
			Direction: glm.Vec3{},
			Destination: glm.Vec3{},
			Len:       0,
		},*/
		{
			From:        glm.Vec3{X: 0, Y: 0, Z: 0},
			To:          glm.Vec3{X: 1, Y: 0, Z: 0},
			Origin:      glm.Vec3{X: 0, Y: 0, Z: 0},
			Direction:   glm.Vec3{X: 1, Y: 0, Z: 0},
			Destination: glm.Vec3{X: 1, Y: 0, Z: 0},
			Len:         1,
		},
		{
			From:        glm.Vec3{X: 1, Y: 1, Z: 1},
			To:          glm.Vec3{X: 1, Y: 1, Z: 1},
			Origin:      glm.Vec3{X: 1, Y: 1, Z: 1},
			Direction:   glm.Vec3{X: 0, Y: 1, Z: 0},
			Destination: glm.Vec3{X: 1, Y: 1, Z: 1},
			Len:         0,
		},
		{
			From:        glm.Vec3{X: 1, Y: 1, Z: 1},
			To:          glm.Vec3{X: 0, Y: 0, Z: 0},
			Origin:      glm.Vec3{X: 1, Y: 1, Z: 1},
			Direction:   glm.Vec3{X: -0.57735026, Y: -0.57735026, Z: -0.57735026},
			Destination: glm.Vec3{X: 0, Y: 0, Z: 0},
			Len:         1.7320508,
		},
		{
			From:        glm.Vec3{X: 1, Y: 2, Z: 3},
			To:          glm.Vec3{X: 9, Y: 8, Z: 7},
			Origin:      glm.Vec3{X: 1, Y: 2, Z: 3},
			Direction:   glm.Vec3{X: 0.74278134, Y: 0.557086, Z: 0.37139067},
			Destination: glm.Vec3{X: 9, Y: 8, Z: 7},
			Len:         10.770329,
		},
	}

	for i, test := range tests {
		ray := NewRayFromTo(test.From, test.To)

		// test origin
		if o := ray.Origin(); !o.EqualThreshold(&test.Origin, 1e-3) {
			t.Errorf("[%d] ray.Origin() = %v, want %v", i, o, test.Origin)
		}

		// test direction
		if d := ray.Direction(); !d.EqualThreshold(&test.Direction, 1e-3) {
			t.Errorf("[%d] ray.Direction() = %v, want %v", i, d, test.Direction)
		}

		// test Len
		if l := ray.Len(); !flops.Eq(l, test.Len) {
			t.Errorf("[%d] ray.Len() = %v, want %v", i, l, test.Len)
		}

		// test destination
		if d := ray.Destination(); !d.EqualThreshold(&test.Destination, 1e-3) {
			t.Errorf("[%d] ray.Destination() = %v, want %v", i, d, test.Destination)
		}
	}

}

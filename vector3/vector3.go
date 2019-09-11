package vec3

import (
	"math"
)

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func Zero() *Vector3 {
	return &Vector3{0.0, 0.0, 0.0}
}

func (v *Vector3) Clone() *Vector3 {
	return &Vector3{v.X, v.Y, v.Z}
}

func (v *Vector3) Assign(v2 *Vector3) *Vector3 {
	v.X = v2.X
	v.Y = v2.Y
	v.Z = v2.Z
	return v
}

func (v *Vector3) Equals(v2 *Vector3) bool {
	return v.X == v2.X && v.Y == v2.Y && v.Z == v2.Z
}

func (v *Vector3) Add(v2 *Vector3) *Vector3 {
	v.X += v2.X
	v.Y += v2.Y
	v.Z += v2.Z
	return v
}

func (v *Vector3) AddScalar(s float64) *Vector3 {
	v.X += s
	v.Y += s
	v.Z += s
	return v
}

func (v *Vector3) Sub(v2 *Vector3) *Vector3 {
	v.X -= v2.X
	v.Y -= v2.Y
	v.Z -= v2.Z
	return v
}

func (v *Vector3) SubScalar(s float64) *Vector3 {
	v.X -= s
	v.Y -= s
	v.Z -= s
	return v
}

func (v *Vector3) MulScalar(s float64) *Vector3 {
	v.X *= s
	v.Y *= s
	v.Z *= s
	return v
}

func (v *Vector3) Dot(v2 *Vector3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v *Vector3) Cross(v2 *Vector3) *Vector3 {
	return &Vector3{v.Y*v2.Z - v.Z*v2.Y, v.Z*v2.X - v.X*v2.Z, v.X*v2.Y - v.Y*v2.X}
}

func (v *Vector3) Length2() float64 {
	return v.Dot(v)
}

func (v *Vector3) Length() float64 {
	return math.Sqrt(v.Length2())
}

func (v *Vector3) Normalize() *Vector3 {
	ilen := 1.0 / v.Length()
	return v.MulScalar(ilen)
}

func (v *Vector3) Abs() *Vector3 {
	return &Vector3{math.Abs(v.X), math.Abs(v.Y), math.Abs(v.Z)}
}

func (v *Vector3) Lerp(v2 *Vector3, t float64) *Vector3 {
	v.X = (1.0-t)*v.X + t*v2.X
	v.Y = (1.0-t)*v.Y + t*v2.Y
	v.Z = (1.0-t)*v.Z + t*v2.Z
	return v
}

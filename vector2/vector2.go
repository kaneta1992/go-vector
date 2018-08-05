package vec

import (
	"fmt"
	"math"
	"strconv"
)

type Vector2 struct {
	X float64
	Y float64
}

func (v *Vector2) Add(v2 *Vector2) *Vector2 {
	return &Vector2{v.X + v2.X, v.Y + v2.Y}
}

func (v *Vector2) AddScalar(s float64) *Vector2 {
	return &Vector2{v.X + s, v.Y + s}
}

func (v *Vector2) Sub(v2 *Vector2) *Vector2 {
	return &Vector2{v.X - v2.X, v.Y - v2.Y}
}

func (v *Vector2) SubScalar(s float64) *Vector2 {
	return &Vector2{v.X - s, v.Y - s}
}

func (v *Vector2) MulScalar(s float64) *Vector2 {
	return &Vector2{v.X * s, v.Y * s}
}

func (v *Vector2) Dot(v2 *Vector2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v *Vector2) Cross(v2 *Vector2) float64 {
	return v.X*v2.Y - v.Y*v2.X
}

func (v *Vector2) Length2() float64 {
	return v.Dot(v)
}

func (v *Vector2) Length() float64 {
	return math.Sqrt(v.Length2())
}

func (v *Vector2) Normalize() *Vector2 {
	ilen := 1.0 / v.Length()
	return v.MulScalar(ilen)
}

func (v *Vector2) Abs() *Vector2 {
	return &Vector2{math.Abs(v.X), math.Abs(v.Y)}
}

func (v *Vector2) ToFloat64() [2]float64 {
	return [2]float64{v.X, v.Y}
}

func Vec2ToFloat64(points []*Vector2) [][2]float64 {
	ret := [][2]float64{}
	for _, v := range points {
		ret = append(ret, v.ToFloat64())
	}
	return ret
}

func Float64ToVec2(points [][2]float64) []*Vector2 {
	ret := []*Vector2{}
	for _, v := range points {
		ret = append(ret, &Vector2{v[0], v[1]})
	}
	return ret
}

func (v *Vector2) ToGLSLString(place int) string {
	str := strconv.Itoa(place)
	return fmt.Sprintf("vec2(%."+str+"f,%."+str+"f)", v.X, v.Y)
}

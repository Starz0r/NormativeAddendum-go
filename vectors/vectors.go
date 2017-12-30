package vectors

import (
	"fmt"
	"reflect"
)

type Vector struct {
	slice  reflect.Value
	typeof reflect.Type
}

func NewVector(z interface{}) *Vector {
	t := reflect.TypeOf(z)

	return &Vector{
		slice:  reflect.MakeSlice(reflect.SliceOf(t), 0, 0),
		typeof: t,
	}
}

func newVector(t reflect.Type, len, cap int) *Vector {
	return &Vector{
		slice:  reflect.MakeSlice(reflect.SliceOf(t), len, cap),
		typeof: t,
	}
}

func (v *Vector) Get(index int) interface{} {
	return v.slice.Index(index)
}

func (v *Vector) Put(element interface{}) {

	if reflect.ValueOf(element).Type() != v.slice.Type().Elem() {
		panic(fmt.Sprintf("Put: cannot put a %T into a slice of %s", element, v.slice.Type().Elem()))
	}

	v.slice = reflect.Append(v.slice, reflect.ValueOf(element))
}

func (v *Vector) Copy() *Vector {
	v2 := newVector(v.typeof, v.slice.Len(), v.slice.Cap())
	reflect.Copy(v2.slice, v.slice)
	return v2
}

func (v *Vector) Cut(i, j int) {
	lastItem := v.slice.Len()
	cutLen := j - i

	reflect.Copy(v.slice.Slice(i, lastItem), v.slice.Slice(j, lastItem))

	for n := v.slice.Len() - cutLen; n < v.slice.Len(); n++ {
		v.slice.Index(n).Set(reflect.Zero(v.typeof))
	}

	v.slice = v.slice.Slice(0, v.slice.Len()-cutLen)
}

func (v *Vector) Delete(i int) {
	reflect.Copy(v.slice.Slice(i, v.slice.Len()), v.slice.Slice(i+1, v.slice.Len()))
	v.slice.Index(v.slice.Len() - 1).Set(reflect.Zero(v.typeof))
	v.slice = v.slice.Slice(0, v.slice.Len()-1)
}

func (v *Vector) DeleteNoPreserveOrder(i int) {
	v.slice.Index(i).Set(v.slice.Index(v.slice.Len() - 1))
	v.slice.Index(v.slice.Len() - 1).Set(reflect.Zero(v.typeof))
	v.slice = v.slice.Slice(0, v.slice.Len()-1)
}
package huisebug

import (
	"fmt"
	"strings"
)

func Print(str string) {
	fmt.Println(str)
}

type Named interface {
	Nm() string
}

// 定义一个类型别名，T 必须是可比较的类型
type ComparableSlice[T comparable] []T

// 查找给定元素 c 在切片中的索引
func (CS ComparableSlice[T]) IndexOf(c T) int {
	for i, v := range CS {
		if v == c {
			return i
		}
	}
	return -1
}

// 定义一个类型别名，T 是任意的类型
type AnySlice[T any] []T

func (ga AnySlice[T]) Find(f func(T) bool) (*T, int) {
	for i, v := range ga {
		if f(v) {
			return &v, i
		}
	}

	return nil, -1
}

func (ga AnySlice[T]) Filter(f func(T) bool) AnySlice[T] {
	r := make([]T, 0)
	for _, v := range ga {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}

func SplitMap[IN any, OUT any](input AnySlice[IN], f func(in IN) (bool, OUT)) (match []OUT, mismatch []OUT) {
	for _, i := range input {
		isMatch, out := f(i)
		if isMatch {
			match = append(match, out)
		} else {
			mismatch = append(mismatch, out)
		}
	}

	return
}

func Map[IN any, OUT any](input AnySlice[IN], f func(in IN) OUT) []OUT {
	var o []OUT = nil
	for _, i := range input {
		o = append(o, f(i))
	}
	return o
}

func Reduce[T any, OUT any](slice []T, initValue OUT, f func(idx int, val OUT, in T) OUT) OUT {
	for i, v := range slice {
		initValue = f(i, initValue, v)
	}
	return initValue
}

func (input AnySlice[T]) SplitTrueFalse(f func(in T) bool) ([]T, []T) {
	r1 := make([]T, 0)
	r2 := make([]T, 0)
	for _, v := range input {
		if f(v) {
			r1 = append(r1, v)
		} else {
			r2 = append(r2, v)
		}
	}
	return r1, r2
}

func FlatMap[IN any, OUT any](input AnySlice[IN], f func(in IN) []OUT) AnySlice[OUT] {
	var o []OUT = nil
	for _, i := range input {
		o = append(o, f(i)...)
	}
	return o
}

type NamedArray[T Named] []T

func (na NamedArray[T]) Names() []string {
	ns := []string{}
	for _, i := range na {
		ns = append(ns, i.Nm())
	}
	return ns
}

func (na NamedArray[T]) Find(name string) *T {
	i := na.IndexByName(name)
	if i < 0 {
		return nil
	}

	return &na[i]
}

func (na NamedArray[T]) IndexOf(item T) int {
	return na.IndexByName(item.Nm())
}

func (na NamedArray[T]) IndexByName(name string) int {
	for i, n := range na {
		if n.Nm() == name {
			return i
		}
	}

	return -1
}

func (na NamedArray[T]) Len() int {
	return len(na)
}

func (na NamedArray[T]) Less(i, j int) bool {
	return strings.Compare(na[i].Nm(), na[j].Nm()) < 0
}

// Swap swaps the elements with indexes i and j.
func (na NamedArray[T]) Swap(i, j int) {
	na[i], na[j] = na[j], na[i]
}

type GSet[K comparable] map[K]struct{}

func (gm GSet[K]) Add(ks ...K) {
	for _, key := range ks {
		gm[key] = struct{}{}
	}
}

func (g1 GSet[K]) Sub(g2 GSet[K]) {
	for key := range g1 {
		if _, has := g2[key]; has {
			delete(g1, key)
		}
	}
}

func NewGSet[K comparable](v []K) GSet[K] {
	s := GSet[K]{}
	s.Add(v...)
	return s
}

func (gm GSet[K]) Values() []K {
	keys := []K{}
	for key := range gm {
		keys = append(keys, key)
	}

	return keys
}

func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

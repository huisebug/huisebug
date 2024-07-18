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
func (CS ComparableSlice[T]) GetIndex(c T) int {
	for i, v := range CS {
		if v == c {
			return i
		}
	}
	return -1
}

// 定义一个类型别名，T 是任意的类型
type AnySlice[T any] []T

// 返回第一个找到的元素及其索引
func (AS AnySlice[T]) Find(f func(T) bool) (*T, int) {
	for i, v := range AS {
		if f(v) {
			// 使用临时变量保存匹配的元素，然后返回其指针和索引
			result := v
			return &result, i
		}
	}

	// 如果找不到匹配的元素，返回 nil 和 -1
	return nil, -1
}

// Filter 方法用于过滤符合条件的元素
// 例如:
// var slice AnySlice[int] = []int{1, 2, 3, 4, 5, 6}
//
//	filtered := slice.Filter(func(v int) bool {
//		return v%2 == 0
//	})
func (AS AnySlice[T]) Filter(f func(T) bool) AnySlice[T] {
	var r AnySlice[T] // 定义一个符合类型别名的切片
	for _, v := range AS {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}

// 方法用于将输入切片中的元素按条件判断后将元素分别放入两个切片
// eligible:符合条件
func (AS AnySlice[T]) SplitSlice(f func(T) bool) (eligibleOuts []T, notEligibleOuts []T) {

	for _, v := range AS {
		if f(v) {
			eligibleOuts = append(eligibleOuts, v)
		} else {
			notEligibleOuts = append(notEligibleOuts, v)
		}
	}
	return
}

// 函数用于将输入切片中的元素按条件处理后将返回值分别放入两个切片
// eligible:符合条件
func SplitSliceOut[IN any, OUT any](AS AnySlice[IN], f func(IN) (bool, OUT)) (eligibleOuts []OUT, notEligibleOuts []OUT) {
	for _, in := range AS {
		eligible, out := f(in)
		if eligible {
			eligibleOuts = append(eligibleOuts, out)
		} else {
			notEligibleOuts = append(notEligibleOuts, out)
		}
	}

	return
}

// 函数用于将输入切片中的元素按条件换算后返回值放到输出切片
// condition:条件
// conversion:换算
func ConditionConversion[IN any, OUT any](AS AnySlice[IN], f func(IN) OUT) []OUT {
	var o []OUT
	for _, i := range AS {
		o = append(o, f(i))
	}
	return o
}

// 函数用于将输入切片中的元素按idx保留后和initValue初始值进行换算后得到返回值
// Retain:保留
// calculation:换算

func RetainCalculation[T any, OUT any](slice []T, initValue OUT, f func(idx int, val OUT, in T) OUT) OUT {
	for i, v := range slice {
		initValue = f(i, initValue, v)
	}
	return initValue
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

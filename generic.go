package huisebug

import (
	"fmt"
	"strings"
)

func Print(str string) {
	fmt.Println(str)
}

// 定义一个类型别名，T 必须是可比较的类型
type ComparableSlice[T comparable] []T

// 查找给定元素 c 在切片中的索引
func (CS ComparableSlice[T]) IndexBy(c T) int {
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

// 函数用于将输入切片中的元素按条件换算后再打散入输出切片
// expand:展开
// calculation:换算
func ExpandcalCulation[IN any, OUT any](AS AnySlice[IN], f func(in IN) []OUT) AnySlice[OUT] {
	var o AnySlice[OUT]
	for _, i := range AS {
		o = append(o, f(i)...)
	}
	return o
}

// 专用于有Name成员的interface
type Namer interface {
	Name() string
}

type NamerSlice[T Namer] []T

// 打印所有的名字
func (NS NamerSlice[T]) Names() []string {
	ns := []string{}
	for _, n := range NS {
		ns = append(ns, n.Name())
	}
	return ns
}

// 获取名字的索引
func (NS NamerSlice[T]) IndexByName(name string) int {
	for i, n := range NS {
		if n.Name() == name {
			return i
		}
	}

	return -1
}

// 使用 IndexByName 方法查找具有指定名称的元素的索引。如果没有找到，该方法返回一个负数。
// 返回指针：如果找到，则返回指向该元素的指针。
func (NS NamerSlice[T]) Find(name string) *T {
	i := NS.IndexByName(name)
	if i < 0 {
		return nil
	}

	return &NS[i]
}

// 调用 item.Name() 获取 item 的名称。这个方法假设 T 类型有一个 Name 方法。
// 查找索引：使用 NS.IndexByName 方法查找该名称在 NamerSlice 切片中的索引。
func (NS NamerSlice[T]) IndexOf(item T) int {
	return NS.IndexByName(item.Name())
}

// 返回一个布尔值，表示 NS[i] 是否应该排在 NS[j] 之前。具体来说，Less 方法返回 true 表示 NS[i] 排在 NS[j] 之前，false 表示 NS[j] 排在 NS[i] 之前。
// 使用 strings.Compare 函数比较两个名称的字典序。strings.Compare(a, b) 的返回值是：
// 小于 0 如果 a 小于 b，
// 等于 0 如果 a 等于 b，
// 大于 0 如果 a 大于 b。
func (NS NamerSlice[T]) Less(i, j int) bool {
	return strings.Compare(NS[i].Name(), NS[j].Name()) < 0
}

func (NS NamerSlice[T]) Len() int {
	return len(NS)
}

// NamerSlice 类型实现了 sort.Interface 接口的三个方法：Len、Less 和 Swap。
// Less 方法定义了排序的比较规则（按名称的字典序升序排序）。
// 通过 sort.Sort 函数对 people 切片进行排序，最终输出按名称升序排序的人员列表。
func (NS NamerSlice[T]) Swap(i, j int) {
	NS[i], NS[j] = NS[j], NS[i]
}

// map 的键是 T 类型，值是 struct{}。使用空结构体作为值可以节省内存，因为空结构体不占用额外空间，仅用于表示键的存在性
// 一个类型为可比较类型的集合
type ComparableSet[T comparable] map[T]struct{}

// NewComparableSet 创建一个新的 ComparableSet 实例
func NewComparableSet[T comparable]() ComparableSet[T] {
	return make(ComparableSet[T])
}

// Add 向集合中添加元素
func (CS ComparableSet[T]) Add(cs ...T) {
	for _, key := range cs {
		CS[key] = struct{}{}
	}
}

// ComparableSlice 转换为 一个新的 ComparableSet 实例
func ComparableSliceConvertComparableSet[T comparable](v []T) ComparableSet[T] {
	s := ComparableSet[T]{}
	s.Add(v...)
	return s
}

// Remove 从集合中删除元素
func (CS ComparableSet[T]) Remove(item T) {
	delete(CS, item)
}

// Contains 检查集合中是否存在元素
func (CS ComparableSet[T]) Contains(item T) bool {
	_, exists := CS[item]
	return exists
}

func (g1 ComparableSet[T]) Sub(g2 ComparableSet[T]) {
	for key := range g1 {
		if _, has := g2[key]; has {
			delete(g1, key)
		}
	}
}

func (CS ComparableSet[T]) Values() []T {
	keys := []T{}
	for key := range CS {
		keys = append(keys, key)
	}

	return keys
}

func Keys[M ~map[T]V, T comparable, V any](m M) []T {
	r := make([]T, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

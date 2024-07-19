package main

import (
	"fmt"
	"strings"

	"github.com/huisebug/huisebug"
)

func main() {
	{
		fmt.Println(huisebug.StringCompare("huisebug")("huisebug"))
		fmt.Println(huisebug.StringCompare("huisebug")("huisebug12"))
	}
	{
		var slice huisebug.AnySlice[int] = []int{1, 2, 3, 4, 5}

		// 查找偶数
		result, index := slice.Find(func(v int) bool {
			return v%2 == 0
		})

		if result != nil {
			fmt.Printf("Found %d at index %d\n", *result, index)
		} else {
			fmt.Println("Element not found")
		}
	}

	{
		// 示例：过滤切片中的偶数
		var slice huisebug.AnySlice[int] = []int{1, 2, 3, 4, 5, 6}

		filtered := slice.Filter(func(v int) bool {
			return v%2 == 0
		})

		fmt.Println(filtered) // 输出：[2 4 6]
	}

	{
		// 示例：将整数切片中的偶数和奇数分组
		var slice huisebug.AnySlice[int] = []int{1, 2, 3, 4, 5}

		// 定义一个条件函数，判断是否为偶数
		isEven := func(n int) bool {
			return n%2 == 0
		}

		// 使用 SplitSlice 方法将切片分成两组
		evens, odds := slice.SplitSlice(isEven)

		fmt.Println("Evens:", evens) // 输出：[2 4]
		fmt.Println("Odds:", odds)   // 输出：[1 3 5]
	}
	{
		// 示例：将整数切片中的偶数放入 eligibleOuts，奇数放入 notEligibleOuts
		var slice huisebug.AnySlice[int] = []int{1, 2, 3, 4, 5}

		isEven := func(n int) (bool, int) {
			return n%2 == 0, n
		}

		eligibleOuts, notEligibleOuts := huisebug.SplitSliceOut(slice, isEven)

		fmt.Println("Eligible elements:", eligibleOuts)        // 输出：[2 4]
		fmt.Println("Not eligible elements:", notEligibleOuts) // 输出：[1 3 5]
	}
	{
		// 示例：将整数切片映射到其平方
		ints := huisebug.AnySlice[int]{1, 2, 3, 4, 5}
		squares := huisebug.ConditionConversion(ints, func(i int) int { return i * i })
		fmt.Println(squares) // 输出: [1, 4, 9, 16, 25]

		// 示例：将字符串切片映射到其长度
		strs := huisebug.AnySlice[string]{"apple", "banana", "cherry"}
		lengths := huisebug.ConditionConversion(strs, func(s string) int { return len(s) })
		fmt.Println(lengths) // 输出: [5, 6, 6]
	}
	{
		// 示例：计算整数切片的前4个元素和
		ints := []int{1, 2, 3, 4, 5}
		sum := huisebug.RetainCalculation(ints, 1, func(idx int, val, in int) int {
			if idx < 4 {
				return val + in
			}
			return val
		})
		fmt.Println("Sum:", sum) // 输出: Sum: 11

		// 示例：换算字符串切片的前3个连接
		strs := []string{"apple", "banana", "cherry", "boy"}
		concatenated := huisebug.RetainCalculation(strs, "", func(idx int, val, in string) string {
			if 0 < idx && idx < 3 {
				return val + ", " + in
			}
			return val + in
		})
		fmt.Println("Concatenated:", concatenated) // 输出: Concatenated: apple, banana, cherry

	}
	{
		// 示例：将整数切片中的每个元素映射到其平方和立方
		ints := huisebug.AnySlice[int]{1, 2, 3}
		results := huisebug.ExpandcalCulation(ints, func(i int) []int {
			return []int{i * i, i * i * i}
		})
		fmt.Println(results) // 输出: [1 1 4 8 9 27]

		// 示例：将字符串切片中的每个字符串映射到其字符切片
		strs := huisebug.AnySlice[string]{"apple", "banana"}
		characters := huisebug.ExpandcalCulation(strs, func(s string) []string {
			return strings.Split(s, "")
		})
		fmt.Println(characters) // 输出: [a p p l e b a n a n a]
	}

	{
		// 创建一个 Person 实例切片
		persons := huisebug.NamerSlice[Person]{
			Person{firstName: "John", lastName: "Doe"},
			Person{firstName: "Jane", lastName: "Smith"},
		}
		// 打印所有人的名字
		fmt.Println(persons.Names()) // 输出: [John.Doe Jane.Smith]
	}
	{
		// 创建一个 ComparableSet 实例
		intSet := huisebug.NewComparableSet[int]()

		// 添加元素
		intSet.Add(1)
		intSet.Add(2)
		intSet.Add(3)

		// 检查元素
		fmt.Println(intSet.Contains(2)) // Output: true
		fmt.Println(intSet.Contains(4)) // Output: false

		// 删除元素
		intSet.Remove(2)
		fmt.Println(intSet.Contains(2)) // Output: false

		// 创建另一个 ComparableSet 实例，处理字符串类型
		stringSet := huisebug.NewComparableSet[string]()
		stringSet.Add("hello")
		stringSet.Add("world")

		// 检查字符串元素
		fmt.Println(stringSet.Contains("hello")) // Output: true
		fmt.Println(stringSet.Contains("go"))    // Output: false
	}
	{
		set1 := huisebug.NewComparableSet[int]()
		set1.Add(1)
		set1.Add(2)
		set1.Add(3)

		set2 := huisebug.NewComparableSet[int]()
		set2.Add(2)
		set2.Add(4)

		// 执行差集操作
		set1.RemoveAll(set2)

		// 打印结果
		for item := range set1 {
			fmt.Println(item) // Output: 1
		}
	}
	{
		// 示例 1：使用 int 键的 map
		intMap := map[int]string{
			1: "one",
			2: "two",
			3: "three",
		}
		intKeys := huisebug.ComparableSetConvertComparableSlice(intMap)
		fmt.Println("Keys of intMap:", intKeys) // Output: Keys of intMap: [1 2 3]

		// 示例 2：使用 string 键的 map
		stringMap := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}
		stringKeys := huisebug.ComparableSetConvertComparableSlice(stringMap)
		fmt.Println("Keys of stringMap:", stringKeys) // Output: Keys of stringMap: [a b c]
	}
}

// Person 实现了 Namer 接口
type Person struct {
	firstName string
	lastName  string
}

func (p Person) Name() string {
	return p.firstName + "." + p.lastName
}

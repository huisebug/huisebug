package huisebug

// 目的是返回一个闭包函数，用来检查传入的字符串是否等于Fixedvalue。这种方法通常用于创建定制的字符串比较函数
// Fixedvalue:固定值
// Transmittingvalue:传递值
// StringCompare("huisebug")("huisebug12")
func StringCompare(f string) func(string) bool {
	return func(t string) bool {
		return f == t
	}
}

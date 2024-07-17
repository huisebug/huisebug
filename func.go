package huisebug

// 目的是返回一个闭包函数，用来检查传入的字符串是否等于 v。这种方法通常用于创建定制的字符串比较函数
// Fixed value
// Transmitting values
// StringCompare("wyf")("wyf1")
func StringCompare(f string) func(string) bool {
	return func(t string) bool {
		return f == t
	}
}

package If

// 返回两个数中较大的那一个
func Judge(First, Second int64) interface{} {
	return If(First >= Second, First, Second)
}

// 输入一个bool类型，true 返回第一个， false 返回第二个
func If(condition bool, First, Second interface{}) interface{} {
	if condition {
		return First
	} else {
		return Second
	}
}

// 求和
func Sum(Sum []int64) (sum int64) {
	for _, v := range Sum {
		sum += v
	}
	return sum
}

// 二分查找
func BinaryCount(Num int64, slice []int64) (int64, int64) {
	var count int64 = 0
	var lens = len(slice)
	for i := 0; i < lens; i++ {
		var length = len(slice)
		if Num < slice[length/2] {
			slice = slice[0 : length/2]
			count += 1
		} else if Num > slice[length/2] {
			slice = slice[length/2 : length]
			count += 1
		} else {
			num := slice[length/2]
			count += 1
			return num, count
		}
	}
	return 0, 0
}

// 找出最大值
func ArryBiggest(Arry []int64) int64 {
	for i := 0; i < len(Arry)-1; i++ {
		Decide(&Arry[i], &Arry[i+1])
	}
	return Last(Arry)
}

// 前一个比后一个大，换位置
func Decide(First, Second *int64) {
	if *First > *Second {
		*First, *Second = *Second, *First
	}
}

// 取最后一个元素
func Last(arry []int64) int64 {
	return arry[len(arry)-1]
}

// 找出最小值
func MiniVal(Arry []int64) int64 {
	for i := 0; i < len(Arry)-1; i++ {
		DecideLeast(&Arry[i], &Arry[i+1])
	}
	return Last(Arry)
}

// 前一个数比后一个数小，换位置
func DecideLeast(First, Second *int64) {
	if *First < *Second {
		*First, *Second = *Second, *First
	}
}

// 从小到大排序
func SortAsc(Arry []int64) []int64 {
	for i := 0; i < len(Arry); i++ {
		for j := len(Arry) - 1; j > i; j-- {
			Decide(&Arry[i], &Arry[j])
		}
	}
	return Arry
}

// 从大到小排列
func SortDesc(Arry []int64) []int64 {
	for i := 0; i < len(Arry); i++ {
		for j := len(Arry) - 1; j > i; j-- {
			DecideLeast(&Arry[i], &Arry[j])
		}
	}
	return Arry
}

// 求平均数
func Average(Arry []int64) float64 {
	sum := Sum(Arry)
	return float64(sum)/float64(len(Arry))
}

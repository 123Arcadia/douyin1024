package utils

// RemoveRepeatArray 数组去重, 注意适用于 []uint
func RemoveRepeatArray(arr []uint) []uint {
	result := []uint{}
	tempMap := map[uint]byte{} // 存放不重复主键
	for _, e := range arr {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result

}

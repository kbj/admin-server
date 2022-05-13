package utils

import "strconv"

// StringArray2intArray 将字符串数组转为uint数组
func StringArray2intArray(source *[]string) (*[]uint, error) {
	res := make([]uint, len(*source))
	for i, s := range *source {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		res[i] = uint(num)
	}
	return &res, nil
}

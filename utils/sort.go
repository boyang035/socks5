package utils

import (
	"strconv"
	"strings"
)

type StringDIYSort []string
func (x StringDIYSort) Len() int { return len(x) }
func (x StringDIYSort) Less(i, j int) bool {
	v1,err := strconv.ParseFloat(strings.Split(x[i]," ")[0], 32)
	if err != nil {
		return true
	}
	v2,err := strconv.ParseFloat(strings.Split(x[j]," ")[0], 32)
	if err != nil {
		return true
	}
	return v1 > v2
}
func (x StringDIYSort) Swap(i, j int) { x[i], x[j] = x[j], x[i] }


type StringToIntSort []string
func (x StringToIntSort) Len() int { return len(x) }
func (x StringToIntSort) Less(i, j int) bool {
	v1,err := strconv.Atoi(x[i])
	if err != nil {
		return true
	}
	v2,err := strconv.Atoi(x[j])
	if err != nil {
		return true
	}
	return v1 < v2
}
func (x StringToIntSort) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

//func StringToIntSortTest (){
//	arr := StringToIntSort{"1","5","11","3","7","4"}
//	sort.Sort(arr)
//	fmt.Println(arr)
//}
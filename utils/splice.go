package utils

/*
removes elements from start(including) to end (not including)
replacing deleted elements with elem
*/
func Splice(arr []string, start int, end int, elem string) []string {
	f1 := make([]string, 0)
	f2 := append(f1, arr[:start]...)
	f3 := append(f2, elem)
	f4 := append(f3, arr[end+1:]...)
	return f4
}

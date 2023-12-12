package array

import "fmt"

func Run() {
	var grades [3]int = [3]int{2, 3, 4}
	fmt.Println(grades)
	grd := [3]int{4, 5, 6}
	fmt.Println(grd)
	nme := []string{"hai", "halo"}
	fmt.Println(nme)
	stt := [...]int{3, 7, 8}
	fmt.Println(stt)
	fmt.Println("---------------")
	var gd [5]int = [5]int{4, 5, 6, 7, 7}
	for index, element := range gd {
		fmt.Println(index, "=>", element)
	}
	fmt.Println("---------------")
	arr := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(arr[1][0])
}

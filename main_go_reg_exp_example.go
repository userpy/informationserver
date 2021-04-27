package main
1

//func array_intersect(a, b []interface{}) []interface{} {
//	return set.NewSetFromSlice(a).Intersect(set.NewSetFromSlice(b)).ToSlice()
//}
func main() {
	//a := []interface{}{1, 2, 3, 4, 5, 7,10}
	//b := []interface{}{10}
	//if len(set.NewSetFromSlice(a).Intersect(set.NewSetFromSlice(b)).ToSlice()) != 0 {
	//	fmt.Printf("Пересекаются\n")
	//}
	var Sevenf [7]int

	five := [5]int{1,2,3,4,5}
	two := [2]int{6,7}

	//this doesn't work as both the inputs and assignment are the wrong type
	seven = append(five,two)

	//this doesn't work as the assignment is still the wrong type
	seven = append(five[:],two[:])

	//this works but I'm not using arrays anymore so may as well use slices everywhere and forget sizing
	seven2 := append(five[:],two[:])
}

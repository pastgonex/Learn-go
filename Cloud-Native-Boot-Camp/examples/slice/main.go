package main

func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	//myArray := [5]int{1, 2, 3, 4, 5}
	//mySlice := myArray[:]
	var mySlice1 []int
	mySlice1 = append(mySlice1, 1)
	mySlice1 = append(mySlice1, 2)
	mySlice1 = append(mySlice1, 8)
	mySlice1 = append(mySlice1, 3)
	mySlice1 = append(mySlice1, 4)
}

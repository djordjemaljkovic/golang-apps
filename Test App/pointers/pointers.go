package main

import "fmt"

func main() {
	age := 32

	ageAddress := &age

	fmt.Println("Age: ", *ageAddress)
	getAdultYears(ageAddress)
	fmt.Println(age)
}

//every execution - new copy of number 32 that is deleted after func finishes
func getAdultYears(ageAddress *int) {
	// return *ageAddress - 18
	*ageAddress = *ageAddress - 18
}

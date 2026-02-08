package main

import (
	"fmt"
)

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	hobbies := [3]string{"Sports", "Cooking", "Reading"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	selectedHobbies := hobbies[0:2]
	fmt.Println(selectedHobbies)

	selectedHobbies = selectedHobbies[1:3]
	fmt.Println(selectedHobbies)

	courseGoals := []string{"Go", ".NET"}
	fmt.Println(courseGoals)

	courseGoals[1] = "React"
	courseGoals = append(courseGoals, "AWS")
	fmt.Println(courseGoals)

	products := []Product{
		{
			"Bread", 1, 2.99,
		},
		{
			"Butter", 2, 4.99,
		},
	}
	fmt.Println(products)

	newProduct := Product{
		"Water", 3, 0.99,
	}
	products = append(products, newProduct)
	fmt.Println(products)
}

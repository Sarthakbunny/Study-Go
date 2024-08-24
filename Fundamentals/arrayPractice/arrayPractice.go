package arrayPractice

import "fmt"

type Product struct {
	title string
	id    int64
	price float64
}

func ArrayPractice() {
	task1 := [3]string{"Hobby1", "Hobby2", "Hobby3"}
	fmt.Println(task1)

	task2_1 := task1[:1]
	task2_2 := task1[1:]
	fmt.Println(task2_1)
	fmt.Println(task2_2)

	task3 := task2_1[:2]
	// task3 := task2_1[0:2]
	fmt.Println(task3)

	task4 := task3[1:3]
	fmt.Println(task4)

	task5 := []string{"Learn Go", "In Depth"}
	fmt.Println(task5)
	task6 := append(task5, "Learn building microservices")
	fmt.Println(task6)

	bonusTask := []Product{{title: "Go", id: 2, price: 1.22}, {title: "Microservice", id: 3, price: 5.55}}
	fmt.Println(bonusTask)
	bonusTask = append(bonusTask, Product{title: "In Depth Leaning", id: 4, price: 8.22})
	fmt.Println(bonusTask)

}

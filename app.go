package main

import (
	"example/my-app/Fundamentals/fileOps"
	"example/my-app/Fundamentals/price"
	"fmt"
)

// func greet(str string, doneChan chan bool) {
// 	fmt.Println(str)
// 	doneChan <- true
// }

// func slowGreet(str string, doneChan chan bool) {
// 	time.Sleep(3 * time.Second)
// 	fmt.Println(str)
// 	doneChan <- true
// 	close(doneChan) //only when we know that the task is most longest amongst all
// }

// func main() {
// 	// dones := make([]chan bool, 4)
// 	done := make(chan bool)

// 	// dones[0] = make(chan bool)
// 	go greet("Hello, world! 1", done)
// 	// dones[1] = make(chan bool)
// 	go greet("Hello, world! 2", done)
// 	// dones[2] = make(chan bool)
// 	go slowGreet("Hello, slow world", done)
// 	// dones[3] = make(chan bool)
// 	go greet("Hello, world! 3", done)

// 	// for _, doneChan := range dones {
// 	// 	<-doneChan
// 	// }

// 	for doneChan := range done {
// 		fmt.Println(doneChan)
// 	}
// }

func main() {
	taxRates := []float64{10, 30, 50}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fileManager := fileOps.New("prices.txt", fmt.Sprintf("converted_%.0f.json", taxRate))
		// cmdManager := cmdmanager.New()
		priceObject := price.NewPrice(taxRate, fileManager)
		go priceObject.Process(doneChans[index], errorChans[index])
	}

	// for _, err := range errorChans {
	// 	<-err
	// }
	// for _, done := range doneChans {
	// 	<-done
	// }

	for index, _ := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}

}

package main

import (
	"fmt"
	"math"
)

//5-28-21



// //5-27-21
// func sayGreeting(n string) {
// 	fmt.Printf("Good morning %v \n", n)
// }

// func sayBye(n string) {
// 	fmt.Printf("Goodbye %v \n", n)
// }

// func cycleNames(n []string, f func(string)) {
// 	for _, v := range n {
// 		f(v)
// 	}
// }

// func circleArea(r float64) float64 {
// 	return math.Pi * r * r
// }

func main() {

	//5-27-21
	//see above main
	// sayGreeting("Brandt")
	// sayGreeting("mario")
	// sayBye("luigi")

	// ff7Names := []string{"cloud", "tifa", "barrett"}

	// cycleNames([]string{"cloud", "tifa", "barrett"}, sayGreeting)
	// cycleNames([]string{"cloud", "tifa", "barrett"}, sayBye)
	// cycleNames(ff7Names, sayGreeting)

	// a1 := circleArea(10.5)
	// a2 := circleArea(15)
	// a3 := circleArea(1)


	// fmt.Println(a1, a2, a3)
	// fmt.Printf("The area of each circle is %0.3f, %0.3f, and %0.3f, respectively \n", a1, a2, a3)





	//5-26-21

	// x := 0

	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++{
	// 	fmt.Println("value of i is:", i)
	// }

	// names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("the value at index %v is %v \n", index, value)
	// }

	// for _, value := range names {
	// 	fmt.Printf("the value is %v \n", value)
	// 	value = "new string"
	// }
	// fmt.Println(names)

	// age := 45

	// if age < 30 {
	// 	fmt.Println("age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("age is less than 40")
	// } else {
	// 	fmt.Println("You're old")
	// }

	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Println("continuing at pos", index)
	// 		continue
	// 	}
	// 	if index > 2 {
	// 		fmt.Println("breaking at postion", index)
	// 		break
	// 	}

	// 	fmt.Printf("the value at position %v is %v \n", index, value)
	// }



	//5-25-21

	// greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting, "hellos"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hey"))
	// var newGreeting = strings.ReplaceAll(greeting, "hello", "hi")
	// fmt.Println(newGreeting)
	// fmt.Println(greeting)

	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.ToUpper(newGreeting))

	// fmt.Println(strings.Index(greeting, "ll"))

	// fmt.Println(strings.Split(greeting, " "))

	// ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	// sort.Ints(ages)
	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 30)
	// fmt.Println(index)

	// names := []string{"Brandt", "Steve", "Gerald", "Gracie"}
	// sort.Strings(names)
	// fmt.Println(names)

	// fmt.Println(sort.SearchStrings(names, "Brandt"))

	//5-24-21

	// var nameOne string = "Brian"
	// var nameTwo = "Doug"
	// var nameThree string

	// fmt.Println(nameOne + " " + nameTwo)

	// nameOne = "Toad"
	// nameThree = "Bowser"

	// fmt.Println(nameOne + " " + nameThree)

	// nameFour := "Yoshi"

	// fmt.Println(nameFour)

	// var ageOne int = 20
	// ageTwo := 32

	// fmt.Println(ageOne, ageTwo)

	// var ages = [3]int{23,24,25}
	// fmt.Println(ages, len(ages))

	// names := [4]string{"Brandt", "Steve", "Gerald", "Gracie"}
	// fmt.Println(names)

	// names[1] = "luigi"
	// ages[2] = 3

}

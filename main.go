package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

// //5-28-21

// func getInitials(n string) (string, string) {
// 	s := strings.ToUpper(n)
// 	names := strings.Split(s, " ")

// 	var initials []string
// 	for _, v := range names {
// 		initials = append(initials, v[:1])
// 	}

// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}
// 	return initials[0], "_"
// }

// //5-31-21

// var score = 99.5

// //6-2-21

// func updateName(x *string) {
// 	*x = "wedge"
// }

// func updateMenu(y map[string]float64) {
// 	y["coffee"] = 2.99
// }

//6-8-21

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		fmt.Println("You chose: ", name, price)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		fmt.Println("You tipped: $", tip)
	case "s":
		fmt.Println("You chose 's'")
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}

}

func main() {

	//6-8-21

	myBill := createBill()
	promptOptions(myBill)

	// //6-3-21, etc.

	// myBill := newBill("my bill")

	// myBill.addItem("onion soup", 4.50)
	// myBill.addItem("veg pie", 8.95)
	// myBill.addItem("toffee pudding", 4.95)
	// myBill.addItem("coffee", 3.25)

	// myBill.updateTip(10)

	// fmt.Println(myBill.format())

	// //6-2-21

	// // group A types - > strings, ints, bools, floate, arrays, structs
	// name := "cloud"

	// // updateName(name)
	// fmt.Println("memory address of name is: ", &name)
	// m := &name
	// // fmt.Println(m)
	// // fmt.Println("value at memory addres: ", *m)
	// fmt.Println(name)
	// updateName(m)
	// fmt.Println(name)

	// // group B types - > slices, maps, functions
	// menu := map[string]float64{
	// 	"pie":       5.95,
	// 	"ice cream": 3.99,
	// }

	// updateMenu(menu)

	// fmt.Println(menu)
	// fmt.Println(menu["coffee"])

	// //6-1-21

	// menu := map[string]float64{
	// 	"soup":           4.99,
	// 	"pie":            7.99,
	// 	"salad":          6.99,
	// 	"toffee pudding": 3.55,
	// }

	// fmt.Println(menu)
	// fmt.Println(menu["pie"])

	// //looping maps
	// for k, v := range menu {
	// 	fmt.Println(k, "-", v)
	// }

	// // ints as key type
	// phonebook := map[int]string{
	// 	2345264: "mario",
	// 	9847643: "peach",
	// 	8348744: "luigi",
	// }

	// fmt.Println(phonebook)
	// fmt.Println(phonebook[2345264])

	// phonebook[2345264] = "bowser"
	// fmt.Println(phonebook)
	// fmt.Println(phonebook[2345264])

	// //5-31-21

	// sayHello("Brandt")

	// for _, v := range points {
	// 	fmt.Println(v)
	// }

	// showScore()

	// //5-28-21
	// fn1, sn1 := getInitials("tifa lockhart")
	// fmt.Println(fn1, sn1)
	// fn2, sn2 := getInitials("Brandt Campbell")
	// fmt.Println(fn2, sn2)
	// fn3, sn3 := getInitials("dick")
	// fmt.Println(fn3, sn3)

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

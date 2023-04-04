package main

import "fmt"

func checkIngredients(water *int, milk *int, coffeeBeans *int, cups *int, waterNeeded, milkNeeded, coffeeBeansNeeded, cupsNeeded int) {
	switch {
	case *water < waterNeeded:
		fmt.Println("Sorry, not enough water!")
		return
	case *milk < milkNeeded:
		fmt.Println("Sorry, not enough milk!")
		return
	case *coffeeBeans < coffeeBeansNeeded:
		fmt.Println("Sorry, not enough coffee beans!")
		return
	case *cups < cupsNeeded:
		fmt.Println("Sorry, not enough cups!")
		return
	}
}

func buy(water, milk, coffeeBeans, cups, money *int) {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	var coffeeType string
	fmt.Scan(&coffeeType)
	switch coffeeType {
	case "1":
		checkIngredients(water, milk, coffeeBeans, cups, 250, 0, 16, 1)
		if *water < 250 || *coffeeBeans < 16 || *cups < 1 {
			return
		}
		*water -= 250
		*coffeeBeans -= 16
		*money += 4
		*cups -= 1
	case "2":
		checkIngredients(water, milk, coffeeBeans, cups, 350, 75, 20, 1)
		if *water < 350 || *milk < 75 || *coffeeBeans < 20 || *cups < 1 {
			return
		}
		*water -= 350
		*milk -= 75
		*coffeeBeans -= 20
		*money += 7
		*cups -= 1
	case "3":
		checkIngredients(water, milk, coffeeBeans, cups, 200, 100, 12, 1)
		if *water < 200 || *milk < 100 || *coffeeBeans < 12 || *cups < 1 {
			return
		}
		*water -= 200
		*milk -= 100
		*coffeeBeans -= 12
		*money += 6
		*cups -= 1
	case "back":
		return
	}
	fmt.Println("I have enough resources, making you a coffee!")
}

func remaining(water, milk, coffeeBeans, cups, money *int) {
	fmt.Println()
	fmt.Println("The coffee machine has:")
	fmt.Println(*water, " ml of water")
	fmt.Println(*milk, " ml of milk")
	fmt.Println(*coffeeBeans, "g of coffee beans")
	fmt.Println(*cups, "disposable cups")
	fmt.Printf("$%d of money", *money)
}

func fill(water, milk, coffeeBeans, cups *int) {
	fmt.Println("Write how many ml of water do you want to add:")
	var waterToAdd int
	fmt.Scan(&waterToAdd)
	*water += waterToAdd

	fmt.Println("Write how many ml of milk do you want to add:")
	var milkToAdd int
	fmt.Scan(&milkToAdd)
	*milk += milkToAdd

	fmt.Println("Write how many grams of coffee beans do you want to add:")
	var coffeeBeansToAdd int
	fmt.Scan(&coffeeBeansToAdd)
	*coffeeBeans += coffeeBeansToAdd

	fmt.Println("Write how many disposable cups of coffee do you want to add:")
	var cupsToAdd int
	fmt.Scan(&cupsToAdd)
	*cups += cupsToAdd
}

func main() {
	// put your code here
	initialWater := 400
	initialMilk := 540
	initialCoffeeBeans := 120
	initialCups := 9
	initialMoney := 550

	water := initialWater
	milk := initialMilk
	coffeeBeans := initialCoffeeBeans
	cups := initialCups
	money := initialMoney

	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	var action string
	fmt.Scan(&action)

	for action != "exit" {
		switch action {
		case "buy":
			buy(&water, &milk, &coffeeBeans, &cups, &money)
		case "fill":
			fill(&water, &milk, &coffeeBeans, &cups)
		case "take":
			fmt.Println("I gave you $", money)
			money = 0
		case "remaining":
			remaining(&water, &milk, &coffeeBeans, &cups, &money)
		}
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		fmt.Scan(&action)
	}
}

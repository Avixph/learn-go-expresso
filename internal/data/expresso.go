package data

import (
	"fmt"
	"github.com/Avixph/learn-go-expresso/internal/action"
	"github.com/Avixph/learn-go-expresso/internal/request"
	"strings"
)

type ExpressoMachine struct {
	TotalWater          int
	TotalMilk           int
	TotalCoffeeBeans    int
	TotalDisposableCups int
	TotalCash           int
}

func (e *ExpressoMachine) ProcessUserAction(act string) {
	switch strings.ToLower(act) {
	case "buy":
		ord := action.GetUserOrder()
		e.PrepareOrder(ord)
		act := action.GetUserAction()
		e.ProcessUserAction(act)
	case "fill":
		water, milk, beans, disposableCups := action.GetInventoryFill()
		e.RefillMachine(water, milk, beans, disposableCups)
		act := action.GetUserAction()
		e.ProcessUserAction(act)
	case "take":
		e.WithdrawCash()
		act := action.GetUserAction()
		e.ProcessUserAction(act)
	case "remaining":
		e.PrintState()
		act := action.GetUserAction()
		e.ProcessUserAction(act)
	case "exit":
		break
	default:
		act := action.GetUserAction()
		defer e.ProcessUserAction(act)
		e.invalidSelection()
	}
}

func (e *ExpressoMachine) PrepareOrder(ord string) {
	resource, ok := e.CheckResources(ord)
	if !ok {
		fmt.Printf("Sorry, not enough %s!\n", resource)
	} else {
		e.BrewCoffee(ord)
		fmt.Println("I have enough resources, making you a coffee!")
	}
}

func (e *ExpressoMachine) CheckResources(ord string) (string, bool) {
	if e.TotalDisposableCups <= 0 {
		return "disposable cups", false
	}
	switch ord {
	case "1":
		if e.TotalWater < waterPerEspresso {
			return "water", false
		}
		if e.TotalCoffeeBeans < beansPerEspresso {
			return "coffee beans", false
		}
		return "", true
	case "2":
		if e.TotalWater < waterPerCappuccino {
			return "water", false
		}
		if e.TotalMilk < milkPerCappuccino {
			return "milk", false
		}
		if e.TotalCoffeeBeans < beansPerCappuccino {
			return "coffee beans", false
		}
		return "", true
	case "3":
		if e.TotalWater < waterPerLatte {
			return "water", false
		}
		if e.TotalMilk < milkPerLatte {
			return "milk", false
		}
		if e.TotalCoffeeBeans < beansPerLatte {
			return "coffee beans", false
		}
		return "", true
	default:
		return "", false
	}
}

func (e *ExpressoMachine) BrewCoffee(ord string) {
	switch ord {
	case "1":
		e.TotalWater -= waterPerEspresso
		e.TotalCoffeeBeans -= beansPerEspresso
		e.TotalCash += costPerEspresso
	case "2":
		e.TotalWater -= waterPerLatte
		e.TotalMilk -= milkPerLatte
		e.TotalCoffeeBeans -= beansPerLatte
		e.TotalCash += costPerLatte
	case "3":
		e.TotalWater -= waterPerCappuccino
		e.TotalMilk -= milkPerCappuccino
		e.TotalCoffeeBeans -= beansPerCappuccino
		e.TotalCash += costPerCappuccino
	case "back":
		act := action.GetUserAction()
		e.ProcessUserAction(act)
	}
	e.TotalDisposableCups--
}

func (e *ExpressoMachine) RefillMachine(ingr ...int) {
	e.TotalWater += ingr[0]
	e.TotalMilk += ingr[1]
	e.TotalCoffeeBeans += ingr[2]
	e.TotalDisposableCups += ingr[3]
}

func (e *ExpressoMachine) WithdrawCash() {
	var cash float64
	cash = float64(e.TotalCash * oneHundredPercent / 100)
	e.TotalCash -= int(cash)
	fmt.Println(request.ReturnWithdrawAmount(cash))
}

func (e *ExpressoMachine) PrintState() {
	fmt.Println()
	fmt.Println("The expresso machine has:")
	fmt.Printf("%d ml of water\n", e.TotalWater)
	fmt.Printf("%d ml of milk\n", e.TotalMilk)
	fmt.Printf("%d g of coffee beans\n", e.TotalCoffeeBeans)
	fmt.Printf("%d disposable cups\n", e.TotalDisposableCups)
	fmt.Printf("➤ $%.2f cash deposit\n", float64(e.TotalCash))
}

func (e *ExpressoMachine) invalidSelection() {
	fmt.Println("Invalid selection")
}

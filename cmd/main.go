package main

import (
	"github.com/Avixph/learn-go-expresso/internal/action"
	"github.com/Avixph/learn-go-expresso/internal/data"
)

func main() {
	expresso := data.ExpressoMachine{TotalWater: 400, TotalMilk: 540, TotalCoffeeBeans: 120, TotalDisposableCups: 9, TotalCash: 550}
	act := action.GetUserAction()
	expresso.ProcessUserAction(act)
}

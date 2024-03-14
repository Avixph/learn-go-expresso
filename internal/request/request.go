package request

import "fmt"

func RequestAction() string {
	return fmt.Sprint("Please choose an action:\n◆ buy\t\t◆ fill\n◆ take\t\t◆ remaining\n◆ exit")
}

func RequestOrder() string {
	return fmt.Sprint("What would you like to purchase?\n1 ◆ espresso\t\t2 ◆ latte\n3 ◆ cappuccino\t\t4 ◆ main menu")
}

func RequestWaterAmount() string {
	return fmt.Sprint("How many ml of water do you want to add:")
}

func RequestMilkAmount() string {
	return fmt.Sprint("How many ml of milk do you want to add:")
}

func RequestBeanAmount() string {
	return fmt.Sprint("How many grams of coffee beans do you want to add:")
}

func RequestCupAmount() string {
	return fmt.Sprint("How many disposable cups do you want to add:")
}

func ReturnWithdrawAmount(csh float64) string {
	return fmt.Sprintf("$%.2f withdrawn", csh)
}

package action

import (
	"fmt"
	"github.com/Avixph/learn-go-expresso/internal/request"
)

func GetUserAction() string {
	var act string
	scanInput(request.RequestAction(), &act)
	return act
}

func GetUserOrder() string {
	var ord string
	scanInput(request.RequestOrder(), &ord)
	return ord
}

func GetInventoryFill() (int, int, int, int) {
	var water, milk, beans, disposableCups int
	scanInput(request.RequestWaterAmount(), &water)
	scanInput(request.RequestMilkAmount(), &milk)
	scanInput(request.RequestBeanAmount(), &beans)
	scanInput(request.RequestCupAmount(), &disposableCups)
	return water, milk, beans, disposableCups
}

func scanInput(rqst string, act any) {
	fmt.Println()
	fmt.Println(rqst)
	_, err := fmt.Scan(act)
	if err != nil {
		return
	}
}

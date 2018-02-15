package telegram

import (
)

func getResponse(command string) string {
	
  switch command {
		case "create": {
			return createOrder()
		}
	  case "list": {
      return ordersList()
	  }
		default: {
			return "Poshel nahui"
		}
	}
}



func createOrder() string {
	return "Order Created"
}

func ordersList() string {
  return "Orders list"
}
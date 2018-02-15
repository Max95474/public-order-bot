package telegram

func getResponse(command string) string {
  
  switch command {
    case "create": {
      return createOrder()
    }
    case "list": {
      return ordersList()
    }
    case "add": {
      return addItem()
    }
    case "remove": {
      return removeItem()
    }
    case "result": {
      return orderResult()
    }
    case "close": {
      return closeOrder()
    }
    default: {
      return "Poshel na hui"
    }
  }
}



func createOrder() string {
  return "Order Created"
}

func ordersList() string {
  return "Orders list"
}

func addItem() string {
  return "Item added to order"
}

func removeItem() string {
  return "Item removed from order"
}

func orderResult() string {
  return "All added items"
}

func closeOrder() string {
  return "Order closed"
}
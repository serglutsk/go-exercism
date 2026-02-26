package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    foo := make(map[string]int)
    foo["quarter_of_a_dozen"] = 3
    foo["half_of_a_dozen"] = 6
    foo["dozen"] = 12
    foo["small_gross"] = 120
    foo["gross"] = 144
    foo["great_gross"] = 1728
    return foo
	panic("Please implement the Units() function")
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    return make(map[string]int)
	panic("Please implement the NewBill() function")
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    quantity, exists := units[unit]
    if !exists {
        return false
    }
    bill[item] += quantity
    
    return true
	panic("Please implement the AddItem() function")
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    unitVal, unitExists := units[unit]
    
    billVal, itemExists := bill[item]

    if !unitExists || !itemExists || billVal < unitVal {
        return false
    }

    
    if billVal == unitVal {
        delete(bill, item)
    } else {
        bill[item] -= unitVal
    }

    return true
    
	panic("Please implement the RemoveItem() function")
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    val, exists := bill[item]
    return val, exists
	panic("Please implement the GetItem() function")
}

package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    if kind =="car" || kind == "truck" {
        return true
    } else {
        return false
    }
	panic("NeedsLicense not implemented")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    var r string
    if option1<option2 {
        r = option1 
    } else {
       r = option2
    }
    return r + " is clearly the better choice."  
	panic("ChooseVehicle not implemented")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    if age<3 {
        return originalPrice * 0.8
    } else if age>=10 {
        return originalPrice * 0.5
    } else {
       return originalPrice * 0.7 
    }
	panic("CalculateResellPrice not implemented")
}

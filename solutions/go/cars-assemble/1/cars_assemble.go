package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64(float64(productionRate)*successRate/100)
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    workingCarsPerHour := float64(productionRate) * successRate / 100
    return int(workingCarsPerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    if carsCount<10 {
        return uint(carsCount*10000)
    }
    bulk := int(carsCount/10)
    lefts := carsCount - bulk*10
    return uint(bulk*95000 + lefts*10000)
	panic("CalculateCost not implemented")
}

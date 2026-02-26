package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    count := 0
    l :=len(birdsPerDay)
    for i :=0;i<l; i++ {
       count += birdsPerDay[i]
    }
    return count
	panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    startDay := (week - 1) * 7
    
    
    endDay := startDay + 7
    
    
    weeklySnapshot := birdsPerDay[startDay:endDay]
    
    
    total := 0
    for _, count := range weeklySnapshot {
        total += count
    }
    
    return total
	panic("Please implement the BirdsInWeek() function")
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    l :=len(birdsPerDay)
    for i:=0;i<l; i+=2 {
        birdsPerDay[i]+=1
    }
    return birdsPerDay
	panic("Please implement the FixBirdCountLog() function")
}

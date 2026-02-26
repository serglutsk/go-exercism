package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    
    layout := "1/2/2006 15:04:05"

    t, err := time.Parse(layout, date)
    if err != nil {
        return time.Time{} 
    }
    return t
	panic("Please implement the Schedule function")
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    
    if err != nil {
        return false 
    }
    now := time.Now()
    return now.After(t)
	panic("Please implement the HasPassed function")
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    
    if err != nil {
        return false 
    }
    h := t.Hour()
    if h>=12 && h<= 18 {
        return true
    }else {
        return false
    }
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    layout := "1/2/2006 15:04:05"

    t, err := time.Parse(layout, date)
    if err != nil {
        return "" 
    }
    outputLayout := "Monday, January 2, 2006, at 15:04"
    return "You have an appointment on " + t.Format(outputLayout) + "."
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    currentYear := time.Now().Year()
    return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
	panic("Please implement the AnniversaryDate function")
}

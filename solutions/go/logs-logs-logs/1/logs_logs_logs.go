package logs
import "unicode/utf8"
// Application identifies the application emitting the given log.
func Application(log string) string {
    r := []rune(log)
    for _, char := range r {
        if char =='❗' {
           return "recommendation" 
        } else if char=='🔍' {
           return "search" 
        } else if char=='☀' {
           return "weather" 
        } 
        
    }
    return "default"
	panic("Please implement the Application() function")
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    r := []rune(log)
    for index, char := range r {
      if char ==oldRune {
          r[index] = newRune
          }
    }
        return string(r)
	panic("Please implement the Replace() function")
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    if utf8.RuneCountInString(log) <= limit {
        return true
    }else {
        return false
    }
	panic("Please implement the WithinLimit() function")
}

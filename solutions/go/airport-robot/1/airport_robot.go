package airportrobot
import "fmt"
type Greeter interface {
    LanguageName()string
    Greet(a string)string
}
func (l German)LanguageName() string {
    return "German"
}
func (l Italian)LanguageName() string {
    return "Italian"
}
func (l Portuguese)LanguageName() string {
    return "Portuguese"
}
func (l German)Greet(a string) string {
    return fmt.Sprintf("Hellp %s!", a)
}
func (l Italian)Greet(a string) string {
    return fmt.Sprintf("Ciao %s!", a)
}
func (l Portuguese)Greet(a string) string {
    return fmt.Sprintf("Olá %s!", a)
}
type German struct{}
type Italian struct{}
type Portuguese struct{}
func SayHello(name string, g Greeter)string {
    
    return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}
// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

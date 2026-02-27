package meteorology
import "fmt"
type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (tu TemperatureUnit) String() string {
   if tu == Celsius {
       return "°C"
   }else {
      return "°F"
   } 
}
type Temperature struct {
	degree int
	unit   TemperatureUnit
}
func (t Temperature) String() string {
    if t.unit == Celsius {
        return fmt.Sprintf("%d %s", t.degree, "°C")
    }else {
        return fmt.Sprintf("%d %s", t.degree, "°F")
    }
}
// Add a String method to the Temperature type

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (su SpeedUnit) String() string {
    if su == KmPerHour {
       return "km/h" 
    }else {
        return "mph"
    }
}
type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
    var m string
    if s.unit == KmPerHour {
        m = "km/h"
    }else {
       m = "mph" 
    }
    return fmt.Sprintf("%d %s", s.magnitude, m)
}
type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}
func (s MeteorologyData) String() string {
    t := s.temperature.String()
    ws := s.windSpeed.String()
    return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity",s.location, t, s.windDirection, ws, s.humidity)
}
// Add a String method to MeteorologyData

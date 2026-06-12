// Package weather provides a simple weather forecast functionality.
package weather

var (
	// CurrentCondition holds the current weather condition for the location specified in CurrentLocation.
	CurrentCondition string
	// CurrentLocation holds the name of the location for which the weather condition is being tracked.
	CurrentLocation string
)

// Forecast takes a city and a weather condition as input, updates the CurrentLocation and CurrentCondition variables, and returns a formatted string describing the current weather condition for the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

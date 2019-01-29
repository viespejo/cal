// (c) 2017 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"time"
)

// Common holidays
var (
	HolyThursday  = NewHolidayFunc(calculateHolyThursday)
	TheThreeKings = NewHoliday(time.January, 6)
	FirstOfMay    = NewHoliday(time.May, 1)
	Asuncion      = NewHoliday(time.August, 15)
	SpainDay      = NewHoliday(time.October, 12)
	AllSaintsDay  = NewHoliday(time.November, 1)
	Constitution  = NewHoliday(time.December, 6)
	Inmaculada    = NewHoliday(time.December, 8)
)

// AddSpainHolidays adds all Spain holidays
// 8 national holidays
func AddSpainHolidays(c *Calendar) {
	c.AddHoliday(
		NewYear,
		GoodFriday,   // Viernes Santos
		FirstOfMay,   // Primero de Mayo
		Asuncion,     // Asunción
		SpainDay,     // Hispanidad
		AllSaintsDay, // Todos los Santos
		Constitution,
		Christmas,
	)
}

// AddMadridHolidays adds all Madrid holidays
// 4 Madrid Autonomous Community holidays
// 2 local holidays
func AddMadridHolidays(c *Calendar) {
	c.AddHoliday(
		TheThreeKings,                // Reyes Magos
		HolyThursday,                 // Jueves Santos
		NewHoliday(time.May, 2),      // Comunidad de Madrid
		NewHoliday(time.May, 15),     // San Isidro (Local - MADRID)
		NewHoliday(time.November, 9), // Almudena (Local - MADRID)
		Inmaculada,                   // Inmaculada
	)
}

func calculateHolyThursday(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// three days before Easter Sunday
	gf := easter.AddDate(0, 0, -3)
	return gf.Month(), gf.Day()
}

package main

import (
	"fmt"
	"math"
)

// Named types
type (
	Feet         float64
	Centimeter   float64
	Minutes      float64
	Seconds      float64
	Milliseconds float64
	Celsius      float64
	Fahrenheit   float64
	Radian       float64
	Degree       float64
	Kilogram     float64
	Pounds       float64
	Gallons      float64
	Liters       float64
)

// Converter struct
type Converter struct {
	feet       Feet
	Centimeter Centimeter
	mins       Minutes
	secs       Seconds
	mSecs      Milliseconds
	cel        Celsius
	feh        Fahrenheit
	rad        Radian
	deg        Degree
	kg         Kilogram
	lbs        Pounds
	gal        Gallons
	L          Liters
}

// FeetToCentimeter : 1ft = 30.48cm
func (cvr Converter) FeetToCentimeter(f Feet) Centimeter {
	return Centimeter(f) * Centimeter(30.48)
}

// CentimeterToFeet : 1cm = 1/30.48ft
func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	return Feet(c / 30.48)
}

// MinutesToSeconds : 1m = 60secs
func (cvr Converter) MinutesToSeconds(m Minutes) Seconds {
	return Seconds(m) * Seconds(60)
}

// SecondsToMinutes : 1s = 1/60mins
func (cvr Converter) SecondsToMinutes(s Seconds) Minutes {
	return Minutes(s / 60)
}

// SecondsToMilliseconds : 1s = 60msecs
func (cvr Converter) SecondsToMilliseconds(s Seconds) Milliseconds {
	return Milliseconds(s) * Milliseconds(60)
}

// MillisecondsToSeconds  :  1msecs = 1/60s
func (cvr Converter) MillisecondsToSeconds(msecs Milliseconds) Seconds {
	return Seconds(msecs / 60)
}

// CelsiusToFahrenheit  :
func (cvr Converter) CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit((c * (9 / 5)) + 32)
}

// FahrenheitToCelsius  :
func (cvr Converter) FahrenheitToCelsius(feh Fahrenheit) Celsius {
	return Celsius((feh - 32) * 5 / 9)
}

// RadianToDegree  :
func (cvr Converter) RadianToDegree(rad Radian) Degree {
	return Degree(rad * (180 / math.Pi))
}

// DegreeToRadian  :
func (cvr Converter) DegreeToRadian(deg Degree) Radian {
	return Radian(deg * (math.Pi / 180))
}

// KilogramToPounds  :
func (cvr Converter) KilogramToPounds(kg Kilogram) Pounds {
	return Pounds(kg * 2.205)
}

// PoundsToKilogram  :
func (cvr Converter) PoundsToKilogram(lbs Pounds) Kilogram {
	return Kilogram(lbs / 2.205)
}

//**************************** Bonus task  *****************

// GallonsToLiters  :
func (cvr Converter) GallonsToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

// LitersToGallons  :
func (cvr Converter) LitersToGallons(l Liters) Gallons {
	return Gallons(l / 3.785)
}

//**************************** End of Bonus task  *****************

func main() {

	cvr := Converter{}
	// fmt.Println(cvr)

	ft := fmt.Sprintf("%f", float64(cvr.CentimeterToFeet(1)))
	fmt.Println(ft + " ft")
	cm := fmt.Sprintf("%f", float64(cvr.FeetToCentimeter(1)))
	fmt.Println(cm + " cm")

	mins := fmt.Sprintf("%f", float64(cvr.SecondsToMinutes(1)))
	fmt.Println(mins + " minutes")
	secs := fmt.Sprintf("%f", float64(cvr.MinutesToSeconds(1)))
	fmt.Println(secs + " seeconds")

	sec := fmt.Sprintf("%f", float64(cvr.MillisecondsToSeconds(1)))
	fmt.Println(sec + " seconds")
	msec := fmt.Sprintf("%f", float64(cvr.SecondsToMilliseconds(1)))
	fmt.Println(msec + " milliseconds")

	feh := fmt.Sprintf("%f", float64(cvr.CelsiusToFahrenheit(0)))
	fmt.Println(feh + " F")
	cel := fmt.Sprintf("%f", float64(cvr.FahrenheitToCelsius(32)))
	fmt.Println(cel + " C")

	deg := fmt.Sprintf("%f", float64(cvr.RadianToDegree(1)))
	fmt.Println(deg + " degree")
	rad := fmt.Sprintf("%f", float64(cvr.DegreeToRadian(1)))
	fmt.Println(rad + " radian")

	lbs := fmt.Sprintf("%f", float64(cvr.KilogramToPounds(1)))
	fmt.Println(lbs + " pounds")
	kg := fmt.Sprintf("%f", float64(cvr.PoundsToKilogram(1)))
	fmt.Println(kg + " Kg")

	lit := fmt.Sprintf("%f", float64(cvr.GallonsToLiters(1)))
	fmt.Println(lit + " L")
	gal := fmt.Sprintf("%f", float64(cvr.LitersToGallons(1)))
	fmt.Println(gal + " gal")

}

package tempconv

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9.0/5.0 + 32.0)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}

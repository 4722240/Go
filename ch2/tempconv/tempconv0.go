package tempconv

type Celsius float32
type Fahrenheit float32

const (
	AbsoluteZeroC	= -273.15
	FreezingC 	= 0
	BoilingC	= 100
)
func Ctof(c Celsius) Fahrenheit {
	return Fahrenheit( c * 9 / 5 + 32 )
}

func Ftoc(f Fahrenheit) Celsius {
	return Celsius( ( f - 32 ) *5 /9)
}
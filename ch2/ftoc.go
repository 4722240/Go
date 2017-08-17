package main

import(
	"fmt"
	"./tempconv"
)

func main() {
	const fleezingF,boilingF  = 32.0,212.0
	fmt.Printf("%g °F = %g °C\n",fleezingF,ftoc(fleezingF))
	fmt.Printf("%g °F = %g °C\n",boilingF,ftoc(boilingF))

	fmt.Println(tempconv.BoilingC - tempconv.FreezingC)
}

func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9
}

package main

import (
	"fmt"
)

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func main() {
	var kelvin float64

	fmt.Println("Digite o valor em Kelvin:")
	_, err := fmt.Scanf("%f", &kelvin)
	if err != nil {
		fmt.Println("Erro ao ler o valor em Kelvin:", err)
		return
	}

	celsius := kelvinToCelsius(kelvin)
	fmt.Printf("%.2f Kelvin é equivalente a %.2f Celsius\n", kelvin, celsius)
	fmt.Printf("A temperatura de ebulição da água em K = %g , temperatura de ebulição da água em °C =%g.", kelvin, celsius)
}

//Criando um Programa em GO para a Conversão de Escalas Termométricas
/* Resultado impresso no terminal.
Digite o valor em Kelvin:
373.15
373.15 Kelvin é equivalente a 100.00 Celsius
A temperatura de ebulição da água em K = 373.15 , temperatura de ebulição da água em °C =100.
*/

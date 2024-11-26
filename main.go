package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	inputFileName := "input.txt"
	outputFileName := "output.txt"

	file, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	output, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error al crear el archivo de salida:", err)
		return
	}
	defer output.Close()

	contadorLinea := 0
	lineaPrecios := ""
	lineaDineroTotal := ""

	for scanner.Scan() {

		linea := scanner.Text()

		switch contadorLinea % 4 {
		case 1:
			lineaPrecios = linea
		case 2:
			lineaDineroTotal = linea

			// fmt.Println("Lineas de precios :", lineaPrecios)
			// fmt.Println("Lineas de dinero total:", lineaDineroTotal)
			listaPrecios, err := convLineaPreciosALista(lineaPrecios) //falta usar el lista precios
			if err != nil {
				fmt.Println("Error al convertir :", err)
				return
			}
			dineroTotalInt, err := strconv.Atoi(lineaDineroTotal) //falta usar dinerototalint
			if err != nil {
				fmt.Println("Error al convertir el string a int:", err)
				return
			}

			sumaPrecios := sumarLista(listaPrecios)

			if sumaPrecios == dineroTotalInt {

				listaPreciosStr := convertirListaIntAStr(listaPrecios)

				listaPreciosConComa := strings.Join(listaPreciosStr, ", ")

				LineaTxt := "Peter deberia comprar los libros de valor: " + listaPreciosConComa + "\n"

				_, err := output.WriteString(LineaTxt + "\n")
				if err != nil {
					fmt.Println("No se pudo escrubir el archivo: ", err)
					return
				}

			}

		}
		contadorLinea++

	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error al leer el archivo: %v\n", err)
	}

}

func convLineaPreciosALista(s string) ([]int, error) {
	partes := strings.Fields(s)

	nums := make([]int, len(partes))

	for i, parte := range partes {
		num, err := strconv.Atoi(parte)
		if err != nil {
			return nil, err
		}
		nums[i] = num
	}
	return nums, nil
}

func sumarLista(lista []int) int {
	suma := 0
	for _, valor := range lista {
		suma += valor
	}
	return suma
}

func convertirListaIntAStr(lista []int) []string {
	strLista := make([]string, len(lista))

	for i, num := range lista {
		strLista[i] = strconv.Itoa(num)
	}
	return strLista
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nombreArchivoEntrada := "input.txt"
	nombreArchivoSalida := "output.txt"

	archivo, err := os.Open(nombreArchivoEntrada)
	if err != nil {
		fmt.Println("Error al abrir el archivo: ", err)
		return
	}

	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)

	output, err := os.Create(nombreArchivoSalida)
	if err != nil {
		fmt.Println("Error al crear el archivo de salida: ", err)
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

			listaPrecios, err := convLineaPreciosALista(lineaPrecios)
			if err != nil {
				fmt.Println("Error al convertir: ", err)
				return
			}

			dineroTotalInt, err := strconv.Atoi(lineaDineroTotal)
			if err != nil {
				fmt.Println("Error al convertir el string a int: ", err)
				return
			}

			valorMenor := 0
			valorMayor := 0
			valor1, valor2 := controlSuma(listaPrecios, dineroTotalInt)

			if valor1 < valor2 {
				valorMenor = valor1
				valorMayor = valor2
			} else {
				valorMenor = valor2
				valorMayor = valor1
			}

			valorMayorStr := strconv.Itoa(valorMayor)
			valorMenorStr := strconv.Itoa(valorMenor)
			lineaTxt := "Peter deberia comprar los libros de valor: " + valorMenorStr + " y " + valorMayorStr + "\n"

			_,err= output.WriteString(lineaTxt + "\n")
			if err != nil {
				fmt.Println("No se pudo escrbir el archivo: ", err)
				return
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

func controlSuma(lista []int, valTotal int) (int, int) {
	var val1, val2 int
	for i := 0; i < len(lista); i++ {
		for j := i + 1; j < len(lista); j++ {
			if lista[i]+lista[j] == valTotal {
				val1, val2 = lista[i], lista[j]
			}
		}
	}

	return val1, val2
}

func convertirListaIntAStr(lista []int) []string {
	strLista := make([]string, len(lista))

	for i, num := range lista {
		strLista[i] = strconv.Itoa(num)
	}

	return strLista
}

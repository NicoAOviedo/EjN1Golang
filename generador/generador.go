package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 	Quiero que cuando corra el programa, en la consola me pregunte cuántos "escenarios" quiero que genere.
	// Cuando le ponga OK, quiero que me diga que está procesando mientras hace su magia.
	// Cuando haya terminado, quiero que me avise en consola que terminó y se termine el programa.
	// Quiero que haya dejado un archivo input-gen.txt en el que me creó los "escenarios".
	// Entiendase un "escenario" como una sucesión de 3 líneas, en las cuales:
	// La primera línea tiene la cantidad de libros que hay disponibles.
	// La segunda línea tiene los precios de cada libro, separados por un espacio.
	// La tercera línea que tiene la cantidad de dinero total.
	// Y obviamente estos escenarios tienen que cumplir los requisitos del pdf original.

	//1. Pedir cantidad de escenarios(4 lineas, 3 + 1 espacio)
	//2. Generar un mensaje al inicio de un proceso que ese proceso comenzo y que esta procesando.
	//3. Avisar por consola que el programa termino.
	//4. Guardar los escenarios en un archivo "input-gen.txt"
	nombreArchivoGenerador := "input-gen.txt"
	archivoGen, err := os.Create(nombreArchivoGenerador)
	if err != nil {
		fmt.Println("Error al crear el archivo generador: ", err)
		return
	}

	defer archivoGen.Close()

	var cantidadPosiblesEscenarios string
	var cantPosInt int
	cantidadValida := false

	for !cantidadValida {
		fmt.Print("Ingresar cantidad de escenarios: ")
		fmt.Scan(&cantidadPosiblesEscenarios)

		cantPosInt, err = strconv.Atoi(cantidadPosiblesEscenarios)
		if err != nil {
			fmt.Println("El numero ingresado no es valido.\nEl valor debe ser un numero entero.")
		} else {
			cantidadValida = true
		}
	}

	for i := 0; i < cantPosInt; i++ {
		cantidadLibros := rand.Intn(10) + 2
		fmt.Println("Cantidad de libros: ", cantidadLibros)
		listaPrecios := make([]int, cantidadLibros)

		for i := 0; i < cantidadLibros; i++ {
			listaPrecios[i] = rand.Intn(50) + 1
		}

		fmt.Println("Lista de precios: ", listaPrecios)

		var precioTotal int
		var indice1 int
		var indice2 int

		for indice1 == indice2 { //redundancia ciclica
			indice1 = rand.Intn(len(listaPrecios))
			indice2 = rand.Intn(len(listaPrecios))
			precioTotal = listaPrecios[indice1] + listaPrecios[indice2]
		}

		fmt.Println("Precio total: ", precioTotal)
		listaAStr := convertirListaIntAStr(listaPrecios)
		lineaStr := strings.Join(listaAStr, " ")
		escenarioATxt := fmt.Sprintf("%v\n%v\n%v\n\n", cantidadLibros, lineaStr, precioTotal)
		_, err := archivoGen.WriteString(escenarioATxt)
		if err != nil {
			fmt.Println("No se pudo escribir el archivo.", err)
			return
		}
	}

	fmt.Println("El programa finalizo. \nCantidad de escenarios: ", cantPosInt)
}

func convertirListaIntAStr(lista []int) []string {
	strLista := make([]string, len(lista))

	for i, num := range lista {
		strLista[i] = strconv.Itoa(num)
	}

	return strLista
}

//Crear un programa con 3 opciones:
// 1. Capturar 5 alumnos con 3 calificaciones cada uno
//2.  Guardar la información capturada en un archivo de texto
//3. Abrir archivo de texto y visualizar información 

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

// Alumno model
type Alumno struct {
	Nombre   string
	Calificaciones []int
}

var alumnos []Alumno

func main() {
	menu()
}

func menu() {
	var input string
	for input != "4" {
		
		input = "0"
		fmt.Println("1.- Ingresar información de Alumnos")
		fmt.Println("2.- Grabar en archivo")
		fmt.Println("3.- Mostrar Archivo")
		fmt.Println("4.- Salir")
		fmt.Print("Seleccione operación")
		fmt.Scanln(&input)
		switch input {
		case "1":
			IngresarAlumnos()
		case "2":
			GrabarArchivo()
		case "3":
			MostrarArchivo()
			Seguir()
		case "4":

		}
	}
}

func IngresarAlumnos() {
	cont := 0
	cont2 := 0
	for len(alumnos) < 5 {
		var alumno Alumno
		fmt.Println("Ingrese nombre del alumno " + strconv.Itoa(cont+1))
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			if str := scanner.Text(); len(str) != 0 {
				alumno.Nombre = str
				cont++
				var calificaciones []int
				for len(calificaciones) < 3 {
					fmt.Println("Ingresar calificación " + strconv.Itoa(cont2+1) + " de alumno " + strconv.Itoa(cont))
					scanner2 := bufio.NewScanner(os.Stdin)
					if scanner2.Scan() {
						if str := scanner2.Text(); len(str) != 0 {
							val, err := strconv.Atoi(str)
							if err != nil {
								fmt.Println("Ingrese un número entero")
							} else {
								calificaciones = append(calificaciones, val)
								cont2++
							}
						}
					}
				}
				cont2 = 0
				alumno.Calificaciones = calificaciones
				alumnos = append(alumnos, alumno)
			}
		}
	}
	MostrarAlumnos(alumnos)
}

func GrabarArchivo() {
	if len(alumnos) > 0 {
		archivo, err := os.Create("test.txt") 
		if err != nil {
			log.Fatalf("Error al crear archivo: %s", err)
		}
		defer archivo.Close() 
		espacio := "              "
		for cont := 0; cont < len(alumnos); cont++ {
			if cont == 0 {
				_, err := archivo.WriteString("Nombre         Calificación 1  Calificación 2  Calificación 3")
				if err != nil {
					log.Fatalf("Error al crear archivo: %s", err)
				}
			}
			alumno := alumnos[cont]
			LugarNombre := ""
			for h := 0; h < 15-len(alumno.Nombre); h++ {
				LugarNombre += " "
			}
			_, err := archivo.WriteString("\n" + alumno.Nombre + LugarNombre + "       " + strconv.Itoa(alumno.Calificaciones[0]) + espacio + strconv.Itoa(alumno.Calificaciones[1]) + espacio + strconv.Itoa(alumno.Calificaciones[2]))
			if err != nil {
				log.Fatalf("Error al crear archivo: %s", err)
			}
		}

		fmt.Println()
		fmt.Println("Archivo Grabado")
		fmt.Println()
	} else {
		fmt.Println("No hay datos para escribir")
		Seguir()
	}
}

func MostrarArchivo() {
	data, err := ioutil.ReadFile("test.txt")
	if err == nil {
		fmt.Printf("\n%s", data)
		fmt.Println()
	} else {
		fmt.Println(err)
		
	}

}

func MostrarAlumnos(alumnos []Alumno) {
	if len(alumnos) != 0 {
		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
		fmt.Println("Estudiantes")
		fmt.Fprintln(w, "\t  Nombre\t  Calificación 1\t  Calificación 2\t  Calificación 3\t")
		for _, alumno := range alumnos {
			fmt.Fprintln(w, ("\t  " + alumno.Nombre + "\t  " + strconv.Itoa(alumno.Calificaciones[0]) + "\t  " + strconv.Itoa(alumno.Calificaciones[1]) + "\t  " + strconv.Itoa(alumno.Calificaciones[2]) + "\t  "))
		}
		w.Flush()
	} else {
		fmt.Println("Not Found")
	}
}

func Seguir() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter paraContinuar")
	scanner.Scan()
}
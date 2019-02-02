package main

import (
	"fmt"

	"github.com/ronaldadrianruizavila/Curso_Go/API/services/heroeservice"
)

func main() {

	fmt.Println("Nombre: Ronal Adrian Ruiz Avila")
	fmt.Println("Pais: Ecuador")
	menuMain()
}

func menuMain() {
	option := 0
	for {
		fmt.Println("\n-----MENU API-----")
		fmt.Println("1. Buscar por nombre")
		fmt.Println("2. Devolver lista")
		fmt.Println("3. salir")
		fmt.Print("Ingrese la opcion: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			searchToName()
			break
		case 2:
			heroeservice.GetHeroes()
		default:
			return
		}
	}
}

func searchToName() {
	option := 0
	nameHeroe := ""

	for {
		fmt.Println("\n\n----BUSQUEDA POR NOMBRE----")
		fmt.Println("1. INGRESAR NOMBRE")
		fmt.Println("2. SALIR")
		fmt.Print("Ingrese la opcion: ")
		fmt.Scanln(&option)

		if option == 1 {

			for {
				fmt.Print("\n-Ingrese el nombre del heroe: ")
				fmt.Println("(Ingrese s para salir)")
				fmt.Scanln(&nameHeroe)

				if len(nameHeroe) > 0 && nameHeroe != "s" {
					heroeservice.GetHeroe(nameHeroe)
				}

				if nameHeroe == "s" {
					break
				}
			}
		}
		if option == 2 {
			break
		}

		if option != 1 && option != 2 {
			return
		}
	}
}

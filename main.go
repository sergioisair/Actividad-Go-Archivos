package main

import (
	"fmt"
  "sort"
  "os"
)

func main() {
	var n int
	var s string
  // INGRESAR DATOS
	fmt.Print("¿Cuántos Strings desea guardar?  ")
	fmt.Scan(&n)
	strings := make([]string,0,n)
  fmt.Println("\nCAPTURAR DATOS\n")
	for i := 0; i < n; i++ {
		fmt.Print("Ingrese el string ",i+1, ": ")
		fmt.Scan(&s)
		strings = append(strings, s)
	}

  // CREAR ARCHIVOS ASC Y DESC
	fileas, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileas.Close()

  filedes, err := os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filedes.Close()

  // IMPRIMIR Y ORDENAMIETNOS
  fmt.Println("\nORIGINAL: ",strings)

  sort.Strings(strings)
	fmt.Println("ASCENDENTE: ",strings)
	for _, s := range strings {
		fileas.WriteString(s + "\n")
	}

	sort.Sort(sort.Reverse(sort.StringSlice(strings)))
	fmt.Println("DESCENDENTE: ",strings)
	for _, s := range strings {
		filedes.WriteString(s + "\n")
	}
}
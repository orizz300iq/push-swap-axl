package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	liste, err := lireNombres(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: arguments invalides")
		os.Exit(1)
	}

	p := Piles{
		a: liste,
	}

	p.trier()
	for _, op := range p.operations { // affichage ligne par ligne car p.operations est un tableau
    	fmt.Println(op)
	}
}
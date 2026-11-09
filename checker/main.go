package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	liste, err := lireNombres(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	p := Piles{
		a: liste,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		op := scanner.Text()

		switch op {
		case "sa":
			p.sa()
		case "sb":
			p.sb()
		case "ss":
			p.ss()
		case "pa":
			p.pa()
		case "pb":
			p.pb()
		case "ra":
			p.ra()
		case "rb":
			p.rb()
		case "rr":
			p.rr()
		case "rra":
			p.rra()
		case "rrb":
			p.rrb()
		case "rrr":
			p.rrr()
		default:
			// Si l'opération n'existe pas, c'est une erreur
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	if p.aEstTriee() && len(p.b) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Piles struct {
	a          []int
	b          []int
	operations []string
}

func (p *Piles) ra() {
	if len(p.a) < 2 {
		return
	}
	p.a = append(p.a[1:], p.a[0])
	p.operations = append(p.operations, "ra")
}

func (p *Piles) sa() {
	if len(p.a) < 2 {
		return
	}
	p.a[0], p.a[1] = p.a[1], p.a[0]
	p.operations = append(p.operations, "sa")
}

func (p *Piles) pb() {
	if len(p.a) == 0 {
		return
	}
	p.b = append([]int{p.a[0]}, p.b...)
	p.a = p.a[1:]
	p.operations = append(p.operations, "pb")
}

func (p *Piles) pa() {
	if len(p.b) == 0 {
		return
	}
	p.a = append([]int{p.b[0]}, p.a...)
	p.b = p.b[1:]
	p.operations = append(p.operations, "pa")
}

func (p *Piles) rra() {
	if len(p.a) < 2 {
		return
	}
	p.a = append([]int{p.a[len(p.a)-1]}, p.a[:len(p.a)-1]...)
	p.operations = append(p.operations, "rra")
}

func (p *Piles) sb() {
	if len(p.b) < 2 {
		return
	}
	p.b[0], p.b[1] = p.b[1], p.b[0]
	p.operations = append(p.operations, "sb")
}

func (p *Piles) rb() {
	if len(p.b) < 2 {
		return
	}
	p.b = append(p.b[1:], p.b[0])
	p.operations = append(p.operations, "rb")
}

func (p *Piles) rrb() {
	if len(p.b) < 2 {
		return
	}
	p.b = append([]int{p.b[len(p.b)-1]}, p.b[:len(p.b)-1]...)
	p.operations = append(p.operations, "rrb")
}

func (p *Piles) afficher() {
	fmt.Println("a :", p.a, " b :", p.b)
}

func (p *Piles) aEstTriee() bool {
	for i := 1; i < len(p.a); i++ {
		if p.a[i] >= p.a[i-1] {
			continue
		} else {
			return false
		}
	}
	return true
}

func (p *Piles) positionDuMinimumA() int {
	min := p.a[0]
	indexmin := 0
	for i := 1; i < len(p.a); i++ {
		if p.a[i] < min {
			min = p.a[i]
			indexmin = i
		}
	}
	return indexmin
}

func (p *Piles) amenerAuSommet(position int) {
	if position < len(p.a)/2 {
		for i := 0; i < position; i++ {
			p.ra()
		}
	} else {
		for i := 0; i < len(p.a)-position; i++ {
			p.rra()
		}
	}
}

func (p *Piles) trierTrois() {
	if p.a[2] > p.a[0] && p.a[2] > p.a[1] {
		if p.a[0] > p.a[1] {
			p.sa() // cas 213, puis cas 123 mais on ne change rien pour le deuxième
		}
	} else if p.a[2] < p.a[0] && p.a[2] < p.a[1] {
		if p.a[0] > p.a[1] {
			p.sa()
			p.rra() // cas 321
		} else {
			p.rra() // cas 231
		}

	} else {
		if p.a[0] > p.a[1] {
			p.ra() // cas 312
		} else {
			p.sa()
			p.ra() // cas 132
		}
	}
}

func (p *Piles) trierCinq() {
	for len(p.a) > 3 {
		posMin := p.positionDuMinimumA()
		p.amenerAuSommet(posMin)
		p.pb()
	}
	p.trierTrois()
	for len(p.b) > 0 {
		p.pa()
	}
}

func (p *Piles) trier() {
	if p.aEstTriee() {
		return
	}

	if len(p.a) == 2 {
		p.sa()
	} else if len(p.a) == 3 {
		p.trierTrois()
	} else if len(p.a) <= 5 {
		p.trierCinq()
	} else {
		p.trierEfficace()
	}
}

func (p *Piles) trierEfficace() {
	if p.aEstTriee() {
		return
	}
	for len(p.a) > 0 {
		posMin := p.positionDuMinimumA()
		p.amenerAuSommet(posMin)
		p.pb()
	}
	for len(p.b) > 0 {
		p.pa()
	}
}

func genererPileAleatoire(taille int) []int {
	slice := make([]int, taille)
	for i := 0; i < taille; i++ {
		slice[i] = rand.Intn(100) // Nombres aléatoires entre 0 et 999
	}
	return slice
}

func (p *Piles) verif(taille int) {
	p.a = genererPileAleatoire(taille)
	p.afficher()
	p.trierEfficace()
	if p.aEstTriee() {
		fmt.Println(p.a)
		fmt.Println(p.operations)
		fmt.Println("Nombre d'opérations :", len(p.operations))

	}
}

func lireNombres(args []string) ([]int, error) {
	tab := []int{}
	elements := []string{}
	for _, arg := range args {
		elements = append(elements, strings.Fields(arg)...)
	}
	for _, elem := range elements {
		nombre, err := strconv.Atoi(elem)
		if err != nil {
			return nil, fmt.Errorf("%q n'est pas un nombre entier", elem)
		} else {
			tab = append(tab, nombre)
		}
	}
	return tab, nil
}
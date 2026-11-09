package main

import (
	"fmt"
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

func (p *Piles) rb() {
	if len(p.b) < 2 {
		return
	}
	p.b = append(p.b[1:], p.b[0])
	p.operations = append(p.operations, "rb")
}

func (p *Piles) rr() {
	if len(p.a) >= 2 {
		p.a = append(p.a[1:], p.a[0])
	}
	if len(p.b) >= 2 {
		p.b = append(p.b[1:], p.b[0])
	}
	p.operations = append(p.operations, "rr")
}

func (p *Piles) sa() {
	if len(p.a) < 2 {
		return
	}
	p.a[0], p.a[1] = p.a[1], p.a[0]
	p.operations = append(p.operations, "sa")
}

func (p *Piles) sb() {
	if len(p.b) < 2 {
		return
	}
	p.b[0], p.b[1] = p.b[1], p.b[0]
	p.operations = append(p.operations, "sb")
}

func (p *Piles) ss() {
	if len(p.a) >= 2 {
		p.a[0], p.a[1] = p.a[1], p.a[0]
	}
	if len(p.b) >= 2 {
		p.b[0], p.b[1] = p.b[1], p.b[0]
	}
	p.operations = append(p.operations, "ss")
}

func (p *Piles) pa() {
	if len(p.b) == 0 {
		return
	}
	p.a = append([]int{p.b[0]}, p.a...)
	p.b = p.b[1:]
	p.operations = append(p.operations, "pa")
}

func (p *Piles) pb() {
	if len(p.a) == 0 {
		return
	}
	p.b = append([]int{p.a[0]}, p.b...)
	p.a = p.a[1:]
	p.operations = append(p.operations, "pb")
}

func (p *Piles) rra() {
	if len(p.a) < 2 {
		return
	}
	p.a = append([]int{p.a[len(p.a)-1]}, p.a[:len(p.a)-1]...)
	p.operations = append(p.operations, "rra")
}

func (p *Piles) rrb() {
	if len(p.b) < 2 {
		return
	}
	p.b = append([]int{p.b[len(p.b)-1]}, p.b[:len(p.b)-1]...)
	p.operations = append(p.operations, "rrb")
}

func (p *Piles) rrr() {
	if len(p.a) >= 2 {
		p.a = append([]int{p.a[len(p.a)-1]}, p.a[:len(p.a)-1]...)
	}
	if len(p.b) >= 2 {
		p.b = append([]int{p.b[len(p.b)-1]}, p.b[:len(p.b)-1]...)
	}
	p.operations = append(p.operations, "rrr")
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

func lireNombres(args []string) ([]int, error) {
	tab := []int{}
	elements := []string{}
	dic := make(map[string]int)
	for _, arg := range args {
		elements = append(elements, strings.Fields(arg)...)
	}
	for _, elem := range elements {
		nombre, err := strconv.Atoi(elem)
		if err != nil {
			return nil, fmt.Errorf("%q n'est pas un nombre entier", elem)
		} else if dic[elem] >= 1 {
			return nil, fmt.Errorf("%q est en double dans la liste donnée", elem)
		} else {
			tab = append(tab, nombre)
			dic[elem] = 1
		}
	}
	return tab, nil
}

package main

import (
	"fmt"
	"strings"
)

const CAPACITY int32 = 10
const GROW_FACTOR = 2

type PaiNatal struct {
	max        float32
	presentes  []Presente
	counter    int32
	disponivel float32
}
type Presente struct {
	descricao     string
	nome          string
	quantidade    int8
	price         float32
	effectiveCost float32
}

/**Cria um novo PaiNatal.
 */
func paiNatal(max float32) PaiNatal {
	return PaiNatal{
		max:        max,
		presentes:  make([]Presente, CAPACITY),
		disponivel: max,
		counter:    0,
	}
}

func (p *PaiNatal) growPresentesSize() {
	aux := p.presentes
	p.presentes = make([]Presente, len(p.presentes)*GROW_FACTOR)
	for index, value := range aux {
		p.presentes[index] = value
	}
}

func (p *PaiNatal) oferece(descricao string, nome string, quantidade int8, precoUnitario float32) {
	totalToSpend := precoUnitario * float32(quantidade)
	if p.disponivel-totalToSpend > 0 {
		p.disponivel -= totalToSpend
		presente := makePresent(descricao, nome, quantidade, precoUnitario)
		if int(p.counter) == len(p.presentes) {
			p.growPresentesSize()
		}
		p.presentes[p.counter] = presente
		p.counter++
	}
}

func makePresent(descricao string, nome string, quantidade int8, precoUnitario float32) Presente {
	return Presente{
		descricao:     descricao,
		nome:          nome,
		quantidade:    quantidade,
		price:         precoUnitario,
		effectiveCost: float32(quantidade) * precoUnitario,
	}
}
func (p *Presente) print() {
	fmt.Printf("Nome: %s ; descricao: %s ; quantidade : %d ; preco : %f \n", p.nome, p.descricao, p.quantidade, p.price)
}

func (p *PaiNatal) print() {
	if p.counter == 0 {
		println("O SACO DE PRESENTES ESTA VAZIO!")
	}
	for index, presente := range p.presentes {
		if int32(index) < p.counter {
			presente.print()
		} else {
			break
		}
	}
}

func (p *PaiNatal) contaOfertas() float32 {
	var total float32 = 0

	for index, presente := range p.presentes {
		if int32(index) < p.counter {
			total += presente.effectiveCost
		} else {
			break
		}
	}
	return total
}

func (p *PaiNatal) maisCaro() *Presente {
	var result *Presente = nil
	if p.counter > 0 {
		result = &p.presentes[0]
	}
	for index, present := range p.presentes {
		if present.effectiveCost > result.effectiveCost {
			result = &present
		}
		if int32(index) >= p.counter {
			break
		}
	}
	return result
}

func (p *PaiNatal) castiga(nome string) {
	var position int32 = -1
	for index, presente := range p.presentes {
		if int32(index) < p.counter {
			if strings.Compare(nome, presente.nome) == 0 {
				position = int32(index)
				break
			}
		} else {
			break
		}
	}
	if position > 0 {
		p.counter--
		p.presentes = append(p.presentes[:position], p.presentes[position+1:]...)
	}
}

func main() {
	println("-> PAI NATAL APP STARTED!")
	pNatal := paiNatal(5000000)
	pNatal.print()
	pNatal.oferece("casa", "daniel", 3, 1000)
	pNatal.oferece("carro", "joao", 3, 2000)
	pNatal.oferece("computador", "mateus", 3, 3000)
	pNatal.oferece("telemovel", "barata", 3, 4000)
	pNatal.oferece("cadeira", "hugo", 3, 5000)
	pNatal.oferece("livro", "nelas", 3, 7000)
	pNatal.print()
	pNatal.maisCaro().print()
	println()
	pNatal.castiga("nelas")
	pNatal.print()
}

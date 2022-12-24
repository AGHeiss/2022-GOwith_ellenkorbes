package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println(p.nome, "diz Oi!")

}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {

	p1 := pessoa{"Augusto", 300}
	p1.falar()

	dizerAlgumaCoisa(&p1)

}

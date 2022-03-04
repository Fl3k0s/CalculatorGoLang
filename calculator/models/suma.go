package models

import "container/list"

type Suma struct {
	numeros list.List
}

func (s Suma) NewSuma() {
	s.numeros = *list.New()
}

package mesas

import "fmt"

type IMesa interface {
	getFormato() string
	setFormato(formato string)
	getPreco() float64
	setPreco(preco float64)
	getMaterial() string
	setMaterial(material string)
	MostrarAtributos()
}

type Mesa struct {
	Formato  string
	Preco    float64
	Material string
}

func (m *Mesa) MostrarAtributos() {
	fmt.Printf("Material: %s, Preço: %f, Formato: %s \n", m.getMaterial(), m.getPreco(), m.getFormato())
}

func (m *Mesa) getFormato() string {
	return m.Formato
}

func (m *Mesa) setFormato(formato string) {
	m.Formato = formato
}

func (m *Mesa) getPreco() float64 {
	return m.Preco
}

func (m *Mesa) setPreco(preco float64) {
	m.Preco = preco
}

func (m *Mesa) getMaterial() string {
	return m.Material
}

func (m *Mesa) setMaterial(material string) {
	m.Material = material
}

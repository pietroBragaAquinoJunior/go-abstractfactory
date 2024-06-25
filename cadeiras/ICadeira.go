package cadeiras

import "fmt"

type ICadeira interface {
	getMaterial() string
	setMaterial(material string)
	getPreco() float64
	setPreco(preco float64)
	MostrarAtributos()
}

type Cadeira struct {
	Material string
	Preco    float64
}

func (c *Cadeira) MostrarAtributos() {
	fmt.Printf("Material: %s, Preço: %f.", c.getMaterial(), c.getPreco())
}

func (c *Cadeira) getMaterial() string {
	return c.Material
}

func (c *Cadeira) setMaterial(material string) {
	c.Material = material
}

func (c *Cadeira) getPreco() float64 {
	return c.Preco
}

func (c *Cadeira) setPreco(preco float64) {
	c.Preco = preco
}

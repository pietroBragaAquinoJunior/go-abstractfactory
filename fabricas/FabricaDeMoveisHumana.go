package fabricas

import (
	"abstractlayout/cadeiras"
	"abstractlayout/mesas"
)

type FabricaDeMoveisHumana struct {
}

func (f *FabricaDeMoveisHumana) FabricarCadeira() cadeiras.ICadeira {
	return &cadeiras.Cadeira{Preco: 49.0, Material: "Pau-Brasil"}
}

func (f *FabricaDeMoveisHumana) FabricarMesa() mesas.IMesa {
	return &mesas.Mesa{Material: "Pau-Brasil", Preco: 99.0, Formato: "Quadrado"}
}

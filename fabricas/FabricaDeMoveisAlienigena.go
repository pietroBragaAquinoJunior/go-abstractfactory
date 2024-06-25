package fabricas

import (
	"abstractlayout/cadeiras"
	"abstractlayout/mesas"
)

type FabricaDeMoveisAlienigena struct {
}

func (f *FabricaDeMoveisAlienigena) FabricarCadeira() cadeiras.ICadeira {
	return &cadeiras.Cadeira{Material: "Gosma Intergalática", Preco: 29999.00}
}

func (f *FabricaDeMoveisAlienigena) FabricarMesa() mesas.IMesa {
	return &mesas.Mesa{Material: "Plasma Elétrico", Preco: 55555.90, Formato: "Espiral"}
}

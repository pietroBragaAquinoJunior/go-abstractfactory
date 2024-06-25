package fabricas

import (
	"abstractlayout/cadeiras"
	"abstractlayout/mesas"
	"fmt"
)

type IFabricaDeMoveis interface {
	FabricarCadeira() cadeiras.ICadeira
	FabricarMesa() mesas.IMesa
}

func CriarFabricaDeMoveis(tipo string) (IFabricaDeMoveis, error) {
	if tipo == "alienigena" {
		return &FabricaDeMoveisAlienigena{}, nil
	}
	if tipo == "humana" {
		return &FabricaDeMoveisHumana{}, nil
	}
	return nil, fmt.Errorf("não foi possível criar a fábrica")
}

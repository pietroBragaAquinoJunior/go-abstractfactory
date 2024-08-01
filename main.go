package main

import (
	"abstractlayout/fabricas"
)

func main() {
	fabricaAlienigena, _ := fabricas.CriarFabricaDeMoveis("alienigena")
	fabricaHumana, _ := fabricas.CriarFabricaDeMoveis("humana")

	mesaHumana := fabricaHumana.FabricarMesa()
	mesaAlienigena := fabricaAlienigena.FabricarMesa()

	cadeiraHumana := fabricaHumana.FabricarCadeira()
	cadeiraAlienigena := fabricaAlienigena.FabricarCadeira()

	mesaHumana.MostrarAtributos()
	mesaAlienigena.MostrarAtributos()
	cadeiraHumana.MostrarAtributos()
	cadeiraAlienigena.MostrarAtributos()
}

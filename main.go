package main

import (
	"abstractlayout/fabricas"
	"fmt"
)

func main() {
	fabricaAlienigena, _ := fabricas.CriarFabricaDeMoveis("alienigena")
	fabricaHumana, _ := fabricas.CriarFabricaDeMoveis("humana")

	mesaHumana := fabricaHumana.FabricarMesa()
	mesaAlienigena := fabricaAlienigena.FabricarMesa()

	cadeiraHumana := fabricaHumana.FabricarCadeira()
	cadeiraAlienigena := fabricaAlienigena.FabricarCadeira()

	mesaHumana.MostrarAtributos()
	fmt.Println()
	mesaAlienigena.MostrarAtributos()
	fmt.Println()
	cadeiraHumana.MostrarAtributos()
	fmt.Println()
	cadeiraAlienigena.MostrarAtributos()
}

package sim

import "pokemon/myinterface"

type Battle struct {
	Gen    int
	sides  []*Side
	effect Effect
	Dex    *ModdedDex
	field  *field
}

//检测是否是接触技能
func (this *Battle) CheckMoveMakesContact(move myinterface.IMove, attacker myinterface.IPoke, defender myinterface.IPoke, announcePads bool) bool {
	if move.GetFlags()["contact"] && attacker.HasItem("protectivepads") {
		// if announcePads {
		// 	this.add("-activate", defender, this.effect.GetFullName())
		// 	this.add("-activate", attacker, "item: Protective Pads")
		// }
		return false
	}
	return move.GetFlags()["contact"]
}

func (this *Battle) getAllActive() []*Pokemon {
	pokemonList := []*Pokemon{}
	for _, side := range this.sides {
		for _, pokemon := range side.Active {
			if pokemon != nil && !pokemon.fainted {
				pokemonList = append(pokemonList, pokemon)
			}
		}
	}

	return pokemonList
}

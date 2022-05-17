package sim

import (
	"pokemon/data"
	"pokemon/myinterface"
)

var dex ModdedDex

type ModdedDex struct {
	abilities DexAbilities
	items     DexItem
}

func (this *ModdedDex) GetNewAbility(id string) myinterface.IAbi {
	switch id {
	case "Ironbarbs":
		return data.NewIronbarbs()
	}

	return data.NewBaseAbi()
}

func (this *ModdedDex) GetNewItem(id string) myinterface.IItem {
	// switch id {
	// case "Ironbarbs":
	// 	return data.NewIronbarbs()
	// }

	//TODO BaseItem
	return nil
}

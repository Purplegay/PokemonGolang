package sim

import (
	"pokemon/myinterface"
)

type DexAbilities struct {
	abilityCache map[string]myinterface.IAbi
}

func (this *DexAbilities) getByID(id string) myinterface.IAbi {
	ability := this.abilityCache[id]
	if ability != nil {
		return ability
	}

	//TODO 重写一个data获取脚本

	// if (dex.data.Aliases.hasOwnProperty(id)) {
	// 	ability = this.get(this.dex.data.Aliases[id]);
	// } else if (id && this.dex.data.Abilities.hasOwnProperty(id)) {
	// 	const abilityData = this.dex.data.Abilities[id] as any;
	// 	const abilityTextData = this.dex.getDescs('Abilities', id, abilityData);
	// 	ability = new Ability({
	// 		name: id,
	// 		...abilityData,
	// 		...abilityTextData,
	// 	});
	// 	if (ability.gen > this.dex.gen) {
	// 		(ability as any).isNonstandard = 'Future';
	// 	}
	// 	if (this.dex.currentMod === 'gen7letsgo' && ability.id !== 'noability') {
	// 		(ability as any).isNonstandard = 'Past';
	// 	}
	// 	if ((this.dex.currentMod === 'gen7letsgo' || this.dex.gen <= 2) && ability.id === 'noability') {
	// 		(ability as any).isNonstandard = null;
	// 	}
	// } else {
	// 	ability = new Ability({
	// 		id, name: id, exists: false,
	// 	});
	// }

	// if (ability.exists) this.abilityCache.set(id, ability);
	// return ability;
	return nil
}

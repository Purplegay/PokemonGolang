package sim

import (
	"pokemon/myinterface"
)

type Pokemon struct {
	battle      Battle
	isActive    bool
	fainted     bool
	transformed bool
	ability     string
	item        string
	volatiles   map[string]interface{}
	m           *m

	abilityState map[string]interface{}
	itemState    map[string]interface{}
}

type m struct {
	gluttonyFlag    bool   // Gen-NEXT
	innate          string // Partners in Crime
	originalSpecies string // Mix and Mega
	abils           map[string]bool
}

func (this *Pokemon) HasItem(items ...string) bool {
	if this.IgnoringItem() {
		return false
	}
	ownItem := this.item
	for _, item := range items {
		if ownItem == item {
			return true
		}
	}
	return false
}

func (this *Pokemon) IgnoringItem() bool {
	condtion1 := this.getItemStateKnockedOff()
	conditon2 := this.battle.Gen >= 5 && !this.isActive
	//klutz 笨拙：携带物品的能力失效（包括负面效果）
	conditon3 := this.hasAbility("klutz") && !this.getItem().IgnoreKlutz()

	conditon4 := this.getVolatilesEmbargo()
	conditon5 := this.getMagicroom()
	return (condtion1) || conditon2 || conditon3 || conditon4 || conditon5
}

func (this *Pokemon) hasAbility(ability ...string) bool {
	if this.ignoringAbility() {
		return false
	}
	ownAbility := this.ability
	for _, abability := range ability {
		if ownAbility == abability {
			return true
		}
	}
	return false

}

func (this *Pokemon) ignoringAbility() bool {
	// Check if any active pokemon have the ability Neutralizing Gas
	neutralizinggas := false
	for _, pokemon := range this.battle.getAllActive() {
		// can't use hasAbility because it would lead to infinite recursion
		//neutralizinggas 化学变化气体：此特性不仅能消除周围宝可梦的特性效果，还能让特性的效果变得无法发动。

		condtion1 := pokemon.getAbilityStateEnding()
		if pokemon.ability == "neutralizinggas" && !pokemon.getVolatilesGastroacid() &&
			!pokemon.transformed && condtion1 {
			neutralizinggas = true
			break
		}
	}
	conditon1 := this.battle.Gen >= 5 && !this.isActive
	conditon2 := (this.getVolatilesGastroacid() || (neutralizinggas && (this.ability != "neutralizinggas" || this.m.abils["ability:neutralizinggas"])) && !this.getAbility().IsPermanent())

	return conditon1 || conditon2

}

func (this *Pokemon) getAbility() myinterface.IAbi {
	return this.battle.Dex.abilities.getByID(this.ability)
}

func (this *Pokemon) getItem() myinterface.IItem {
	return this.battle.Dex.items.getByID(this.ability)
}

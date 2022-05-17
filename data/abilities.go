package data

import (
	"pokemon/myinterface"
)

type BaseAbi struct {
	IsPer bool
}

func NewBaseAbi() myinterface.IAbi {
	return &BaseAbi{}
}

func (this *BaseAbi) IsPermanent() bool {
	return this.IsPer
}

func (this *BaseAbi) onDamagingHit(battle myinterface.IBattle, damage int32, target myinterface.IPoke, source myinterface.IPoke, move myinterface.IMove) {

}

type Ironbarbs struct {
	*BaseAbi
}

func NewIronbarbs() myinterface.IAbi {
	return &Ironbarbs{BaseAbi: &BaseAbi{}}
}

//拥有铁刺特性的宝可梦受到近身攻击时，攻击方损失HP最大值的1/8。
// 连续攻击技的每次攻击都将发动铁刺。
// 当攻击方受到铁刺的伤害而濒死时，以下招式的特效不会发动：
// 龙尾、仰投、高速旋转、虫食、啄食、清醒、清醒拍打。
func (this *Ironbarbs) onDamagingHit(battle myinterface.IBattle, damage int32, target myinterface.IPoke, source myinterface.IPoke, move myinterface.IMove) {
	if battle.CheckMoveMakesContact(move, source, target, true) {
		//this.damage(source.baseMaxhp/8, source, target)
	}
}

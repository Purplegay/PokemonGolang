package sim

//魔法空间Magicroom：包括施放回合在内，魔法空间存在5回合，期间场上所有宝可梦的道具没有效果。
func (this *Pokemon) getMagicroom() bool {
	if v, ok := this.battle.field.pseudoWeather["magicroom"]; ok {
		if realV, ok2 := v.(bool); ok2 {
			return realV
		}

	}
	return false
}

package sim

func (this *Pokemon) getMagicroom() bool {
	if v, ok := this.battle.field.pseudoWeather["magicroom"]; ok {
		if realV, ok2 := v.(bool); ok2 {
			return realV
		}

	}
	return false
}

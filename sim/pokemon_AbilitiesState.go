package sim

func (this *Pokemon) getAbilityStateEnding() bool {
	if v, ok := this.itemState["ending"]; ok {
		if realV, ok2 := v.(bool); ok2 {
			return realV
		}

	}
	return false
}

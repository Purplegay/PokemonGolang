package sim

func (this *Pokemon) getAbilityStateKnockedOff() bool {
	if v, ok := this.itemState["knockedOff"]; ok {
		if realV, ok2 := v.(bool); ok2 {
			return realV
		}

	}
	return false
}

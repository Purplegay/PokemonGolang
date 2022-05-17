package sim

func (this *Pokemon) getVolatilesEmbargo() bool {
	if v, ok := this.volatiles["embargo"]; ok {
		if realV, ok2 := v.(bool); ok2 {
			return realV
		}

	}
	return false
}

func (this *Pokemon) getVolatilesGastroacid() bool {
	if v, ok := this.volatiles["gastroacid"]; ok {
		if realV, ok2 := v.(bool); ok2 {
			return realV
		}

	}
	return false
}

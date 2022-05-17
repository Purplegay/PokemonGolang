package sim

//embargo 扣押阻止对手使用携带道具。
func (this *Pokemon) getVolatilesEmbargo() bool {
	if v, ok := this.volatiles["embargo"]; ok {
		if realV, ok2 := v.(bool); ok2 {
			return realV
		}

	}
	return false
}

//gastroacid 胃液：使目标陷入无特性状态
func (this *Pokemon) getVolatilesGastroacid() bool {
	if v, ok := this.volatiles["gastroacid"]; ok {
		if realV, ok2 := v.(bool); ok2 {
			return realV
		}

	}
	return false
}

package sim

type ActiveMove struct {
	Flags map[string]bool
}

func (this *ActiveMove) GetFlags() map[string]bool {
	return this.Flags
}

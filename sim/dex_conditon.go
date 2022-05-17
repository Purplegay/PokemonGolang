package sim

type EventMethods interface {
	onDamagingHit(battle *Battle, damage int32, target *Pokemon, source *Pokemon, move *ActiveMove)
}

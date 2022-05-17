package myinterface

type IBattle interface {
	CheckMoveMakesContact(move IMove, attacker IPoke, defender IPoke, announcePads bool) bool
}

package cards

// RankOrder is an iterface that is used compare two ranks
type RankOrder interface {
	Len() int
	Less(i int, j int) bool
	Swap(i int, j int)
}

// Ranks is a slice of Rank objects
type Ranks []Rank

// Len returns the number of Rank objects in the array.
func (rs Ranks) Len() int {
	return len(rs)
}

// Less compares two Rank objects and returns true if the first
// one is less than the second one.
func (rs Ranks) Less(i int, j int) bool {
	return rs[i] < rs[j]
}

// Exchanges two Rank objects in the array.
func (rs Ranks) Swap(i int, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

// PinochleRanks is a slice of Rank objects in a Pinochle deck
type PinochleRanks []Rank

// Less compares two Rank objects in a Pinochle deck and returns true
// if the first one is less than the second one.
func (prs PinochleRanks) Less(i int, j int) bool {
	var rankI, rankJ int
	rankI = pinochleRank(prs[i])
	rankJ = pinochleRank(prs[j])
	return rankI < rankJ
}

func pinochleRank(rank Rank) int {
	switch rank {
	case NINE, JACK, QUEEN, KING:
		return int(rank)
	case TEN:
		return int(KING) + 1
	case ACE:
		return int(KING) + 2
	default:
		return -1
	}
}

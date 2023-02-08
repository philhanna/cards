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

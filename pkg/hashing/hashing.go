package hashing

// HashFunction is a type alias for a function that takes a byte slice and returns a uint64.
type HashFunction func([]byte) uint64

// XorDjb2Hash XORes the byte slice and calc its djb2 hash.
func XorDjb2Hash(s []byte) uint64

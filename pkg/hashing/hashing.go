package hashing

// HashFunction is a type alias for a function that takes a byte slice and returns a uint64.
type HashFunction func([]byte) uint64

// DJB2 is an implementation of the djb2 hash function. Ref: http://www.cse.yorku.ca/~oz/hash.html
func DJB2(s []byte) uint64 {
	var hash uint64 = 5381
	for _, c := range s {
		hash = ((hash << 5) + hash) + uint64(c)
	}
	return hash
}

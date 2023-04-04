package hashing

type djb2 struct{}

var _ Hasher = (*djb2)(nil)

// NewDjb2 returns a new djb2 hasher.
// Algorithm taken from http://www.cse.yorku.ca/~oz/hash.html
func NewDjb2() Hasher {
	return &djb2{}
}

// HashByteString hashes a byte string using the djb2 algorithm.
func (d *djb2) HashByteString(s []byte) int64 {
	var hash int64 = 5381
	for _, c := range s {
		hash = ((hash << 5) + hash) + int64(c)
	}
	return hash
}

// HashString hashes a string using the djb2 algorithm.
func (d *djb2) HashString(s string) int64 {
	return d.HashByteString([]byte(s))
}

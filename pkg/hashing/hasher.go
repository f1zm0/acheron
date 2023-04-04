package hashing

type Hasher interface {
	// HashString hashes a string using the djb2 algorithm.
	HashString(s string) int64

	// HashByteString hashes a byte string using the djb2 algorithm.
	HashByteString(s []byte) int64
}

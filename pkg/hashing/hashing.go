package hashing

type HashFunction func([]byte) int64

// DJB2 hashing algorithm; ref: http://www.cse.yorku.ca/~oz/hash.html
func DJB2(s []byte) int64 {
	var hash int64 = 5381
	for _, c := range s {
		hash = ((hash << 5) + hash) + int64(c)
	}
	return hash
}

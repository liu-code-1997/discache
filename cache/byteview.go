package cache

type ByteView struct {
	b []byte
}

// return view's length
func (v ByteView) Len() int {
	return len(v.b)
}

// return a copy of the data
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// return the data as a string
func (v ByteView) String() string {
	return string(v.b)
}
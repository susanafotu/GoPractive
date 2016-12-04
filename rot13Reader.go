func (r13 rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r13.r.Read(b)
	for i := range b {
		b[i] = rot13(b[i])
	}
	return 
}


package polynormal

var poly258 = [...]int64{
	0x11931,
	0x12109,
	0x13559,
	0x15935,
	0x16fed,
	0x1755d,
	0x17bbd,
	0x18103,
	0x1a38b,
	0x1ad6b,
	0x1b7db,
	0x1c107,
	0x1d557,
	0x1e38f,
	0x1ed6f,
	0x1f93f}

var P258 []*Prime

func init() {
	P258 = make([]*Prime, len(poly258))
	for i, p := range poly258 {
		P258[i] = NewPrime(p, false)
	}
}

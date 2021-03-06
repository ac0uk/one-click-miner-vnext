package pools

type Pool interface {
	GetPendingPayout() uint64
	GetStratumUrl() string
	GetUsername() string
	GetPassword() string
	GetName() string
	GetID() int
}

func GetPools(addr string, testnet bool) []Pool {
	if testnet {
		return []Pool{
			NewP2Pool(addr),
		}
	}
	return []Pool{
		NewHashalot(addr),
		//NewSuprnova(addr),
		NewP2Pool(addr),
	}
}

func GetPool(pool int, addr string, testnet bool) Pool {
	pools := GetPools(addr, testnet)
	for _, p := range pools {
		if p.GetID() == pool {
			return p
		}
	}
	return pools[0]
}

package types

const (
	// ModuleName defines the module name
	ModuleName = "hackathonaccountrecovery"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_hackathonaccountrecovery"
	MemCountKey = "cnt_hackathonaccountrecovery"
)

var (
	ParamsKey = []byte("p_hackathonaccountrecovery")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

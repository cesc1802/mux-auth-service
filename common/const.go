package common

const (
	DbTypePet    = 1
	DbTypeUser   = 2
	DbTypeNFTPet = 3
	DbTypeTx     = 4

	CurrentUser = "user"
)

const DateTimeFmt = "2006-01-02 15:04:05.999999"

type Requester interface {
	GetUserId() int
}

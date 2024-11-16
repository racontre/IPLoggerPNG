package iplogger

type IPLoggerService interface {
	InsertIP(ip string) error
	GetIPList(count int) ([]string, error)
}
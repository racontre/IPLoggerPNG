package iplogger

type InmemoryLoggerService struct {
	Ips []string
}

func (m *InmemoryLoggerService) InsertIP(ip string) error {
	m.Ips = append(m.Ips, ip)
	return nil
}

func (m *InmemoryLoggerService) GetIPList(count int) ([]string, error) {
	if len(m.Ips) > count {
		return m.Ips[len(m.Ips)-count:], nil
	}
	return m.Ips, nil
}
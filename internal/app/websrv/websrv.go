package websrv

type WEBsrv struct {
	config *Config
}

func New() *WEBsrv {
	return &WEBsrv{
		config: config,
	}
}
func (s *WEBsrv) Start() error {
	return nil
}

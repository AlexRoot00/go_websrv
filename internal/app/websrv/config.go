package websrv

//config
type Config struct{
	BindAddr string 'toml:"bind_addr"'
}
//newconfig
func NewConfig *Config{
	return &Config {
		BindAddr:"*:8080",
	}
}
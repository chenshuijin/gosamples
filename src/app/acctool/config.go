package main

type DBConf struct {
	DriverName   string `json:"drivername,omitempty"`
	DataSource   string `json:"datasource,omitempty"`
	MaxIdleConns int    `json:"maxidleconns,omitempty"`
	MaxOpenConns int    `json:"maxopenconns,omitempty"`
	Enable       bool   `json:"enable,omitempty"`
}

type Config struct {
	Pprof struct {
		Enable bool   `json:"enable,omitempty"`
		Host   string `json:"host,omitempty"`
	}
	DBs map[string]DBConf `json:"dbs,omitempty"`
	Log struct {
		Level  string `json:"level"`
		Path   string `json:"path"`
		File   string `json:"file"`
		Format string `json:"format"`
	}
	KeyStore string `json:"keystore"`
}

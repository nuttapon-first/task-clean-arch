package configs

type Configs struct {
	App      Gin
	Database SQL
}

type Gin struct {
	Host string
	Port string
}

type SQL struct {
	DriverName string
	DbName     string
}

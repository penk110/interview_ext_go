package conf

var defaultConf *Conf

type Conf struct {
	Storage *Storage
	Casbin  *Casbin
}

type Storage struct {
	Driver string `json:"driver"`
	DSN    string
}

type Casbin struct {
	ResourcePath string `yaml:"resourcePath"`
}

func Init(cfgPath string) (*Conf, error) {
	// load from yaml

	defaultConf = &Conf{
		Storage: &Storage{
			Driver: "mysql",
			DSN:    "dev:Dev@3306@tcp(192.168.3.11:3306)/casbin?charset=utf8mb4&parseTime=True&loc=Local",
		},
		Casbin: &Casbin{
			ResourcePath: "./conf/resources",
		},
	}

	return defaultConf, nil
}

func DefaultConf() *Conf {
	return defaultConf
}

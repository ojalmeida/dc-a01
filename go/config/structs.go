package config

type Configuration struct {
	Performance struct {
		MaxNumberOfWorkers int `yaml:"max-number-of-workers"`
	} `yaml:"performance"`

	Server struct {
		Addr string
		Port int
	}

	Network struct {
		DB struct {
			User         string `yaml:"user"`
			Password     string `yaml:"password"`
			Host         string `yaml:"host"`
			Port         int    `yaml:"port"`
			DatabaseName string `yaml:"database-name"`
		} `yaml:"db"`
	} `yaml:"network"`

	Timeout struct {
		Read  int `yaml:"read"`
		Write int `yaml:"write"`
		Idle  int `yaml:"idle"`
		Wait  int `yaml:"wait"`
	} `yaml:"timeout"`
}

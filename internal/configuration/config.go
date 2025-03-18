package configuration

type Config struct {
	Logging *LoggingConfig `yaml:"logging"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Output string `yaml:"output"`
}

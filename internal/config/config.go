package config

type Config struct {
	Vault   VaultConfig   `yaml:"vault"`
	Project ProjectConfig `yaml:"project"`
	LLM     LLMConfig     `yaml:"llm"`
	Logging LoggingConfig `yaml:"logging"`
}

type VaultConfig struct {
	Path string `yaml:"path"`
}

type ProjectConfig struct {
	Path string `yaml:"path"`
}
type LLMConfig struct {
	Provider string `yaml:"provider"`
	Model    string `yaml:"model"`
}

type LoggingConfig struct {
	Level string `yaml:"level"`
}

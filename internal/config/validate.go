package config

import "fmt"

func (c *Config) Validate() error {
	if c.Project.Path == "" {
		return fmt.Errorf("Project.path is required")
	}
	if c.Vault.Path == "" {
		return fmt.Errorf("Vault.path is required")
	}
	if c.LLM.Provider == "" {
		return fmt.Errorf("llm.provider is required")
	}
	return nil
}
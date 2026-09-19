package once

import (
	"sync"
)

type Config struct {
	Host string
	Port int
}

type ConfigLoader struct {
	o         sync.Once
	config    Config
	loadCount int
}

func NewConfigLoader() *ConfigLoader {
	return &ConfigLoader{}
}

func (l *ConfigLoader) Load() *Config {
	l.o.Do(func() {
		l.config = Config{Host: "localhost", Port: 8080}
		l.loadCount++
	})
	return &l.config
}

func (l *ConfigLoader) LoadCount() int {
	return l.loadCount
}

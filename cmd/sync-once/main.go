package main

import "sync"

type ConfigLoader struct {
	once   sync.Once
	values map[string]string
	err    error
}

func (c *ConfigLoader) Load() (error, map[string]string) {
	c.once.Do(func() {
		//if map key not set set to one, otherwise increment.
	})
	//result := make(map[string]string)

	return nil, c.values
}

package conf

func (c *Conf) HandleShowDatabases() []interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	dbs := make([]interface{}, 0, len(c.schemas))
	for key := range c.schemas {
		dbs = append(dbs, key)
	}

	return dbs
}

package vici

// Stats returns IKE daemon statistics and load information.
func (c *ViciClient) Stats() (msg map[string]interface{}, err error) {
	msg, err = c.Request("stats", nil)
	return
}

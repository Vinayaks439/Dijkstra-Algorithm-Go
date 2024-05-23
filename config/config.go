package config

type Config struct {
	NodeStart string
	NodeEnd   string
	Cost      int
}

type ConfigFile struct {
	Links []struct {
		From string `json:"from"`
		To   string `json:"to"`
		Cost int    `json:"cost"`
	} `json:"links"`
}

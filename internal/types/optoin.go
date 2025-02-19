package types

import "time"

type Options struct {
	Webserver WebserverOpts `yaml:"webserver"`
	Consumer  ConsumerOpts  `yaml:"consumer"`
}
type WebserverOpts struct {
	Address string `yaml:"address"`
}
type ConsumerOpts struct {
	Topics           []string      `yaml:"topics"`
	BootstrapServers string        `yaml:"bootstrap_servers"`
	Group            string        `yaml:"group"`
	AutoOffsetReset  string        `yaml:"auto_offset_reset"`
	EnableAutoCommit bool          `yaml:"enable_auto_commit"`
	PollTimeout      time.Duration `yaml:"poll_timeout"`
	SessionTimeoutMs int           `yaml:"session_timeout_ms"`
	FetchMinBytes    int           `yaml:"fetch_min_bytes"`
}

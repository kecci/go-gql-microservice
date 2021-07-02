package config

import (
	"time"
)

// Config struct define configuration service
type Config struct {
	Server     Server           `yaml:"server"`
	Database   DBConfig         `yaml:"database"`
	Redis      RedisConfig      `yaml:"redis"`
	APIBaseURL APIBaseURLConfig `yaml:"api_base_url"`
}

// RedisConfig struct define configuration Redis
type RedisConfig struct {
	Engine string `json:"engine" yaml:"engine"`
	// Redis server address
	Address string `json:"address" yaml:"address"`

	// Maximum number of idle connections in the pool.
	// Only for redigo engine
	MaxIdle int `json:"maxidle" yaml:"maxidle"`

	// Maximum number of connections allocated by the pool at a given time.
	// When zero, there is no limit on the number of connections in the pool.
	MaxActive int `json:"maxactive" yaml:"maxactive" default:"50"`

	// Close connections after remaining idle for this duration. If the value
	// is zero, then idle connections are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	Timeout int `json:"timeout" yaml:"timeout"`

	// IdlePingPeriod defines period in seconds after which the connection need
	// to be check(PING) before being used.
	// Default is 10 seconds
	IdlePingPeriod int `json:"idle_ping_period" yaml:"idle_ping_period" default:"10"`

	// pool wait time in millisecond.
	// If > 0 and the pool is at the MaxActive limit, then Get() waits
	// for a connection to be returned to the pool before returning. It only waits
	// for PoolWaitMs millisecond
	PoolWaitMs int `json:"pool_wait_ms" yaml:"pool_wait_ms" default:"1000"`

	// NoPingOnCreate is a flag to indicate whether it will be do ping check on `New` or not.
	// If true: client will do redis PING on `New`, make sure that the server is up.
	NoPingOnCreate bool `json:"no_ping_on_create" yaml:"no_ping_on_create"`

	// Optional password. Must match the password specified in the
	// requirepass server configuration option.
	// THIS OPTION IS CURRENTLY ONLY AVAILABLE FOR goredis engine
	Password string

	// Database to be selected after connecting to the server.
	// THIS OPTION IS CURRENTLY ONLY AVAILABLE FOR goredis engine
	DB int
}

// DBConfig struct define configuration database
type DBConfig struct {
	Driver                string        `json:"driver" yaml:"driver"`
	MasterDSN             string        `json:"master" yaml:"master"`
	FollowerDSN           string        `json:"follower" yaml:"follower"`
	MaxOpenConnections    int           `json:"max_open_conns" yaml:"max_open_conns"`
	MaxIdleConnections    int           `json:"max_idle_conns" yaml:"max_idle_conns"`
	ConnectionMaxLifetime time.Duration `json:"conn_max_lifetime" yaml:"conn_max_lifetime"`

	// number of retry during Connect
	// won't be used if `NoPingOnOpen`=true
	Retry int `json:"retry" yaml:"retry"`

	// no Ping when openning DB connection, useful if we don't care whether the server is up or not
	NoPingOnOpen bool `json:"no_ping_on_open" yaml:"no_ping_on_open"`
}

type Server struct {
	GQL ServerGQL `yaml:"gql"`
}

type ServerGQL struct {
	Address        string `yaml:"address"`
	WriteTimeout   int    `yaml:"write_timeout"`
	ReadTimeout    int    `yaml:"read_timeout"`
	IdleTimeout    int    `yaml:"idle_timeout"`
	MaxHeaderBytes int    `yaml:"max_header_bytes"`
}

type APIBaseURLConfig struct {
	SpecialSite     string `yaml:"special_site"`
	CampaignService string `yaml:"campaign_service"`
	FlashSale       string `yaml:"flash_sale"`
}

package configs

import "time"

type JwtConfig struct {
	Secret       string
	TimeToExpire time.Duration
}

package sslcheck

import "time"

type Result struct {
	Address string
	Valid   bool
	Error   string
	Expiry  time.Time
}

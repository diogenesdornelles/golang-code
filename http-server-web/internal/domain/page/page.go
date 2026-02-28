package page

import "time"

type HomePage struct {
	Title       string
	Message     string
	DataCriacao time.Time
	IsSignedIn  bool
}

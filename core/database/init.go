package database

import (
	"sync"
)

var (
	// once once
	once sync.Once
	// err err
	err error
)

// InitAll InitAll
func InitAll() {
	once.Do(func() {
		InitMysql()
	})
}

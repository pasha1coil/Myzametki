package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Zametki struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}

type Users struct {
	IDU      int
	Username string
}

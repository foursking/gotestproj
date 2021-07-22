package dao

import "context"

// Dao dao interface
type Dao interface {
	Ping(context.Context) error
	Close() error
}

type dao struct {
}

// New creates dao
func New() (d Dao, err error) {
	d = &dao{}
	return
}

// Ping ping dao
func (d *dao) Ping(ctx context.Context) error {
	return nil
}

// Close close dao
func (d *dao) Close() error {
	return nil
}

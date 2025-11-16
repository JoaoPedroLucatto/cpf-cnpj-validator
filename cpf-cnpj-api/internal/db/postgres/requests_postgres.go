package postgres

import "sync/atomic"

func (pg *Postgres) IncrementRequest() {
	atomic.AddInt64(&pg.requestCount, 1)
}

func (pg *Postgres) CountRequests() (int, error) {
	return int(atomic.LoadInt64(&pg.requestCount)), nil
}

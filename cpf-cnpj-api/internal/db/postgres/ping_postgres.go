package postgres

func (db *Postgres) PingDatabase() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	return nil
}

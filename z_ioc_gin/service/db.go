package service

type DB struct {
	DSN string
}

func NewDB() *DB {
	return &DB{
		DSN: "mysql",
	}
}

func (db *DB) GetDSN() string {
	return db.DSN
}

package database

func DbConnection(dsn string) string {
	if dsn != "" {
		return "connecting to the database " + dsn
	}

	return "error connecting to the database" + dsn
}

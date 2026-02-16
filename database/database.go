package database

var connection string

// init ini pasti di eksekusi saat diimpor
func init()  {
	connection = "postgreSQL"
}

func GetDatabase()string{
	return connection
}
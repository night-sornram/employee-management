package common_utils

type database struct {
	host     string
	port     int
	user     string
	pass     string
	dbname   string
	timeZone string
}

func initDatabase(port string) database {
	var dbname string
	switch port {
	case "8080":
		dbname = "employee"
	case "8081":
		dbname = "attendance"
	case "8082":
		dbname = "leave"
	case "8083":
		dbname = "notification"
	}

	return database{
		host:     "localhost",
		port:     5432,
		user:     "postgres",
		pass:     "password",
		dbname:   dbname,
		timeZone: "Asia/Bangkok",
	}
}

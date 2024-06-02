package common_utils

import (
	"fmt"
)

func ConnectDB(port string) string {
	config := initDatabase(port)
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s TimeZone=%s sslmode=disable", config.host, config.port, config.user, config.pass, config.dbname, config.timeZone)
}

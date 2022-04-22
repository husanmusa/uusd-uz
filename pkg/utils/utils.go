package utils

import (
	"database/sql"
	"time"
)

func StringToNullTime(s string) (nt sql.NullTime) {
	if s != "" {
		nt.Valid = true
		nt.Time, _ = time.Parse("2006-01-02", s)

		return nt
	}

	return nt
}

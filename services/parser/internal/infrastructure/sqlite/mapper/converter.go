package sqlitemap

import "database/sql"

func toNullString(value string) sql.NullString {
	if value == "" {
		return sql.NullString{Valid: false}
	}

	return sql.NullString{String: value, Valid: true}
}

func toNullFloat64(value float64) sql.NullFloat64 {
	if value == 0 {
		return sql.NullFloat64{Valid: false}
	}

	return sql.NullFloat64{Float64: value, Valid: true}
}

func toString(value sql.NullString) string {
	if value.Valid {
		return value.String
	}

	return ""
}

func toFloat64(value sql.NullFloat64) float64 {
	if value.Valid {
		return value.Float64
	}

	return 0
}

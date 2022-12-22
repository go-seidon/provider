package typeconv

import (
	"database/sql"
	"time"
)

func SqlBool(i sql.NullBool) *bool {
	if !i.Valid {
		return nil
	}
	return Bool(i.Bool)
}

func SqlBoolVal(i *bool) sql.NullBool {
	if i == nil {
		return sql.NullBool{}
	}
	return sql.NullBool{
		Bool:  BoolVal(i),
		Valid: true,
	}
}

func SqlFloat64(i sql.NullFloat64) *float64 {
	if !i.Valid {
		return nil
	}
	return Float64(i.Float64)
}

func SqlFloat64Val(i *float64) sql.NullFloat64 {
	if i == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{
		Float64: Float64Val(i),
		Valid:   true,
	}
}

func SqlInt32(i sql.NullInt32) *int32 {
	if !i.Valid {
		return nil
	}
	return Int32(i.Int32)
}

func SqlInt32Val(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: Int32Val(i),
		Valid: true,
	}
}

func SqlInt64(i sql.NullInt64) *int64 {
	if !i.Valid {
		return nil
	}
	return Int64(i.Int64)
}

func SqlInt64Val(i *int64) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: Int64Val(i),
		Valid: true,
	}
}

func SqlString(i sql.NullString) *string {
	if !i.Valid {
		return nil
	}
	return String(i.String)
}

func SqlStringVal(i *string) sql.NullString {
	if i == nil {
		return sql.NullString{}
	}
	return sql.NullString{
		String: StringVal(i),
		Valid:  true,
	}
}

func SqlTime(i sql.NullTime) *time.Time {
	if !i.Valid {
		return nil
	}
	return Time(i.Time)
}

func SqlTimeVal(i *time.Time) sql.NullTime {
	if i == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  TimeVal(i),
		Valid: true,
	}
}

func SqlUnixMilli(i sql.NullInt64) *time.Time {
	if !i.Valid {
		return nil
	}
	return Time(time.UnixMilli(i.Int64))
}

func SqlTimeMilli(i *time.Time) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: i.UnixMilli(),
		Valid: true,
	}
}

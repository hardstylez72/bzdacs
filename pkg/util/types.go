package util

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Timestamp struct {
	time.Time
}

///implement encoding.JSON.Marshaler interface
func (v *Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String)
}

func (v *Timestamp) UnmarshalJSON(data []byte) error {
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	return nil
}

///

type JsonNullString struct {
	sql.NullString
}

///implement encoding.JSON.Marshaler interface
func (v *JsonNullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullString) UnmarshalJSON(data []byte) error {
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}
	return nil
}

///

type JsonNullTime struct {
	sql.NullTime
}

///implement encoding.JSON.Marshaler interface
func (v *JsonNullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullTime) UnmarshalJSON(data []byte) error {
	var x *time.Time
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Time = *x
	} else {
		v.Valid = false
	}
	return nil
}

///

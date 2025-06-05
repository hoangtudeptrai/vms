package model

import (
	"database/sql/driver"
	"encoding/json"
)

/********* Keyvalue ***********/
func (sla *User) Scan(src interface{}) error {
	// For gorm:"type:jsonb"
	return json.Unmarshal(src.([]byte), &sla)
	// For gorm:"type:text"
	//return json.Unmarshal([]byte(src.(string)), &sla)
}

func (sla User) Value() (driver.Value, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}

/********* Keyvalue ***********/
func (sla *Assignment) Scan(src interface{}) error {
	// For gorm:"type:jsonb"
	return json.Unmarshal(src.([]byte), &sla)
	// For gorm:"type:text"
	//return json.Unmarshal([]byte(src.(string)), &sla)
}

func (sla Assignment) Value() (driver.Value, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}

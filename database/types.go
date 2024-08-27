package database

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
)

// SliceString GORM NOT SUPPORT SLICE, so custom.
type SliceString []string

// UserModel UserModel-Only User can be added into, not accept channel user.
type UserModel struct {
	gorm.Model
	UserID   int64  `gorm:"primaryKey;<-:create"` // userid cannot change.
	Username string // username may not exist yet, maybe more than one,
	// but no one sure if @username is the really only one.
	FullName    string      // User Full name + LastName + FirstName.
	FromChannel SliceString // what you are active in bot's service.
	UserPhoto   int64       // userPhoto shown as a file id, no required to save a raw file, get it from Telegram Official.
}

// Scan SLiceString Default Scanner / Receiver.
func (t *SliceString) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

// Value SLiceString Default Scanner / Receiver.
func (t SliceString) Value() (driver.Value, error) {
	return json.Marshal(t)
}

package models

type User struct {
	ID        string `gorm:"type:varchar(36);primaryKey"` // UUID almacenado como string
	FirstName string `gorm:"size:100;not null"`
	LastName  string `gorm:"size:100;not null"`
	BirthDate string `gorm:"size:20;not null"`
	Gender    string `gorm:"size:20;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"size:20;not null;default:customer"`
	Points    int    `gorm:"default:10"`
	Level     int    `gorm:"default:1"`
}

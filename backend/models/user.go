package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"` // Identificador único del usuario
	FirstName string    `gorm:"size:100;not null"`                               // Nombre del usuario (campo obligatorio)
	LastName  string    `gorm:"size:100;not null"`                               // Apellidos del usuario (campo obligatorio)
	BirthDate string    `gorm:"size:20;not null"`
	Gender    string    `gorm:"size:20;not null"`
	Email     string    `gorm:"size:100;unique;not null"`          // Correo electrónico (único y obligatorio)
	Password  string    `gorm:"not null"`                          // Contraseña hasheada (campo obligatorio)
	Role      string    `gorm:"size:20;not null;default:customer"` // Rol del usuario (ej. 'manager' o 'customer')
	Points    int       `gorm:"default:0"`
	Level     int       `gorm:"default:1"` // Nivel de experiencia del usuario

}

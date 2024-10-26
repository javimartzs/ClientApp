package models

import (
	"time"

	"github.com/google/uuid"
)

// Promotion representa una promoción en el sistema.
// Algunas promociones requieren un nivel mínimo de experiencia para ser accesibles.
type Promotion struct {
	ID            uint       `gorm:"primaryKey"`        // Identificador único de la promoción
	Title         string     `gorm:"size:100;not null"` // Título de la promoción
	Description   string     `gorm:"size:255"`          // Descripción de la promoción
	RequiredLevel int        `gorm:"default:1"`         // Nivel mínimo necesario para acceder a la promoción
	StartDate     time.Time  `gorm:"not null"`          // Fecha de inicio de la promoción
	EndDate       time.Time  `gorm:"not null"`          // Fecha de fin de la promoción
	UserID        *uuid.UUID `gorm:"type:uuid"`         // ID opcional del usuario al que está dirigida la promoción
}

package models

import (
	"time"

	"github.com/google/uuid"
)

type PointsTransaction struct {
	ID        uint      `gorm:"primaryKey"`     // Identificador único de la transacción
	UserID    uuid.UUID `gorm:"type:uuid"`      // ID del usuario que recibió los puntos
	Points    int       `gorm:"not null"`       // Cantidad de puntos asignados
	CreatedAt time.Time `gorm:"autoCreateTime"` // Fecha en la que se realizó la transacción
}

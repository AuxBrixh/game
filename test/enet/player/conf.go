package player

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/google/uuid"
)

type Config struct {
	Uuid uuid.UUID
	Username string
	Name string
	Position mgl64.Vec2
	Speed int8
	Health, Maxhealth int16
}
package repository

import (
	"github.com/google/uuid"
	"github.com/lipzy13/dakas-backend.git/internal/domain"
	"gorm.io/gorm"
)

type GerobakRepository interface {
	CreateGerobak(gerobak *domain.Gerobak) error
	GetGerobakByID(id uuid.UUID) (*domain.Gerobak, error)
	GetAllGerobaks() ([]domain.Gerobak, error)
}

type gerobakRepository struct {
	db *gorm.DB
}

func NewGerobakRepository(db *gorm.DB) GerobakRepository {
	return &gerobakRepository{db: db}
}

func (g *gerobakRepository) CreateGerobak(gerobak *domain.Gerobak) error {
	gerobak.ID = uuid.New()
	return g.db.Create(gerobak).Error
}

func (g *gerobakRepository) GetGerobakByID(id uuid.UUID) (*domain.Gerobak, error) {
	var gerobak domain.Gerobak
	if err := g.db.First(&gerobak, "gerobak_id=?", id).Error; err != nil {
		return nil, err
	}
	return &gerobak, nil
}

func (g *gerobakRepository) GetAllGerobaks() ([]domain.Gerobak, error) {
	var gerobaks []domain.Gerobak
	if err := g.db.Find(&gerobaks).Error; err != nil {
		return nil, err
	}
	return gerobaks, nil
}

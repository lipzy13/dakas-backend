package service

import (
	"github.com/google/uuid"
	"github.com/lipzy13/dakas-backend.git/internal/domain"
	"github.com/lipzy13/dakas-backend.git/internal/repository"
)

type GerobakService interface {
	CreateGerobak(gerobak *domain.Gerobak) error
	GetGerobakById(id uuid.UUID) (*domain.Gerobak, error)
	GetAllGerobaks() ([]domain.Gerobak, error)
}

type gerobakService struct {
	repo repository.GerobakRepository
}

func NewGerobakService(repo repository.GerobakRepository) GerobakService {
	return &gerobakService{repo: repo}
}

func (g *gerobakService) CreateGerobak(gerobak *domain.Gerobak) error {
	return g.repo.CreateGerobak(gerobak)
}

func (g *gerobakService) GetGerobakById(id uuid.UUID) (*domain.Gerobak, error) {
	return g.repo.GetGerobakByID(id)
}

func (g *gerobakService) GetAllGerobaks() ([]domain.Gerobak, error) {
	return g.repo.GetAllGerobaks()
}

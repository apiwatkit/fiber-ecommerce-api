package services

import (
	"context"

	"fiber-ecommerce-api/internal/core/domain/entities"
	"fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"fiber-ecommerce-api/internal/core/domain/ports/services"
)

type statsService struct {
	statsRepo repositories.StatsRepository
}

func NewStatsService(statsRepo repositories.StatsRepository) services.StatsService {
	return &statsService{
		statsRepo: statsRepo,
	}
}

func (s *statsService) GetSalesStats(ctx context.Context) (*entities.SalesStats, error) {
	return s.statsRepo.GetSalesStats(ctx)
}

func (s *statsService) GetProductStats(ctx context.Context) (*entities.ProductStats, error) {
	return s.statsRepo.GetProductStats(ctx)
}

func (s *statsService) GetUserStats(ctx context.Context) (*entities.UserStats, error) {
	return s.statsRepo.GetUserStats(ctx)
}
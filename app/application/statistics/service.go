package statistics

import (
	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
)

type service struct {
	repo gateways.StatisticsRepo
}

func NewService(repo gateways.StatisticsRepo) gateways.StatisticsService {
	return &service{
		repo: repo,
	}
}

// FetchAdminBestSelling implements gateways.StatisticsService.
func (s *service) FetchAdminBestSelling(limit int, offset int, minDate string, maxDate string) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAdminBestSelling(limit, offset, minDate, maxDate)
}

// FetchAdminRevenue implements gateways.StatisticsService.
func (s *service) FetchAdminRevenue() *presenters.AdminRevenue {
	return s.repo.ReadAdminRevenue()
}

// FetchAdminRevenueOverview implements gateways.StatisticsService.
func (s *service) FetchAdminRevenueOverview(filter string) ([]*presenters.AdminOverview, error) {
	return s.repo.ReadAdminRevenueOverview(filter)
}

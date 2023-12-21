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
func (s *service) FetchAdminRevenueOverview(filter string) ([]presenters.AdminOverview, error) {
	return s.repo.ReadAdminRevenueOverview(filter)
}

// FetchAdminCompanyOverview implements gateways.StatisticsService.
func (s *service) FetchAdminCompanyOverview(companyId int, filter string) (*presenters.AdminCompanyOverview, error) {
	return s.repo.ReadAdminCompanyOverview(companyId, filter)
}

// FetchCompanyIncidentOverview implements gateways.StatisticsService.
func (s *service) FetchCompanyIncidentOverview(companyId int, filter string) (*presenters.CompanyIncidentOverview, error) {
	return s.repo.ReadCompanyIncidentOverview(companyId, filter)
}

// FetchCompanyMonthRevenue implements gateways.StatisticsService.
func (s *service) FetchCompanyMonthRevenue(companyId int) (*presenters.CompanyMonthRevenue, error) {
	return s.repo.ReadCompanyMonthRevenue(companyId)
}

// FetchCompanyRevenueOverview implements gateways.StatisticsService.
func (s *service) FetchCompanyRevenueOverview(companyId int, filter string) ([]presenters.CompanyRevenueOverview, error) {
	return s.repo.ReadCompanyRevenueOverview(companyId, filter)
}

// FetchCompanyTripOverview implements gateways.StatisticsService.
func (s *service) FetchCompanyTripOverview(companyId int, filter string) (*presenters.CompanyTripOverview, error) {
	return s.repo.ReadCompanyTripOverview(companyId, filter)
}

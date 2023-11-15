package terminal

import (
	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/ent"
)

type service struct {
	repo gateways.TerminalRepo
}

func NewService(repo gateways.TerminalRepo) gateways.TerminalService {
	return &service{
		repo: repo,
	}
}

// Create implements gateways.TerminalService.
func (s *service) Create(companyId int, request *requeststructs.TerminalRequest) (*ent.Terminal, error) {
	return s.repo.Insert(companyId, request)
}

// Fetch implements gateways.TerminalService.
func (s *service) Fetch(id int) (*ent.Terminal, error) {
	return s.repo.Read(id)
}

// FetchAll implements gateways.TerminalService.
func (s *service) FetchAll(limit int, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset)
}

// FetchAllByCompany implements gateways.TerminalService.
func (s *service) FetchAllByCompany(companyId int, limit int, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllByCompany(companyId, limit, offset)
}

// Remove implements gateways.TerminalService.
func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}

// Update implements gateways.TerminalService.
func (s *service) Update(id int, request *requeststructs.TerminalRequest) (*ent.Terminal, error) {
	return s.repo.Update(id, request)
}

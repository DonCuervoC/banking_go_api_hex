package service

import (
	"time"

	"github.com/DonCuervoC/banking_go_api_hex/domain"
	"github.com/DonCuervoC/banking_go_api_hex/dto"
	"github.com/DonCuervoC/banking_go_api_hex/errs"
)

// primary port
type AccountService interface {
	NewAccount(dto.NewAccountRequestDto) (*dto.NewAccountResponseDto, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequestDto) (*dto.NewAccountResponseDto, *errs.AppError) {

	err := req.Validate()

	if err != nil {
		return nil, err
	}

	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02T15:04:05Z07:00"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      true,
	}

	newAccount, err := s.repo.SaveAccount(a)

	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountResponseDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}

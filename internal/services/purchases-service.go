package services

import "github.com/afmireski/garchop-api/internal/ports"

type PurchasesService struct {
	repository ports.PurchaseRepositoryPort
	cartRepository ports.CartsRepositoryPort
}

func NewPurchasesService(repository ports.PurchaseRepositoryPort, cartRepository ports.CartsRepositoryPort) *PurchasesService {
	return &PurchasesService{
		repository: repository,
		cartRepository: cartRepository,
	}
}


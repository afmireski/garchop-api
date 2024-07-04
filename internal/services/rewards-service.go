package services

import (
	"github.com/afmireski/garchop-api/internal/entities"
	"github.com/afmireski/garchop-api/internal/ports"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type RewardsService struct {
	rewardsRepository ports.RewardsRepositoryPort
}

func NewRewardsService(rewardsRepository ports.RewardsRepositoryPort) *RewardsService {
	return &RewardsService{
		rewardsRepository: rewardsRepository,
	}
}

func (r *RewardsService) ListAllRewards() ([]entities.Reward, *customErrors.InternalError) {

	repositoryData, err := r.rewardsRepository.FindAll(myTypes.Where{})

	if err != nil {
		return nil, customErrors.NewInternalError("a failure occurred when try to find the rewards", 500, []string{err.Error()})
	}

	rewards := entities.BuildRewardsFromModels(repositoryData)

	return rewards, nil
}

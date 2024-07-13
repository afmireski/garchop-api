package services

import (
	"github.com/afmireski/garchop-api/internal/entities"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/validators"

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

func validateNewRewardInput(input myTypes.NewRewardInput) *customErrors.InternalError {
	if !validators.IsValidNumericId(input.TierId) {
		return customErrors.NewInternalError("invalid tier id", 400, []string{})
	}

	return nil
}

func (s *RewardsService) NewReward(input myTypes.NewRewardInput) *customErrors.InternalError {
	if inputErr := validateNewRewardInput(input); inputErr != nil {
		return inputErr
	}

	data := ports.CreateRewardInput{
		TierId:             input.TierId,
		Name:               input.Name,
		Description:        input.Description,
		ExperienceRequired: input.ExperienceRequired,
		Type:               input.Type,
		Prize:              input.Prize,
	}

	_, err := s.rewardsRepository.Create(data)

	if err != nil {
		return customErrors.NewInternalError("a failure occurred when trying to create a reward", 500, []string{})
	}

	return nil
}

func (s *RewardsService) ListAllRewards() ([]entities.Reward, *customErrors.InternalError) {

	repositoryData, err := s.rewardsRepository.FindAll(myTypes.Where{})

	if err != nil {
		return nil, customErrors.NewInternalError("a failure occurred when try to find the rewards", 500, []string{err.Error()})
	}

	rewards := entities.BuildRewardsFromModels(repositoryData)

	return rewards, nil
}

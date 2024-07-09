package services

import (
	"github.com/afmireski/garchop-api/internal/entities"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/validators"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type RewardsService struct {
	rewardsRepository     ports.RewardsRepositoryPort
	userRewardsRepository ports.UserRewardsRepositoryPort
}

func NewRewardsService(rewardsRepository ports.RewardsRepositoryPort, userRewardsRepository ports.UserRewardsRepositoryPort) *RewardsService {
	return &RewardsService{
		rewardsRepository:     rewardsRepository,
		userRewardsRepository: userRewardsRepository,
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

func (r *RewardsService) validateClaimRewardInput(input myTypes.UserRewardInput) *customErrors.InternalError {

	if !validators.IsValidUuid(input.UserId) {
		return customErrors.NewInternalError("invalid user_id", 400, []string{"the user_id must be a valid uuid"})
	} else if !validators.IsValidUuid(input.RewardId) {
		return customErrors.NewInternalError("invalid reward_id", 400, []string{"the reward_id must be a valid uuid"})
	}

	return nil
}

func (r *RewardsService) ClaimReward(input myTypes.UserRewardInput) *customErrors.InternalError {
	validationErr := r.validateClaimRewardInput(input); if validationErr != nil {
		return validationErr
	}

	_, err := r.userRewardsRepository.Create(myTypes.UserRewardInput{
		UserId: input.UserId,
		RewardId: input.RewardId,
	}); if err != nil {
		return customErrors.NewInternalError("a failure occurred when try to claim the reward", 500, []string{err.Error()})
	}

	return nil
}
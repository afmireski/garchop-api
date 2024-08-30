package services

import (
	"time"

	"github.com/afmireski/garchop-api/internal/entities"
	"github.com/afmireski/garchop-api/internal/models"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/validators"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type RewardsService struct {
	rewardsRepository      ports.RewardsRepositoryPort
	userRewardsRepository  ports.UserRewardsRepositoryPort
	userPokemonsRepository ports.UserPokemonRepositoryPort
}

func NewRewardsService(rewardsRepository ports.RewardsRepositoryPort, userRewardsRepository ports.UserRewardsRepositoryPort, userPokemonsRepository ports.UserPokemonRepositoryPort) *RewardsService {
	return &RewardsService{
		rewardsRepository:      rewardsRepository,
		userRewardsRepository:  userRewardsRepository,
		userPokemonsRepository: userPokemonsRepository,
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

	data := myTypes.CreateRewardInput{
		TierId:             input.TierId,
		Name:               input.Name,
		Description:        input.Description,
		ExperienceRequired: input.ExperienceRequired,
		Type:               input.Type,
		Prize:              input.Prize,
	}

	_, err := s.rewardsRepository.Create(data)

	if err != nil {
		return customErrors.NewInternalError("a failure occurred when trying to create a reward", 500, []string{err.Error()})
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

func (s *RewardsService) RemoveReward(rewardId string) *customErrors.InternalError {
	if !validators.IsValidUuid(rewardId) {
		return customErrors.NewInternalError("invalid reward_id", 400, []string{"the reward_id must be a valid uuid"})
	}

	data := myTypes.AnyMap{
		"deleted_at": time.Now(),
	}
	_, err := s.rewardsRepository.Update(rewardId, data, myTypes.Where{})

	if err != nil {
		return customErrors.NewInternalError("a failure occurred when try to delete the reward", 500, []string{err.Error()})
	}

	return nil
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
	validationErr := r.validateClaimRewardInput(input)
	if validationErr != nil {
		return validationErr
	}

	reward, findRewardErr := r.rewardsRepository.FindById(input.RewardId, myTypes.Where{})

	if findRewardErr != nil {
		return customErrors.NewInternalError("reward not found", 404, []string{findRewardErr.Error()})
	}

	_, err := r.userRewardsRepository.Create(myTypes.UserRewardInput{
		UserId:   input.UserId,
		RewardId: input.RewardId,
	})
	if err != nil {
		return customErrors.NewInternalError("a failure occurred when try to claim the reward", 500, []string{err.Error()})
	}

	prizeErr := r.getRewardPrize(*reward, input.UserId)
	if prizeErr != nil {
		return prizeErr
	}

	return nil
}

func (r *RewardsService) getRewardPrize(reward models.RewardModel, userId string) *customErrors.InternalError {
	if reward.PrizeType == "pokemon" {
		pokemonId := reward.Prize["pokemon_id"].(string)

		_, err := r.userPokemonsRepository.Upsert(myTypes.UserPokemonData{
			UserId:    userId,
			PokemonId: pokemonId,
			Quantity:  1,
		})
		if err != nil {
			return customErrors.NewInternalError("a failure occurred when try to get the prize", 500, []string{err.Error()})
		}

	}

	return nil
}

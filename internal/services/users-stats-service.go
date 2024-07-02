package services

import (
	"github.com/afmireski/garchop-api/internal/models"
	"github.com/afmireski/garchop-api/internal/ports"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type UsersStatsService struct {
	repository   ports.UserStatsRepository
	tiersService *TiersService
}

func NewUsersStatsService(repository ports.UserStatsRepository, tiersService *TiersService) *UsersStatsService {
	return &UsersStatsService{
		repository,
		tiersService,
	}
}

func (s *UsersStatsService) LevelUp(userId string, newTierId uint) *customErrors.InternalError {
	_, err := s.repository.Update(userId, myTypes.AnyMap{"tier_id": newTierId}, myTypes.Where{})

	if err != nil {
		return customErrors.NewInternalError("a failure occurred during user level up", 500, []string{err.Error()})
	}

	return nil
}

func (s *UsersStatsService) GainExperience(userId string, currentExperience uint, gainedXp uint) (*uint, *customErrors.InternalError) {
	experienceSum := currentExperience + gainedXp

	_, err := s.repository.Update(userId, myTypes.AnyMap{"experience": experienceSum}, myTypes.Where{}); if err != nil {
		return nil, customErrors.NewInternalError("a failure occurred during user experience gain", 500, []string{err.Error()})
	}

	return &experienceSum, nil
}

func (s *UsersStatsService) canLevelUp(currentTierId uint, newExperienceAmount uint) (*uint, *customErrors.InternalError) {
	nextTier, findNextTierErr := s.tiersService.FindNextTier(int(currentTierId))

	if findNextTierErr != nil {
		return nil, customErrors.NewInternalError("a failure occurred when try to find the next tier", 500, []string{findNextTierErr.Error()})
	}

	if nextTier.MinimalExperience > newExperienceAmount {
		return nil, nil
	}

	return &nextTier.Id, nil
}


func (s *UsersStatsService) calculateGainedXpFromItems(items []models.ItemModel) uint {
	gainedXp := uint(0)
	for _, item := range items {
		gainedXp += item.Pokemon.Experience
	}

	return gainedXp
} 

func (s *UsersStatsService) ComputateExperienceFromItems(userId string, currentTierId uint, currentExperience uint, items []models.ItemModel) *customErrors.InternalError {

	gainedXp := s.calculateGainedXpFromItems(items)

	newExperience, gainXpErr := s.GainExperience(userId,currentExperience, gainedXp); if gainXpErr != nil {
		return gainXpErr
	}

	newTierId, canLevelUpErr := s.canLevelUp(currentTierId, *newExperience); if canLevelUpErr != nil {
		return canLevelUpErr
	}

	if newTierId != nil {
		return s.LevelUp(userId, *newTierId)
	}

	return nil
	
}

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

func (s *UsersStatsService) GainExperience(userId string, currentTierId uint, currentExperience uint, items []models.ItemModel) *customErrors.InternalError {

	gainedXp := s.calculateGainedXp(items)
	nextTier, findNextTierErr := s.tiersService.FindNextTier(int(currentTierId))

	if findNextTierErr != nil {
		return customErrors.NewInternalError("a failure occurred when try to find the next tier", 500, []string{findNextTierErr.Error()})
	}
	experienceSum := currentExperience + gainedXp

	_, err := s.repository.Update(userId, myTypes.AnyMap{"experience": experienceSum}, myTypes.Where{}); if err != nil {
		return customErrors.NewInternalError("a failure occurred during user experience gain", 500, []string{err.Error()})
	}

	if nextTier != nil && experienceSum >= nextTier.MinimalExperience {
		s.LevelUp(userId, nextTier.Id)
	}

	return nil
}


func (s *UsersStatsService) calculateGainedXp(items []models.ItemModel) uint {
	var gainedXp uint
	gainedXp = 0
	for _, item := range items {
		gainedXp += item.Pokemon.Experience
	}

	return gainedXp
} 

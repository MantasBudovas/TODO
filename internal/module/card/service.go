package card

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"app/internal/entity"
	"app/pkg/logger"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	repo  *Repository
	redis *redis.Client
}

func NewService(repo *Repository, redis *redis.Client) *Service {
	return &Service{repo: repo, redis: redis}
}

func (s *Service) CreateCard(card *entity.Card) error {
	if card.Priority == "" {
		card.Priority = low
	}

	if card.Priority != low && card.Priority != medium && card.Priority != high {
		return errors.New("invalid priority: must be Low, Medium, or High")
	}

	if err := s.repo.Save(card); err != nil {
		return err
	}

	s.redis.Del(context.Background(), "all_cards")

	logger.Info("Created new card with priority: " + card.Priority)
	return nil
}

func (s *Service) GetCards() ([]entity.Card, error) {
	ctx := context.Background()
	cacheKey := "all_cards"

	val, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		logger.Info("Cache Hit! Returning data from Redis.")
		var cards []entity.Card
		json.Unmarshal([]byte(val), &cards)
		return cards, nil
	}

	logger.Info("Cache Miss. Fetching from DB...")
	cards, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(cards)
	s.redis.Set(ctx, cacheKey, data, 5*time.Minute)

	return cards, nil
}

func (s *Service) ModifyCard(id int, input *entity.Card) (*entity.Card, error) {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err // Likely record not found
	}

	existing.Completed = input.Completed

	if input.Name != "" {
		existing.Name = input.Name
	}
	if input.Priority != "" {
		existing.Priority = input.Priority
	}
	if input.DueDate != nil {
		existing.DueDate = input.DueDate
	}

	if existing.Priority != low && existing.Priority != medium && existing.Priority != high {
		return nil, errors.New("invalid priority")
	}

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}

	s.redis.Del(context.Background(), "all_cards")

	return existing, nil
}

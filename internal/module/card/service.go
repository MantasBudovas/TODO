package card

import (
	"context"
	"encoding/json"
	"time"

	"app/internal/entity"
	"app/pkg/logger"

	"github.com/redis/go-redis/v9"
)

const cacheKey = "all_cards"

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
		return ErrInvalidPriority
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

	val, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		logger.Info("Returning data from Redis.")
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
		return nil, err
	}

	if input.Name != "" {
		existing.Name = input.Name
	}
	if input.Completed != nil {
		existing.Completed = input.Completed
	}
	if input.DueDate != nil {
		existing.DueDate = input.DueDate
	}
	if input.Priority != "" {
		existing.Priority = input.Priority
	}

	if existing.Priority != low && existing.Priority != medium && existing.Priority != high {
		return nil, ErrInvalidPriority
	}

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}

	s.redis.Del(context.Background(), cacheKey)

	return existing, nil
}

package domain

import (
	//"errors"
	//"app/internal/models"
	"app/pkg/config"
	"app/pkg/telegram"
	"fmt"
	"sync"
)

type Services struct {
	Telegram   *telegram.Telegram
	Config     *config.Config
	TestSpeedI int64
	mu         sync.Mutex
}

func (s *Services) TestSpeed() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.TestSpeedI++
	return fmt.Sprintf("123 %d", s.TestSpeedI), nil
	//return fmt.Errorf("ошибка: %s", "что-то пошло не так555")

}

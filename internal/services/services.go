package services

import (
	//"errors"
	//"app/internal/models"
	"app/pkg/config"
	"app/pkg/telegram"
	"fmt"
)

type Services struct {
	Telegram   *telegram.Telegram
	Config     *config.Config
	TestSpeedI int64
}

func (s *Services) TestSpeed() (string, error) {
	s.TestSpeedI++
	return fmt.Sprintf("123 %d", s.TestSpeedI), nil
	//return fmt.Errorf("ошибка: %s", "что-то пошло не так555")

}

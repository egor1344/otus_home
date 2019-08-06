package calendar

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

var Logger *zap.SugaredLogger

// Инициализация логера
func InitLogger() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	Logger = logger.Sugar()
	return nil
}

// Чтение конфинга
func ReadConfigFile(name, path string) {
	viper.SetConfigName(name)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		Logger.Fatal("Read config failed", err)
	}
}

// Структура события
type Event struct {
	id          int
	Date        time.Time
	Name        string
	Description string
}

// Структура календаря
type Calendar struct {
	idSeq  int // Счетчик id(костыль пока бд не подключим)
	Events map[int]*Event
}

// Инициализация каледаря
func (c *Calendar) Init() {
	c.Events = make(map[int]*Event)
}

// Добавление события в календаре
func (c *Calendar) AddEvent(event *Event) error {
	_, ok := c.Events[event.id]
	if ok {
		return errors.New("Event with this id has already been created")
	}
	c.Events[event.id] = event
	if event.id > c.idSeq {
		c.idSeq = event.id
	}
	return nil
}

// Обновление события в календаре
func (c *Calendar) UpdateEvent(eventId int, newEvent *Event) error {
	if eventId != 0 {
		_, ok := c.Events[eventId]
		if !ok {
			return errors.New("Event with id not created")
		}
		c.Events[eventId] = newEvent
		return nil
	}
	_, ok := c.Events[newEvent.id]
	if !ok {
		return errors.New("Event with id not created")
	}
	c.Events[newEvent.id] = newEvent
	return nil
}

// Удаление события в календаря
func (c *Calendar) DeleteEvent(eventId int) error {
	delete(c.Events, eventId)
	return nil
}

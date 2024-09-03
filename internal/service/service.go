package service

import "fmt"

// Service структура, представляющая основной сервис
type Service struct {
	Name string
}

// NewService создает новый сервис
func NewService(name string) *Service {
	return &Service{Name: name}
}

// PrintName выводит имя сервиса в консоль
func (s *Service) PrintName() {
	fmt.Printf("Service Name from Service Struct: %s\n", s.Name)
}

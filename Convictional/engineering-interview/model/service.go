package model

import (
	"sync"
)

// Service is an in memory storage interface for Products
type Service struct {
	ProductsMap map[string]*Product
	mu          sync.RWMutex
}

// NewService makes a new empty map of Products
func NewService() (*Service, error) {

	return &Service{
		ProductsMap: make(map[string]*Product),
	}, nil
}

// AddProduct adds a Product to the map
func (s *Service) AddProduct(w *Product) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ProductsMap[string(w.ID)] = w
}

// GetProduct returns a Product
func (s *Service) GetProduct(id string) (w *Product) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	w = s.ProductsMap[id]
	return w
}

func (s *Service) GetSize() (w int) {
	return len(s.ProductsMap)
}

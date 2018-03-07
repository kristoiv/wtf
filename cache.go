package wtf

type TodoListServiceCache struct {
	TodoListService
	itemsCache []Item
	cacheValid bool
}

func (s *TodoListServiceCache) Add(title string) (*Item, error) {
	s.cacheValid = false
	return s.TodoListService.Add(title)
}

func (s *TodoListServiceCache) SetChecked(id ItemID, checked bool) error {
	s.cacheValid = false
	return s.TodoListService.SetChecked(id, checked)
}

func (s *TodoListServiceCache) Remove(id ItemID) error {
	s.cacheValid = false
	return s.TodoListService.Remove(id)
}

func (s *TodoListServiceCache) Items() ([]Item, error) {
	if s.cacheValid {
		return s.itemsCache, nil
	}

	items, err := s.TodoListService.Items()
	s.itemsCache = items
	s.cacheValid = true

	return items, err
}

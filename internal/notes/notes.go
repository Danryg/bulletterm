package notes

import "time"

type Note struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Store struct {
	notes []Note
	path  string
}

func NewStore(path string) (*Store, error) {
	s := &Store{path: path}
	if err := s.load(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Store) All() []Note {
	result := make([]Note, len(s.notes))
	copy(result, s.notes)
	return result
}

func (s *Store) Add(title, body string) (Note, error) {
	now := time.Now()
	n := Note{
		ID:        newID(),
		Title:     title,
		Body:      body,
		CreatedAt: now,
		UpdatedAt: now,
	}
	s.notes = append(s.notes, n)
	return n, s.save()
}

func (s *Store) Update(id, title, body string) error {
	for i, n := range s.notes {
		if n.ID == id {
			s.notes[i].Title = title
			s.notes[i].Body = body
			s.notes[i].UpdatedAt = time.Now()
			return s.save()
		}
	}
	return ErrNotFound
}

func (s *Store) Delete(id string) error {
	for i, n := range s.notes {
		if n.ID == id {
			s.notes = append(s.notes[:i], s.notes[i+1:]...)
			return s.save()
		}
	}
	return ErrNotFound
}

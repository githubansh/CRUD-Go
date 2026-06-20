package repository

import (
	"strconv"
	"sync"

	"user-service/internal/model"
)

// UserRepository stores users in memory.
// A mutex keeps the map safe if multiple requests arrive at the same time.
type UserRepository struct {
	mu     sync.RWMutex
	users  map[string]model.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make(map[string]model.User),
		nextID: 1,
	}
}

func (r *UserRepository) Create(user model.User) model.User {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = strconv.Itoa(r.nextID)
	r.users[user.ID] = user
	r.nextID++

	return user
}

func (r *UserRepository) GetByID(id string) (model.User, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, found := r.users[id]
	return user, found
}

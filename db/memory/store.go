package memory

type Repository struct {
	count int
	store map[int]locationDao
}

func NewRepository() *Repository {
	return &Repository{
		count: 0,
		store: make(map[int]locationDao),
	}
}

func (r *Repository) NewID() int {
	r.count = r.count + 1
	return r.count
}

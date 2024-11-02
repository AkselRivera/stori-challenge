package migration

import "github.com/AkselRivera/stori-challenge/migration-service/pkg/domain"

func (r *Repository) InsertMany(t []*domain.Transaction) error {
	result := r.Client.Omit("id").Create(t)

	return result.Error
}

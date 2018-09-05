package types

import (
	"time"

	"github.com/kamilsk/form-api/pkg/domain"
)

// Account TODO issue#173
type Account struct {
	ID        domain.ID  `db:"id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Users     []*User    `db:"-"`
}

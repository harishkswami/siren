package namespace

import (
	"time"
)

//go:generate mockery --name=Repository -r --case underscore --with-expecter --structname NamespaceRepository --filename namespace_repository.go --output=./mocks
type Repository interface {
	Migrate() error
	List() ([]*EncryptedNamespace, error)
	Create(*EncryptedNamespace) error
	Get(uint64) (*EncryptedNamespace, error)
	Update(*EncryptedNamespace) error
	Delete(uint64) error
}

type Namespace struct {
	Id          uint64                 `json:"id"`
	Urn         string                 `json:"urn"`
	Name        string                 `json:"name"`
	Provider    uint64                 `json:"provider"`
	Credentials map[string]interface{} `json:"credentials"`
	Labels      map[string]string      `json:"labels"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

type EncryptedNamespace struct {
	*Namespace
	Credentials string
}
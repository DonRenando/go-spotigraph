package dgraph

import (
	"context"

	"github.com/dgraph-io/dgo"
	"github.com/dolan-in/dgman"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"
)

type dgoUserRepository struct {
	c *dgo.Dgraph
}

func NewUserRepository() repositories.User {
	return &dgoUserRepository{}
}

// -----------------------------------------------------------------------------

type dgoUser struct {
	UID       string `json:"uid,omitempty"`
	ID        string `json:"xid,omitempty" dgraph:"index=exact unique"`
	Principal string `json:"prn,omitempty" dgraph:"index=exact unique"`
}

func (u *dgoUser) NodeType() string {
	return "user"
}

func fromUser(entity *models.User) *dgoUser {
	return &dgoUser{
		ID:        entity.ID,
		Principal: entity.Principal,
	}
}

func (u *dgoUser) ToModel() *models.User {
	return &models.User{
		ID:        u.UID,
		Principal: u.Principal,
	}
}

// -----------------------------------------------------------------------------

func (r *dgoUserRepository) Create(ctx context.Context, entity *models.User) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Defer to adapter
	return dgman.Create(ctx, r.c.NewTxn(), fromUser(entity), dgman.MutateOptions{CommitNow: true})
}

func (r *dgoUserRepository) Get(ctx context.Context, id string) (*models.User, error) {
	var entity user

	// Get with a filter
	err := dgman.Get(ctx, r.c.NewTxn(), &entity).
		Vars("getNode($id: string)", map[string]string{"$id", id}).
		Filter("eq(xid, $id)").
		All(1).
		Node()
	if err != nil {
		if err == dgman.ErrNodeNotFound {
			return nil, db.ErrNoResult
		}
		return nil, err
	}

	// Return result
	return entity.ToModel(), nil
}

func (r *dgoUserRepository) Update(ctx context.Context, entity *models.User) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Initialize a transaction
	txn := r.c.NewTxn()

	// Convert to node
	node := fromUser(entity)

	// Retreive entity
	err := dgman.Get(ctx, txn, &node).
		Vars("getNode($id: string)", map[string]string{"$id", id}).
		Filter("eq(xid, $id)").
		All(1).
		Node()
	if err != nil {
		if err == dgman.ErrNodeNotFound {
			return nil, db.ErrNoResult
		}
		return nil, err
	}

	// Update node

	// Mutate node
	return dgman.Update(ctx, txn, &node, dgman.MutateOptions{CommitNow: true})
}

func (r *dgoUserRepository) Delete(ctx context.Context, id string) error {
	_, err := dgman.Delete(ctx, r.c.NewTxn(), &dgoUser{}, dgman.MutateOptions{CommintNow: true}).
		Vars("getNode($id: string)", map[string]string{"$id", id}).
		Filter("eq(xid, $id)").
		Nodes()
	return err
}

func (r *dgoUserRepository) Search(ctx context.Context, filter *repositories.UserSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.User, int, error) {
	panic("Not implemented")
}

func (r *dgoUserRepository) FindByPrincipal(ctx context.Context, prn string) (*models.User, error) {
	var entity user

	// Get with a filter
	err := dgman.Get(ctx, r.c.NewTxn(), &entity).
		Vars("getNode($prn: string)", map[string]string{"$prn", prn}).
		Filter("eq(prn, $prn)").
		All(1).
		Node()
	if err != nil {
		if err == dgman.ErrNodeNotFound {
			return nil, db.ErrNoResult
		}
		return nil, err
	}

	// Return result
	return entity.ToModel(), nil
}

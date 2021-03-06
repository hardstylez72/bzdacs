package user

import (
	"context"
	"github.com/pkg/errors"
)

type secRepository struct {
	rep Repository
}

func NewSecurityGuard(rep Repository) *secRepository {
	return &secRepository{rep: rep}
}

func (r *secRepository) Insert(ctx context.Context, user *SysUser) (*SysUser, error) {
	encPassword, err := Encrypt(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = encPassword

	u, err := r.rep.Insert(ctx, user)
	if err != nil {
		return nil, err
	}
	u.Password = ""

	return u, nil
}
func (r *secRepository) GetByParams(ctx context.Context, login, password string) (*SysUser, error) {

	encPassword, err := Encrypt(password)
	if err != nil {
		return nil, err
	}

	u, err := r.rep.GetByParams(ctx, login, encPassword)
	if err != nil {
		return nil, err
	}

	if !ComparePasswords(u.Password, []byte(password)) {
		return nil, errors.New("invalid password")
	}

	u.Password = ""

	return u, nil
}
func (r *secRepository) GetById(ctx context.Context, id, namespaceId int) (*SysUser, error) {
	u, err := r.rep.GetById(ctx, id, namespaceId)
	if err != nil {
		return nil, err
	}

	u.Password = ""

	return u, nil
}

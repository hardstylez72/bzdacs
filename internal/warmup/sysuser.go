package warmup

import (
	"context"
	"errors"
	"github.com/hardstylez72/bzdacs/internal/user"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
)

func registerUser(ctx context.Context, sysUserRepository user.Repository, login, password string) error {

	_, err := sysUserRepository.Insert(ctx, &user.SysUser{
		Login:    login,
		Password: password,
	})
	if err != nil {
		if !errors.Is(err, storage.EntityAlreadyExist) {
			return err
		}

	}
	return nil
}

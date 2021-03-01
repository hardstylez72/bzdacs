package sysuser

import (
	"context"
	"errors"
	"github.com/hardstylez72/bzdacs/generated/client"
	sysuser "github.com/hardstylez72/bzdacs/generated/client/sys_user"
	"github.com/hardstylez72/bzdacs/generated/models"
	"github.com/hardstylez72/bzdacs/tests"
)

type User struct {
	Id       int64
	Login    string
	Password int64
	tests.Times
}

var (
	ErrAlreadyRegistered = errors.New("user has been already registered")
)

func Create(ctx context.Context, client *client.BZDACS, login, password string) error {
	_, err := client.SysUser.SysUserRegister(
		&sysuser.SysUserRegisterParams{
			Req: &models.SysUserRegisterRequest{
				Login:    &login,
				Password: &password,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return wrapErr(err)
	}

	return nil
}

func wrapErr(err error) error {
	_, ok := err.(*sysuser.SysUserRegisterUnprocessableEntity)
	if ok {
		return ErrAlreadyRegistered
	}
	return err
}

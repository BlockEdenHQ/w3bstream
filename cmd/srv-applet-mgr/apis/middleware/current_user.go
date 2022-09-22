package middleware

import (
	"context"
	"reflect"

	"github.com/iotexproject/Bumblebee/conf/jwt"
	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/pkg/errors/status"
	"github.com/iotexproject/w3bstream/pkg/types"

	"github.com/iotexproject/w3bstream/pkg/models"
	"github.com/iotexproject/w3bstream/pkg/modules/account"
)

type ContextAccountAuth struct {
	httpx.MethodGet
}

var contextAccountAuthKey = reflect.TypeOf(ContextAccountAuth{}).String()

func (r *ContextAccountAuth) ContextKey() string { return contextAccountAuthKey }

func (r *ContextAccountAuth) Output(ctx context.Context) (interface{}, error) {
	v, ok := jwt.AuthFromContext(ctx).(string)
	if !ok {
		panic("not an account id")
	}
	ca, err := account.GetAccountByAccountID(ctx, v)
	if err != nil {
		return nil, err
	}
	return &CurrentAccount{*ca}, nil
}

func CurrentAccountFromContext(ctx context.Context) *CurrentAccount {
	return ctx.Value(contextAccountAuthKey).(*CurrentAccount)
}

type CurrentAccount struct {
	models.Account
}

func (v *CurrentAccount) ValidateProjectPerm(ctx context.Context, prjID string) (*models.Project, error) {
	d := types.MustDBExecutorFromContext(ctx)
	a := CurrentAccountFromContext(ctx)
	m := &models.Project{RelProject: models.RelProject{ProjectID: prjID}}

	if err := m.FetchByProjectID(d); err != nil {
		return nil, status.NotFound.StatusErr().WithDesc("project not found")
	}
	if a.AccountID != m.AccountID {
		return nil, status.Unauthorized.StatusErr().WithDesc("no project permission")
	}
	return m, nil
}

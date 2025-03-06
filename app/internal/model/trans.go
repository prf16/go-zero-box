package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Func func(ctx context.Context, session sqlx.Session) error

// Trans 数据库事物支持 .
func Trans(ctx context.Context, db sqlx.SqlConn, fns ...Func) (err error) {
	err = db.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		for _, fn := range fns {
			if err = fn(ctx, session); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

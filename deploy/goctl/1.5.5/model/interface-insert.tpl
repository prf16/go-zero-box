WithSession(session sqlx.Session) *default{{.upperStartCamelObject}}Model
Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error)
InsertWithUpdate(ctx context.Context, data *{{.upperStartCamelObject}}) error
InsertWithUpdateField(ctx context.Context, data *{{.upperStartCamelObject}}, fields []string) error
BatchInsert(ctx context.Context, list []*{{.upperStartCamelObject}}) (sql.Result,error)
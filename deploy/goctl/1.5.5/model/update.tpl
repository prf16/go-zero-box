func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, newData.{{.upperStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}

{{end}}	{{.keys}}
    _, {{if .containsIndexCache}}err{{else}}err:{{end}}= m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
    _,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return err
}

func (m *default{{.upperStartCamelObject}}Model) UpdateField(ctx context.Context, data *{{.upperStartCamelObject}}, fields []string) error {
	if data.Id == 0 || len(fields) == 0 {
		return nil
	}

	var fieldVals []interface{}
	for _, column := range fields {
		fieldVal := reflect.ValueOf(data).Elem().FieldByName(tools.UderscoreToUpperCamelCase(column)).Interface()
		fieldVals = append(fieldVals, fieldVal)
	}

	fieldVals = append(fieldVals, data.Id)

	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, strings.Join(fields, "=?,") + "=?")
	_, err := m.conn.ExecCtx(ctx, query, fieldVals...)
	if err != nil {
		return err
	}

	return nil
}

func (m *default{{.upperStartCamelObject}}Model) Select(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*{{.upperStartCamelObject}}, error) {
	query, values, err := rowBuilder.From(m.table).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*{{.upperStartCamelObject}}
	err = m.conn.QueryRowsPartialCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) First(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*{{.upperStartCamelObject}}, error) {
	query, values, err := rowBuilder.From(m.table).ToSql()
	if err != nil {
		return nil, err
	}

	var resp {{.upperStartCamelObject}}
	err = m.conn.QueryRowPartialCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) Count(ctx context.Context, rowBuilder squirrel.SelectBuilder) (int64, error) {
	query, values, err := rowBuilder.RemoveColumns().From(m.table).RemoveOffset().RemoveLimit().Columns("count(*) as c").ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) Avg(ctx context.Context, rowBuilder squirrel.SelectBuilder, field string) (float64, error) {
	query, values, err := rowBuilder.RemoveColumns().From(m.table).Columns("IFNULL(AVG(" + field + "),0)").ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) Sum(ctx context.Context, rowBuilder squirrel.SelectBuilder, field string) (float64, error) {
	query, values, err := rowBuilder.RemoveColumns().From(m.table).Columns("IFNULL(SUM(" + field + "),0)").ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) Max(ctx context.Context, rowBuilder squirrel.SelectBuilder, field string) (float64, error) {
	query, values, err := rowBuilder.RemoveColumns().From(m.table).Columns("IFNULL(MAX(" + field + "),0)").ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) Min(ctx context.Context, rowBuilder squirrel.SelectBuilder, field string) (float64, error) {
	query, values, err := rowBuilder.RemoveColumns().From(m.table).Columns("IFNULL(Min(" + field + "),0)").ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) Aggregate(ctx context.Context, rowBuilder squirrel.SelectBuilder) (float64, error) {
	query, values, err := rowBuilder.From(m.table).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) DeleteFilter(ctx context.Context, where squirrel.Sqlizer) (sql.Result, error) {
	query, values, err := squirrel.Delete(m.table).Where(where).ToSql()
	if err != nil {
		return nil, err
	}

	resp, err := m.conn.ExecCtx(ctx, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) UpdateFilter(ctx context.Context, updateData map[string]interface{}, where squirrel.Sqlizer) (sql.Result, error) {
	rowBuilder := squirrel.Update(m.table)
	for column, value := range updateData {
		rowBuilder = rowBuilder.Set(fmt.Sprintf("`%s`", column), value)
	}
	rowBuilder = rowBuilder.Where(where)
	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	resp, err := m.conn.ExecCtx(ctx, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) SelectGroup(ctx context.Context, rowBuilder squirrel.SelectBuilder, resp interface{}) error {
	resp = tools.GetPointer(resp)

	query, values, err := rowBuilder.From(m.table).ToSql()
	if err != nil {
		return err
	}

	err = m.conn.QueryRowsPartialCtx(ctx, resp, query, values...)
	switch err {
	case nil:
		return nil
	case sqlc.ErrNotFound:
		return nil
	default:
		return err
	}
}

func (m *default{{.upperStartCamelObject}}Model) SelectGroupCount(ctx context.Context, rowBuilder squirrel.SelectBuilder) (int64, error) {
	var resp int64
	query, values, err := rowBuilder.From(m.table).RemoveOffset().RemoveLimit().ToSql()
	if err != nil {
		return resp, err
	}

	query = fmt.Sprintf("SELECT COUNT(*) as c FROM ( %s ) t", query)

	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, nil
	default:
		return resp, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) FirstCustom(ctx context.Context, rowBuilder squirrel.SelectBuilder, resp interface{}) error {
	resp = tools.GetPointer(resp)

	query, values, err := rowBuilder.From(m.table).ToSql()
	if err != nil {
		return err
	}

	err = m.conn.QueryRowPartialCtx(ctx, resp, query, values...)
	switch err {
	case nil:
		return nil
	case sqlc.ErrNotFound:
		return nil
	default:
		return err
	}
}

func (m *default{{.upperStartCamelObject}}Model) SelectCustom(ctx context.Context, rowBuilder squirrel.SelectBuilder, resp interface{}) error {
	resp = tools.GetPointer(resp)

	query, values, err := rowBuilder.From(m.table).ToSql()
	if err != nil {
		return err
	}

	err = m.conn.QueryRowsPartialCtx(ctx, resp, query, values...)
	switch err {
	case nil:
		return nil
	case sqlc.ErrNotFound:
		return nil
	default:
		return err
	}
}

func (m *default{{.upperStartCamelObject}}Model) Customs(ctx context.Context, rowBuilder squirrel.SelectBuilder, resp interface{}) error {
	resp = tools.GetPointer(resp)

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return err
	}

	err = m.conn.QueryRowsPartialCtx(ctx, resp, query, values...)
	switch err {
	case nil:
		return nil
	case sqlc.ErrNotFound:
		return nil
	default:
		return err
	}
}

func (m *default{{.upperStartCamelObject}}Model) Custom(ctx context.Context, rowBuilder squirrel.SelectBuilder, resp interface{}) error {
	resp = tools.GetPointer(resp)

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return err
	}

	err = m.conn.QueryRowPartialCtx(ctx, resp, query, values...)
	switch err {
	case nil:
		return nil
	case sqlc.ErrNotFound:
		return nil
	default:
		return err
	}
}
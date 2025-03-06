update.tpl

func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, updateData map[string]interface{}, where squirrel.Sqlizer) (sql.Result, error) {
	rowBuilder := squirrel.Update(m.table)
	for column, value := range updateData {
		rowBuilder = rowBuilder.Set(column, value)
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
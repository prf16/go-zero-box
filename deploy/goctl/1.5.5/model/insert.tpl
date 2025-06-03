func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error) {
    m.Init(data)
    data.CreatedAt = time.Now()
	{{if .withCache}}{{.keys}}
    ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
    ret,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return ret,err
}

func (m *default{{.upperStartCamelObject}}Model) InsertWithUpdate(ctx context.Context, data *{{.upperStartCamelObject}}) error {
	if data.Id == 0 {
	    m.Init(data)
		data.CreatedAt = time.Now()
        {{if .withCache}}{{.keys}}
        ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
            query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
            return conn.ExecCtx(ctx, query, {{.expressionValues}})
        }, {{.keyValues}}){{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
        ret,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
        id, err := ret.LastInsertId()
        if err != nil {
            return err
        }
        data.Id = uint64(id)
        return err
	} else {
        {{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, newData.{{.upperStartCamelPrimaryKey}})
        if err!=nil{
            return err
        }

        {{end}}	{{.keys}}
        _, {{if .containsIndexCache}}err{{else}}err:{{end}}= m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
            query := fmt.Sprintf("update %s set %s where `id` = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
            return conn.ExecCtx(ctx, query, {{.expressionValues}}, data.Id)
        }, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where `id` = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
        _,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}, data.Id){{end}}
        return err
	}
}

func (m *default{{.upperStartCamelObject}}Model) InsertWithUpdateField(ctx context.Context, data *{{.upperStartCamelObject}}, fields []string) error {
	if data.Id == 0 {
	    m.Init(data)
		data.CreatedAt = time.Now()
        {{if .withCache}}{{.keys}}
        ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
            query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
            return conn.ExecCtx(ctx, query, {{.expressionValues}})
        }, {{.keyValues}}){{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
        ret,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
        id, err := ret.LastInsertId()
        if err != nil {
            return err
        }
        data.Id = uint64(id)
        return err
	} else {
        err := m.UpdateField(ctx, data, fields)
        return err
	}
}

func (m *default{{.upperStartCamelObject}}Model) BatchInsert(ctx context.Context, list []*{{.upperStartCamelObject}}) (sql.Result,error) {
	rowBuilder := squirrel.Insert(m.table).Columns(strings.Split({{.lowerStartCamelObject}}RowsExpectAutoSet, ",")...)

	for _, data := range list {
        m.Init(data)
        data.CreatedAt = time.Now()
		rowBuilder = rowBuilder.Values({{.expressionValues}})
	}

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

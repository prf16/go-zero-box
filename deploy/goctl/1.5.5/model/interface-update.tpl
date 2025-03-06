Update(ctx context.Context, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error
UpdateField(ctx context.Context, data *{{.upperStartCamelObject}}, fields []string) error

Select(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*{{.upperStartCamelObject}}, error)
First(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*{{.upperStartCamelObject}}, error)
Count(ctx context.Context, rowBuilder squirrel.SelectBuilder) (int64, error)
Avg(ctx context.Context, rowBuilder squirrel.SelectBuilder, field string) (float64, error)
Sum(ctx context.Context, rowBuilder squirrel.SelectBuilder, field string) (float64, error)
Max(ctx context.Context, rowBuilder squirrel.SelectBuilder, field string) (float64, error)
Min(ctx context.Context, rowBuilder squirrel.SelectBuilder, field string) (float64, error)
Aggregate(ctx context.Context, rowBuilder squirrel.SelectBuilder) (float64, error)
DeleteFilter(ctx context.Context, where squirrel.Sqlizer) (sql.Result, error)
UpdateFilter(ctx context.Context, updateData map[string]interface{}, where squirrel.Sqlizer) (sql.Result, error)
SelectGroup(ctx context.Context, rowBuilder squirrel.SelectBuilder, resp interface{}) error
SelectGroupCount(ctx context.Context, rowBuilder squirrel.SelectBuilder) (int64, error)
FirstCustom(ctx context.Context, rowBuilder squirrel.SelectBuilder, resp interface{}) error
SelectCustom(ctx context.Context, rowBuilder squirrel.SelectBuilder, resp interface{}) error
Customs(ctx context.Context, rowBuilder squirrel.SelectBuilder, resp interface{}) error
Custom(ctx context.Context, rowBuilder squirrel.SelectBuilder, resp interface{}) error
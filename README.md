# bks

```go
s:=map[string]interface{}{
}
_, err = config.EsClient.Index().
Index("").
BodyJson(s).
Do(context.Background())
```

```go
boolQuery := elastic.NewBoolQuery()
match := elastic.NewMatchQuery("name", req.Name)
boolQuery.Must(match)
rangeQuery := elastic.NewRangeQuery("price").Gte(req.MinPrice).Lte(req.MaxPrice)
boolQuery.Must(rangeQuery)
rangeQuery := elastic.NewRangeQuery("expiration").Gte(req.MinTime).Lte(req.MaxTime)
boolQuery.Must(rangeQuery)
sort := elastic.NewFieldSort("_score").Desc()
if req.Stort != "" {
split := strings.Split(req.Stort, "_")
if len(split) == 2 && split[0] == "price" {
sort = elastic.NewFieldSort("price")
if split[1] == "desc" {
sort.Desc()
} else {
sort.Asc()
}
}
}
highlight:=elastic.NewHighlight()
highlight = highlight.Field().
PreTags("<span style=\"color:red;\">").
PostTags("</span>")
res, err := config.EsClient.Search("").
Query().
Size().
From().
SortBy().
Highlight().
Do(context.Background())
var list []*
for _, hit := range res.Hits.Hits {
p := new()
err := json.Unmarshal(hit.Source, p)
if err != nil {
return nil, errors.New("解析失败")
}
h,ok:=hit.Highlight[]
if ok{
p.Name = strings.join(h," ")
}
list = append(list, p)
}
```

```go
nowDay := time.Now().Truncate(24 * time.Hour)
if parse.Before(nowDay) {
return nil, errors.New("")
}
```

```go
defer func() {
if r := recover(); r != nil {
tx.Rollback()
}
}()

err = tx.Commit().Error
if err != nil {
return nil, errors.New("事务提交失败")
}
```

## License

MIT License

package student

import "github.com/olivere/elastic/v7"

type request struct {
	queryParams
}

type queryParams struct {
	id int32
}

const (
	idField = "id"
)

func New() *request {
	return &request{}
}

func (r *request) Id(id int32) *request {
	r.id = id
	return r
}

func (r *request) buildQuery() *elastic.BoolQuery {
	query := elastic.NewBoolQuery()
	if r.queryParams.id != 0 {
		query.Must(elastic.NewTermQuery(idField, r.queryParams.id))
	}
	return query
}

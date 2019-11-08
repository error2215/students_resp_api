package student

import (
	"encoding/json"

	"github.com/olivere/elastic/v7"
)

type request struct {
	queryParams
}

type queryParams struct {
	id       int32
	bodyJSON json.RawMessage
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

func (r *request) BodyJSON(bodyJSON json.RawMessage) *request {
	r.bodyJSON = bodyJSON
	return r
}

func (r *request) buildSearchQuery() *elastic.BoolQuery {
	query := elastic.NewBoolQuery()
	if r.queryParams.id != 0 {
		query.Must(elastic.NewTermQuery(idField, r.queryParams.id))
	}
	return query
}

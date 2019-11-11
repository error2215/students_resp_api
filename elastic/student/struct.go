package student

import (
	"github.com/olivere/elastic/v7"

	model "students_rest_api/models/student"
)

type request struct {
	queryParams
}

type queryParams struct {
	id   int32
	body *model.Student
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

func (r *request) Body(body *model.Student) *request {
	r.body = body
	return r
}

func (r *request) buildSearchQuery() *elastic.BoolQuery {
	query := elastic.NewBoolQuery()
	if r.queryParams.id != 0 {
		query.Must(elastic.NewTermQuery(idField, r.queryParams.id))
	}
	return query
}

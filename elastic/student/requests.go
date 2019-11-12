package student

import (
	"context"
	"encoding/json"
	"students_rest_api/models/student"

	"github.com/olivere/elastic/v7"

	"github.com/iostrovok/go-convert"

	"students_rest_api/config"
	client "students_rest_api/elastic"

	log "github.com/sirupsen/logrus"
)

func (r *request) GetStudent() *student.Student {
	hits, err := client.GetClient().Search().
		Query(r.buildSearchQuery()).
		Size(1).
		Index(config.GlobalConfig.StudentsIndex).
		Do(context.Background())
	if err != nil {
		log.WithField("method", "GetStudent").Error(err)
	}

	var res *student.Student
	if hits.TotalHits() == 0 {
		return nil
	}
	err = json.Unmarshal(hits.Hits.Hits[0].Source, &res)
	if err != nil {
		log.WithField("method", "GetStudent").Error(err)
		return nil
	}
	return res
}

func (r *request) ListStudents() []*student.Student {
	hits, err := client.GetClient().Search().
		Query(r.buildSearchQuery()).
		Size(500).
		Index(config.GlobalConfig.StudentsIndex).
		Do(context.Background())
	if err != nil {
		log.WithField("method", "ListStudents").Error(err)
	}

	var res []*student.Student
	if hits.TotalHits() == 0 {
		return res
	}
	for _, hit := range hits.Hits.Hits {
		singleRes := &student.Student{}
		err = json.Unmarshal(hit.Source, &singleRes)
		if err != nil {
			log.WithField("method", "ListStudents").Error(err)
			break
		}
		res = append(res, singleRes)
	}
	return res
}

func (r *request) CreateStudent() error {
	r.body.Id = convert.Int32(getNewId())
	_, err := client.GetClient().Index().
		Index(config.GlobalConfig.StudentsIndex).
		BodyJson(r.body).
		Id(getNewId()).
		Refresh("true").
		Do(context.Background())
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *request) UpdateStudent() error {
	_, err := client.GetClient().Index().
		Index(config.GlobalConfig.StudentsIndex).
		BodyJson(r.body).
		Id(convert.String(r.body.Id)).
		Refresh("true").
		Do(context.Background())
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *request) DeleteStudent() error {
	_, err := client.GetClient().Delete().
		Index(config.GlobalConfig.StudentsIndex).
		Id(convert.String(r.id)).
		Refresh("true").
		Do(context.Background())
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func getNewId() string {
	hits, err := client.GetClient().Search().
		Index(config.GlobalConfig.StudentsIndex).
		Query(elastic.NewBoolQuery()).
		Sort("id", false).
		Size(1).
		Do(context.Background())
	if err != nil {
		log.Error(err)
	}
	if hits.TotalHits() == 0 {
		return "0"
	}
	id := hits.Hits.Hits[0].Id
	return convert.String(convert.Int32(id) + 1)
}

package student

import (
	"context"
	"encoding/json"
	"strconv"

	"students_rest_api/config"
	client "students_rest_api/elastic"
	"students_rest_api/models"

	log "github.com/sirupsen/logrus"
)

func (r *request) GetStudent() *models.Student {
	hits, err := client.GetClient().Search().
		Query(r.buildSearchQuery()).
		Size(1).
		Index(config.GlobalConfig.StudentsIndex).
		Do(context.Background())
	if err != nil {
		log.WithField("method", "GetStudent").Error(err)
	}

	var res *models.Student
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

func (r *request) ListStudents() []*models.Student {
	hits, err := client.GetClient().Search().
		Query(r.buildSearchQuery()).
		Size(500).
		Index(config.GlobalConfig.StudentsIndex).
		Do(context.Background())
	if err != nil {
		log.WithField("method", "ListStudents").Error(err)
	}

	var res []*models.Student
	if hits.TotalHits() == 0 {
		return res
	}
	for _, hit := range hits.Hits.Hits {
		singleRes := &models.Student{}
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
	_, err := client.GetClient().Index().
		Index(config.GlobalConfig.StudentsIndex).
		BodyJson(r.bodyJSON).
		Id(strconv.Itoa(int(r.id))).
		Refresh("true").
		Do(context.Background())
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

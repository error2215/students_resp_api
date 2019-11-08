package student

import (
	"context"
	"encoding/json"

	"students_rest_api/config"
	client "students_rest_api/elastic"
	"students_rest_api/models"

	log "github.com/sirupsen/logrus"
)

func (r *request) GetStudent() *models.Student {

}

func (r *request) ListStudents() []*models.Student {
	hits, err := client.GetClient().Search().
		Query(r.buildQuery()).
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

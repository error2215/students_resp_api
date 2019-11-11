package elastic

import (
	"students_rest_api/config"

	"github.com/olivere/elastic/v7"

	log "github.com/sirupsen/logrus"
)

var client *elastic.Client

func init() {
	var err error
	client, err = elastic.NewClient(
		elastic.SetURL(config.GlobalConfig.ElasticAddress),
	)
	if err != nil {
		log.WithField("method", "elastic.client.init").Fatal(err)
	}

	//lessonsExists, err := client.IndexExists(config.GlobalConfig.LessonsIndex).Do(context.Background())
	//if err != nil {
	//	log.WithField("method", "elastic.client.init").Fatal(err)
	//}
	//if !lessonsExists {
	//	// need for sorting, if there are no this fields in start mapping, app will break down, other part will be created automatically
	//	mapping := `{"mappings":{"properties":{"time":{"type":"date"}, "id":{"type": "long"}}}}`
	//	_, err = client.CreateIndex(config.GlobalConfig.LessonsIndex).BodyString(mapping).Do(context.Background())
	//	if err != nil {
	//		log.WithField("method", "elastic.client.init").Fatal(err)
	//	}
	//}
	//studentsExists, err := client.IndexExists(config.GlobalConfig.StudentsIndex).Do(context.Background())
	//if err != nil {
	//	log.WithField("method", "elastic.client.init").Fatal(err)
	//}
	//if !studentsExists {
	//	// need for sorting, if there is no this field in start mapping, app will break down, other part will be created automatically
	//	mapping := `{"mappings":{"properties":{"id":{"type":"long"}}}}`
	//	_, err = client.CreateIndex(config.GlobalConfig.StudentsIndex).BodyString(mapping).Do(context.Background())
	//	if err != nil {
	//		log.WithField("method", "elastic.client.init").Fatal(err)
	//	}
	//}

	log.Info("Connection to ES cluster finished. Address: " + config.GlobalConfig.ElasticAddress)
}

func GetClient() *elastic.Client {
	return client
}

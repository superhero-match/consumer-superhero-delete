package es

import (
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v7"
)

// GetDocumentID fetches document source id, it is needed for update function.
func (es *ES) GetDocumentID(id string) (string, error) {
	q := elastic.NewTermQuery("superhero_id", id)

	searchResult, err := es.Client.Search().
		Index(es.Index).
		Query(q).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return "", err
	}

	if searchResult.TotalHits() > 0 {
		for _, hit := range searchResult.Hits.Hits {
			return hit.Id, nil
		}
	}

	return "", errors.New("no result")
}

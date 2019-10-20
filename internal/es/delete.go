package es

import (
	"context"
	"fmt"
)

// DeleteSuperhero saves newly registered superhero in Elasticsearch.
func (es *ES) DeleteSuperhero(id string) error {
	sourceID, err := es.GetDocumentID(id)
	if err != nil {
		return err
	}

	// Delete tweet with specified ID
	res, err := es.Client.Delete().
		Index(es.Index).
		Id(sourceID).
		Do(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("deleted id: ", res.Id)

	return nil
}

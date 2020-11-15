package _1_first_iteration

import (
	"context"
	"encoding/json"
	"errors"

	"gopkg.in/olivere/elastic.v6"
)

type Search struct {
	esClient *elastic.Client
}

func (s *Search) Do(ctx context.Context, customerName string) ([]*Order, error) {
	termQuery := elastic.NewTermQuery("customer", customerName)
	results, err := s.esClient.Search().
		Index("orderHistory").
		Query(termQuery).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	if results.Hits.TotalHits == 0 {
		return nil, errors.New("no results")
	}

	return s.unmarshalResults(results)
}

func (s *Search) unmarshalResults(results *elastic.SearchResult) ([]*Order, error) {
	var output []*Order

	for _, hit := range results.Hits.Hits {
		order := &Order{}
		err := json.Unmarshal(*hit.Source, order)
		if err != nil {
			return nil, err
		}

		output = append(output, order)
	}

	return output, nil
}

type Order struct {
	OrderID  string
	Customer string
}

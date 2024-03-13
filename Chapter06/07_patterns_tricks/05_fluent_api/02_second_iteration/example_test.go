package _2_second_iteration

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/olivere/elastic.v6"
)

func TestSearch_Do(t *testing.T) {
	// inputs
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	customerName := "Tricia"

	scenarios := []struct {
		desc              string
		historySearchStub historySearchFunc
		expected          []*Order
		expectAnErr       bool
	}{
		{
			desc: "Happy path",
			historySearchStub: func(ctx context.Context, esClient *elastic.Client, termQuery *elastic.TermQuery) (*elastic.SearchResult, error) {
				order := &Order{
					OrderID:  "ABC123",
					Customer: "Tricia",
				}
				payload, err := json.Marshal(order)
				require.NoError(t, err)

				source := json.RawMessage(payload)

				result := &elastic.SearchResult{
					Hits: &elastic.SearchHits{
						TotalHits: 1,
						Hits: []*elastic.SearchHit{
							{
								Source: &source,
							},
						},
					},
				}

				return result, nil
			},
			expected: []*Order{
				{
					OrderID:  "ABC123",
					Customer: "Tricia",
				},
			},
			expectAnErr: false,
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			// call object under test
			search := &Search{
				historySearch: scenario.historySearchStub,
			}

			result, resultErr := search.Do(ctx, customerName)

			// validation
			require.Equal(t, scenario.expectAnErr, resultErr != nil, "expected error. err: %s", resultErr)
			assert.Equal(t, scenario.expected, result, "expected result")
		})
	}
}

package query

import (
	"error-group/database"
	"error-group/eg"
	"errors"
	"sync"
)

type Result struct {
	data []database.Data
	err  error
}

func DistributedQuery(shards []database.Database, query string) (*[]database.Data, error) {
	_ = query
	mutex := sync.Mutex{}
	eg, ctx := eg.NewErrorGroup()
	results := make([]database.Data, 0, len(shards))

	for _, shard := range shards {
		eg.Go(func() error {
			resultCh := make(chan Result)

			go func() {
				data, err := shard.Get()
				resultCh <- Result{data, err}
			}()

			select {
			case res := <-resultCh:
				if res.err != nil {
					return res.err
				}

				mutex.Lock()
				results = append(results, res.data...)
				mutex.Unlock()
			case <-ctx.Done():
				return eg.Error
			}

			return nil
		})
	}

	err := eg.Wait()
	if err != nil {
		return nil, errors.New("ERROR")
	}

	return &results, nil
}

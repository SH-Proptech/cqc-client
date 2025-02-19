package cqc

import (
	"context"
	"fmt"
)

// GetProviders fetches a single page of providers.
func (c *Client) GetProviders(path string) (*ProvidersResponse, error) {
	providersResponse, err := Get[ProvidersResponse](c, path)
	if err != nil {
		return nil, fmt.Errorf("failed to get providers: %v", err)
	}
	return providersResponse, nil
}

// GetProvider fetches details of a single provider by ID.
func (c *Client) GetProvider(id string) (*Provider, error) {
	providerDetails, err := Get[Provider](c, pathProviders+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("failed to get provider details for ID %s: %v", id, err)
	}
	return providerDetails, nil
}

// GetAllProviders fetches all providers and calls a callback function for each provider.
func (c *Client) GetAllProviders(ctx context.Context, concurrency, pageSize, startFrom int, callback func(*Provider) error) error {

	if startFrom < 1 {
		startFrom = 1
	}

	nextPage := fmt.Sprintf("%s?perPage=%d&page=%d", pathProviders, pageSize, startFrom)
	workerPool := NewWorkerPool(concurrency)
	workerPool.Start(ctx)

	for nextPage != "" {
		// Fetch a page of providers
		providersResponse, err := c.GetProviders(nextPage)
		if err != nil {
			return err
		}

		// Submit tasks for each provider
		for _, providerItem := range providersResponse.Providers {
			providerID := providerItem.ProviderID
			workerPool.Submit(func(ctx context.Context) error {
				provider, err := c.GetProvider(providerID)
				if err != nil {
					return err
				}
				// Call the callback function
				return callback(provider)
			})
		}

		// Move to the next page (if exists)
		nextPage = providersResponse.NextPageUri
	}

	// Stop the worker pool and collect errors
	workerPool.Stop()
	errors := workerPool.CollectErrors()

	if len(errors) > 0 {
		return fmt.Errorf("errors occurred during provider fetching: %v", errors)
	}

	return nil
}

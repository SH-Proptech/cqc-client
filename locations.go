package cqc

import (
	"context"
	"fmt"

	"github.com/sh-proptech/pimp/pkg/logger"
)

// GetLocations fetches a single page of locations.
func (c *Client) GetLocations(path string) (*LocationsResponse, error) {
	locationsResponse, err := Get[LocationsResponse](c, path)
	if err != nil {
		return nil, fmt.Errorf("failed to get locations: %v", err)
	}
	return locationsResponse, nil
}

// GetLocation fetches details of a single location by ID.
func (c *Client) GetLocation(id string) (*Location, error) {
	locationDetails, err := Get[Location](c, pathLocations+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("failed to get location details for ID %s: %v", id, err)
	}
	if locationDetails.LocationID == "" {
		logger.Sugar.Info("Location not found:", id)
		logger.Pjson(locationDetails)
	}
	return locationDetails, nil
}

// GetAllLocations fetches all locations and retrieves full details concurrently.
func (c *Client) GetAllLocations(ctx context.Context, concurrency, pageSize, startFrom int, callback func(*Location) error) error {

	if startFrom < 1 {
		startFrom = 1
	}

	nextPage := fmt.Sprintf("%s?perPage=%d&page=%d", pathLocations, pageSize, startFrom)
	workerPool := NewWorkerPool(concurrency)
	workerPool.Start(ctx)

	for nextPage != "" {
		// Fetch a page of locations
		locationsResponse, err := c.GetLocations(nextPage)
		if err != nil {
			return err
		}

		// Submit tasks for each location
		for _, locationItem := range locationsResponse.Locations {
			locationID := locationItem.LocationID
			workerPool.Submit(func(ctx context.Context) error {
				location, err := c.GetLocation(locationID)
				if err != nil {
					return err
				}
				// Call the callback function
				return callback(location)
			})
		}

		// Move to the next page (if exists)
		nextPage = locationsResponse.NextPageUri
	}

	// Stop the worker pool and collect errors
	workerPool.Stop()
	errors := workerPool.CollectErrors()

	if len(errors) > 0 {
		return fmt.Errorf("errors occurred during location fetching: %v", errors)
	}

	return nil
}

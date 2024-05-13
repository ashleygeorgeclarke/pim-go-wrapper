package pim

import (
	"context"
	"fmt"
	"net/http"
)

type (
	MatrixDimensions service
	MatrixDimension  struct {
		DimensionID    int64             `json:"dimension_id"`
		DimensionName  map[string]string `json:"dimension_name,omitempty"`
		DimensionOrder int               `json:"dimension_order"`
		Status         string            `json:"status"`
		AddedByChangedBy
	}
)

func (s *MatrixDimensions) Read(ctx context.Context, opts *ListOptions) (*[]MatrixDimension, *http.Response, error) {
	urlStr := "matrix/dimension"
	u, err := addOptions(urlStr, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	dataResp := new([]MatrixDimension)
	resp, err := s.client.Do(ctx, req, dataResp)
	return dataResp, resp, err
}

func (s *MatrixDimensions) Delete(ctx context.Context, dimensionID int) (*IDResponse, *http.Response, error) {
	u := fmt.Sprintf("matrix/dimension/%d", dimensionID)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, nil, err
	}

	id := new(IDResponse)
	resp, err := s.client.Do(ctx, req, id)
	return id, resp, err
}

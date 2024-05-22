package pim

import (
	"context"
	"net/http"
)

func (p *ProductDeletedIds) Get(ctx context.Context) ([]int, *http.Response, error) {
	u, err := addOptions(p.Path, nil)
	if err != nil {
		return nil, nil, err
	}
	req, err := p.service.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}
	dataResp := &[]int{}
	resp, err := p.client.Do(ctx, req, dataResp)
	return *dataResp, resp, err
}

package pim

import (
	"context"
	"net/http"

	"github.com/erply/pim-go-wrapper/pkg/pim/models"
)

func (p *Product) Get(ctx context.Context, opts *ListOptions) (*[]models.Product, *http.Response, error) {
	u, err := addOptions(p.Path, opts)
	if err != nil {
		return nil, nil, err
	}
	req, err := p.service.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}
	dataResp := &[]models.Product{}
	resp, err := p.client.Do(ctx, req, dataResp)
	return dataResp, resp, err
}

func (p *Product) Post(ctx context.Context, product *models.ProductRequest, trimInput bool, generateCodeAutomatically bool, generateCode2Automatically bool) (*models.IDResponse, *http.Response, error) {
	req, err := p.service.client.NewRequest(http.MethodPost, p.Path, product)
	if err != nil {
		return nil, nil, err
	}
	id := new(models.IDResponse)
	resp, err := p.client.Do(ctx, req, id)
	return id, resp, err
}

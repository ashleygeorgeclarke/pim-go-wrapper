package pim

import (
	"context"
	"net/http"

	"github.com/erply/pim-go-wrapper/pkg/pim/models"
)

func (p *ProductLinkedProducts) Post(ctx context.Context, product *models.LinkedProductCreateRequest) (*models.IDResponse, *http.Response, error) {
	req, err := p.service.client.NewRequest(http.MethodPost, p.Path, product)
	if err != nil {
		return nil, nil, err
	}
	id := new(models.IDResponse)
	resp, err := p.client.Do(ctx, req, id)
	return id, resp, err
}

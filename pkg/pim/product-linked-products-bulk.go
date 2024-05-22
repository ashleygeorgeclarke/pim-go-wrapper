package pim

import (
	"context"
	"net/http"

	"github.com/erply/pim-go-wrapper/pkg/pim/models"
)

func (p *ProductLinkedProductsBulk) Post(ctx context.Context, links *[]models.LinkedProductCreateRequest) (*models.IDResponse, *http.Response, error) {
	panic("implement me")
	bulkBody := createBulkRequestStruct(links)
	req, err := p.service.client.NewRequest(http.MethodPost, p.Path, bulkBody)
	if err != nil {
		return nil, nil, err
	}
	id := new(models.IDResponse)
	resp, err := p.client.Do(ctx, req, id)
	return id, resp, err
}

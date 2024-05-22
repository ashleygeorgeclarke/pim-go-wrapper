package pim

import (
	"context"
	"net/http"

	"github.com/erply/pim-go-wrapper/pkg/pim/models"
)

func (p *ProductBulk) Post(ctx context.Context, products *[]models.ProductRequest) (*models.BulkResponseWithResults, *http.Response, error) {
	bulkBody := createBulkRequestStruct(products)

	req, err := p.service.client.NewRequest(http.MethodPost, p.Path, bulkBody)
	if err != nil {
		return nil, nil, err
	}
	result := new(models.BulkResponseWithResults)
	resp, err := p.client.Do(ctx, req, result)
	return result, resp, err
}

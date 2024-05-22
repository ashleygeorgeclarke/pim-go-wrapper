package pim

import (
	"context"
	"net/http"
)

func (p *ProductCodes) Get(ctx context.Context) (*AutoCodes, *http.Response, error) {
	urlStr := "product/codes"
	req, err := p.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	code := new(AutoCodes)
	resp, err := p.client.Do(ctx, req, code)
	return code, resp, err
}

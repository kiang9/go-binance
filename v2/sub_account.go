package binance

import (
	"context"
	"encoding/json"
)

type CreateSubAccountService struct {
	c          *Client
	subAccount string
}

func (s *CreateSubAccountService) Do(ctx context.Context, opts ...RequestOption) (string, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/sub-account/virtualSubAccount",
		secType:  secTypeSigned,
	}
	r.setFormParam("subAccountString", s.subAccount)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}
	var res = struct {
		Email string `json:"email"`
	}{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return "", err
	}
	return res.Email, nil
}

func (s *CreateSubAccountService) SetSubAccount(sa string) *CreateSubAccountService {
	s.subAccount = sa
	return s
}

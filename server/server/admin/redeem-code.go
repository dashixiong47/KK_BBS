package admin

import (
	"errors"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/server/data"
)

type RedeemCodeServer struct {
}

func (r *RedeemCodeServer) Create(quantity, _type, number int) ([]string, error) {
	var docs models.RedeemCode
	code, err := data.CreateRedeemCode()
	if err != nil {
		return nil, errors.New("create_redeem_code_error")
	}
	docs.Code = code
	return []string{code}, nil

}

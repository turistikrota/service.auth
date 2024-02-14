package claims

import (
	"context"

	"github.com/turistikrota/service.auth/config"
	"github.com/turistikrota/service.auth/protos/account"
	"github.com/turistikrota/service.auth/protos/business"
	"github.com/turistikrota/service.shared/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func callAccountClaims(ctx context.Context, cnf config.Rpc, userUUID string) ([]jwt.UserClaimAccount, error) {
	var opt grpc.DialOption
	if !cnf.AccountUsesSsl {
		opt = grpc.WithTransportCredentials(insecure.NewCredentials())
	} else {
		opt = grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	}
	conn, err := grpc.Dial(cnf.AccountHost, opt)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := account.NewAccountListServiceClient(conn)
	response, err := c.ListAsClaim(ctx, &account.AccountListAsClaimRequest{
		UserId: userUUID,
	})
	if err != nil {
		return nil, err
	}
	result := make([]jwt.UserClaimAccount, 0, len(response.Accounts))
	for _, acc := range response.Accounts {
		result = append(result, jwt.UserClaimAccount{
			ID:   acc.Id,
			Name: acc.Name,
		})
	}
	return result, nil
}

func callBusinessClaims(ctx context.Context, cnf config.Rpc, userUUID string) ([]jwt.UserClaimBusiness, error) {
	var opt grpc.DialOption
	if !cnf.BusinessUsesSsl {
		opt = grpc.WithTransportCredentials(insecure.NewCredentials())
	} else {
		opt = grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	}
	conn, err := grpc.Dial(cnf.BusinessHost, opt)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := business.NewBusinessListServiceClient(conn)
	response, err := c.ListAsClaim(ctx, &business.BusinessListAsClaimRequest{
		UserId: userUUID,
	})
	if err != nil {
		return nil, err
	}
	result := make([]jwt.UserClaimBusiness, 0, len(response.Business))
	for _, acc := range response.Business {
		result = append(result, jwt.UserClaimBusiness{
			UUID:        acc.Uuid,
			AccountName: acc.AccountName,
			NickName:    acc.NickName,
			Roles:       acc.Roles,
		})
	}
	return result, nil
}

func Fetch(ctx context.Context, cnf config.Rpc, userUUID string) ([]jwt.UserClaimAccount, []jwt.UserClaimBusiness, error) {
	acc, err := callAccountClaims(ctx, cnf, userUUID)
	if err != nil {
		return nil, nil, err
	}
	bus, err := callBusinessClaims(ctx, cnf, userUUID)
	if err != nil {
		return nil, nil, err
	}
	return acc, bus, nil
}

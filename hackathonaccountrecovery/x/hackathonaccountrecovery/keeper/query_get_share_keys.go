package keeper

import (
	"context"
	"fmt"
	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashicorp/vault/shamir"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetShareKeys(goCtx context.Context, req *types.QueryGetShareKeysRequest) (*types.QueryGetShareKeysResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	secret := []byte(req.Key)
	totalShares := int(req.Totalshares)
	threshold := int(req.Minshares)

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	shares, err := shamir.Split(secret, totalShares, threshold)
	if err != nil {
		log.Fatalf("Error splitting secret: %v", err)
	}

	return &types.QueryGetShareKeysResponse{Secretkey: sharesToHexStrings(shares)}, nil
}

// Helper function to convert byte slices to hex string
func sharesToHexStrings(shares [][]byte) []string {
	hexShares := make([]string, len(shares))
	for i, share := range shares {
		hexShares[i] = fmt.Sprintf("%x", share) // convert to hex string
	}
	return hexShares
}

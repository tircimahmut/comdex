package keeper

import (
	"context"
	"github.com/comdex-official/comdex/x/auctionsV2/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = QueryServer{}

type QueryServer struct {
	Keeper
}

func NewQueryServer(k Keeper) types.QueryServer {
	return &QueryServer{
		Keeper: k,
	}
}

func (q QueryServer) Auction(c context.Context, req *types.QueryAuctionRequest) (res *types.QueryAuctionResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		ctx  = sdk.UnwrapSDKContext(c)
		item types.Auction
	)
	if req.History {
		auctionHistorical, _ := q.GetAuctionHistorical(ctx, req.AuctionId)
		item = *auctionHistorical.AuctionHistorical
	} else {
		item, err = q.GetAuction(ctx, req.AuctionId)
	}
	if err != nil {
		return nil, err
	}

	return &types.QueryAuctionResponse{
		Auction: item,
	}, nil
}

func (q QueryServer) Auctions(c context.Context, req *types.QueryAuctionsRequest) (*types.QueryAuctionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		items []types.Auction
		ctx   = sdk.UnwrapSDKContext(c)
		key   []byte
	)
	if req.History {
		key = types.AuctionHistoricalKeyPrefix
	} else {
		key = types.AuctionKeyPrefix
	}
	pagination, err := query.FilteredPaginate(
		prefix.NewStore(q.Store(ctx), key),
		req.Pagination,
		func(_, value []byte, accumulate bool) (bool, error) {
			var item types.Auction
			if err := q.cdc.Unmarshal(value, &item); err != nil {
				return false, err
			}

			if accumulate {
				items = append(items, item)
			}

			return true, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAuctionsResponse{
		Auctions:   items,
		Pagination: pagination,
	}, nil
}

func (q QueryServer) Bids(c context.Context, req *types.QueryBidsRequest) (*types.QueryBidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		ctx   = sdk.UnwrapSDKContext(c)
		key   []byte
		items []types.Bid
	)
	if req.History {
		//TODO: add historical key
		//key = types.HistoryUserAuctionTypeKey(req.Bidder, req.AppId, types.DutchString)
	} else {
		key = types.UserBidKeyPrefix

	}

	pagination, err := query.FilteredPaginate(
		prefix.NewStore(q.Store(ctx), key),
		req.Pagination,
		func(_, value []byte, accumulate bool) (bool, error) {
			var item types.Bid
			if err := q.cdc.Unmarshal(value, &item); err != nil {
				return false, err
			}

			if accumulate {
				items = append(items, item)
			}

			return true, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryBidsResponse{
		Bidder:     req.Bidder,
		Bids:       items,
		Pagination: pagination,
	}, nil
}

func (q QueryServer) AuctionParams(c context.Context, req *types.QueryAuctionParamsRequest) (*types.QueryAuctionParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		ctx  = sdk.UnwrapSDKContext(c)
		item types.AuctionParams
	)

	item, found := q.GetAuctionParams(ctx)
	if !found {
		return nil, types.ErrAuctionParamsNotFound
	}

	return &types.QueryAuctionParamsResponse{
		AuctionParams: item,
	}, nil
}

func (q QueryServer) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: q.GetParams(ctx)}, nil
}

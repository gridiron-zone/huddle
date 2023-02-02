package app

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/codec"

	wasmhuddle "github.com/gridiron-zone/huddle/cosmwasm"
	postskeeper "github.com/gridiron-zone/huddle/x/posts/keeper"
	postswasm "github.com/gridiron-zone/huddle/x/posts/wasm"
	profileskeeper "github.com/gridiron-zone/huddle/x/profiles/keeper"
	profileswasm "github.com/gridiron-zone/huddle/x/profiles/wasm"
	reactionskeeper "github.com/gridiron-zone/huddle/x/reactions/keeper"
	reactionswasm "github.com/gridiron-zone/huddle/x/reactions/wasm"
	relationshipskeeper "github.com/gridiron-zone/huddle/x/relationships/keeper"
	relationshipswasm "github.com/gridiron-zone/huddle/x/relationships/wasm"
	reportskeeper "github.com/gridiron-zone/huddle/x/reports/keeper"
	reportswasm "github.com/gridiron-zone/huddle/x/reports/wasm"
	subspaceskeeper "github.com/gridiron-zone/huddle/x/subspaces/keeper"
	subspaceswasm "github.com/gridiron-zone/huddle/x/subspaces/wasm"
)

const (
	// DefaultHuddleInstanceCost is how much SDK gas we charge each time we load a WASM instance
	DefaultHuddleInstanceCost uint64 = 60_000
	// DefaultHuddleCompileCost is how much SDK gas is charged *per byte* for compiling WASM code
	DefaultHuddleCompileCost uint64 = 2
)

// HuddleWasmGasRegister is defaults plus a custom compile amount
func HuddleWasmGasRegister() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultHuddleInstanceCost
	gasConfig.CompileCost = DefaultHuddleCompileCost

	return gasConfig
}

func NewHuddleWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(HuddleWasmGasRegister())
}

// NewHuddleCustomQueryPlugin initialize the custom querier to handle huddle queries for contracts
func NewHuddleCustomQueryPlugin(
	cdc codec.Codec,
	profilesKeeper profileskeeper.Keeper,
	subspacesKeeper subspaceskeeper.Keeper,
	relationshipsKeeper relationshipskeeper.Keeper,
	postsKeeper postskeeper.Keeper,
	reportsKeeper reportskeeper.Keeper,
	reactionsKeeper reactionskeeper.Keeper,
) wasm.QueryPlugins {
	queriers := map[string]wasmhuddle.Querier{
		wasmhuddle.QueryRouteProfiles:      profileswasm.NewProfilesWasmQuerier(profilesKeeper, cdc),
		wasmhuddle.QueryRouteSubspaces:     subspaceswasm.NewSubspacesWasmQuerier(subspacesKeeper, cdc),
		wasmhuddle.QueryRouteRelationships: relationshipswasm.NewRelationshipsWasmQuerier(relationshipsKeeper, cdc),
		wasmhuddle.QueryRoutePosts:         postswasm.NewPostsWasmQuerier(postsKeeper, cdc),
		wasmhuddle.QueryRouteReports:       reportswasm.NewReportsWasmQuerier(reportsKeeper, cdc),
		wasmhuddle.QueryRouteReactions:     reactionswasm.NewReactionsWasmQuerier(reactionsKeeper, cdc),
		// add other modules querier here
	}

	querier := wasmhuddle.NewQuerier(queriers)

	return wasm.QueryPlugins{
		Custom: querier.QueryCustom,
	}
}

// NewHuddleCustomMessageEncoder initialize the custom message encoder to huddle app for contracts
func NewHuddleCustomMessageEncoder(cdc codec.Codec) wasm.MessageEncoders {
	// Initialization of custom Huddle messages for contracts
	parserRouter := wasmhuddle.NewParserRouter()
	parsers := map[string]wasmhuddle.MsgParserInterface{
		wasmhuddle.WasmMsgParserRouteProfiles:      profileswasm.NewWasmMsgParser(cdc),
		wasmhuddle.WasmMsgParserRouteSubspaces:     subspaceswasm.NewWasmMsgParser(cdc),
		wasmhuddle.WasmMsgParserRouteRelationships: relationshipswasm.NewWasmMsgParser(cdc),
		wasmhuddle.WasmMsgParserRoutePosts:         postswasm.NewWasmMsgParser(cdc),
		wasmhuddle.WasmMsgParserRouteReports:       reportswasm.NewWasmMsgParser(cdc),
		wasmhuddle.WasmMsgParserRouteReactions:     reactionswasm.NewWasmMsgParser(cdc),
		// add other modules parsers here
	}

	parserRouter.Parsers = parsers
	return wasm.MessageEncoders{
		Custom: parserRouter.ParseCustom,
	}
}

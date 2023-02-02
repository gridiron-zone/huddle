package provider

import (
	"github.com/gridiron-zone/huddle/app/huddle/cmd/chainlink/builder"
	multibuilder "github.com/gridiron-zone/huddle/app/huddle/cmd/chainlink/builder/multi"
	singlebuilder "github.com/gridiron-zone/huddle/app/huddle/cmd/chainlink/builder/single"
	multigetter "github.com/gridiron-zone/huddle/app/huddle/cmd/chainlink/getter/multi"
	singlegetter "github.com/gridiron-zone/huddle/app/huddle/cmd/chainlink/getter/single"
)

// DefaultChainLinkJSONBuilderProvider returns the default ChainLinkJSONBuilder provider implementation
func DefaultChainLinkJSONBuilderProvider(owner string, isSingleAccount bool) builder.ChainLinkJSONBuilder {
	if isSingleAccount {
		return singlebuilder.NewAccountChainLinkJSONBuilder(owner, singlegetter.NewChainLinkJSONInfoGetter())
	}
	return multibuilder.NewAccountChainLinkJSONBuilder(multigetter.NewChainLinkJSONInfoGetter())
}

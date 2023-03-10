package types

// DONTCOVER

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// NewQueryProfileRequest returns a new QueryProfileRequest containing the given data
func NewQueryProfileRequest(user string) *QueryProfileRequest {
	return &QueryProfileRequest{
		User: user,
	}
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (r *QueryProfileResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	if r.Profile != nil {
		var profile authtypes.AccountI
		return unpacker.UnpackAny(r.Profile, &profile)
	}
	return nil
}

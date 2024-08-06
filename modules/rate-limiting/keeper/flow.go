package keeper

import (
	"github.com/cosmos/ibc-apps/modules/rate-limiting/v8/types"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// The total value on a given path (aka, the denominator in the percentage calculation)
// is the total supply of the given denom
func (k Keeper) GetChannelValue(ctx sdk.Context, denom string) sdkmath.Int {
	return k.bankKeeper.GetSupply(ctx, denom).Amount
}

// Adds an amount to the flow in either the SEND or RECV direction
func (k Keeper) UpdateFlow(rateLimit types.RateLimit, direction types.PacketDirection, amount sdkmath.Int) error {
	switch direction {
	case types.PACKET_SEND:
		return rateLimit.Flow.AddOutflow(amount, *rateLimit.Quota)
	case types.PACKET_RECV:
		return rateLimit.Flow.AddInflow(amount, *rateLimit.Quota)
	default:
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid packet direction (%s)", direction.String())
	}
}

// Checks whether the given packet will exceed the rate limit
// Called by OnRecvPacket and OnSendPacket
func (k Keeper) CheckRateLimitAndUpdateFlow(
	ctx sdk.Context,
	direction types.PacketDirection,
	packetInfo RateLimitedPacketInfo,
) (updatedFlow bool, err error) {
	denom := packetInfo.Denom
	channelId := packetInfo.ChannelID
	amount := packetInfo.Amount

	// First check if the denom is blacklisted
	if k.IsDenomBlacklisted(ctx, denom) {
		err := errorsmod.Wrapf(types.ErrDenomIsBlacklisted, "denom %s is blacklisted", denom)
		EmitTransferDeniedEvent(ctx, types.EventBlacklistedDenom, denom, channelId, direction, amount, err)
		return false, err
	}

	// If there's no rate limit yet for this denom, no action is necessary
	rateLimit, found := k.GetRateLimit(ctx, denom, channelId)
	if !found {
		return false, nil
	}

	// Check if the sender/receiver pair is whitelisted
	// If so, return a success without modifying the quota
	if k.IsAddressPairWhitelisted(ctx, packetInfo.Sender, packetInfo.Receiver) {
		return false, nil
	}

	// Update the flow object with the change in amount
	if err := k.UpdateFlow(rateLimit, direction, amount); err != nil {
		// If the rate limit was exceeded, emit an event
		EmitTransferDeniedEvent(ctx, types.EventRateLimitExceeded, denom, channelId, direction, amount, err)
		return false, err
	}

	// If there's no quota error, update the rate limit object in the store with the new flow
	k.SetRateLimit(ctx, rateLimit)

	return true, nil
}

// If a SendPacket fails or times out, undo the outflow increment that happened during the send
func (k Keeper) UndoSendPacket(ctx sdk.Context, channelId string, sequence uint64, denom string, amount sdkmath.Int) error {
	rateLimit, found := k.GetRateLimit(ctx, denom, channelId)
	if !found {
		return nil
	}

	// If the packet was sent during this quota, decrement the outflow
	// Otherwise, it can be ignored
	if k.CheckPacketSentDuringCurrentQuota(ctx, channelId, sequence) {
		rateLimit.Flow.Outflow = rateLimit.Flow.Outflow.Sub(amount)
		k.SetRateLimit(ctx, rateLimit)

		k.RemovePendingSendPacket(ctx, channelId, sequence)
	}

	return nil
}

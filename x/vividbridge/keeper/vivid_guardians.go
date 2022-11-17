package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"vivid-bridge/x/vividbridge/types"
)

// GetVividGuardiansCount get the total number of vividGuardians
func (k Keeper) GetVividGuardiansCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VividGuardiansCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetVividGuardiansCount set the total number of vividGuardians
func (k Keeper) SetVividGuardiansCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VividGuardiansCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendVividGuardians appends a vividGuardians in the store with a new id and update the count
func (k Keeper) AppendVividGuardians(
	ctx sdk.Context,
	vividGuardians types.VividGuardians,
) uint64 {
	// Create the vividGuardians
	count := k.GetVividGuardiansCount(ctx)

	// Set the ID of the appended value
	vividGuardians.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VividGuardiansKey))
	appendedValue := k.cdc.MustMarshal(&vividGuardians)
	store.Set(GetVividGuardiansIDBytes(vividGuardians.Id), appendedValue)

	// Update vividGuardians count
	k.SetVividGuardiansCount(ctx, count+1)

	return count
}

// SetVividGuardians set a specific vividGuardians in the store
func (k Keeper) SetVividGuardians(ctx sdk.Context, vividGuardians types.VividGuardians) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VividGuardiansKey))
	b := k.cdc.MustMarshal(&vividGuardians)
	store.Set(GetVividGuardiansIDBytes(vividGuardians.Id), b)
}

// GetVividGuardians returns a vividGuardians from its id
func (k Keeper) GetVividGuardians(ctx sdk.Context, id uint64) (val types.VividGuardians, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VividGuardiansKey))
	b := store.Get(GetVividGuardiansIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVividGuardians removes a vividGuardians from the store
func (k Keeper) RemoveVividGuardians(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VividGuardiansKey))
	store.Delete(GetVividGuardiansIDBytes(id))
}

// GetAllVividGuardians returns all vividGuardians
func (k Keeper) GetAllVividGuardians(ctx sdk.Context) (list []types.VividGuardians) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VividGuardiansKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VividGuardians
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetVividGuardiansIDBytes returns the byte representation of the ID
func GetVividGuardiansIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetVividGuardiansIDFromBytes returns ID in uint64 format from a byte array
func GetVividGuardiansIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

package genesis

import (
	"context"

	cbor "github.com/ipfs/go-ipld-cbor"

	"github.com/filecoin-project/go-state-types/big"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/actors/builtin/cron"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupCronActor(ctx context.Context, bs bstore.Blockstore, av actors.Version) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)
	st, err := cron.MakeState(adt.WrapStore(ctx, cbor.NewCborStore(bs)), av)
	if err != nil {
		return nil, err
	}

	statecid, err := cst.Put(ctx, st.GetState())
	if err != nil {
		return nil, err
	}

	actcid, err := builtin.GetCronActorCodeID(av)
	if err != nil {
		return nil, err
	}

	act := &types.Actor{
		Code:    actcid,
		Head:    statecid,
		Balance: big.Zero(),
	}

	return act, nil
}

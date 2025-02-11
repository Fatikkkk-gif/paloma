package keeper

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/VolumeFi/whoops"
	sdk "github.com/cosmos/cosmos-sdk/types"
	wasmutil "github.com/palomachain/paloma/util/wasm"
	"github.com/palomachain/paloma/x/scheduler/types"
)

type ExecuteJobWasmEvent struct {
	JobID   string `json:"job_id"`
	Payload []byte `json:"payload"`
	Sender  string `json:"sender"`
}

func (e ExecuteJobWasmEvent) valid() error {
	if len(e.JobID) == 0 {
		return whoops.String("you must provide a jobID")
	}
	if len(e.Payload) == 0 {
		return whoops.String("payload bytes is empty")
	}
	// todo: add more in the future
	return nil
}

func (k Keeper) UnmarshallJob(msg []byte) (ExecuteJobWasmEvent, error) {
	var executeMsg ExecuteJobWasmEvent
	err := json.Unmarshal(msg, &executeMsg)

	hexString := hex.EncodeToString(executeMsg.Payload)

	executeMsg.Payload = []byte(fmt.Sprintf("{\"hexPayload\":\"%s\"}", hexString))

	return executeMsg, err
}

func (k Keeper) ExecuteWasmJobEventListener() wasmutil.MessengerFnc {
	return func(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
		executeMsg, err := k.UnmarshallJob(msg.Custom)
		if err != nil {
			return nil, nil, err
		}

		k.Logger(ctx).Info("got a request to schedule a job via CosmWasm smart contract", "job_id", executeMsg.JobID)
		if err = executeMsg.valid(); err != nil {
			return nil, nil, whoops.Wrap(err, types.ErrWasmExecuteMessageNotValid)
		}

		err = k.ScheduleNow(ctx, executeMsg.JobID, executeMsg.Payload, nil, contractAddr)
		if err != nil {
			return nil, nil, err
		}

		return nil, nil, nil
	}
}

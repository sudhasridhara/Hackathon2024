package hackathonaccountrecovery

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"hackathonaccountrecovery/testutil/sample"
	hackathonaccountrecoverysimulation "hackathonaccountrecovery/x/hackathonaccountrecovery/simulation"
	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"
)

// avoid unused import issue
var (
	_ = hackathonaccountrecoverysimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgAddAccountInfo = "op_weight_msg_add_account_info"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddAccountInfo int = 100

	opWeightMsgUpdatetotalshares = "op_weight_msg_updatetotalshares"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatetotalshares int = 100

	opWeightMsgInitiateDeathValidation = "op_weight_msg_initiate_death_validation"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitiateDeathValidation int = 100

	opWeightMsgValidateDeath = "op_weight_msg_validate_death"
	// TODO: Determine the simulation weight value
	defaultWeightMsgValidateDeath int = 100

	opWeightMsgCreateMultisite = "op_weight_msg_create_multisite"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMultisite int = 100

	opWeightMsgUpdateMultisite = "op_weight_msg_update_multisite"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMultisite int = 100

	opWeightMsgUpdateEmailId = "op_weight_msg_update_email_id"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateEmailId int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	hackathonaccountrecoveryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&hackathonaccountrecoveryGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAddAccountInfo int
	simState.AppParams.GetOrGenerate(opWeightMsgAddAccountInfo, &weightMsgAddAccountInfo, nil,
		func(_ *rand.Rand) {
			weightMsgAddAccountInfo = defaultWeightMsgAddAccountInfo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddAccountInfo,
		hackathonaccountrecoverysimulation.SimulateMsgAddAccountInfo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatetotalshares int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatetotalshares, &weightMsgUpdatetotalshares, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatetotalshares = defaultWeightMsgUpdatetotalshares
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatetotalshares,
		hackathonaccountrecoverysimulation.SimulateMsgUpdatetotalshares(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgInitiateDeathValidation int
	simState.AppParams.GetOrGenerate(opWeightMsgInitiateDeathValidation, &weightMsgInitiateDeathValidation, nil,
		func(_ *rand.Rand) {
			weightMsgInitiateDeathValidation = defaultWeightMsgInitiateDeathValidation
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInitiateDeathValidation,
		hackathonaccountrecoverysimulation.SimulateMsgInitiateDeathValidation(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgValidateDeath int
	simState.AppParams.GetOrGenerate(opWeightMsgValidateDeath, &weightMsgValidateDeath, nil,
		func(_ *rand.Rand) {
			weightMsgValidateDeath = defaultWeightMsgValidateDeath
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgValidateDeath,
		hackathonaccountrecoverysimulation.SimulateMsgValidateDeath(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateMultisite int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateMultisite, &weightMsgCreateMultisite, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMultisite = defaultWeightMsgCreateMultisite
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMultisite,
		hackathonaccountrecoverysimulation.SimulateMsgCreateMultisite(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMultisite int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateMultisite, &weightMsgUpdateMultisite, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMultisite = defaultWeightMsgUpdateMultisite
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMultisite,
		hackathonaccountrecoverysimulation.SimulateMsgUpdateMultisite(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateEmailId int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateEmailId, &weightMsgUpdateEmailId, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateEmailId = defaultWeightMsgUpdateEmailId
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateEmailId,
		hackathonaccountrecoverysimulation.SimulateMsgUpdateEmailId(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddAccountInfo,
			defaultWeightMsgAddAccountInfo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hackathonaccountrecoverysimulation.SimulateMsgAddAccountInfo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdatetotalshares,
			defaultWeightMsgUpdatetotalshares,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hackathonaccountrecoverysimulation.SimulateMsgUpdatetotalshares(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgInitiateDeathValidation,
			defaultWeightMsgInitiateDeathValidation,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hackathonaccountrecoverysimulation.SimulateMsgInitiateDeathValidation(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgValidateDeath,
			defaultWeightMsgValidateDeath,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hackathonaccountrecoverysimulation.SimulateMsgValidateDeath(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateMultisite,
			defaultWeightMsgCreateMultisite,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hackathonaccountrecoverysimulation.SimulateMsgCreateMultisite(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateMultisite,
			defaultWeightMsgUpdateMultisite,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hackathonaccountrecoverysimulation.SimulateMsgUpdateMultisite(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateEmailId,
			defaultWeightMsgUpdateEmailId,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hackathonaccountrecoverysimulation.SimulateMsgUpdateEmailId(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

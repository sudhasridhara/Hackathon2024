package hackathonaccountrecovery

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "hackathonaccountrecovery/api/hackathonaccountrecovery/hackathonaccountrecovery"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "ShowAccount",
					Use:            "show-account [owner]",
					Short:          "Query show-account",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}},
				},

				{
					RpcMethod:      "ListAccount",
					Use:            "list-account",
					Short:          "Query list-account",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "GetShareKeys",
					Use:            "get-share-keys [key] [totalshares] [minshares]",
					Short:          "Query get-share-keys",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "key"}, {ProtoField: "totalshares"}, {ProtoField: "minshares"}},
				},

				{
					RpcMethod:      "GetSecretKey",
					Use:            "get-secret-key [accountOwner]",
					Short:          "Query get-secret-key",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "accountOwner"}},
				},

				{
					RpcMethod:      "GetMappedAccounts",
					Use:            "get-mapped-accounts [account]",
					Short:          "Query get-mapped-accounts",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "account"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "AddAccountInfo",
					Use:            "add-account-info [owner] [totalshares] [minshares] [emailId]",
					Short:          "Send a addAccountInfo tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}, {ProtoField: "totalshares"}, {ProtoField: "minshares"}, {ProtoField: "emailId"}},
				},
				{
					RpcMethod:      "Updatetotalshares",
					Use:            "updatetotalshares [owner] [totalshares] [minshares] [id]",
					Short:          "Send a updatetotalshares tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}, {ProtoField: "totalshares"}, {ProtoField: "minshares"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "InitiateDeathValidation",
					Use:            "initiate-death-validation [owner] [death-date] [id]",
					Short:          "Send a initiate-death-validation tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}, {ProtoField: "deathDate"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "ValidateDeath",
					Use:            "validate-death [owner] [isvalidated] [id]",
					Short:          "Send a validate-death tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}, {ProtoField: "isvalidated"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateMultisite",
					Use:            "create-multisite [account] [id] [multisign]",
					Short:          "Send a create-multisite tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "account"}, {ProtoField: "id"}, {ProtoField: "multisign"}},
				},
				{
					RpcMethod:      "UpdateMultisite",
					Use:            "update-multisite [account] [id] [multisign-address] [multisign-share-key]",
					Short:          "Send a update-multisite tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "account"}, {ProtoField: "id"}, {ProtoField: "multisignAddress"}, {ProtoField: "multisignShareKey"}},
				},
				{
					RpcMethod:      "UpdateEmailId",
					Use:            "update-email-id [account] [emailid]",
					Short:          "Send a update-email-id tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "account"}, {ProtoField: "emailid"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

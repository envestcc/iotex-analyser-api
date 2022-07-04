// Code generated by proroc-gen-graphql, DO NOT EDIT.
package api

import (
	"context"

	"github.com/graphql-go/graphql"
	pagination "github.com/iotexproject/iotex-analyser-api/api/pagination"
	"github.com/pkg/errors"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	gql__type_RewardDistribution                          *graphql.Object      // message RewardDistribution in api_hermes.proto
	gql__type_HermesResponse                              *graphql.Object      // message HermesResponse in api_hermes.proto
	gql__type_HermesRequest                               *graphql.Object      // message HermesRequest in api_hermes.proto
	gql__type_HermesMetaResponse                          *graphql.Object      // message HermesMetaResponse in api_hermes.proto
	gql__type_HermesMetaRequest                           *graphql.Object      // message HermesMetaRequest in api_hermes.proto
	gql__type_HermesDistribution                          *graphql.Object      // message HermesDistribution in api_hermes.proto
	gql__type_HermesByVoterResponse_Delegate              *graphql.Object      // message HermesByVoterResponse.Delegate in api_hermes.proto
	gql__type_HermesByVoterResponse                       *graphql.Object      // message HermesByVoterResponse in api_hermes.proto
	gql__type_HermesByVoterRequest                        *graphql.Object      // message HermesByVoterRequest in api_hermes.proto
	gql__type_HermesByDelegateVoterInfo                   *graphql.Object      // message HermesByDelegateVoterInfo in api_hermes.proto
	gql__type_HermesByDelegateResponse                    *graphql.Object      // message HermesByDelegateResponse in api_hermes.proto
	gql__type_HermesByDelegateRequest                     *graphql.Object      // message HermesByDelegateRequest in api_hermes.proto
	gql__type_HermesByDelegateDistributionRatio           *graphql.Object      // message HermesByDelegateDistributionRatio in api_hermes.proto
	gql__type_HermesAverageStatsResponse_AveragePerEpoch  *graphql.Object      // message HermesAverageStatsResponse.AveragePerEpoch in api_hermes.proto
	gql__type_HermesAverageStatsResponse                  *graphql.Object      // message HermesAverageStatsResponse in api_hermes.proto
	gql__type_HermesAverageStatsRequest                   *graphql.Object      // message HermesAverageStatsRequest in api_hermes.proto
	gql__input_RewardDistribution                         *graphql.InputObject // message RewardDistribution in api_hermes.proto
	gql__input_HermesResponse                             *graphql.InputObject // message HermesResponse in api_hermes.proto
	gql__input_HermesRequest                              *graphql.InputObject // message HermesRequest in api_hermes.proto
	gql__input_HermesMetaResponse                         *graphql.InputObject // message HermesMetaResponse in api_hermes.proto
	gql__input_HermesMetaRequest                          *graphql.InputObject // message HermesMetaRequest in api_hermes.proto
	gql__input_HermesDistribution                         *graphql.InputObject // message HermesDistribution in api_hermes.proto
	gql__input_HermesByVoterResponse_Delegate             *graphql.InputObject // message HermesByVoterResponse.Delegate in api_hermes.proto
	gql__input_HermesByVoterResponse                      *graphql.InputObject // message HermesByVoterResponse in api_hermes.proto
	gql__input_HermesByVoterRequest                       *graphql.InputObject // message HermesByVoterRequest in api_hermes.proto
	gql__input_HermesByDelegateVoterInfo                  *graphql.InputObject // message HermesByDelegateVoterInfo in api_hermes.proto
	gql__input_HermesByDelegateResponse                   *graphql.InputObject // message HermesByDelegateResponse in api_hermes.proto
	gql__input_HermesByDelegateRequest                    *graphql.InputObject // message HermesByDelegateRequest in api_hermes.proto
	gql__input_HermesByDelegateDistributionRatio          *graphql.InputObject // message HermesByDelegateDistributionRatio in api_hermes.proto
	gql__input_HermesAverageStatsResponse_AveragePerEpoch *graphql.InputObject // message HermesAverageStatsResponse.AveragePerEpoch in api_hermes.proto
	gql__input_HermesAverageStatsResponse                 *graphql.InputObject // message HermesAverageStatsResponse in api_hermes.proto
	gql__input_HermesAverageStatsRequest                  *graphql.InputObject // message HermesAverageStatsRequest in api_hermes.proto
)

func Gql__type_RewardDistribution() *graphql.Object {
	if gql__type_RewardDistribution == nil {
		gql__type_RewardDistribution = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_RewardDistribution",
			Fields: graphql.Fields{
				"voterEthAddress": &graphql.Field{
					Type: graphql.String,
				},
				"voterIotexAddress": &graphql.Field{
					Type: graphql.String,
				},
				"amount": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_RewardDistribution
}

func Gql__type_HermesResponse() *graphql.Object {
	if gql__type_HermesResponse == nil {
		gql__type_HermesResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_HermesResponse",
			Fields: graphql.Fields{
				"hermesDistribution": &graphql.Field{
					Type: graphql.NewList(Gql__type_HermesDistribution()),
				},
			},
		})
	}
	return gql__type_HermesResponse
}

func Gql__type_HermesRequest() *graphql.Object {
	if gql__type_HermesRequest == nil {
		gql__type_HermesRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_HermesRequest",
			Fields: graphql.Fields{
				"startEpoch": &graphql.Field{
					Type: graphql.Int,
				},
				"epochCount": &graphql.Field{
					Type: graphql.Int,
				},
				"rewardAddress": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_HermesRequest
}

func Gql__type_HermesMetaResponse() *graphql.Object {
	if gql__type_HermesMetaResponse == nil {
		gql__type_HermesMetaResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_HermesMetaResponse",
			Fields: graphql.Fields{
				"exist": &graphql.Field{
					Type: graphql.Boolean,
				},
				"numberOfDelegates": &graphql.Field{
					Type: graphql.Int,
				},
				"numberOfRecipients": &graphql.Field{
					Type: graphql.Int,
				},
				"totalRewardDistributed": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_HermesMetaResponse
}

func Gql__type_HermesMetaRequest() *graphql.Object {
	if gql__type_HermesMetaRequest == nil {
		gql__type_HermesMetaRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_HermesMetaRequest",
			Fields: graphql.Fields{
				"startEpoch": &graphql.Field{
					Type: graphql.Int,
				},
				"epochCount": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_HermesMetaRequest
}

func Gql__type_HermesDistribution() *graphql.Object {
	if gql__type_HermesDistribution == nil {
		gql__type_HermesDistribution = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_HermesDistribution",
			Fields: graphql.Fields{
				"delegateName": &graphql.Field{
					Type: graphql.String,
				},
				"rewardDistribution": &graphql.Field{
					Type: graphql.NewList(Gql__type_RewardDistribution()),
				},
				"stakingIotexAddress": &graphql.Field{
					Type: graphql.String,
				},
				"voterCount": &graphql.Field{
					Type: graphql.Int,
				},
				"waiveServiceFee": &graphql.Field{
					Type: graphql.Boolean,
				},
				"refund": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_HermesDistribution
}

func Gql__type_HermesByVoterResponse_Delegate() *graphql.Object {
	if gql__type_HermesByVoterResponse_Delegate == nil {
		gql__type_HermesByVoterResponse_Delegate = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_HermesByVoterResponse_Delegate",
			Fields: graphql.Fields{
				"delegateName": &graphql.Field{
					Type: graphql.String,
				},
				"fromEpoch": &graphql.Field{
					Type: graphql.Int,
				},
				"toEpoch": &graphql.Field{
					Type: graphql.Int,
				},
				"amount": &graphql.Field{
					Type: graphql.String,
				},
				"actHash": &graphql.Field{
					Type: graphql.String,
				},
				"timestamp": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_HermesByVoterResponse_Delegate
}

func Gql__type_HermesByVoterResponse() *graphql.Object {
	if gql__type_HermesByVoterResponse == nil {
		gql__type_HermesByVoterResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_HermesByVoterResponse",
			Fields: graphql.Fields{
				"exist": &graphql.Field{
					Type: graphql.Boolean,
				},
				"delegates": &graphql.Field{
					Type: graphql.NewList(Gql__type_HermesByVoterResponse_Delegate()),
				},
				"count": &graphql.Field{
					Type: graphql.Int,
				},
				"totalRewardReceived": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_HermesByVoterResponse
}

func Gql__type_HermesByVoterRequest() *graphql.Object {
	if gql__type_HermesByVoterRequest == nil {
		gql__type_HermesByVoterRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_HermesByVoterRequest",
			Fields: graphql.Fields{
				"startEpoch": &graphql.Field{
					Type: graphql.Int,
				},
				"epochCount": &graphql.Field{
					Type: graphql.Int,
				},
				"voterAddress": &graphql.Field{
					Type: graphql.String,
				},
				"pagination": &graphql.Field{
					Type: pagination.Gql__type_Pagination(),
				},
			},
		})
	}
	return gql__type_HermesByVoterRequest
}

func Gql__type_HermesByDelegateVoterInfo() *graphql.Object {
	if gql__type_HermesByDelegateVoterInfo == nil {
		gql__type_HermesByDelegateVoterInfo = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_HermesByDelegateVoterInfo",
			Fields: graphql.Fields{
				"voterAddress": &graphql.Field{
					Type: graphql.String,
				},
				"fromEpoch": &graphql.Field{
					Type: graphql.Int,
				},
				"toEpoch": &graphql.Field{
					Type: graphql.Int,
				},
				"amount": &graphql.Field{
					Type: graphql.String,
				},
				"actHash": &graphql.Field{
					Type: graphql.String,
				},
				"timestamp": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_HermesByDelegateVoterInfo
}

func Gql__type_HermesByDelegateResponse() *graphql.Object {
	if gql__type_HermesByDelegateResponse == nil {
		gql__type_HermesByDelegateResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_HermesByDelegateResponse",
			Fields: graphql.Fields{
				"exist": &graphql.Field{
					Type: graphql.Boolean,
				},
				"count": &graphql.Field{
					Type: graphql.Int,
				},
				"voterInfoList": &graphql.Field{
					Type: graphql.NewList(Gql__type_HermesByDelegateVoterInfo()),
				},
				"totalRewardsDistributed": &graphql.Field{
					Type: graphql.String,
				},
				"distributionRatio": &graphql.Field{
					Type: graphql.NewList(Gql__type_HermesByDelegateDistributionRatio()),
				},
			},
		})
	}
	return gql__type_HermesByDelegateResponse
}

func Gql__type_HermesByDelegateRequest() *graphql.Object {
	if gql__type_HermesByDelegateRequest == nil {
		gql__type_HermesByDelegateRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_HermesByDelegateRequest",
			Fields: graphql.Fields{
				"startEpoch": &graphql.Field{
					Type: graphql.Int,
				},
				"epochCount": &graphql.Field{
					Type: graphql.Int,
				},
				"delegateName": &graphql.Field{
					Type: graphql.String,
				},
				"pagination": &graphql.Field{
					Type: pagination.Gql__type_Pagination(),
				},
			},
		})
	}
	return gql__type_HermesByDelegateRequest
}

func Gql__type_HermesByDelegateDistributionRatio() *graphql.Object {
	if gql__type_HermesByDelegateDistributionRatio == nil {
		gql__type_HermesByDelegateDistributionRatio = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_HermesByDelegateDistributionRatio",
			Fields: graphql.Fields{
				"epochNumber": &graphql.Field{
					Type: graphql.Int,
				},
				"blockRewardRatio": &graphql.Field{
					Type: graphql.Float,
				},
				"epochRewardRatio": &graphql.Field{
					Type: graphql.Float,
				},
				"foundationBonusRatio": &graphql.Field{
					Type: graphql.Float,
				},
			},
		})
	}
	return gql__type_HermesByDelegateDistributionRatio
}

func Gql__type_HermesAverageStatsResponse_AveragePerEpoch() *graphql.Object {
	if gql__type_HermesAverageStatsResponse_AveragePerEpoch == nil {
		gql__type_HermesAverageStatsResponse_AveragePerEpoch = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_HermesAverageStatsResponse_AveragePerEpoch",
			Fields: graphql.Fields{
				"delegateName": &graphql.Field{
					Type: graphql.String,
				},
				"rewardDistribution": &graphql.Field{
					Type: graphql.String,
				},
				"totalWeightedVotes": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_HermesAverageStatsResponse_AveragePerEpoch
}

func Gql__type_HermesAverageStatsResponse() *graphql.Object {
	if gql__type_HermesAverageStatsResponse == nil {
		gql__type_HermesAverageStatsResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_HermesAverageStatsResponse",
			Fields: graphql.Fields{
				"exist": &graphql.Field{
					Type: graphql.Boolean,
				},
				"averagePerEpoch": &graphql.Field{
					Type: graphql.NewList(Gql__type_HermesAverageStatsResponse_AveragePerEpoch()),
				},
			},
		})
	}
	return gql__type_HermesAverageStatsResponse
}

func Gql__type_HermesAverageStatsRequest() *graphql.Object {
	if gql__type_HermesAverageStatsRequest == nil {
		gql__type_HermesAverageStatsRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_HermesAverageStatsRequest",
			Fields: graphql.Fields{
				"startEpoch": &graphql.Field{
					Type: graphql.Int,
				},
				"epochCount": &graphql.Field{
					Type: graphql.Int,
				},
				"rewardAddress": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_HermesAverageStatsRequest
}

func Gql__input_RewardDistribution() *graphql.InputObject {
	if gql__input_RewardDistribution == nil {
		gql__input_RewardDistribution = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_RewardDistribution",
			Fields: graphql.InputObjectConfigFieldMap{
				"voterEthAddress": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"voterIotexAddress": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"amount": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_RewardDistribution
}

func Gql__input_HermesResponse() *graphql.InputObject {
	if gql__input_HermesResponse == nil {
		gql__input_HermesResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_HermesResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"hermesDistribution": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_HermesDistribution()),
				},
			},
		})
	}
	return gql__input_HermesResponse
}

func Gql__input_HermesRequest() *graphql.InputObject {
	if gql__input_HermesRequest == nil {
		gql__input_HermesRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_HermesRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"startEpoch": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"epochCount": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"rewardAddress": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_HermesRequest
}

func Gql__input_HermesMetaResponse() *graphql.InputObject {
	if gql__input_HermesMetaResponse == nil {
		gql__input_HermesMetaResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_HermesMetaResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"exist": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"numberOfDelegates": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"numberOfRecipients": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"totalRewardDistributed": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_HermesMetaResponse
}

func Gql__input_HermesMetaRequest() *graphql.InputObject {
	if gql__input_HermesMetaRequest == nil {
		gql__input_HermesMetaRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_HermesMetaRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"startEpoch": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"epochCount": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_HermesMetaRequest
}

func Gql__input_HermesDistribution() *graphql.InputObject {
	if gql__input_HermesDistribution == nil {
		gql__input_HermesDistribution = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_HermesDistribution",
			Fields: graphql.InputObjectConfigFieldMap{
				"delegateName": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"rewardDistribution": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_RewardDistribution()),
				},
				"stakingIotexAddress": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"voterCount": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"waiveServiceFee": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"refund": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_HermesDistribution
}

func Gql__input_HermesByVoterResponse_Delegate() *graphql.InputObject {
	if gql__input_HermesByVoterResponse_Delegate == nil {
		gql__input_HermesByVoterResponse_Delegate = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_HermesByVoterResponse_Delegate",
			Fields: graphql.InputObjectConfigFieldMap{
				"delegateName": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"fromEpoch": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"toEpoch": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"amount": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"actHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"timestamp": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_HermesByVoterResponse_Delegate
}

func Gql__input_HermesByVoterResponse() *graphql.InputObject {
	if gql__input_HermesByVoterResponse == nil {
		gql__input_HermesByVoterResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_HermesByVoterResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"exist": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"delegates": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_HermesByVoterResponse_Delegate()),
				},
				"count": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"totalRewardReceived": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_HermesByVoterResponse
}

func Gql__input_HermesByVoterRequest() *graphql.InputObject {
	if gql__input_HermesByVoterRequest == nil {
		gql__input_HermesByVoterRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_HermesByVoterRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"startEpoch": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"epochCount": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"voterAddress": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.InputObjectFieldConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
		})
	}
	return gql__input_HermesByVoterRequest
}

func Gql__input_HermesByDelegateVoterInfo() *graphql.InputObject {
	if gql__input_HermesByDelegateVoterInfo == nil {
		gql__input_HermesByDelegateVoterInfo = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_HermesByDelegateVoterInfo",
			Fields: graphql.InputObjectConfigFieldMap{
				"voterAddress": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"fromEpoch": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"toEpoch": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"amount": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"actHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"timestamp": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_HermesByDelegateVoterInfo
}

func Gql__input_HermesByDelegateResponse() *graphql.InputObject {
	if gql__input_HermesByDelegateResponse == nil {
		gql__input_HermesByDelegateResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_HermesByDelegateResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"exist": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"count": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"voterInfoList": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_HermesByDelegateVoterInfo()),
				},
				"totalRewardsDistributed": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"distributionRatio": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_HermesByDelegateDistributionRatio()),
				},
			},
		})
	}
	return gql__input_HermesByDelegateResponse
}

func Gql__input_HermesByDelegateRequest() *graphql.InputObject {
	if gql__input_HermesByDelegateRequest == nil {
		gql__input_HermesByDelegateRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_HermesByDelegateRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"startEpoch": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"epochCount": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"delegateName": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.InputObjectFieldConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
		})
	}
	return gql__input_HermesByDelegateRequest
}

func Gql__input_HermesByDelegateDistributionRatio() *graphql.InputObject {
	if gql__input_HermesByDelegateDistributionRatio == nil {
		gql__input_HermesByDelegateDistributionRatio = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_HermesByDelegateDistributionRatio",
			Fields: graphql.InputObjectConfigFieldMap{
				"epochNumber": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"blockRewardRatio": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				"epochRewardRatio": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				"foundationBonusRatio": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
			},
		})
	}
	return gql__input_HermesByDelegateDistributionRatio
}

func Gql__input_HermesAverageStatsResponse_AveragePerEpoch() *graphql.InputObject {
	if gql__input_HermesAverageStatsResponse_AveragePerEpoch == nil {
		gql__input_HermesAverageStatsResponse_AveragePerEpoch = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_HermesAverageStatsResponse_AveragePerEpoch",
			Fields: graphql.InputObjectConfigFieldMap{
				"delegateName": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"rewardDistribution": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"totalWeightedVotes": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_HermesAverageStatsResponse_AveragePerEpoch
}

func Gql__input_HermesAverageStatsResponse() *graphql.InputObject {
	if gql__input_HermesAverageStatsResponse == nil {
		gql__input_HermesAverageStatsResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_HermesAverageStatsResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"exist": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"averagePerEpoch": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_HermesAverageStatsResponse_AveragePerEpoch()),
				},
			},
		})
	}
	return gql__input_HermesAverageStatsResponse
}

func Gql__input_HermesAverageStatsRequest() *graphql.InputObject {
	if gql__input_HermesAverageStatsRequest == nil {
		gql__input_HermesAverageStatsRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_HermesAverageStatsRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"startEpoch": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"epochCount": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"rewardAddress": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_HermesAverageStatsRequest
}

// graphql__resolver_HermesService is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_HermesService struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_HermesService creates pointer of service struct
func new_graphql_resolver_HermesService(conn *grpc.ClientConn) *graphql__resolver_HermesService {
	return &graphql__resolver_HermesService{
		conn:        conn,
		host:        "localhost:50051",
		dialOptions: []grpc.DialOption{},
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_HermesService) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
	// If x.conn is not nil, user injected their own connection
	if x.conn != nil {
		return x.conn, func() {}, nil
	}

	// Otherwise, this handler opens connection with specified host
	conn, err := grpc.DialContext(ctx, x.host, x.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}

// GetQueries returns acceptable graphql.Fields for Query.
func (x *graphql__resolver_HermesService) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"Hermes": &graphql.Field{
			Type: Gql__type_HermesResponse(),
			Args: graphql.FieldConfigArgument{
				"startEpoch": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"epochCount": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"rewardAddress": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req HermesRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for Hermes")
				}
				client := NewHermesServiceClient(conn)
				resp, err := client.Hermes(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC Hermes")
				}
				return resp, nil
			},
		},
		"HermesByVoter": &graphql.Field{
			Type: Gql__type_HermesByVoterResponse(),
			Args: graphql.FieldConfigArgument{
				"startEpoch": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"epochCount": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"voterAddress": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.ArgumentConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req HermesByVoterRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for HermesByVoter")
				}
				client := NewHermesServiceClient(conn)
				resp, err := client.HermesByVoter(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC HermesByVoter")
				}
				return resp, nil
			},
		},
		"HermesByDelegate": &graphql.Field{
			Type: Gql__type_HermesByDelegateResponse(),
			Args: graphql.FieldConfigArgument{
				"startEpoch": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"epochCount": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"delegateName": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.ArgumentConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req HermesByDelegateRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for HermesByDelegate")
				}
				client := NewHermesServiceClient(conn)
				resp, err := client.HermesByDelegate(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC HermesByDelegate")
				}
				return resp, nil
			},
		},
		"HermesMeta": &graphql.Field{
			Type: Gql__type_HermesMetaResponse(),
			Args: graphql.FieldConfigArgument{
				"startEpoch": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"epochCount": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req HermesMetaRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for HermesMeta")
				}
				client := NewHermesServiceClient(conn)
				resp, err := client.HermesMeta(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC HermesMeta")
				}
				return resp, nil
			},
		},
		"HermesAverageStats": &graphql.Field{
			Type: Gql__type_HermesAverageStatsResponse(),
			Args: graphql.FieldConfigArgument{
				"startEpoch": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"epochCount": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"rewardAddress": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req HermesAverageStatsRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for HermesAverageStats")
				}
				client := NewHermesServiceClient(conn)
				resp, err := client.HermesAverageStats(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC HermesAverageStats")
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_HermesService) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterHermesServiceGraphqlHandler with *grpc.ClientConn manually.
func RegisterHermesServiceGraphql(mux *runtime.ServeMux) error {
	return RegisterHermesServiceGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, you can specify automatic opening connection with ServiceOption directive:
//
// service HermesService {
//    option (graphql.service) = {
//        host: "host:port"
//        insecure: true or false
//    };
//
//    ...with RPC definitions
// }
func RegisterHermesServiceGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return mux.AddHandler(new_graphql_resolver_HermesService(conn))
}

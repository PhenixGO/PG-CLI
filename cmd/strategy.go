package cmd

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	pb "github.com/pg-cli/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var pg pb.PhenixGOClient

var (
	strategyName string
	strategyID   string
	parameter    string
	group        string
	limit        uint64
	offset       uint64
	startTime    string
	endTime      string
	apikey1      string
	apikey2      string
)

func init() {
	rootCmd.AddCommand(startStrategyCmd)
	startStrategyCmd.Flags().StringVar(&strategyName, "strategy-name", "", "Strategy Name")
	startStrategyCmd.Flags().StringVar(&parameter, "parameter", "", "Strategy Parameter")
	startStrategyCmd.Flags().StringVar(&group, "group", "", "Strategy Group")
	startStrategyCmd.Flags().StringVar(&apikey1, "apikey1", "", "The first set of exchange api key")
	startStrategyCmd.Flags().StringVar(&apikey2, "apikey2", "", "The second set of exchange api key")

	rootCmd.AddCommand(stopStrategyCmd)
	stopStrategyCmd.Flags().StringVar(&strategyID, "strategy-id", "", "Strategy ID")

	rootCmd.AddCommand(listStrategyCmd)
	listStrategyCmd.Flags().StringVar(&group, "group", "", "Strategy Group")
	listStrategyCmd.Flags().Uint64Var(&limit, "limit", 100, "Pagination Limit")
	listStrategyCmd.Flags().Uint64Var(&offset, "offset", 0, "Pagination Offset")

	rootCmd.AddCommand(predictedPerfCmd)
	predictedPerfCmd.Flags().StringVar(&strategyName, "strategy-name", "", "Strategy Name")
	predictedPerfCmd.Flags().StringVar(&parameter, "parameter", "", "Strategy Parameter")

	rootCmd.AddCommand(perfCmd)
	perfCmd.Flags().StringVar(&strategyID, "strategy-id", "", "Strategy ID")
	perfCmd.Flags().StringVar(&startTime, "start-time", "", "Query range start time. format in RFC3339(2006-01-02T15:04:05Z07:00)")
	perfCmd.Flags().StringVar(&endTime, "end-time", "", "Query range end time. format in RFC3339(2006-01-02T15:04:05Z07:00)")
}

func connect() {
	config := &tls.Config{
		InsecureSkipVerify: false,
	}
	conn, err := grpc.Dial(
		viper.GetString("endpoint"),
		grpc.WithTransportCredentials(credentials.NewTLS(config)),
	)
	if err != nil {
		cobra.CheckErr(err)
	}
	pg = pb.NewPhenixGOClient(conn)
}

var startStrategyCmd = &cobra.Command{
	Use:   "start",
	Short: "Start strategy with parameters",
	Run: func(cmd *cobra.Command, args []string) {

		connect()
		md := metadata.Pairs("X-API-KEY", viper.GetString("token"))
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		if strategyName == "" {
			cobra.CheckErr(fmt.Errorf("strategy-name flag should not be empty"))
		}
		if parameter == "" {
			cobra.CheckErr(fmt.Errorf("parameter flag should not be empty"))
		}
		if apikey1 == "" {
			cobra.CheckErr(fmt.Errorf("apikey1 flag should not be empty"))
		}

		spfparam := pb.SPFParameter{}
		err := json.Unmarshal([]byte(parameter), &spfparam)
		if err != nil {
			cobra.CheckErr(err)
		}

		param, err := anypb.New(&spfparam)
		if err != nil {
			cobra.CheckErr(err)
		}

		var pgroup *wrapperspb.StringValue
		if group != "" {
			pgroup = wrapperspb.String(group)
		}

		papikey1 := pb.ExchangeAPIKey{}
		err = json.Unmarshal([]byte(apikey1), &papikey1)
		if err != nil {
			cobra.CheckErr(err)
		}

		var papikey2 *pb.ExchangeAPIKey
		if apikey2 != "" {
			t := pb.ExchangeAPIKey{}
			err = json.Unmarshal([]byte(apikey2), &t)
			if err != nil {
				cobra.CheckErr(err)
			}

			papikey2 = &t
		}

		_, err = pg.StartSubUserStrategy(ctx, &pb.StartSubUserStrategyRequest{
			StrategyName: strategyName,
			Parameter:    param,
			Group:        pgroup,
			Apikey1:      &papikey1,
			Apikey2:      papikey2,
		})

		if err != nil {
			cobra.CheckErr(err)
		}
	},
}

var stopStrategyCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop strategy with strategy id",
	Run: func(cmd *cobra.Command, args []string) {

		connect()
		md := metadata.Pairs("X-API-KEY", viper.GetString("token"))
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		if strategyID == "" {
			cobra.CheckErr(fmt.Errorf("strategy-id flag should not be empty"))
		}

		_, err := pg.StopSubUserStrategy(ctx, &pb.StopSubUserStrategyRequest{
			StrategyId: strategyID,
		})

		if err != nil {
			cobra.CheckErr(err)
		}
	},
}

var listStrategyCmd = &cobra.Command{
	Use:   "list",
	Short: "Get list of strategy with parameters",
	Run: func(cmd *cobra.Command, args []string) {

		connect()
		md := metadata.Pairs("X-API-KEY", viper.GetString("token"))
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		var pStrategyID *wrapperspb.StringValue
		if strategyID != "" {
			pStrategyID = wrapperspb.String(strategyID)
		}
		var pStrategyName *wrapperspb.StringValue
		if strategyName != "" {
			pStrategyName = wrapperspb.String(strategyName)
		}
		var pGroup *wrapperspb.StringValue
		if group != "" {
			pGroup = wrapperspb.String(group)
		}

		_, err := pg.GetSubUserStrategy(ctx, &pb.GetSubUserStrategyRequest{
			Filter: &pb.StrategyFilter{
				StrategyId:   pStrategyID,
				StrategyName: pStrategyName,
				Group:        pGroup,
			},
			Page: &pb.Pagination{
				Limit:  limit,
				Offset: offset,
			},
		})

		if err != nil {
			cobra.CheckErr(err)
		}
	},
}

var predictedPerfCmd = &cobra.Command{
	Use:   "predict",
	Short: "Get strategy predicted performance by query input",
	Run: func(cmd *cobra.Command, args []string) {

		connect()
		md := metadata.Pairs("X-API-KEY", viper.GetString("token"))
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		if strategyName == "" {
			cobra.CheckErr(fmt.Errorf("strategy-name flag should not be empty"))
		}
		if parameter == "" {
			cobra.CheckErr(fmt.Errorf("parameter flag should not be empty"))
		}

		spfparam := pb.SPFParameter{}
		err := json.Unmarshal([]byte(parameter), &spfparam)
		if err != nil {
			cobra.CheckErr(err)
		}

		param, err := anypb.New(&spfparam)
		if err != nil {
			cobra.CheckErr(err)
		}
		_, err = pg.GetPredictedPerformance(ctx,
			&pb.GetPredictedPerformanceRequest{
				StrategyName: strategyName,
				Parameter:    param,
			})

		if err != nil {
			cobra.CheckErr(err)
		}
	},
}

var perfCmd = &cobra.Command{
	Use:   "performance",
	Short: "Get strategy performance",
	Run: func(cmd *cobra.Command, args []string) {

		connect()
		md := metadata.Pairs("X-API-KEY", viper.GetString("token"))
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		if strategyID == "" {
			cobra.CheckErr(fmt.Errorf("strategy-name flag should not be empty"))
		}

		var pStartTime *timestamppb.Timestamp
		if startTime != "" {
			t, err := time.Parse(time.RFC3339, startTime)
			if err != nil {
				cobra.CheckErr(err)
			}
			pStartTime = timestamppb.New(t)
		}

		var pEndTime *timestamppb.Timestamp
		if endTime != "" {
			t, err := time.Parse(time.RFC3339, startTime)
			if err != nil {
				cobra.CheckErr(err)
			}
			pEndTime = timestamppb.New(t)
		}

		_, err := pg.GetPerformance(ctx, &pb.GetPerformanceRequest{
			StrategyId: strategyID,
			StartTime:  pStartTime,
			EndTime:    pEndTime,
		})

		if err != nil {
			cobra.CheckErr(err)
		}
	},
}

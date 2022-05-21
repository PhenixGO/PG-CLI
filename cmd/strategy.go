package cmd

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"os"
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
	tag          string
	status       string
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
	startStrategyCmd.Flags().StringVar(&tag, "tag", "", "Strategy Tag")
	startStrategyCmd.Flags().StringVar(&apikey1, "apikey1", "", "The first set of exchange api key")
	startStrategyCmd.Flags().StringVar(&apikey2, "apikey2", "", "The second set of exchange api key")

	rootCmd.AddCommand(stopStrategyCmd)
	stopStrategyCmd.Flags().StringVar(&strategyID, "strategy-id", "", "Strategy ID")

	rootCmd.AddCommand(listStrategyCmd)
	listStrategyCmd.Flags().StringVar(&tag, "tag", "", "Strategy Tag")
	listStrategyCmd.Flags().StringVar(&status, "status", "", "Strategy Status")
	listStrategyCmd.Flags().Uint64Var(&limit, "limit", 100, "Pagination Limit")
	listStrategyCmd.Flags().Uint64Var(&offset, "offset", 0, "Pagination Offset")

	rootCmd.AddCommand(profitCmd)
	profitCmd.Flags().StringVar(&strategyID, "strategy-id", "", "Strategy ID")
	profitCmd.Flags().StringVar(&startTime, "start-time", "", "Query range start time. format in RFC3339(2006-01-02T15:04:05Z07:00)")
	profitCmd.Flags().StringVar(&endTime, "end-time", "", "Query range end time. format in RFC3339(2006-01-02T15:04:05Z07:00)")
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

		baParam := pb.BitginArbitrageParameter{}
		err := json.Unmarshal([]byte(parameter), &baParam)
		if err != nil {
			cobra.CheckErr(err)
		}

		param, err := anypb.New(&baParam)
		if err != nil {
			cobra.CheckErr(err)
		}

		var pTag *wrapperspb.StringValue
		if tag != "" {
			pTag = wrapperspb.String(tag)
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

		s, err := pg.StartUserStrategy(ctx, &pb.StartUserStrategyRequest{
			StrategyName: strategyName,
			Parameter:    param,
			Tag:          pTag,
			Apikey1:      &papikey1,
			Apikey2:      papikey2,
		})

		if err != nil {
			cobra.CheckErr(err)
		}

		msg, _ := json.MarshalIndent(s, "", "  ")
		fmt.Fprintln(os.Stderr, "Result:", string(msg))

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

		s, err := pg.StopUserStrategy(ctx, &pb.StopUserStrategyRequest{
			StrategyId: strategyID,
		})

		if err != nil {
			cobra.CheckErr(err)
		}

		msg, _ := json.MarshalIndent(s, "", "  ")
		fmt.Fprintln(os.Stderr, "Result:", string(msg))
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
		var pTag *wrapperspb.StringValue
		if tag != "" {
			pTag = wrapperspb.String(tag)
		}
		var pStatus []*wrapperspb.StringValue
		if status != "" {
			slice := []string{}
			err := json.Unmarshal([]byte(status), &slice)
			if err != nil {
				cobra.CheckErr(err)
			}
			for _, s := range slice {
				pStatus = append(pStatus, wrapperspb.String(s))
			}
		}
		s, err := pg.GetUserStrategy(ctx, &pb.GetStrategyRequest{
			Filter: &pb.StrategyFilter{
				StrategyId:   pStrategyID,
				StrategyName: pStrategyName,
				Tag:          pTag,
				Status:       pStatus,
			},
			Page: &pb.Pagination{
				Limit:  limit,
				Offset: offset,
			},
		})

		if err != nil {
			cobra.CheckErr(err)
		}

		msg, _ := json.MarshalIndent(s, "", "  ")
		fmt.Fprintln(os.Stderr, "Result:", string(msg))
	},
}

var profitCmd = &cobra.Command{
	Use:   "profit",
	Short: "Get strategy profit",
	Run: func(cmd *cobra.Command, args []string) {

		connect()
		md := metadata.Pairs("X-API-KEY", viper.GetString("token"))
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		if strategyID == "" {
			cobra.CheckErr(fmt.Errorf("strategy-id flag should not be empty"))
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

		s, err := pg.GetStrategyProfits(ctx, &pb.GetStrategyProfitsRequest{
			StrategyId: strategyID,
			StartTime:  pStartTime,
			EndTime:    pEndTime,
		})

		if err != nil {
			cobra.CheckErr(err)
		}

		msg, _ := json.MarshalIndent(s, "", "  ")
		fmt.Fprintln(os.Stderr, "Result:", string(msg))
	},
}

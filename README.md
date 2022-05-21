# pg-cli

PhenixGO is a CLI tool for run PhenixGO Strategy functions.
This Application demonstrate functionalities for cryptocurrency quantitative
tradingstrategy platfrom Phenixgo. If any question please contact with
service@phenixgo.net

###Usage:
  phenixgo [flags]
  phenixgo [command]

### Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        Get list of strategy with parameters
  profit      Get strategy profit
  start       Start strategy with parameters
  stop        Stop strategy with strategy id

### Flags:
  --config string   config file (default is $PWD/.config.yaml)
  -h, --help            help for phenixgo

### Example:
  * go run ./main.go start --strategy-name ba --tag {{user_id}} --apikey1 "{\"public_key\":\"{{max_api_key}}\",\"private_key\":\"{{max_secret_key}}\",\"exchange\":\"max\"}" --apikey2 "{\"public_key\":\"{{bitgin_api_key}}\",\"private_key\":\"{{bitgin_secret_key}}\",\"exchange\":\"bitgin\"}" --parameter "{\"usdt_amount\":1000,\"min_rate\":0.004}"
  * go run ./main.go stop --strategy-id {{strategy_id}}
  * go run ./main.go list --status "[\"Running\"]"
  * go run ./main.go profit --strategy-id {{strategy_id}} --start-time 2006-01-02T15:04:05Z

Use "phenixgo [command] --help" for more information about a command.

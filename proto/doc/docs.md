# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [generic.proto](#generic-proto)
    - [Pagination](#generic-Pagination)
  
- [model.proto](#model-proto)
    - [BAState](#model-BAState)
    - [BitginArbitrageParameter](#model-BitginArbitrageParameter)
    - [ExchangeAPIKey](#model-ExchangeAPIKey)
    - [LendingOffer](#model-LendingOffer)
    - [LendingParameter](#model-LendingParameter)
    - [LendingState](#model-LendingState)
    - [Notification](#model-Notification)
    - [Profit](#model-Profit)
    - [SPFParameter](#model-SPFParameter)
    - [Strategy](#model-Strategy)
  
    - [Event](#model-Event)
  
- [phenixgo.proto](#phenixgo-proto)
    - [GetStrategyProfitsRequest](#phenixgo-GetStrategyProfitsRequest)
    - [GetStrategyProfitsResponse](#phenixgo-GetStrategyProfitsResponse)
    - [GetStrategyRequest](#phenixgo-GetStrategyRequest)
    - [GetStrategyResponse](#phenixgo-GetStrategyResponse)
    - [GetStrategyStateRequest](#phenixgo-GetStrategyStateRequest)
    - [GetStrategyStateResponse](#phenixgo-GetStrategyStateResponse)
    - [StartUserStrategyRequest](#phenixgo-StartUserStrategyRequest)
    - [StartUserStrategyResponse](#phenixgo-StartUserStrategyResponse)
    - [StopUserStrategyRequest](#phenixgo-StopUserStrategyRequest)
    - [StopUserStrategyResponse](#phenixgo-StopUserStrategyResponse)
    - [StrategyFilter](#phenixgo-StrategyFilter)
    - [UpdateUserStrategyRequest](#phenixgo-UpdateUserStrategyRequest)
    - [UpdateUserStrategyResponse](#phenixgo-UpdateUserStrategyResponse)
  
    - [PhenixGO](#phenixgo-PhenixGO)
  
- [Scalar Value Types](#scalar-value-types)



<a name="generic-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## generic.proto



<a name="generic-Pagination"></a>

### Pagination
Pagination represents a search query, with offset and limit options to  indicate which results to include in the response.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| offset | [uint64](#uint64) |  | indicate response from the index. |
| limit | [uint64](#uint64) |  | indicate response number of data to return. |





 

 

 

 



<a name="model-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model.proto



<a name="model-BAState"></a>

### BAState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [string](#string) |  |  |
| sold | [string](#string) |  |  |
| bought | [string](#string) |  |  |
| updated_at | [string](#string) |  |  |






<a name="model-BitginArbitrageParameter"></a>

### BitginArbitrageParameter
BitginArbitrageParameter represents parameters for arbitrage between bitgin and the other exchanges.

Note: for invoking StartUserStrategy for Bitgin Arbitrage need to fill the sell side and the buy side API key in the StartUserStrategyRequest (apikey1: Max, apikey2: Bitgin).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| usdt_amount | [double](#double) |  | used to specified USDT amount to run (must be prepare in max and large than 500). |
| min_rate | [double](#double) |  | used to specified the min profit for place order price (must large than 0.15). |






<a name="model-ExchangeAPIKey"></a>

### ExchangeAPIKey
ExchangeAPIKey represents exchange key pairs and subaccount to be use in the strategy.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| public_key | [string](#string) |  | exchange key pair, the one can be public. |
| private_key | [string](#string) |  | exchange key pair, the one is private. |
| exchange | [string](#string) |  | represent this key pair is belonging to the exchange. |
| sub_account | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | specify sub-account if needed. |






<a name="model-LendingOffer"></a>

### LendingOffer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| rate | [string](#string) |  |  |
| period | [int32](#int32) |  |  |
| position_pair | [string](#string) |  |  |
| created_at | [string](#string) |  |  |






<a name="model-LendingParameter"></a>

### LendingParameter
LendingParameter represents parameters for bitfinex lending strategy.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currency | [string](#string) |  | specified lending currency, fill with capital currency symbol (ex. USD, USDT). |
| max_amount | [double](#double) |  | represent max lending amount for running strategy. The sum of offer would be constrained by this value. |
| keep_amount | [double](#double) |  | specified funding balance should left this amount. |






<a name="model-LendingState"></a>

### LendingState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| active_offers | [LendingOffer](#model-LendingOffer) | repeated |  |
| open_offers | [LendingOffer](#model-LendingOffer) | repeated |  |
| available_balance | [string](#string) |  |  |
| updated_at | [string](#string) |  |  |






<a name="model-Notification"></a>

### Notification



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#model-Event) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="model-Profit"></a>

### Profit
Profit represents the data form for the profit.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| profit_id | [string](#string) |  | unique profit id for a profit data. |
| currency | [string](#string) |  | profit currency |
| type | [string](#string) |  | description for the profit |
| amount | [double](#double) |  | amount for the profit currency |
| rate | [double](#double) |  | profit rate means the percentage of investing amount |
| time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | profit accounting time |






<a name="model-SPFParameter"></a>

### SPFParameter
Not ready SPFParameter represent spot perpetual funding strategy.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| initial_balance | [double](#double) |  |  |
| currency | [string](#string) |  |  |
| spread | [double](#double) |  |  |
| size | [double](#double) |  |  |






<a name="model-Strategy"></a>

### Strategy
Strategy represents the data form for the strategy.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| strategy_id | [string](#string) |  | unique strategy id for a strategy data. |
| strategy_name | [string](#string) |  | name |
| parameter | [google.protobuf.Any](#google-protobuf-Any) |  | parameters for specified strategy (BitginArbitrageParameter, LendingParameter...). |
| status | [string](#string) |  | strategy status (running, stopping, stopped, completed, error). |
| error_message | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | strategy error message |
| user_id | [string](#string) |  | user id |
| tag | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | tag for strategy instance |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | strategy create time |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | strategy update time |
| closed_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | strategy end time |





 


<a name="model-Event"></a>

### Event


| Name | Number | Description |
| ---- | ------ | ----------- |
| EventUnknown | 0 |  |
| EventStatus | 1 |  |
| EventProfit | 2 |  |


 

 

 



<a name="phenixgo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## phenixgo.proto



<a name="phenixgo-GetStrategyProfitsRequest"></a>

### GetStrategyProfitsRequest
GetStrategyProfitsRequest represent a query, that specify time range to indicate profit data to return.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| strategy_id | [string](#string) |  | strategy id means profit from. |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | query range start. |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | query range end. |






<a name="phenixgo-GetStrategyProfitsResponse"></a>

### GetStrategyProfitsResponse
GetStrategyProfitsResponse returns profit data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| profits | [model.Profit](#model-Profit) | repeated | profit data. |






<a name="phenixgo-GetStrategyRequest"></a>

### GetStrategyRequest
GetStrategyRequest represents a search query, with options StrategyFilter and Pagination to indicate which results to include in the response.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filter | [StrategyFilter](#phenixgo-StrategyFilter) |  | condition of the request |
| page | [generic.Pagination](#generic-Pagination) |  | range of the response |






<a name="phenixgo-GetStrategyResponse"></a>

### GetStrategyResponse
GetStrategyResponse returns strategy array for the matches query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| strategies | [model.Strategy](#model-Strategy) | repeated | strategy data. |






<a name="phenixgo-GetStrategyStateRequest"></a>

### GetStrategyStateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| strategy_id | [string](#string) |  |  |






<a name="phenixgo-GetStrategyStateResponse"></a>

### GetStrategyStateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [google.protobuf.Any](#google-protobuf-Any) |  |  |






<a name="phenixgo-StartUserStrategyRequest"></a>

### StartUserStrategyRequest
StartUserStrategyRequest represents request for starting a strategy.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| strategy_name | [string](#string) |  | strategy name. |
| parameter | [google.protobuf.Any](#google-protobuf-Any) |  | strategy parameters. (ex. LendingParameter) |
| tag | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | strategy tag for classfying. |
| apikey1 | [model.ExchangeAPIKey](#model-ExchangeAPIKey) |  | first exchange api key. (for all strategy) |
| apikey2 | [model.ExchangeAPIKey](#model-ExchangeAPIKey) |  | second exchange api key. (for cross exchange strategy. ex: Bitgin Arbitrage) |






<a name="phenixgo-StartUserStrategyResponse"></a>

### StartUserStrategyResponse
StartUserStrategyResponse returns created strategy information.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| strategy | [model.Strategy](#model-Strategy) |  | strategy data. |






<a name="phenixgo-StopUserStrategyRequest"></a>

### StopUserStrategyRequest
StopUserStrategyRequest represents request for stopping a strategy.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| strategy_id | [string](#string) |  | specified target strategy id. |
| is_force | [bool](#bool) |  |  |






<a name="phenixgo-StopUserStrategyResponse"></a>

### StopUserStrategyResponse
StopUserStrategyResponse returns strategy id if stop successfully.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| strategy_id | [string](#string) |  | stopping strategy id. |






<a name="phenixgo-StrategyFilter"></a>

### StrategyFilter
StrategyFilter represents a search query, with options strategy_id, strategy_name  and tag to indicate which results to include in the responose.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| strategy_id | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | strategy_id. |
| strategy_name | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | strategy_name. |
| tag | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | tag for classify. |
| status | [google.protobuf.StringValue](#google-protobuf-StringValue) | repeated |  |






<a name="phenixgo-UpdateUserStrategyRequest"></a>

### UpdateUserStrategyRequest
UpdateUserStrategyRequest represents request for updating a strategy.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| strategy_id | [string](#string) |  | specified target strategy id. |
| parameter | [google.protobuf.Any](#google-protobuf-Any) |  | specified update parameters. |






<a name="phenixgo-UpdateUserStrategyResponse"></a>

### UpdateUserStrategyResponse
UpdateUserStrategyResponse returns strategy id if update successfully.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| strategy_id | [string](#string) |  | updated strategy id. |





 

 

 


<a name="phenixgo-PhenixGO"></a>

### PhenixGO
PhenixGO GRPC functions

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| StartUserStrategy | [StartUserStrategyRequest](#phenixgo-StartUserStrategyRequest) | [StartUserStrategyResponse](#phenixgo-StartUserStrategyResponse) | StartUserStrategy function for initiate a strategy. |
| UpdateUserStrategy | [UpdateUserStrategyRequest](#phenixgo-UpdateUserStrategyRequest) | [UpdateUserStrategyResponse](#phenixgo-UpdateUserStrategyResponse) | UpdateUserStrategy function for updating a running strategy. |
| StopUserStrategy | [StopUserStrategyRequest](#phenixgo-StopUserStrategyRequest) | [StopUserStrategyResponse](#phenixgo-StopUserStrategyResponse) | StopUserStrategy function for stopping a running strategy. |
| GetUserStrategy | [GetStrategyRequest](#phenixgo-GetStrategyRequest) | [GetStrategyResponse](#phenixgo-GetStrategyResponse) | GetUserStrategy function for stopping a running strategy. |
| GetStrategyProfits | [GetStrategyProfitsRequest](#phenixgo-GetStrategyProfitsRequest) | [GetStrategyProfitsResponse](#phenixgo-GetStrategyProfitsResponse) | GetStrategyProfits function for retrieving strategy profit history. |
| GetStrategyState | [GetStrategyStateRequest](#phenixgo-GetStrategyStateRequest) | [GetStrategyStateResponse](#phenixgo-GetStrategyStateResponse) | GetStrategyState function for retrieving current strategy state. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |


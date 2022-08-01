# Error

## API Error

| Situation | Msg |
| --- | --- |
| api key(s) does not exist | api key not found |
| api key is not available | invalid exchange API key |
| the specified market not in the exchange | market not found |
| parameter does not match rule | invalid strategy parameter |
| strategy does not support the specified exchange | exchange not supported |
| third party exchange connect failed | exchange unavailable |
| program run into abnormal state | unexpected error |

## Runtime Error

| Situation | ErrMsg |
| --- | --- |
| run into unexpected state | unexpected error |
| determined order size less than the exchange condition | order size too small |
| exchange balance not afford to place the order | insufficient balance |
| api key is not available after running | invalid exchange API key |
| exchange request failed | exchange request failed |
| user call stop with force | force terminated |


# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/logging/logging.proto](#api/logging/logging.proto)
    - [GetLevelRequest](#onos.lib.go.logging.GetLevelRequest)
    - [GetLevelResponse](#onos.lib.go.logging.GetLevelResponse)
    - [SetLevelRequest](#onos.lib.go.logging.SetLevelRequest)
    - [SetLevelResponse](#onos.lib.go.logging.SetLevelResponse)
  
    - [Level](#onos.lib.go.logging.Level)
    - [ResponseStatus](#onos.lib.go.logging.ResponseStatus)
  
    - [logger](#onos.lib.go.logging.logger)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api/logging/logging.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/logging/logging.proto



<a name="onos.lib.go.logging.GetLevelRequest"></a>

### GetLevelRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logger_name | [string](#string) |  | logger name |






<a name="onos.lib.go.logging.GetLevelResponse"></a>

### GetLevelResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| level | [Level](#onos.lib.go.logging.Level) |  |  |






<a name="onos.lib.go.logging.SetLevelRequest"></a>

### SetLevelRequest
SetLevelRequest request for setting a logger level


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logger_name | [string](#string) |  | logger name |
| level | [Level](#onos.lib.go.logging.Level) |  | logger level |






<a name="onos.lib.go.logging.SetLevelResponse"></a>

### SetLevelResponse
SetLevelResponse response for setting a logger level


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response_status | [ResponseStatus](#onos.lib.go.logging.ResponseStatus) |  |  |





 


<a name="onos.lib.go.logging.Level"></a>

### Level
Logger level

| Name | Number | Description |
| ---- | ------ | ----------- |
| DEBUG | 0 | Debug log level |
| INFO | 1 | Info log level |
| WARN | 2 | Warn log level |
| ERROR | 3 | Error log level |
| DPANIC | 4 | DPanic log level |
| PANIC | 5 | Panic log level |
| FATAL | 6 | Fatal log level |



<a name="onos.lib.go.logging.ResponseStatus"></a>

### ResponseStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 |  |
| FAILED | 1 |  |
| PRECONDITION_FAILED | 2 |  |


 

 


<a name="onos.lib.go.logging.logger"></a>

### logger
logger service provides rpc functions to controller a logger remotely

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SetLevel | [SetLevelRequest](#onos.lib.go.logging.SetLevelRequest) | [SetLevelResponse](#onos.lib.go.logging.SetLevelResponse) | Sets a logger level |
| GetLevel | [GetLevelRequest](#onos.lib.go.logging.GetLevelRequest) | [GetLevelResponse](#onos.lib.go.logging.GetLevelResponse) | Gets a logger level |

 



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


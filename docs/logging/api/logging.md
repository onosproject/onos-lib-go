# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [pkg/logging/api/logging.proto](#pkg/logging/api/logging.proto)
    - [SetLevelRequest](#onos.lib.go.logger.SetLevelRequest)
    - [SetLevelResponse](#onos.lib.go.logger.SetLevelResponse)
    - [SetSinkRequest](#onos.lib.go.logger.SetSinkRequest)
    - [SetSinkResponse](#onos.lib.go.logger.SetSinkResponse)
  
    - [Level](#onos.lib.go.logger.Level)
    - [ResponseStatus](#onos.lib.go.logger.ResponseStatus)
  
  
    - [logger](#onos.lib.go.logger.logger)
  

- [Scalar Value Types](#scalar-value-types)



<a name="pkg/logging/api/logging.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pkg/logging/api/logging.proto



<a name="onos.lib.go.logger.SetLevelRequest"></a>

### SetLevelRequest
SetLevelRequest request for setting a logger level


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logger_name | [string](#string) |  | logger name |
| level | [Level](#onos.lib.go.logger.Level) |  | logger level |






<a name="onos.lib.go.logger.SetLevelResponse"></a>

### SetLevelResponse
SetLevelResponse response for setting a logger level


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response_status | [ResponseStatus](#onos.lib.go.logger.ResponseStatus) |  |  |






<a name="onos.lib.go.logger.SetSinkRequest"></a>

### SetSinkRequest
SetSinkRequest request for setting a sink


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logger_name | [string](#string) |  |  |






<a name="onos.lib.go.logger.SetSinkResponse"></a>

### SetSinkResponse
SetSinkResponse response for setting a sink


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response_status | [ResponseStatus](#onos.lib.go.logger.ResponseStatus) |  |  |





 


<a name="onos.lib.go.logger.Level"></a>

### Level
Logger level

| Name | Number | Description |
| ---- | ------ | ----------- |
| DEBUG | 0 | Debug log level |
| INFO | 1 | Info log level |
| ERROR | 2 | Error log level |
| DPANIC | 3 | DPanic log level |
| PANIC | 4 | Panic log level |
| FATAL | 5 | Fatal log level |



<a name="onos.lib.go.logger.ResponseStatus"></a>

### ResponseStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 |  |
| FAILED | 1 |  |
| PRECONDITION_FAILED | 2 |  |


 

 


<a name="onos.lib.go.logger.logger"></a>

### logger
logger service provides rpc functions to controller a logger remotely

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SetLevel | [SetLevelRequest](#onos.lib.go.logger.SetLevelRequest) | [SetLevelResponse](#onos.lib.go.logger.SetLevelResponse) | Sets a logger level |
| SetSink | [SetSinkRequest](#onos.lib.go.logger.SetSinkRequest) | [SetSinkResponse](#onos.lib.go.logger.SetSinkResponse) | Sets a sink for a logger |

 



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


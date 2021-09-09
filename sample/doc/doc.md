# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [sample/proto/pancake.proto](#sample/proto/pancake.proto)
    - [BakeRequest](#pancake.baker.BakeRequest)
    - [BakeResponse](#pancake.baker.BakeResponse)
    - [Pancake](#pancake.baker.Pancake)
    - [Report](#pancake.baker.Report)
    - [Report.BakeCount](#pancake.baker.Report.BakeCount)
    - [ReportRequest](#pancake.baker.ReportRequest)
    - [ReportResponse](#pancake.baker.ReportResponse)
  
    - [Pancake.Menu](#pancake.baker.Pancake.Menu)
  
    - [BakePancakeService](#pancake.baker.BakePancakeService)
  
- [sample/proto/sample.proto](#sample/proto/sample.proto)
    - [Email](#sample.Email)
    - [Person](#sample.Person)
    - [SampleRequest](#sample.SampleRequest)
    - [SampleResponse](#sample.SampleResponse)
  
    - [Email.EmailDomainType](#sample.Email.EmailDomainType)
  
    - [AddressService](#sample.AddressService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="sample/proto/pancake.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## sample/proto/pancake.proto



<a name="pancake.baker.BakeRequest"></a>

### BakeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| menu | [Pancake.Menu](#pancake.baker.Pancake.Menu) |  |  |






<a name="pancake.baker.BakeResponse"></a>

### BakeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pancake | [Pancake](#pancake.baker.Pancake) |  |  |






<a name="pancake.baker.Pancake"></a>

### Pancake



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baker_name | [string](#string) |  | name of chef |
| menu | [Pancake.Menu](#pancake.baker.Pancake.Menu) |  |  |
| techinical_score | [float](#float) |  | bakeing score |
| create_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | creation timestamp |






<a name="pancake.baker.Report"></a>

### Report
Report on how much panques were baked


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bake_counts | [Report.BakeCount](#pancake.baker.Report.BakeCount) | repeated |  |






<a name="pancake.baker.Report.BakeCount"></a>

### Report.BakeCount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| menu | [Pancake.Menu](#pancake.baker.Pancake.Menu) |  |  |
| count | [int32](#int32) |  |  |






<a name="pancake.baker.ReportRequest"></a>

### ReportRequest







<a name="pancake.baker.ReportResponse"></a>

### ReportResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| report | [Report](#pancake.baker.Report) |  |  |





 


<a name="pancake.baker.Pancake.Menu"></a>

### Pancake.Menu
pacake menu

| Name | Number | Description |
| ---- | ------ | ----------- |
| MENU_UNSPECIFIED | 0 |  |
| MENU_CLASSIC | 1 |  |
| MENU_BANANA | 2 |  |
| MENU_BACON_AND_CHEESE | 3 |  |
| MENU_BERRY | 4 |  |


 

 


<a name="pancake.baker.BakePancakeService"></a>

### BakePancakeService
Service about Pancake

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Bake | [BakeRequest](#pancake.baker.BakeRequest) | [BakeResponse](#pancake.baker.BakeResponse) | bake pancake service on the specified menu |
| Report | [ReportRequest](#pancake.baker.ReportRequest) | [ReportResponse](#pancake.baker.ReportResponse) | Get the total number of panques for each menu |

 



<a name="sample/proto/sample.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## sample/proto/sample.proto



<a name="sample.Email"></a>

### Email



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | email address |
| domain | [Email.EmailDomainType](#sample.Email.EmailDomainType) |  | domain type |






<a name="sample.Person"></a>

### Person



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | sequencial id |
| name | [string](#string) |  | his/her name |
| age | [string](#string) |  | his/her age |
| emails | [Email](#sample.Email) | repeated | his/her email |






<a name="sample.SampleRequest"></a>

### SampleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | sequencial id |






<a name="sample.SampleResponse"></a>

### SampleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person | [Person](#sample.Person) |  | person info |





 


<a name="sample.Email.EmailDomainType"></a>

### Email.EmailDomainType


| Name | Number | Description |
| ---- | ------ | ----------- |
| EMAIL_DOMAIN_TYPE_UNSPECIFIED | 0 |  |
| EMAIL_DOMAIN_TYPE_GMAIL | 1 | google |
| EMAIL_DOMAIN_TYPE_YAHOO_JAPAN | 2 | yahoo japan |
| EMAIL_DOMAIN_TYPE_OUTLOOK | 3 | ms outlook |
| EMAIL_DOMAIN_TYPE_OTHER | 4 | other domains |


 

 


<a name="sample.AddressService"></a>

### AddressService
Service that receives an id and returns Person information

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Search | [SampleRequest](#sample.SampleRequest) | [SampleResponse](#sample.SampleResponse) |  |

 



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


# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [geobufpb/geobuf.proto](#geobufpb_geobuf-proto)
    - [Data](#geobuf-Data)
    - [Data.Feature](#geobuf-Data-Feature)
    - [Data.FeatureCollection](#geobuf-Data-FeatureCollection)
    - [Data.Geometry](#geobuf-Data-Geometry)
    - [Data.Value](#geobuf-Data-Value)
  
    - [Data.Geometry.Type](#geobuf-Data-Geometry-Type)
  
- [Scalar Value Types](#scalar-value-types)



<a name="geobufpb_geobuf-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## geobufpb/geobuf.proto



<a name="geobuf-Data"></a>

### Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keys | [string](#string) | repeated | global arrays of unique keys |
| dimensions | [uint32](#uint32) |  | max coordinate dimensions, default 2 |
| precision | [uint32](#uint32) |  | number of digits after decimal point for coordinates, default 6 |
| feature_collection | [Data.FeatureCollection](#geobuf-Data-FeatureCollection) |  |  |
| feature | [Data.Feature](#geobuf-Data-Feature) |  |  |
| geometry | [Data.Geometry](#geobuf-Data-Geometry) |  |  |






<a name="geobuf-Data-Feature"></a>

### Data.Feature



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| geometry | [Data.Geometry](#geobuf-Data-Geometry) |  |  |
| id | [string](#string) |  |  |
| int_id | [sint64](#sint64) |  |  |
| values | [Data.Value](#geobuf-Data-Value) | repeated | unique values |
| properties | [uint32](#uint32) | repeated | pairs of key/value indexes |
| custom_properties | [uint32](#uint32) | repeated | arbitrary properties |






<a name="geobuf-Data-FeatureCollection"></a>

### Data.FeatureCollection



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| features | [Data.Feature](#geobuf-Data-Feature) | repeated |  |
| values | [Data.Value](#geobuf-Data-Value) | repeated |  |
| custom_properties | [uint32](#uint32) | repeated |  |






<a name="geobuf-Data-Geometry"></a>

### Data.Geometry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [Data.Geometry.Type](#geobuf-Data-Geometry-Type) |  |  |
| lengths | [uint32](#uint32) | repeated | coordinate structure in lengths |
| coords | [sint64](#sint64) | repeated | delta-encoded integer values |
| geometries | [Data.Geometry](#geobuf-Data-Geometry) | repeated |  |
| values | [Data.Value](#geobuf-Data-Value) | repeated |  |
| custom_properties | [uint32](#uint32) | repeated |  |






<a name="geobuf-Data-Value"></a>

### Data.Value



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| string_value | [string](#string) |  |  |
| double_value | [double](#double) |  |  |
| pos_int_value | [uint64](#uint64) |  |  |
| neg_int_value | [uint64](#uint64) |  |  |
| bool_value | [bool](#bool) |  |  |
| json_value | [bytes](#bytes) |  |  |





 


<a name="geobuf-Data-Geometry-Type"></a>

### Data.Geometry.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| EMPTY | 0 |  |
| POINT | 1 |  |
| MULTIPOINT | 2 |  |
| LINESTRING | 3 |  |
| MULTILINESTRING | 4 |  |
| POLYGON | 5 |  |
| MULTIPOLYGON | 6 |  |
| GEOMETRYCOLLECTION | 7 |  |


 

 

 



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


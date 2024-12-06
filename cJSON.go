package cjson

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type CJSON struct {
	Next        *CJSON
	Prev        *CJSON
	Child       *CJSON
	Type        c.Int
	Valuestring *int8
	Valueint    c.Int
	Valuedouble float64
	String      *int8
}

type Hooks struct {
	MallocFn unsafe.Pointer
	FreeFn   unsafe.Pointer
}
type Bool c.Int
//go:linkname Version C.cJSON_Version
func Version() *int8
// llgo:link (*Hooks).InitHooks C.cJSON_InitHooks
func (p *Hooks) InitHooks() {
}
//go:linkname Parse C.cJSON_Parse
func Parse(value *int8) *CJSON
//go:linkname ParseWithLength C.cJSON_ParseWithLength
func ParseWithLength(value *int8, buffer_length uintptr) *CJSON
//go:linkname ParseWithOpts C.cJSON_ParseWithOpts
func ParseWithOpts(value *int8, return_parse_end **int8, require_null_terminated Bool) *CJSON
//go:linkname ParseWithLengthOpts C.cJSON_ParseWithLengthOpts
func ParseWithLengthOpts(value *int8, buffer_length uintptr, return_parse_end **int8, require_null_terminated Bool) *CJSON
// llgo:link (*CJSON).Print C.cJSON_Print
func (p *CJSON) Print() *int8 {
	return nil
}
// llgo:link (*CJSON).PrintUnformatted C.cJSON_PrintUnformatted
func (p *CJSON) PrintUnformatted() *int8 {
	return nil
}
// llgo:link (*CJSON).PrintBuffered C.cJSON_PrintBuffered
func (p *CJSON) PrintBuffered(prebuffer c.Int, fmt Bool) *int8 {
	return nil
}
// llgo:link (*CJSON).PrintPreallocated C.cJSON_PrintPreallocated
func (p *CJSON) PrintPreallocated(buffer *int8, length c.Int, format Bool) Bool {
	return 0
}
// llgo:link (*CJSON).Delete C.cJSON_Delete
func (p *CJSON) Delete() {
}
// llgo:link (*CJSON).GetArraySize C.cJSON_GetArraySize
func (p *CJSON) GetArraySize() c.Int {
	return 0
}
// llgo:link (*CJSON).GetArrayItem C.cJSON_GetArrayItem
func (p *CJSON) GetArrayItem(index c.Int) *CJSON {
	return nil
}
// llgo:link (*CJSON).GetObjectItem C.cJSON_GetObjectItem
func (p *CJSON) GetObjectItem(string *int8) *CJSON {
	return nil
}
// llgo:link (*CJSON).GetObjectItemCaseSensitive C.cJSON_GetObjectItemCaseSensitive
func (p *CJSON) GetObjectItemCaseSensitive(string *int8) *CJSON {
	return nil
}
// llgo:link (*CJSON).HasObjectItem C.cJSON_HasObjectItem
func (p *CJSON) HasObjectItem(string *int8) Bool {
	return 0
}
//go:linkname GetErrorPtr C.cJSON_GetErrorPtr
func GetErrorPtr() *int8
// llgo:link (*CJSON).GetStringValue C.cJSON_GetStringValue
func (p *CJSON) GetStringValue() *int8 {
	return nil
}
// llgo:link (*CJSON).GetNumberValue C.cJSON_GetNumberValue
func (p *CJSON) GetNumberValue() float64 {
	return 0
}
// llgo:link (*CJSON).IsInvalid C.cJSON_IsInvalid
func (p *CJSON) IsInvalid() Bool {
	return 0
}
// llgo:link (*CJSON).IsFalse C.cJSON_IsFalse
func (p *CJSON) IsFalse() Bool {
	return 0
}
// llgo:link (*CJSON).IsTrue C.cJSON_IsTrue
func (p *CJSON) IsTrue() Bool {
	return 0
}
// llgo:link (*CJSON).IsBool C.cJSON_IsBool
func (p *CJSON) IsBool() Bool {
	return 0
}
// llgo:link (*CJSON).IsNull C.cJSON_IsNull
func (p *CJSON) IsNull() Bool {
	return 0
}
// llgo:link (*CJSON).IsNumber C.cJSON_IsNumber
func (p *CJSON) IsNumber() Bool {
	return 0
}
// llgo:link (*CJSON).IsString C.cJSON_IsString
func (p *CJSON) IsString() Bool {
	return 0
}
// llgo:link (*CJSON).IsArray C.cJSON_IsArray
func (p *CJSON) IsArray() Bool {
	return 0
}
// llgo:link (*CJSON).IsObject C.cJSON_IsObject
func (p *CJSON) IsObject() Bool {
	return 0
}
// llgo:link (*CJSON).IsRaw C.cJSON_IsRaw
func (p *CJSON) IsRaw() Bool {
	return 0
}
//go:linkname CreateNull C.cJSON_CreateNull
func CreateNull() *CJSON
//go:linkname CreateTrue C.cJSON_CreateTrue
func CreateTrue() *CJSON
//go:linkname CreateFalse C.cJSON_CreateFalse
func CreateFalse() *CJSON
// llgo:link Bool.CreateBool C.cJSON_CreateBool
func (p Bool) CreateBool() *CJSON {
	return nil
}
//go:linkname CreateNumber C.cJSON_CreateNumber
func CreateNumber(num float64) *CJSON
//go:linkname CreateString C.cJSON_CreateString
func CreateString(string *int8) *CJSON
//go:linkname CreateRaw C.cJSON_CreateRaw
func CreateRaw(raw *int8) *CJSON
//go:linkname CreateArray C.cJSON_CreateArray
func CreateArray() *CJSON
//go:linkname CreateObject C.cJSON_CreateObject
func CreateObject() *CJSON
//go:linkname CreateStringReference C.cJSON_CreateStringReference
func CreateStringReference(string *int8) *CJSON
// llgo:link (*CJSON).CreateObjectReference C.cJSON_CreateObjectReference
func (p *CJSON) CreateObjectReference() *CJSON {
	return nil
}
// llgo:link (*CJSON).CreateArrayReference C.cJSON_CreateArrayReference
func (p *CJSON) CreateArrayReference() *CJSON {
	return nil
}
//go:linkname CreateIntArray C.cJSON_CreateIntArray
func CreateIntArray(numbers *c.Int, count c.Int) *CJSON
//go:linkname CreateFloatArray C.cJSON_CreateFloatArray
func CreateFloatArray(numbers *float32, count c.Int) *CJSON
//go:linkname CreateDoubleArray C.cJSON_CreateDoubleArray
func CreateDoubleArray(numbers *float64, count c.Int) *CJSON
//go:linkname CreateStringArray C.cJSON_CreateStringArray
func CreateStringArray(strings **int8, count c.Int) *CJSON
// llgo:link (*CJSON).AddItemToArray C.cJSON_AddItemToArray
func (p *CJSON) AddItemToArray(item *CJSON) Bool {
	return 0
}
// llgo:link (*CJSON).AddItemToObject C.cJSON_AddItemToObject
func (p *CJSON) AddItemToObject(string *int8, item *CJSON) Bool {
	return 0
}
// llgo:link (*CJSON).AddItemToObjectCS C.cJSON_AddItemToObjectCS
func (p *CJSON) AddItemToObjectCS(string *int8, item *CJSON) Bool {
	return 0
}
// llgo:link (*CJSON).AddItemReferenceToArray C.cJSON_AddItemReferenceToArray
func (p *CJSON) AddItemReferenceToArray(item *CJSON) Bool {
	return 0
}
// llgo:link (*CJSON).AddItemReferenceToObject C.cJSON_AddItemReferenceToObject
func (p *CJSON) AddItemReferenceToObject(string *int8, item *CJSON) Bool {
	return 0
}
// llgo:link (*CJSON).DetachItemViaPointer C.cJSON_DetachItemViaPointer
func (p *CJSON) DetachItemViaPointer(item *CJSON) *CJSON {
	return nil
}
// llgo:link (*CJSON).DetachItemFromArray C.cJSON_DetachItemFromArray
func (p *CJSON) DetachItemFromArray(which c.Int) *CJSON {
	return nil
}
// llgo:link (*CJSON).DeleteItemFromArray C.cJSON_DeleteItemFromArray
func (p *CJSON) DeleteItemFromArray(which c.Int) {
}
// llgo:link (*CJSON).DetachItemFromObject C.cJSON_DetachItemFromObject
func (p *CJSON) DetachItemFromObject(string *int8) *CJSON {
	return nil
}
// llgo:link (*CJSON).DetachItemFromObjectCaseSensitive C.cJSON_DetachItemFromObjectCaseSensitive
func (p *CJSON) DetachItemFromObjectCaseSensitive(string *int8) *CJSON {
	return nil
}
// llgo:link (*CJSON).DeleteItemFromObject C.cJSON_DeleteItemFromObject
func (p *CJSON) DeleteItemFromObject(string *int8) {
}
// llgo:link (*CJSON).DeleteItemFromObjectCaseSensitive C.cJSON_DeleteItemFromObjectCaseSensitive
func (p *CJSON) DeleteItemFromObjectCaseSensitive(string *int8) {
}
// llgo:link (*CJSON).InsertItemInArray C.cJSON_InsertItemInArray
func (p *CJSON) InsertItemInArray(which c.Int, newitem *CJSON) Bool {
	return 0
}
// llgo:link (*CJSON).ReplaceItemViaPointer C.cJSON_ReplaceItemViaPointer
func (p *CJSON) ReplaceItemViaPointer(item *CJSON, replacement *CJSON) Bool {
	return 0
}
// llgo:link (*CJSON).ReplaceItemInArray C.cJSON_ReplaceItemInArray
func (p *CJSON) ReplaceItemInArray(which c.Int, newitem *CJSON) Bool {
	return 0
}
// llgo:link (*CJSON).ReplaceItemInObject C.cJSON_ReplaceItemInObject
func (p *CJSON) ReplaceItemInObject(string *int8, newitem *CJSON) Bool {
	return 0
}
// llgo:link (*CJSON).ReplaceItemInObjectCaseSensitive C.cJSON_ReplaceItemInObjectCaseSensitive
func (p *CJSON) ReplaceItemInObjectCaseSensitive(string *int8, newitem *CJSON) Bool {
	return 0
}
// llgo:link (*CJSON).Duplicate C.cJSON_Duplicate
func (p *CJSON) Duplicate(recurse Bool) *CJSON {
	return nil
}
// llgo:link (*CJSON).Compare C.cJSON_Compare
func (p *CJSON) Compare(b *CJSON, case_sensitive Bool) Bool {
	return 0
}
//go:linkname Minify C.cJSON_Minify
func Minify(json *int8)
// llgo:link (*CJSON).AddNullToObject C.cJSON_AddNullToObject
func (p *CJSON) AddNullToObject(name *int8) *CJSON {
	return nil
}
// llgo:link (*CJSON).AddTrueToObject C.cJSON_AddTrueToObject
func (p *CJSON) AddTrueToObject(name *int8) *CJSON {
	return nil
}
// llgo:link (*CJSON).AddFalseToObject C.cJSON_AddFalseToObject
func (p *CJSON) AddFalseToObject(name *int8) *CJSON {
	return nil
}
// llgo:link (*CJSON).AddBoolToObject C.cJSON_AddBoolToObject
func (p *CJSON) AddBoolToObject(name *int8, boolean Bool) *CJSON {
	return nil
}
// llgo:link (*CJSON).AddNumberToObject C.cJSON_AddNumberToObject
func (p *CJSON) AddNumberToObject(name *int8, number float64) *CJSON {
	return nil
}
// llgo:link (*CJSON).AddStringToObject C.cJSON_AddStringToObject
func (p *CJSON) AddStringToObject(name *int8, string *int8) *CJSON {
	return nil
}
// llgo:link (*CJSON).AddRawToObject C.cJSON_AddRawToObject
func (p *CJSON) AddRawToObject(name *int8, raw *int8) *CJSON {
	return nil
}
// llgo:link (*CJSON).AddObjectToObject C.cJSON_AddObjectToObject
func (p *CJSON) AddObjectToObject(name *int8) *CJSON {
	return nil
}
// llgo:link (*CJSON).AddArrayToObject C.cJSON_AddArrayToObject
func (p *CJSON) AddArrayToObject(name *int8) *CJSON {
	return nil
}
// llgo:link (*CJSON).SetNumberHelper C.cJSON_SetNumberHelper
func (p *CJSON) SetNumberHelper(number float64) float64 {
	return 0
}
// llgo:link (*CJSON).SetValuestring C.cJSON_SetValuestring
func (p *CJSON) SetValuestring(valuestring *int8) *int8 {
	return nil
}
//go:linkname Malloc C.cJSON_malloc
func Malloc(size uintptr) unsafe.Pointer
//go:linkname Free C.cJSON_free
func Free(object unsafe.Pointer)

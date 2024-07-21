package ctx

import (
	_ "embed"
	"sync"

	"github.com/buke/quickjs-go"
)

//go:embed node_fs.js
var node_fs string

//go:embed node_path.js
var node_path string

//go:embed node_process.js
var node_process string

//go:embed prelude.js
var prelude string

var Default = sync.OnceValue(func() *quickjs.Context {
	rt := quickjs.NewRuntime(quickjs.WithModuleImport(true))
	ctx := rt.NewContext()

	_, err := ctx.LoadModule(node_fs, "node:fs")
	if err != nil {
		panic(err)
	}
	_, err = ctx.LoadModule("export * from 'node:fs'; export { default } from 'node:fs';", "fs")
	if err != nil {
		panic(err)
	}
	_, err = ctx.LoadModule(node_path, "node:path")
	if err != nil {
		panic(err)
	}
	_, err = ctx.LoadModule("export * from 'node:path'; export { default } from 'node:path';", "path")
	if err != nil {
		panic(err)
	}
	_, err = ctx.LoadModule(node_path, "node:process")
	if err != nil {
		panic(err)
	}
	_, err = ctx.LoadModule("export * from 'node:process'; export { default } from 'node:process';", "path")
	if err != nil {
		panic(err)
	}

	p, err := ctx.Eval(prelude)
	if err != nil {
		panic(err)
	}
	_, err = ctx.Await(p)
	if err != nil {
		panic(err)
	}

	return ctx
})

func Array() *quickjs.Array {
	return Default().Array()
}
func ArrayBuffer(binaryData []byte) quickjs.Value {
	return Default().ArrayBuffer(binaryData)
}
func AsyncFunction(asyncFn func(ctx *quickjs.Context, this quickjs.Value, promise quickjs.Value, args []quickjs.Value) quickjs.Value) quickjs.Value {
	return Default().AsyncFunction(asyncFn)
}
func Atom(v string) quickjs.Atom {
	return Default().Atom(v)
}
func AtomIdx(idx int64) quickjs.Atom {
	return Default().AtomIdx(idx)
}
func Await(v quickjs.Value) (quickjs.Value, error) {
	return Default().Await(v)
}
func BigInt64(v int64) quickjs.Value {
	return Default().BigInt64(v)
}
func BigUint64(v uint64) quickjs.Value {
	return Default().BigUint64(v)
}
func Bool(b bool) quickjs.Value {
	return Default().Bool(b)
}
func Compile(code string, opts ...quickjs.EvalOption) ([]byte, error) {
	return Default().Compile(code, opts...)
}
func CompileFile(filePath string, opts ...quickjs.EvalOption) ([]byte, error) {
	return Default().CompileFile(filePath, opts...)
}
func Error(err error) quickjs.Value {
	return Default().Error(err)
}
func Eval(code string, opts ...quickjs.EvalOption) (quickjs.Value, error) {
	return Default().Eval(code, opts...)
}
func EvalBytecode(buf []byte) (quickjs.Value, error) {
	return Default().EvalBytecode(buf)
}
func EvalFile(filePath string, opts ...quickjs.EvalOption) (quickjs.Value, error) {
	return Default().EvalFile(filePath, opts...)
}
func Exception() error {
	return Default().Exception()
}
func Float64(v float64) quickjs.Value {
	return Default().Float64(v)
}
func Function(fn func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value) quickjs.Value {
	return Default().Function(fn)
}
func Globals() quickjs.Value {
	return Default().Globals()
}
func Int32(v int32) quickjs.Value {
	return Default().Int32(v)
}
func Int64(v int64) quickjs.Value {
	return Default().Int64(v)
}
func Invoke(fn quickjs.Value, this quickjs.Value, args ...quickjs.Value) quickjs.Value {
	return Default().Invoke(fn, this, args...)
}
func LoadModule(code string, moduleName string) (quickjs.Value, error) {
	return Default().LoadModule(code, moduleName)
}
func LoadModuleBytecode(buf []byte, moduleName string) (quickjs.Value, error) {
	return Default().LoadModuleBytecode(buf, moduleName)
}
func LoadModuleFile(filePath string, moduleName string) (quickjs.Value, error) {
	return Default().LoadModuleFile(filePath, moduleName)
}
func Loop() {
	Default().Loop()
}
func Map() *quickjs.Map {
	return Default().Map()
}
func Null() quickjs.Value {
	return Default().Null()
}
func Object() quickjs.Value {
	return Default().Object()
}
func ParseJSON(v string) quickjs.Value {
	return Default().ParseJSON(v)
}
func Runtime() *quickjs.Runtime {
	return Default().Runtime()
}
func Set() *quickjs.Set {
	return Default().Set()
}
func SetInterruptHandler(handler quickjs.InterruptHandler) {
	Default().SetInterruptHandler(handler)
}
func String(v string) quickjs.Value {
	return Default().String(v)
}
func Throw(v quickjs.Value) quickjs.Value {
	return Default().Throw(v)
}
func ThrowError(err error) quickjs.Value {
	return Default().ThrowError(err)
}
func ThrowInternalError(format string, args ...interface{}) quickjs.Value {
	return Default().ThrowInternalError(format, args...)
}
func ThrowRangeError(format string, args ...interface{}) quickjs.Value {
	return Default().ThrowRangeError(format, args...)
}
func ThrowReferenceError(format string, args ...interface{}) quickjs.Value {
	return Default().ThrowReferenceError(format, args...)
}
func ThrowSyntaxError(format string, args ...interface{}) quickjs.Value {
	return Default().ThrowSyntaxError(format, args...)
}
func ThrowTypeError(format string, args ...interface{}) quickjs.Value {
	return Default().ThrowTypeError(format, args...)
}
func Uint32(v uint32) quickjs.Value {
	return Default().Uint32(v)
}
func Undefined() quickjs.Value {
	return Default().Undefined()
}
func Uninitialized() quickjs.Value {
	return Default().Uninitialized()
}

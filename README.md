# wax - An WebAssembly host environment and tools

## How to use

Before starting using `wax` or other tools, you need to build them.

```
./scripts/build.sh
```

The build script builds `wax`, `wadisasm`, `wadump` or other tools, as well as examples in `examples` directory.


### `wax` command - executes .wasm files

You can invoke a function in a .wasm file and get its result:

```
./cmd/wax/wax -f "add" -a "i32:1" -a "i32:2" examples/go/add/main.wasm
```

Options:

- `-f`: name of a function to invoke. need to be an exported function.
- `-a`: argument. specify as `type:value`. you can specify `-a` multiple times.

The result will be shown in the following format:

```
0:i32:3
```

The first part (number `0`) is index of the returned value. (Due to the current version of WebAssembly spec, functions can have up to 1 returned value. So this is for the future extension.)
The second part (`i32`) is the type of the returned value.
The third part (number `3`) is the returned value of the function invocation.


### `wadump` command - dump .wasm module

`wadump` command reads a .wasm file and dump its parsed structure in JSON format:

```
./cmd/wadump/wadump examples/go/add/main.wasm
```

The result will be JSON string like the following:

```
{"Preamble":{"MagicNumber":1836278016,"Version":1},"Sections":[{"ID":1,"Size":30,"Content":"BmAAAX9gA39/fwF/YAAAYAF/AGACf38AYAJ/
...(snip)...
```

More readable `jq`ed version:

```
{
  "Preamble": {
    "MagicNumber": 1836278016,
    "Version": 1
  },
  "Sections": [
    {
      "ID": 1,
      "Size": 30,
      "Content": "BmAAAX9gA39/fwF/YAAAYAF/AGACf38AYAJ/fwF/",
      "FuncTypes": [
        {
          "ParamTypes": "",
          "ReturnTypes": "fw=="
        },
        {
          "ParamTypes": "f39/",
          "ReturnTypes": "fw=="
        },

        :
      (snip)
        :
  ]
}

```

### `wadisasm` command - disassemble .wasm module

`wadisasm` command reads a .wasm file and output disassembled list for a specified function or all functions in the module:

```
./cmd/wadisasm/wadisasm -n add examples/go/add/main.wasm
```

The output is like the following:

```
func:12
20 01 local.get 00000001
20 00 local.get 00000000
6a    i32.add
0b    end
```

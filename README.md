# jsonify

Test project for manipulating json and yaml with go

## Build

```bash
make
```

## convert


### JSON to YAML

```bash
echo '{"foo":"show_value_of_foo","bar": [ "buz", "cuz", "duz" ], "baz" : { "caz" : "fuz"}}' | jsonify
```

### YAML to JSON

```bash
echo '---
foo: show_value_of_foo
bar:
- buz
- cuz
- duz
baz:
  caz: fuz
' | jsonify
```

```bash
echo '---
foo: show_value_of_foo
bar:
- buz
- cuz
- duz
baz:
  caz: fuz
' | jsonify convert --noindent
```

```
{"bar":["buz","cuz","duz"],"baz":{"caz":"fuz"},"foo":"show_value_of_foo"}
```

## Convert files

```bash
jsonify convert ./tests/test.json
```

```bash
jsonify convert ./tests/test.yaml
```



```bash
jsonify convert --noindent tests/test.yaml
```

```
{"bar":["buz","cuz","duz"],"baz":{"caz":"fuz"},"flag":true,"foo":"show_value_of_foo","yyy":["one",2,true,"4",{"key":"value"},[1,"2",{"things":[{"complicated":true,"couldbe":"maybe","notreally":false}]}]],"zzz":{"buz":[1,2,3]}}
```

## Path

```bash
echo '{"foo":"show_value_of_foo","bar": [ "buz", "cuz", "duz" ], "baz" : { "caz" : "fuz"}}' | jsonify path -p "$.baz.caz"
```

Produces

```
fuz
```


```bash
jsonify path -p "$.baz.caz" tests/test.yaml
```
Produces

```
fuz
```

## Inspect

An attempt to print a go struct from json or yaml

```bash
echo '{"foo":"show_value_of_foo","bar": [ "buz", "cuz", "duz" ], "baz" : { "caz" : "fuz"}}' | jsonify inspect
```

Produces

```
package main

// Foo struct generated
type Foo struct {
	Foo	string			`json:"foo" yaml:"foo"`
	Bar	[]interface{}		`json:"bar" yaml:"bar"`
	Baz	map[string]interface{}	`json:"baz" yaml:"baz"`
	Baz	struct {
		Caz string `json:"caz" yaml:"caz"`
	}
}


```


```bash
jsonify inspect tests/test.yaml
```

Produces

```
package main

// Foo struct generated
type Foo struct {
	Baz	map[string]interface{}	`json:"baz" yaml:"baz"`
	Flag	bool			`json:"flag" yaml:"flag"`
	Foo	string			`json:"foo" yaml:"foo"`
	Yyy	[]interface{}		`json:"yyy" yaml:"yyy"`
	Zzz	map[string]interface{}	`json:"zzz" yaml:"zzz"`
	Bar	[]interface{}		`json:"bar" yaml:"bar"`
	Zzz	struct {
		Buz []interface{} `json:"buz" yaml:"buz"`
	}

	Baz	struct {
		Caz string `json:"caz" yaml:"caz"`
	}
}
```

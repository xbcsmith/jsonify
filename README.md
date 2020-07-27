# jsonify

A CLI for manipulating json and yaml

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

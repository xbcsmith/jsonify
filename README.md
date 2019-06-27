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

## Convert files

```bash
jsonify convert ./tests/test.json
```

```bash
jsonify convert ./tests/test.yaml
```


## Path

```bash
echo '{"foo":"show_value_of_foo","bar": [ "buz", "cuz", "duz" ], "baz" : { "caz" : "fuz"}}' | jsonify path -p "$.baz.caz"
```

```bash
jsonify path -p "$.baz.caz" tests/test.yaml
```

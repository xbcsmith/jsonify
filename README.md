# jsonify

Test project for manipulating json and yaml with go

## Build

```bash
make
```

## convert

### JSON to YAML

```bash
echo '{"foo":"show_value_of_foo","bar": [ "buz", "cuz", "duz" ], "baz" : { "caz" : "fuz"}}' | jsonify convert
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
' | jsonify convert
```

## Path

```bash
echo '{"foo":"show_value_of_foo","bar": [ "buz", "cuz", "duz" ], "baz" : { "caz" : "fuz"}}' | jsonify path -p "$.baz.caz"
```

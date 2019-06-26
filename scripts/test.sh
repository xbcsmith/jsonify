#!/bin/bash

BINARY="./bin/jsonify"

echo "Testing jsonify convert json to yaml"
echo
echo '{"foo":"show_value_of_foo","bar": [ "buz", "cuz", "duz" ], "baz" : { "caz" : "fuz"}}' | $BINARY convert
echo

echo "Testing jsonify inspect"
echo
echo '{"flag": true, "foo": {"baz": [1, 2, 3]}, "list": ["one", 2, true, "4", {"key": "value"}, [1, "2", false]]}' | $BINARY inspect
echo

echo "Testing jsonify convert yaml to json"
echo
echo '---
foo: show_value_of_foo
bar:
- buz
- cuz
- duz
baz:
  caz: fuz
' | $BINARY convert
echo

echo "Testing jsonify path"
echo
echo '{
    "store": {
        "book": [
            {
                "category": "reference",
                "author": "Nigel Rees",
                "title": "Sayings of the Century",
                "price": 8.95
            },
            {
                "category": "fiction",
                "author": "Evelyn Waugh",
                "title": "Sword of Honour",
                "price": 12.99
            },
            {
                "category": "fiction",
                "author": "Herman Melville",
                "title": "Moby Dick",
                "isbn": "0-553-21311-3",
                "price": 8.99
            },
            {
                "category": "fiction",
                "author": "J. R. R. Tolkien",
                "title": "The Lord of the Rings",
                "isbn": "0-395-19395-8",
                "price": 22.99
            }
        ],
        "bicycle": {
            "color": "red",
            "price": 19.95
        }
    },
    "expensive": 10
}' | $BINARY path -p "$.store.book[-1].price"
echo

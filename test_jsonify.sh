#!/bin/bash

go build

if [ ! $? == 0 ];then exit 0;fi

echo "Testing jsonify key foo"
echo
echo '{"foo":"show_value_of_foo","bar":"buz"}' | ./jsonify key foo
echo

echo "Testing jsonify key bar (COMPLEX)"
echo
echo '{"foo":"show_value_of_foo","bar": [ "buz", "cuz", "duz" ], "baz" : { "caz" : "fuz"}}' | ./jsonify key bar
echo

echo "Testing jsonify key baz (COMPLEX)"
echo
echo '{"foo":"show_value_of_foo","bar": [ "buz", "cuz", "duz" ], "baz" : { "caz" : "fuz"}}' | ./jsonify key baz
echo

echo "Testing jsonify inspect"
echo
echo '{"flag": true, "foo": {"baz": [1, 2, 3]}, "list": ["one", 2, true, "4", {"key": "value"}, [1, "2", false]]}' | ./jsonify inspect
echo

echo "Testing jsonify val buz"
echo
echo '{"foo":"baz","show_key_of_buz":"buz"}' | ./jsonify val buz
echo 

echo "Testing jsonify keys"
echo
echo '{"foo":"baz","bar":"buz","show_all_keys":"keys"}' | ./jsonify keys
echo

echo "Testing jsonify vals"
echo
echo '{"foo":"baz","bar":"buz","val":"show_all_values"}' | ./jsonify vals
echo

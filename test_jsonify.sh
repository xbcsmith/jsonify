#!/bin/bash

echo "jsonify -K foo"
echo '{"foo":"show_value_of_foo","bar":"buz"}' | ./jsonify -K foo
echo

echo "jsonify --key foo"
echo '{"foo":"show_value_of_foo","bar":"buz"}' | ./jsonify --key foo
echo

echo "jsonify -V buz"
echo '{"foo":"baz","show_key_of_buz":"buz"}' | ./jsonify -V buz
echo 

echo "jsonify --val buz"
echo '{"foo":"baz","show_key_of_buz":"buz"}' | ./jsonify -val buz
echo 

echo "jsonify --keys"
echo '{"foo":"baz","bar":"buz","show_all_keys":"keys"}' | ./jsonify --keys
echo

echo "jsonify -k"
echo '{"foo":"baz","bar":"buz","show_all_keys":"keys"}' | ./jsonify --k
echo

echo "jsonify --vals"
echo '{"foo":"baz","bar":"buz","val":"show_all_values"}' | ./jsonify --vals
echo

#!/bin/bash

new_directory=$(./hummingbird)

if [[ "$new_directory" != "x" ]]; then
    # Change to the new directory
    cd "$new_directory"
fi

clear
ls -l

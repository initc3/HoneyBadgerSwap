#!/usr/bin/env bash

app=$1
IDs=$2


# Set space as the delimiter
IFS=','

#Read the split words into an array based on space delimiter
read -a strarr <<< "$IDs"

# Print each value of the array by using the loop
for id in "${strarr[@]}";
do
  pkill -f "python3 -m ratel.src.python.$app.run $id"
  python3 -m ratel.src.python.$app.run $id &
done
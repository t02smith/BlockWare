#/bin/bash

contract="0x750cf6392175f94ff5014803a0bb6b79de543337"

for port in 7050 7051 
do
  ./toolkit -profile 2 -contract $contract -port $port &
done

./toolkit -profile 2 -contract $contract -port 7052 
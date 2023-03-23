#!/bin/bash

contract="0x750cf6392175f94ff5014803a0bb6b79de543337"
ports=(7050 7051 7052 7053 7054 7055)

for port in ${ports[@]} 
do
  ./toolkit -profile 2 -contract $contract -port $port &
done

./toolkit -profile 2 -contract $contract -port 9052 
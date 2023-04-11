#!/bin/bash

contract="0x750cf6392175f94ff5014803a0bb6b79de543337"
ports=(7051 7052 7053 7054 7055 7056 7057 7058 7059 7060 7061 7062 7063 7064)

for port in ${ports[@]} 
do
  ./toolkit -profile 2 -contract $contract -port $port &
done

./toolkit -profile 2 -contract $contract -port 7050
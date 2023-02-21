
ganache:
	ganache-cli -m "you know nothing jon snow"

compile-contracts:
	make compile-contract NAME=library

compile-contract:
	solcjs --abi ./contracts/$(NAME).sol -o build/abi
	solcjs --bin ./contracts/$(NAME).sol -o build/bin
	abigen --bin=./build/bin/contracts_$(NAME)_sol_$(NAME).bin --abi=./build/abi/contracts_$(NAME)_sol_$(NAME).abi --pkg=$(NAME) --out=./build/contracts/$(NAME)/$(NAME).go

compile-library:
	solc --abi ./contracts/library.sol -o build/abi --overwrite
	solc --bin ./contracts/library.sol -o build/bin --overwrite
	abigen --bin=./build/bin/Library.bin --abi=./build/abi/Library.abi --pkg=library --out=./build/contracts/library/library.go

setup-profiles:
	cp ./build/bin/blockware.exe ./test/profiles/1
	cp ./build/bin/blockware.exe ./test/profiles/2
	cp ./build/bin/blockware.exe ./test/profiles/3
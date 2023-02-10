
package-win:
	fyne package -os windows -icon ./static/img/icon.png

package-lin:
	fyne package -os linux -icon ./static/img/icon.png

ganache:
	ganache-cli -m "you know nothing jon snow"

compile-contracts:
	make compile-contract NAME=library

compile-contract:
	solcjs --abi ./contracts/$(NAME).sol -o build/abi
	solcjs --bin ./contracts/$(NAME).sol -o build/bin
	abigen --bin=./build/bin/contracts_$(NAME)_sol_$(NAME).bin --abi=./build/abi/contracts_$(NAME)_sol_$(NAME).abi --pkg=$(NAME) --out=./build/contracts/$(NAME)/$(NAME).go

compile-library:
	solc --abi ./contracts/library.sol -o build/abi
	solc --bin ./contracts/library.sol -o build/bin
	abigen --bin=./build/bin/Library.bin --abi=./build/abi/Library.abi --pkg=library --out=./build/contracts/library/library.go

package-win:
	fyne package -os windows -icon ./static/img/icon.png

package-lin:
	fyne package -os linux -icon ./static/img/icon.png

ganache:
	ganache-cli -m "you know nothing jon snow"

compile-contracts:
	make compile-contract NAME=store

compile-contract:
	cd build/contracts && mkdir $(NAME) && cd ../..
	solcjs --abi ./contracts/$(NAME).sol -o build/abi
	solcjs --bin ./contracts/$(NAME).sol -o build/bin
	abigen --bin=./build/bin/contracts_$(NAME)_sol_$(NAME).bin --abi=./build/abi/contracts_$(NAME)_sol_$(NAME).abi --pkg=$(NAME) --out=./build/contracts/$(NAME)/$(NAME).go
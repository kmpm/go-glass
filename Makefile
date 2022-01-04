



.PHONY: build
build: setup
	cd frontend; npm run build


.PHONY: setup
setup: frontend/node_modules
	@echo "setup done"

frontend/node_modules:
	cd frontend; npm install

.PHONY: start
start:
	cd frontend; npm start


.PHONY: test
test:
	@echo "TODO: Test frontend and backend"

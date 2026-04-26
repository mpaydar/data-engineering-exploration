
setup:
	@echo "Step 1: Generating the logs "
	python generate.py 
process: ${setup}
	@echo "Step 2: Filtering logs"
	python filter.py

run: setup process
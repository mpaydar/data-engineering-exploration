
setup:
	@echo "Step 1: Generating the logs "
	python python_exploration/generate.py
process: setup
	@echo "Step 2: Filtering logs"
	python python_exploration/filter.py

run: setup process
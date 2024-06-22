build_run_docker:
	docker build --tag asl-counters .
	docker run -d -p 8080:8080 asl-counters
	@echo "Starting local server: https://"$$(ipconfig getifaddr en0)":8080"

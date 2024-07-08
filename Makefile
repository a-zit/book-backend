test_cover:
	go test -coverprofile=coverage.out ./module/... && go tool cover -html=coverage.out
	CGO_ENABLED=1 go test ./module/... -coverprofile=coverage.out -json > test-report.json
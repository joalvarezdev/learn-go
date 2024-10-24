GO = go run

hello:
	@$(GO) 1-hello_world/main.go

values:
	@$(GO) 2-simple_values/main.go

variables:
	@$(GO) 3-variables/main.go

constants:
	@$(GO) 4-constants/main.go

loops:
	@$(GO) 5-loops/main.go
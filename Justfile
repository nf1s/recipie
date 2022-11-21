run:
  @go run .

build:
  @go build -o recipie

debug:
	@arch -arm64 dlv debug .

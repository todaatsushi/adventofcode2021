module github.com/todaatsushi/adventofcode2021

go 1.18

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v1.4.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

require internal/readings v1.0.0

replace internal/readings => ./internal/readings

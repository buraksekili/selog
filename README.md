# selog

Simple colorful logger for Go.

## Installation

`go get github.com/buraksekili/selog`

## Usage

```go
l := log.New(os.Stdout, "TEST", log.LstdFlags)
logger := selog.NewLogger(l)
logger.Success("IT WORKED %s\n", "burak"
```
![selog](https://user-images.githubusercontent.com/32663655/107861477-05982800-6e57-11eb-8d2a-7e763a24ed0b.png)


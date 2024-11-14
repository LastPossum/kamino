module testing.fun/kamino/benchmarks

go 1.21

toolchain go1.23.0

require (
	github.com/LastPossum/kamino v0.0.1
	github.com/barkimedes/go-deepcopy v0.0.0-20220514131651-17c30cfc62df
	github.com/jinzhu/copier v0.4.0
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/vmihailenco/msgpack v4.0.4+incompatible
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

replace github.com/LastPossum/kamino => ../

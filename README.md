# Learn mocking with Gomock

## Installation
We need to install the `gomock` package `github.com/golang/mock/gomock` as well as the `mockgen` code generation tool `github.com/golang/mock/mockgen`.
Install both packages using go install
`go install github.com/golang/mock/gomock`
`go install github.com/golang/mock/mockgen`

If you have `$GOPATH/bin` in your `$PATH`, you can verify that the binary was installed successfully by running
`mockgen`.
If `$GOPATH/bin` is not in your `$PATH` you have to call `mockgen` via `$GOPATH/bin/mockgen`. In our example below
we assume we have `$GOPATH/bin` in our `$PATH`.

## Basic Usage
Usage of *GoMock* follows four basic steps:
1. Use `mockgen` to generate a mock for the interface you wish to mock.
2. In your test, create an instance of `gomock.Controller` and pass it to your mock object's constructor to obtain a mock object.
3. Call `EXPECT()` on your mocks to set up their expectations and return values.
4. Call `Finish()` on the mock controller to assert the mock's expectations.

## Generating mocks
You can invoke the `mockgen` tool manually (e.g `mockgen -destination=mocks/mock_doer.go -package=mocks github.com/basebandit/mocking/doer Doer`) or use the `go:generate` (e.g. `go generate ./...`) tool to invoke it for you.
`mockgen` takes the following arguments:
- `-destination=mocks/mock_doer.go`:put the generated mocks in the file `mocks/mock_doer.go`
- `-package=mocks`:put the generated mocks in the package `mocks`
- `github.com/basebandit/mocking/doer`:generate mocks for this package.
- `Doer`:generate mocks for this interface. This argument is required - we need to specify the interfaces
to generate mocks for explicitly. We *can* however specify multiple interfaces here as a comma-separated list(e.g `Doer1,Doer2`).


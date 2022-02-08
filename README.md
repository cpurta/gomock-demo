# GoMock Demo

This is a simple demo repo to show off how one can use [gomock](https://github.com/golang/mock)
in place of creating your own mock(s). This repo is meant for educational purposes and is
not meant to be a serious implemenatation of a storage client.

## Project structure

`runner` - Meant to be an abstracted way to run the program that is using an injected
storage client, currently for this demo it will read from the examples text files in the
`testing` folder.

`storage` - This holds the interface for a `StorageClient` and each sub-folder is an
implementation of the interface.

`testing` - Holds 3 example text files which are read from in the main execution of the
program.

## GoMock Generation

If you have modified the `storage/interface.go` file at all you will need to regenerate
the gomock generated mocks. First, ensure that you have `mockgen` installed:

```
$ go install github.com/golang/mock/mockgen@v1.6.0
```

Once installed you can then use the `go generate` command from the project root.

```
$ go generate ./...
```

You should then see the newly generated mocks in the `storage/mock_storage` folder.

## License

MIT
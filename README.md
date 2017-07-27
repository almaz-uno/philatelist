# Philatelist
Philatelist is a RESTful server, that allows searching images in web by address or Google placeid. For this, it uses Google API at now.

## Building

Before compiling, please:

1. Tune GO environment (setup `$GOROOT`, `$GOPATH` environment variables)
2. Clone repository in `$GOPATH/src`

To compile, just type:
```sh
cd $GOPATH/src/bitbucket.org/CuredPlumbum/philatelist
make distr
```

Assembled files will be located at `$GOPATH/src/bitbucket.org/CuredPlumbum/philatelist/build`

To lint, test and coverage, type:
```sh
cd $GOPATH/src/bitbucket.org/CuredPlumbum/philatelist
make lint gocov-report
```


## Running

After compiling, copy and correct [philatelist.yaml](philatelist.yaml) file.
Most of parameters can be specified in yml-file or via command line interface.
Please, consider to review [philatelist.yaml](philatelist.yaml) for each parameter description.
Also, you can get additional information, just use `--help` CLI switch.

Please, use your own Google API key. Visit https://support.google.com/googleapi/answer/6158862 for information.

After running, you can use browser to inspect API:

For example (assume, the service is listening on localhost:9080):

1. http://localhost:9080/v1/images/address-text?query=Auckland
2. http://localhost:9080/v1/images/google-place-id?placeid=ChIJyWEHuEmuEmsRm9hTkapTCrk

# Writing RESTful client

Please, consider to review Swagger specification file [philatelist.swagger.yaml](philatelist.swagger.yaml) for service description 
or client creation.

Philatelist supports test invocation. Just review [test.go](cmd/test.go) file. It can be base for client creation. 


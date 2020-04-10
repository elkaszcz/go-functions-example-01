# golang functions

Based on https://codelabs.developers.google.com/codelabs/cloud-functions-go-http/#0

## Deploy
> First you need to `cd samplefunctions`

```console
gcloud functions deploy HelloWorld \
  --runtime go111 \
  --trigger-http \
  --allow-unauthenticated \
  --region=europe-west3 
```

## Test functions
> First you need to `cd samplefunctions`

```console
go test -v
```

## Run local server
```console
./run.sh
```
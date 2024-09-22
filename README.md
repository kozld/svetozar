# ton-api
grpc proxy into tonlib-go

## Methods
    rpc FetchTransactions (FetchTransactionsRequest) returns (FetchTransactionsResponse) {} 
    rpc GetAccountState (GetAccountStateRequest) returns (GetAccountStateResponse) {}       
    rpc GetBetSeed (GetBetSeedRequest) returns (GetBetSeedResponse) {}                     
    rpc GetSeqno (GetSeqnoRequest) returns (GetSeqnoResponse) {}                           
    rpc SendMessage(SendMessageRequest) returns (SendMessageResponse) {}     

## build
```docker build -t ton-api .```

## run (development)
```docker run -e CONTRACT_ADDR='kQBgMySA9l-X185xzrEm20M0INDbi0AavYOog9p0133yq-ZJ' --name ton-api ton-api```

## ENV VARS
    * TONLIB_CFG_PATH - path to tonlib.config.json. Default value is: `/usr/local/bin/app/tonlib.config.json.example`, the path to embedded version of config.
    * LISTEN_PORT - default value is `5400`.
    * CONTRACT_ADDR - required variable, no default value.


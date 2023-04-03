# NFT Feed Go APIs

## Code Structure
```
├── etc
│   └── loopx-api.yaml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── handler
│   │   ├── comment
│   │   │   ├── delete_handler.go
│   │   │   ├── read_handler.go
│   │   │   └── update_handler.go
│   │   ├── feeds
│   │   │   ├── list_handler.go
│   │   │   └── search_handler.go
│   │   ├── like
│   │   │   ├── delete_handler.go
│   │   │   ├── read_handler.go
│   │   │   └── update_handler.go
│   │   ├── routes.go
│   │   └── user
│   │       ├── account_handler.go
│   │       ├── avatar_handler.go
│   │       ├── captcha_handler.go
│   │       ├── follower_handler.go
│   │       ├── following_handler.go
│   │       ├── login_handler.go
│   │       ├── logout_handler.go
│   │       ├── me_handler.go
│   │       ├── notification_handler.go
│   │       └── settings_handler.go
│   ├── logic
│   │   ├── comment
│   │   │   ├── delete_logic.go
│   │   │   ├── read_logic.go
│   │   │   └── update_logic.go
│   │   ├── feeds
│   │   │   ├── list_logic.go
│   │   │   └── search_logic.go
│   │   ├── like
│   │   │   ├── delete_logic.go
│   │   │   ├── read_logic.go
│   │   │   └── update_logic.go
│   │   └── user
│   │       ├── account_logic.go
│   │       ├── avatar_logic.go
│   │       ├── captcha_logic.go
│   │       ├── follower_logic.go
│   │       ├── following_logic.go
│   │       ├── login_logic.go
│   │       ├── logout_logic.go
│   │       ├── me_logic.go
│   │       ├── notification_logic.go
│   │       └── settings_logic.go
│   ├── middleware
│   │   └── auth_middleware.go
│   ├── svc
│   │   └── service_context.go
│   └── types
│       └── types.go
└── loopx.go

```

## Installs
```bash
# for Go 1.15 and earlier
GO111MODULE=on go get -u github.com/zeromicro/go-zero/tools/goctl@master

# for Go 1.16 and later
go install github.com/zeromicro/go-zero/tools/goctl@master

# generate typescript api sdk
octl api ts --api api/loopx.api --dir sdk/ts

# generate the mysql and redis code
goctl model mysql
```

## key logic
Complete the webapi using axios, fetch or the xhr method. The handler is generated automatically, you only need to complete the code logic below:

```bash
internal/logic
├── comment
│   ├── delete_logic.go
│   ├── read_logic.go
│   └── update_logic.go
├── feeds
│   ├── list_logic.go
│   └── search_logic.go
├── like
│   ├── delete_logic.go
│   ├── read_logic.go
│   └── update_logic.go
└── user
    ├── account_logic.go
    ├── avatar_logic.go
    ├── captcha_logic.go
    ├── follower_logic.go
    ├── following_logic.go
    ├── login_logic.go
    ├── logout_logic.go
    ├── me_logic.go
    ├── notification_logic.go
    └── settings_logic.g
```

## auth middleware

1. complete auth logic ```internal/middleware/auth_middleware.go`
2. and then add the auth middleware in `internal/handler/routes.go` or add globally

## References
1. [go-zero](https://go-zero.dev)
2. [https://github.com/aws/aws-sdk-go.git](https://github.com/aws/aws-sdk-go.git)
3. [loopx_openapi2.0.json](https://github.com/deltafi-trade/looopx-api/blob/main/doc/loopx_openapi2.0.json)
4. [ts-sdk](https://github.com/deltafi-trade/looopx-api/blob/main/sdk/ts)
5. [golang wallet protocol](https://github.com/chaors/PublicBlockChain_go/blob/master/part10-address_Prototype/goLang%E5%85%AC%E9%93%BE%E5%AE%9E%E6%88%98%E4%B9%8B%E9%92%B1%E5%8C%85%26%E5%9C%B0%E5%9D%80.md)
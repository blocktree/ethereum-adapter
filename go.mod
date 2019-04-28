module github.com/blocktree/ethereum-adapter

go 1.12

require (
	github.com/allegro/bigcache v1.2.0 // indirect
	github.com/asdine/storm v2.1.2+incompatible
	github.com/astaxie/beego v1.11.1
	github.com/blocktree/go-owcdrivers v1.0.4
	github.com/blocktree/go-owcrypt v1.0.1
	github.com/blocktree/openwallet v1.4.1
	github.com/ethereum/go-ethereum v1.8.25
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/imroc/req v0.2.3
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/tidwall/gjson v1.2.1
)

//replace github.com/blocktree/openwallet => ../../openwallet

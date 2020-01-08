module github.com/blocktree/ethereum-adapter

go 1.12

require (
	github.com/asdine/storm v2.1.2+incompatible
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/go-owcdrivers v1.2.0
	github.com/blocktree/go-owcrypt v1.1.1
	github.com/blocktree/openwallet v1.7.0
	github.com/ethereum/go-ethereum v1.9.9
	github.com/imroc/req v0.2.4
	github.com/shopspring/decimal v0.0.0-20200105231215-408a2507e114
	github.com/tidwall/gjson v1.3.5
)

//replace github.com/blocktree/openwallet => ../../openwallet

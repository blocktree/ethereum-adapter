module github.com/blocktree/ethereum-adapter

go 1.12

require (
	github.com/astaxie/beego v1.12.1
	github.com/blocktree/openwallet/v2 v2.4.3
	github.com/blocktree/quorum-adapter v1.6.1
	github.com/ethereum/go-ethereum v1.10.17
)

//replace github.com/blocktree/quorum-adapter => ../quorum-adapter
//replace github.com/blocktree/openwallet/v2 => ../../openwallet

module github.com/blocktree/ethereum-adapter

go 1.12

require (
	github.com/astaxie/beego v1.12.1
	github.com/blocktree/go-owaddress v1.1.7
	github.com/blocktree/go-owcdrivers v1.2.0
	github.com/blocktree/go-owcrypt v1.1.2
	github.com/blocktree/openwallet/v2 v2.0.6
	github.com/blocktree/quorum-adapter v1.3.2
	github.com/ethereum/go-ethereum v1.9.9
)

//replace github.com/blocktree/quorum-adapter => ../quorum-adapter

//replace github.com/blocktree/openwallet/v2 => ../../openwallet

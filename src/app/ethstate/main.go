package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/syndtr/goleveldb/leveldb"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	app = cli.NewApp()
)

func init() {
	app.Action = run
}

func main() {
	db, err := NewLDBDatabase()

	if err != nil {
		fmt.Println("err:", err)
	}
	h := core.GetHeadHeaderHash(db)
	fmt.Println("h:", h)
}

func dbtest() {
	db, err := leveldb.OpenFile("/home/ubuntu/.ethereum/geth/chaindata/", nil)
	if err != nil {
		fmt.Println("err:", err)
	}
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		kv := make(map[string]interface{})
		fmt.Println("key:", rlp.Decode(bytes.NewReader(key), &kv))
		fmt.Println("value:", value)
		break
	}
	db.Close()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func gogogog() {
	fmt.Println("yes")
	h := common.HexToHash("0xaf38600cf2bc881eb3872a9ee97b40708c456d7a00c903d31233c9c24aa6a0ba")
	ldb, err := NewLDBDatabase()
	if err != nil {
		fmt.Println("db err:", err)
	}
	sdb := NewDatabase(ldb)
	statedb, err := NewStateDb(h, sdb)
	if err != nil {
		fmt.Println("state db err:", err)
	}
	fmt.Println("statedb:", statedb)

}

func run(ctx *cli.Context) error {
	return nil
}

func makeFullNode(ctx *cli.Context) *node.Node {
	return nil
}

func NewDatabase(db ethdb.Database) state.Database {
	return state.NewDatabase(db)
}

func NewStateDb(root common.Hash, db state.Database) (*state.StateDB, error) {
	statedb, err := state.New(root, db)
	return statedb, err
}

func NewLDBDatabase() (*ethdb.LDBDatabase, error) {
	return ethdb.NewLDBDatabase("/home/ubuntu/.ethereum/geth/chaindata/", 768, 512)
	//	return ethdb.NewLDBDatabase("/home/ubuntu/.ethereum/geth/chaindata/", 768, 512)
}

func NewBlockChain(db ethdb.Database, cacheConfig *core.CacheConfig, chainConfig *params.ChainConfig, engine consensus.Engine, vmConfig vm.Config) (*core.BlockChain, error) {
	return core.NewBlockChain(db, cacheConfig, chainConfig, engine, vmConfig)
}

func MakeChain(chainDb ethdb.Database) {
	cache := &core.CacheConfig{
		Disabled:      false,
		TrieNodeLimit: eth.DefaultConfig.TrieCache,
		TrieTimeLimit: eth.DefaultConfig.TrieTimeout,
	}
	fmt.Println("cache:", cache)
	vmcfg := vm.Config{EnablePreimageRecording: false}
	fmt.Println("vmcfg:", vmcfg)
}

//{ChainID: 1 Homestead: 1150000 DAO: 1920000 DAOSupport: true EIP150: 2463000 EIP155: 2675000 EIP158: 2675000 Byzantium: 4370000 Constantinople: <nil> Engine: ethash}

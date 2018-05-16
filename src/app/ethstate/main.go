package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/dashboard"
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
	log.Println("gogo")
	//	makeConfigNode()
	//	BlockDb()
	dbtest()
}

func statedbtest() {

}

func BlockDb() {
	db, err := NewLDBDatabase()

	if err != nil {
		fmt.Println("err:", err)
	}
	h := core.GetHeadHeaderHash(db)
	fmt.Println("h:", h)
	n := core.GetBlockNumber(db, h)
	n = 5581000
	h = common.HexToHash("0x188a5bd86db8a5951e5a202e1f5e5e48fd115d914adc6709bf18a53fc18bf804")
	fmt.Println("n:", n)
	b := core.GetBlock(db, h, n)
	fmt.Println("b.root:", b.Root())
	ndb := NewDatabase(db)
	sdb, err := NewStateDb(b.Root(), ndb)
	if err != nil {
		fmt.Println("sdb err:", err)
		return
	}
	addr := common.HexToAddress("0x06012c8cf97BEaD5deAe237070F9587f8E7A266d")
	fmt.Println("addr:", addr)
	sobj := sdb.GetOrNewStateObject(addr)
	fmt.Println("sobj:", sobj)
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
		//		kv := new(interface{})
		fmt.Println("key:", string(key))
		fmt.Println("value:", value)
		break

	}
	headerPrefix := []byte("h") // headerPrefix + num (uint64 big endian) + hash -> header
	numSuffix := []byte("n")    // headerPrefix + num (uint64 big endian) + numSuffix -> hash
	hh, _ := db.Get(append(append(headerPrefix, encodeBlockNumber(1)...), numSuffix...), nil)
	if len(hh) == 0 {
		fmt.Println("empty")
	}
	hdata, _ := db.Get(append(append(headerPrefix, encodeBlockNumber(1)...), hh...), nil)
	header := new(types.Header)
	if err = rlp.Decode(bytes.NewReader(hdata), header); err != nil {
		fmt.Println("decode err:", err)
	}
	data, err := json.Marshal(header)
	if err != nil {
		fmt.Println("marshal err:", err)
	}
	fmt.Println("data:", string(data))
	fmt.Println("header:", header)

	data, _ = db.Get([]byte("TrieSync"), nil)
	fmt.Println("data:", new(big.Int).SetBytes(data).Uint64())
	db.Close()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func encodeBlockNumber(number uint64) []byte {
	enc := make([]byte, 8)
	binary.BigEndian.PutUint64(enc, number)
	return enc
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

type ethstatsConfig struct {
	URL string `toml:",omitempty"`
}
type gethConfig struct {
	Eth eth.Config
	//	Shh       whisper.Config
	Node      node.Config
	Ethstats  ethstatsConfig
	Dashboard dashboard.Config
}

//{ChainID: 1 Homestead: 1150000 DAO: 1920000 DAOSupport: true EIP150: 2463000 EIP155: 2675000 EIP158: 2675000 Byzantium: 4370000 Constantinople: <nil> Engine: ethash}
func makeConfigNode() {
	cfg := gethConfig{
		Eth: eth.DefaultConfig,
		//		Shh:       whisper.DefaultConfig,
		Node:      node.DefaultConfig,
		Dashboard: dashboard.DefaultConfig,
	}
	cfg.Node.Name = "geth"
	stack, err := node.New(&cfg.Node)
	if err != nil {
		fmt.Println("err:", err)
	}
	utils.RegisterEthService(stack, &cfg.Eth)
	utils.StartNode(stack)

	var ethereum *eth.Ethereum
	err = stack.Service(&ethereum)
	if err != nil {
		fmt.Println("ethereum service not running:", err)
	}

	bc := ethereum.BlockChain()
	cb := bc.CurrentFastBlock()
	fmt.Println("bc.root:", cb.Root().Hex())
	fmt.Println("bc.body:", cb.Body())
	fmt.Println("bc.num:", cb.Number())
	ch := bc.CurrentHeader()
	fmt.Println("ch.root:", ch.Root)
	fmt.Println("ch.d:", ch.Difficulty)
	fmt.Println("ch.Num:", ch.Number)

	//	sdb, err := state.New(bc.GetBlockByNumber(5000000).Root(), bcsdb)
	sdb, err := bc.State()
	if err != nil {
		fmt.Println("sdb:", sdb)
	}
	tdb := sdb.Database().TrieDB()
	ns := tdb.Nodes()
	fmt.Println("nodes:", len(ns))

	addr := common.HexToAddress("0x06012c8cf97BEaD5deAe237070F9587f8E7A266d")

	sobj := sdb.GetOrNewStateObject(addr)
	fmt.Println("sobj:", sobj)
	fmt.Println("codehash:", sobj.CodeHash())
	fmt.Println("nonce:", sobj.Nonce())
	fmt.Println("exist:", sdb.Exist(addr))
	fmt.Println("code:", sdb.GetCode(addr))
	fmt.Println("trie:", sdb.StorageTrie(addr))
	//	stack.Wait()

}
func tryimport() {
	cfg := gethConfig{
		Eth: eth.DefaultConfig,
		//		Shh:       whisper.DefaultConfig,
		Node:      node.DefaultConfig,
		Dashboard: dashboard.DefaultConfig,
	}
	cfg.Node.Name = "geth"
	stack, err := node.New(&cfg.Node)
	if err != nil {
		fmt.Println("err:", err)
	}
	utils.RegisterEthService(stack, &cfg.Eth)
	utils.StartNode(stack)

	var ethereum *eth.Ethereum
	err = stack.Service(&ethereum)
	if err != nil {
		fmt.Println("ethereum service not running:", err)
	}

	bc := ethereum.BlockChain()
	cb := bc.CurrentFastBlock()
	fmt.Println("bc.root:", cb.Root().Hex())
	fmt.Println("bc.body:", cb.Body())
	fmt.Println("bc.num:", cb.Number())
	ch := bc.CurrentHeader()
	fmt.Println("ch.root:", ch.Root)
	fmt.Println("ch.d:", ch.Difficulty)
	fmt.Println("ch.Num:", ch.Number)

	//	sdb, err := state.New(bc.GetBlockByNumber(5000000).Root(), bcsdb)
	sdb, err := bc.State()
	if err != nil {
		fmt.Println("sdb:", sdb)
	}
	tdb := sdb.Database().TrieDB()
	ns := tdb.Nodes()
	fmt.Println("nodes:", len(ns))

	//	addr := common.HexToAddress("0x06012c8cf97BEaD5deAe237070F9587f8E7A266d")

}

package main

import (
	"app/couchdb-client/couchdb"
	"fmt"
	"time"
)

var compositeKeySep = []byte{0x00}

func main() {
	fmt.Println("begin...")
	//dbcondef,err := couchdb.CreateConnectionDefinition("http://10.70.21.173:5984", "", "", 2, 2, time.Duration(3000))
	dbinst, err := couchdb.CreateCouchInstance("10.70.21.173:5984", "", "", 1, 1, time.Duration(300000000))
	if err != nil {
		fmt.Println("err1:", err)
	}
	db, err := couchdb.CreateCouchDatabase(*dbinst, "channel0")
	if err != nil {
		fmt.Println("err2:", err)
	}
	info, dbret, err := db.GetDatabaseInfo()
	if err != nil {
		fmt.Println("err3:", err)
	}
	fmt.Println("info:", info)
	fmt.Println("dbret:", dbret)
	key := constructCompositeKey("sany-cc", "financeleasecontract:0000002")
	fmt.Printf("key:%s\n", key)
	doc, revision, err := db.ReadDoc(string(key))
	fmt.Printf("doc:%s\n", doc)
	fmt.Println("revision:", revision)
	fmt.Println("err:", err)
}

func constructCompositeKey(ns, key string) []byte {
	compositeKey := []byte(ns)
	compositeKey = append(compositeKey, compositeKeySep...)
	compositeKey = append(compositeKey, []byte(key)...)
	return compositeKey
}

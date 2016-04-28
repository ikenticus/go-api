package main

import (
//	"fmt"
	"github.com/couchbase/gocb"
)

func DocStoreUpsert(docKey string, docJson string) {
	var bucket string = "default"
	var docStore string = "couchbase://localhost"

	myCluster, _ := gocb.Connect(docStore)
	myBucket, _ := myCluster.OpenBucket(bucket, "")

	myBucket.Upsert(docKey, &docJson, 0)
}

func DocStoreGet(docKey string) interface{} {
	var bucket string = "default"
	var docStore string = "couchbase://localhost"

	myCluster, _ := gocb.Connect(docStore)
	myBucket, _ := myCluster.OpenBucket(bucket, "")

	var docJson interface{}
	myBucket.Get(docKey, &docJson)

	return docJson
}
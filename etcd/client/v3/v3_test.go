package etcd_v3

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/trist725/mgsu/util"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func Test_Base(t *testing.T) {

	Init(Endpoint)
	defer Close()

	key := "/service/config/"
	value := "{ \"address\": \"172.30.0.0/16\", \"port\": 555}"

	_, err := put(key, value)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("put key=[%s], value=[%s]\n", key, value)

	resp, err := get(key)
	if err != nil {
		t.Error(err)
		return
	}
	for _, kv := range resp.Kvs {
		fmt.Printf("get key=[%s], value=[%s]\n", kv.Key, kv.Value)
	}

	if _, err := delete_(key); err != nil {
		t.Error(err)
	} else {
		fmt.Printf("delete key=[%s]\n", key)
	}

}

func Test_GetWithRev(t *testing.T) {
	endpoint := os.Getenv("ETCD_ENDPOINT")
	if endpoint == "" {
		t.Error("not found env ETCD_ENDPOINT or ETCD_ENDPOINT is empty string")
		return
	}

	fmt.Printf("endpoint=[%s]\n", endpoint)

	Init(endpoint)
	defer Close()

	presp, err := gClient.Put(context.TODO(), "foo", "bar1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(presp.Header.Revision)
	_, err = gClient.Put(context.TODO(), "foo", "bar2")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := gClient.Get(ctx, "foo", clientv3.WithRev(presp.Header.Revision))
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func Test_PutGetDelete(t *testing.T) {
	endpoint := os.Getenv("ETCD_ENDPOINT")
	if endpoint == "" {
		t.Error("not found env ETCD_ENDPOINT or ETCD_ENDPOINT is empty string")
		return
	}

	fmt.Printf("endpoint=[%s]\n", endpoint)

	if err := Init(endpoint); err != nil {
		t.Error(err)
		return
	}
	defer Close()

	begin := time.Now()

	k := "test:key:1:2:3"

	for i := 0; i < 1000; i++ {
		v := util.GenRandomString(8)

		err := Put(k, v)

		v, err = Get(k)
		if err != nil {
			t.Error(err)
			return
		}

		//fmt.Println(v)

		_, err = Delete(k)
		if err != nil {
			t.Error(err)
			return
		}

		//fmt.Println(n)
	}

	fmt.Printf("consume=%v\n", time.Now().Sub(begin))
}

func Test_PutGetDeleteWithPrefix(t *testing.T) {
	endpoint := os.Getenv("ETCD_ENDPOINT")
	if endpoint == "" {
		t.Error("not found env ETCD_ENDPOINT or ETCD_ENDPOINT is empty string")
		return
	}

	fmt.Printf("endpoint=[%s]\n", endpoint)

	Init(endpoint)
	defer Close()

	for i := range make([]int, 3) {
		err := Put(fmt.Sprintf("key_%d", i), fmt.Sprintf("value_%d", i))
		if err != nil {
			log.Fatal(err)
		}
	}

	kvs, err := GetWithPrefix("key_")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(kvs)

	n, err := DeleteWithPrefix("key_")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(n)
}

func Test_Watch(t *testing.T) {
	endpoint := os.Getenv("ETCD_ENDPOINT")
	if endpoint == "" {
		t.Error("not found env ETCD_ENDPOINT or ETCD_ENDPOINT is empty string")
		return
	}

	fmt.Printf("endpoint=[%s]\n", endpoint)

	Init(endpoint)
	defer Close()

	k := "key"
	//v := "value"

	rch := WatchWithPrefix(k)
	go WatchLoop(rch, func(ev *clientv3.Event) {
		fmt.Printf("%v\n", ev)
	})

	time.Sleep(1 * time.Second)

	err := Put(k, util.GenRandomString(8))
	if err != nil {
		t.Error(err)
		return
	}

	_, err = Delete(k)
	if err != nil {
		t.Error(err)
		return
	}

	time.Sleep(5 * time.Second)
}

func Test_PutWithLease(t *testing.T) {
	endpoint := os.Getenv("ETCD_ENDPOINT")
	if endpoint == "" {
		t.Error("not found env ETCD_ENDPOINT or ETCD_ENDPOINT is empty string")
		return
	}

	fmt.Printf("endpoint=[%s]\n", endpoint)

	Init(endpoint)
	defer Close()

	leaseID, err := GrantLease(5)
	if err != nil {
		t.Error(err)
		return
	}

	k := "key"
	v := "value"

	err = PutWithLease(k, v, leaseID)
	if err != nil {
		t.Error(err)
		return
	}
}

func Test_KeepAlive(t *testing.T) {
	endpoint := os.Getenv("ETCD_ENDPOINT")
	if endpoint == "" {
		t.Error("not found env ETCD_ENDPOINT or ETCD_ENDPOINT is empty string")
		return
	}

	fmt.Printf("endpoint=[%s]\n", endpoint)

	Init(endpoint)
	defer Close()

	k := "foo"
	v := "bar"

	leaseID, err := GrantLease(5)

	err = PutWithLease(k, v, leaseID)
	if err != nil {
		t.Error(err)
		return
	}

	// the key 'foo' will be kept forever
	ch, err := KeepAliveLease(leaseID)
	if err != nil {
		t.Error(err)
		return
	}

	ka := <-ch
	fmt.Println("ttl:", ka.TTL)

	time.Sleep(5 * time.Second)
}

//func Test_KeepAliveOnce(t *testing.T) {
//	endpoint := os.Getenv("ETCD_ENDPOINT")
//	if endpoint == "" {
//		t.Error("not found env ETCD_ENDPOINT or ETCD_ENDPOINT is empty string")
//		return
//	}
//
//	fmt.Printf("endpoint=[%s]\n", endpoint)
//
//	Init(endpoint)
//	defer Close()
//
//	k := "foo"
//	v := "bar"
//
//	leaseID, err := GrantLease(5)
//
//	err = PutWithLease(k, v, leaseID)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	time.Sleep(4 * time.Second)
//
//	// to renew the lease only once
//	resp, err := KeepAliveLeaseOnce(leaseID)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	fmt.Println("ttl:", resp.TTL)
//}

func Test_Txn(t *testing.T) {
	endpoint := os.Getenv("ETCD_ENDPOINT")
	if endpoint == "" {
		t.Error("not found env ETCD_ENDPOINT or ETCD_ENDPOINT is empty string")
		return
	}

	fmt.Printf("endpoint=[%s]\n", endpoint)

	Init(endpoint)
	defer Close()

	kvc := clientv3.NewKV(gClient)

	_, err := kvc.Put(context.TODO(), "key", "xyz")
	if err != nil {
		t.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = kvc.Txn(ctx).
		If(clientv3.Compare(clientv3.Value("key"), ">", "abc")). // txn value comparisons are lexical
		Then(clientv3.OpPut("key", "XYZ")).                      // this runs, since 'xyz' > 'abc'
		Else(clientv3.OpPut("key", "ABC")).
		Commit()
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := kvc.Get(context.TODO(), "key")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
	// Output: key : XYZ
}

package etcd_v3

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	dialTimeout    = 3 * time.Second
	requestTimeout = 3 * time.Second
)

var (
	Endpoint string = "ip:12379"
	gClient  *clientv3.Client
)

func Init(endpoint string) error {
	var err error
	gClient, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: dialTimeout,
		Username:    "root",
		Password:    "pwd",
	})
	if gClient == nil || err != nil {
		return err
	}
	return nil
}

func Close() {
	gClient.Close()
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// minimum lease TTL is 5-second
func GrantLease(timeout int64) (clientv3.LeaseID, error) {
	resp, err := gClient.Grant(context.TODO(), timeout)
	if err != nil {
		return 0, err
	}
	return resp.ID, nil
}

// func RevokeLease(leaseID clientv3.LeaseID) (*clientv3.LeaseRevokeResponse, error) {
//	resp, err := gClient.Revoke(context.TODO(), leaseID)
//	return resp, err
// }

func KeepAliveLease(leaseID clientv3.LeaseID) (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	rch, err := gClient.KeepAlive(context.TODO(), leaseID)
	return rch, err
}

// func KeepAliveLeaseOnce(leaseID clientv3.LeaseID) (*clientv3.LeaseKeepAliveResponse, error) {
//	resp, err := gClient.KeepAliveOnce(context.TODO(), leaseID)
//	return resp, err
// }

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func put(key string, value string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := gClient.Put(ctx, key, value, opts...)
	cancel()
	return resp, err
}

func Put(key string, value string) error {
	_, err := put(key, value)
	return err
}

func PutWithLease(key string, value string, leaseID clientv3.LeaseID) error {
	_, err := put(key, value, clientv3.WithLease(leaseID))
	return err
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func get(key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := gClient.Get(ctx, key, opts...)
	cancel()
	return resp, err
}

type ErrNotFoundKey struct {
	Key string
}

func (e ErrNotFoundKey) Error() string {
	return fmt.Sprintf("not found key=[%s]", e.Key)
}

func Get(key string) (string, error) {
	resp, err := get(key)
	if err != nil {
		return "", err
	}
	if len(resp.Kvs) > 0 {
		return string(resp.Kvs[0].Value), nil
	}
	return "", &ErrNotFoundKey{Key: key}
}

func GetWithPrefix(prefix string) (map[string]string, error) {
	kvs := map[string]string{}
	resp, err := get(prefix, clientv3.WithPrefix())
	if err != nil {
		return kvs, err
	}
	for _, kv := range resp.Kvs {
		kvs[string(kv.Key)] = string(kv.Value)
	}
	return kvs, nil
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func delete_(key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := gClient.Delete(ctx, key, opts...)
	cancel()
	return resp, err
}

func Delete(key string) (int64, error) {
	resp, err := delete_(key)
	if err != nil {
		return 0, err
	}
	return resp.Deleted, err
}

func DeleteWithPrefix(prefix string) (int64, error) {
	resp, err := delete_(prefix, clientv3.WithPrefix())
	if err != nil {
		return 0, err
	}
	return resp.Deleted, nil
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func Compact(rev int64, opts ...clientv3.CompactOption) (*clientv3.CompactResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := gClient.Compact(ctx, rev, opts...)
	cancel()
	return resp, err
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func do(op clientv3.Op) (clientv3.OpResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := gClient.Do(ctx, op)
	cancel()
	return resp, err
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func txn(ctx context.Context) clientv3.Txn {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	t := gClient.Txn(ctx)
	cancel()
	return t
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func watch(key string, opts ...clientv3.OpOption) clientv3.WatchChan {
	rch := gClient.Watch(context.Background(), key, opts...)
	return rch
}

func Watch(key string) clientv3.WatchChan {
	return watch(key)
}

func WatchWithPrefix(prefix string) clientv3.WatchChan {
	return watch(prefix, clientv3.WithPrefix())
}

type WatchCallback func(ev *clientv3.Event)

func WatchLoop(rch clientv3.WatchChan, cb WatchCallback) {
	for {
		select {
		case resp := <-rch:
			for _, ev := range resp.Events {
				cb(ev)
			}
		}
	}
}

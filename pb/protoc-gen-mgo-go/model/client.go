package model

import (
	"context"
	"fmt"

	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	officialOpts "go.mongodb.org/mongo-driver/mongo/options"
)

var SC = NewSimpleClient()

type SimpleClient struct {
	cli *qmgo.QmgoClient
	uri string
}

func NewSimpleClient() (sc *SimpleClient) {
	sc = &SimpleClient{}
	return
}

func (sc *SimpleClient) Init(ctx context.Context, uri string, conf *qmgo.Config, o ...options.ClientOptions) (err error) {
	sc.uri = uri

	sc.cli, err = qmgo.Open(ctx, conf, o...)
	if err != nil {
		err = fmt.Errorf("connect to %s fail, %s", sc.uri, err)
		return
	}
	if err = sc.cli.Ping(3); err != nil {
		err = fmt.Errorf("ping to %s fail, %s", sc.uri, err)
		return
	}

	for _, seq := range seqs {
		err = sc.EnsureCounter(TblCounters, seq)
		if err != nil {
			err = fmt.Errorf("ensure counters [%s] error, %s", seq, err)
			return
		}
	}

	for tbl, indexes := range uniqueIndexes {
		for _, index := range indexes {
			err = sc.EnsureUniqueIndex(tbl, index)
			if err != nil {
				err = fmt.Errorf("ensure table[%s] unique index[%+v] error, %s", tbl, index, err)
				return
			}
		}
	}

	for tbl, is := range indexes {
		for _, index := range is {
			err = sc.EnsureIndexes(index)
			if err != nil {
				err = fmt.Errorf("ensure table[%s] index[%+v] error, %s", tbl, index, err)
				return
			}
		}
	}

	return
}

func (sc *SimpleClient) Release() {
	if sc.cli != nil {
		if err := sc.cli.Close(context.Background()); err != nil {
			panic(err)
		}
		sc.cli = nil
	}
}

func (sc SimpleClient) DBName() string {
	return sc.cli.GetDatabaseName()
}

func (sc *SimpleClient) NextSeq(id string) (int, error) {
	var res struct {
		Seq int
	}
	err := sc.cli.Database.Collection(TblCounters).Find(context.Background(), bson.M{"_id": id}).Apply(qmgo.Change{
		Update:    bson.M{"$inc": bson.M{"seq": 1}},
		ReturnNew: true,
	}, &res)
	return res.Seq, err
}

func (sc *SimpleClient) EnsureCounter(collection, id string) error {
	_, err := sc.cli.Database.Collection(collection).InsertOne(context.Background(), bson.M{
		"_id": id,
		"seq": 0,
	})
	if qmgo.IsDup(err) {
		return nil
	} else {
		return err
	}
}

func (sc *SimpleClient) EnsureIndexes(key []string) error {
	return sc.cli.CreateIndexes(context.Background(), []options.IndexModel{{Key: key}})
}

func (sc *SimpleClient) EnsureUniqueIndex(collection string, key []string) error {
	spares := true
	unique := true
	return sc.cli.Database.Collection(collection).CreateOneIndex(context.Background(), options.IndexModel{
		Key: key,
		IndexOptions: &officialOpts.IndexOptions{
			Sparse: &spares,
			Unique: &unique,
		},
	})
}

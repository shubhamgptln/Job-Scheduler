package scheduler

import (
	"context"
	etcd "github.com/job-scheduler/etcd/repository"
	clientv3 "github.com/coreos/etcd/client/v3"
	"time"
)

var (
	dialTimeout    = 2 * time.Second
	requestTimeout = 10 * time.Second
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
	cli, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints: []string{"127.0.0.1:2379"},
	})
	defer cli.Close()
	kv := clientv3.NewKV(cli)

	etcd.GetSingleValueDemo(ctx, kv)
	etcd.GetMultipleValuesWithPaginationDemo(ctx, kv)
	WatchDemo(ctx, cli, kv)
	etcd.LeaseDemo(ctx, cli, kv)
}



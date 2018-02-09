package cmd

import (
	"reflect"
	"testing"
	"time"
)

func TestOption(testing *testing.T) {
	localKubeServer := NewLocalkubeServer()
	etcdDataDirectory := "/home/jimma/tmp/etcd"
	testing.Logf("Etcd data directory: %s", etcdDataDirectory)
	etcd, err := localKubeServer.NewEtcd(etcdDataDirectory)
	if err != nil {
		panic(err)
	}
	logPrint(testing, etcd.Config)
	logPrint(testing, localKubeServer)
	// Start etcd first
	etcd.Start()
	time.Sleep(time.Hour)
}
func logPrint(testing *testing.T, t interface{}) {
	testing.Logf("----------------------%s--------------------\n", reflect.TypeOf(t))

	s := reflect.ValueOf(t).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		testing.Logf("%s %s = %v\n", typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

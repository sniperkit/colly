package memkv

import (
	"testing"
	"github.com/v2pro/quokka/kvstore"
	"github.com/stretchr/testify/require"
)

func Test_scan_metadata(t *testing.T) {
	ResetKVStore()
	should := require.New(t)
	kvstore.SetMetadata("server_1.2.3.4:8000", []byte("1.2.3.4:8000"))
	kvstore.SetMetadata("server_1.2.3.4:8001", []byte("1.2.3.4:8001"))
	iter, err := kvstore.ScanMetadata("server_", "server"+string([]byte{'_' + 1}))
	should.Nil(err)
	batch, err := iter()
	should.Nil(err)
	should.Equal("1.2.3.4:8000", string(batch[0].Value))
	should.Equal("1.2.3.4:8001", string(batch[1].Value))
	batch, err = iter()
	should.Nil(err)
	should.Nil(batch)
}

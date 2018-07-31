package tfstate

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomas.obenaus/inframapper/test/mock_tfstate_iface"
	"github.com/thomas.obenaus/inframapper/tfstate/tfstate_iface"
	"github.com/thomas.obenaus/inframapper/trace"
)

func TestSMNew(t *testing.T) {
	sm := NewStateLoader()
	require.NotNil(t, sm)
}

func TestSMLoad(t *testing.T) {
	sm := NewStateLoader()
	require.NotNil(t, sm)

	tfstate, err := sm.Load("ssss")
	assert.Error(t, err)
	assert.Nil(t, tfstate)

	tfstate, err = sm.Load("../examples/statefiles/instance.tfstate")
	assert.NoError(t, err)
	assert.NotNil(t, tfstate)

}

func TestSMLoadRemote(t *testing.T) {

	keys := make([]string, 0)
	keys = append(keys, "f1")
	keys = append(keys, "f2")
	keys = append(keys, "f3")

	remoteCfg := tfstate_iface.RemoteConfig{BucketName: "foo", Keys: keys}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockS3DownloaderAPI := mock_tfstate_iface.NewMockS3DownloaderAPI(mockCtrl)

	sl := tfStateLoader{tracer: trace.Off()}
	sl.loadRemoteStateImpl(remoteCfg, mockS3DownloaderAPI)
}

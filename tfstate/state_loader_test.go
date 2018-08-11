package tfstate

import (
	"fmt"
	"io"
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomas.obenaus/inframapper/test/mock_tfstate_iface"
	"github.com/thomas.obenaus/inframapper/tfstate/iface"
	"github.com/thomas.obenaus/inframapper/trace"
)

type emptyStateBuffer struct {
	buf []byte
	m   sync.Mutex
	foo float64
}

var dummyStateData = `{
	"version": 3,
	"terraform_version": "0.11.7",
	"serial": 3,
	"lineage": "3e6f20dc-3dfa-b8df-882b-1ccbbfe9c46c",
	"modules": [
			{
					"path": [
							"root"
					],
					"outputs": {},
					"resources": {},
					"depends_on": []
			}
	]
}`

func (b *emptyStateBuffer) WriteAt(p []byte, pos int64) (n int, err error) {
	return 0, nil
}

func (b *emptyStateBuffer) Bytes() []byte {
	return []byte(emptyStateData)
}

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

func TestSMLoadRemoteSuccess(t *testing.T) {

	keys := make([]string, 0)
	keys = append(keys, "f1")
	keys = append(keys, "f2")
	keys = append(keys, "f3")

	remoteCfg := iface.RemoteConfig{BucketName: "foo", Keys: keys}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockS3DownloaderAPI := mock_iface.NewMockS3DownloaderAPI(mockCtrl)
	var buffer []byte
	wBuffer := aws.NewWriteAtBuffer(buffer)
	call := mockS3DownloaderAPI.EXPECT().Download(wBuffer, gomock.Any()).Times(3)

	// write data into the buffer to avoid the error
	call.Do(func(wBuf io.WriterAt, oi *s3.GetObjectInput, fn ...func(*s3manager.Downloader)) {
		wBuf.WriteAt([]byte(dummyStateData), 0)
	})

	sl := tfStateLoader{tracer: trace.Off()}
	stateList, err := sl.loadRemoteStateImpl(remoteCfg, mockS3DownloaderAPI)

	assert.NoError(t, err)
	assert.NotEmpty(t, stateList)
}

func TestSMLoadRemoteFailNoData(t *testing.T) {

	keys := make([]string, 0)
	keys = append(keys, "f1")

	remoteCfg := iface.RemoteConfig{BucketName: "foo", Keys: keys}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockS3DownloaderAPI := mock_iface.NewMockS3DownloaderAPI(mockCtrl)
	mockS3DownloaderAPI.EXPECT().Download(gomock.Any(), gomock.Any()).Return(int64(0), fmt.Errorf("fail"))

	sl := tfStateLoader{tracer: trace.Off()}
	stateList, err := sl.loadRemoteStateImpl(remoteCfg, mockS3DownloaderAPI)

	assert.Error(t, err)
	assert.Empty(t, stateList)
}

func TestSMLoadRemoteFailParse(t *testing.T) {

	keys := make([]string, 0)
	keys = append(keys, "f1")

	remoteCfg := iface.RemoteConfig{BucketName: "foo", Keys: keys}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockS3DownloaderAPI := mock_iface.NewMockS3DownloaderAPI(mockCtrl)
	var buffer []byte
	wBuffer := aws.NewWriteAtBuffer(buffer)
	call := mockS3DownloaderAPI.EXPECT().Download(wBuffer, gomock.Any()).Times(1)

	// write invalif data into the buffer to force a parse error
	call.Do(func(wBuf io.WriterAt, oi *s3.GetObjectInput, fn ...func(*s3manager.Downloader)) {
		wBuf.WriteAt([]byte("uhuuhu"), 0)
	})

	sl := tfStateLoader{tracer: trace.Off()}
	stateList, err := sl.loadRemoteStateImpl(remoteCfg, mockS3DownloaderAPI)

	assert.Error(t, err)
	assert.Empty(t, stateList)
}

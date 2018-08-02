package tfstate

import (
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
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

func TestSMLoadRemote(t *testing.T) {

	keys := make([]string, 0)
	keys = append(keys, "f1")
	//keys = append(keys, "f2")
	//keys = append(keys, "f3")

	remoteCfg := iface.RemoteConfig{BucketName: "foo", Keys: keys}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockS3DownloaderAPI := mock_iface.NewMockS3DownloaderAPI(mockCtrl)
	//wBuffer := &emptyStateBuffer{}
	//var buffer []byte
	//wBuffer := aws.NewWriteAtBuffer(buffer)

	call := mockS3DownloaderAPI.EXPECT().Download(gomock.Any(), &s3.GetObjectInput{
		Bucket: aws.String(remoteCfg.BucketName),
		Key:    aws.String("f1"),
	})

	require.NotNil(t, call)

	call.Do(func(id int) {
	})

	//mockS3DownloaderAPI.EXPECT().Download(wBuffer, &s3.GetObjectInput{
	//	Bucket: aws.String(remoteCfg.BucketName),
	//	Key:    aws.String("f2"),
	//}).Return(int64(0), fmt.Errorf("fail"))

	sl := tfStateLoader{tracer: trace.Off()}
	stateList, err := sl.loadRemoteStateImpl(remoteCfg, mockS3DownloaderAPI)

	assert.NoError(t, err)
	assert.NotEmpty(t, stateList)
}

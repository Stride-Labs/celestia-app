package testnode

import (
	"fmt"
	"testing"
	"time"

	"github.com/celestiaorg/celestia-app/app"
	"github.com/celestiaorg/celestia-app/app/encoding"
	"github.com/celestiaorg/celestia-app/pkg/appconsts"
	appns "github.com/celestiaorg/celestia-app/pkg/namespace"
	blobtypes "github.com/celestiaorg/celestia-app/x/blob/types"
	abci "github.com/cometbft/cometbft/abci/types"
	tmrand "github.com/cometbft/cometbft/libs/rand"
	"github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

type IntegrationTestSuite struct {
	suite.Suite

	accounts []string
	cctx     Context
	rpc      *http.HTTP
}

func (s *IntegrationTestSuite) SetupSuite() {
	if testing.Short() {
		s.T().Skip("skipping full node integration test in short mode.")
	}
	t := s.T()

	accounts := make([]string, 40)
	for i := 0; i < 40; i++ {
		accounts[i] = tmrand.Str(10)
	}

	ecfg := encoding.MakeConfig(app.ModuleEncodingRegisters...)
	blobGenState := blobtypes.DefaultGenesis()
	blobGenState.Params.GovMaxSquareSize = uint64(appconsts.DefaultSquareSizeUpperBound)

	cfg := DefaultConfig().
		WithAccounts(accounts).
		WithGenesisOptions(SetBlobParams(ecfg.Codec, blobGenState.Params))

	cctx, rpcAddr, _ := NewNetwork(t, cfg)
	s.cctx = cctx
	var err error
	s.rpc, err = http.New(rpcAddr, "/websocket")
	require.NoError(t, err)
	s.accounts = accounts
}

// The "_Flaky" suffix indicates that the test may fail non-deterministically especially when executed in CI.
func (s *IntegrationTestSuite) Test_Liveness_Flaky() {
	require := s.Require()
	err := s.cctx.WaitForNextBlock()
	require.NoError(err)
	// check that we're actually able to set the consensus params
	var params *coretypes.ResultConsensusParams
	// this query can be flaky with fast block times, so we repeat it multiple
	// times in attempt to decrease flakiness
	for i := 0; i < 40; i++ {
		params, err = s.rpc.ConsensusParams(s.cctx.GoContext(), nil)
		if err == nil && params != nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	require.NoError(err)
	require.NotNil(params)
	_, err = s.cctx.WaitForHeight(20)
	require.NoError(err)
}

func (s *IntegrationTestSuite) Test_PostData() {
	require := s.Require()
	_, err := s.cctx.PostData(s.accounts[0], flags.BroadcastSync, appns.RandomBlobNamespace(), tmrand.Bytes(100000))
	require.NoError(err)
}

func (s *IntegrationTestSuite) Test_FillBlock() {
	require := s.Require()

	for squareSize := 2; squareSize <= appconsts.DefaultGovMaxSquareSize; squareSize *= 2 {
		resp, err := s.cctx.FillBlock(squareSize, s.accounts, flags.BroadcastSync)
		require.NoError(err)

		err = s.cctx.WaitForBlocks(3)
		require.NoError(err, squareSize)

		res, err := QueryWithoutProof(s.cctx.Context, resp.TxHash)
		require.NoError(err, squareSize)
		require.Equal(abci.CodeTypeOK, res.TxResult.Code, squareSize)

		b, err := s.cctx.Client.Block(s.cctx.GoContext(), &res.Height)
		require.NoError(err, squareSize)
		require.Equal(uint64(squareSize), b.Block.SquareSize, squareSize)
	}
}

func (s *IntegrationTestSuite) Test_FillBlock_InvalidSquareSizeError() {
	tests := []struct {
		name        string
		squareSize  int
		expectedErr error
	}{
		{
			name:        "when squareSize less than 2",
			squareSize:  0,
			expectedErr: fmt.Errorf("unsupported squareSize: 0"),
		},
		{
			name:        "when squareSize is greater than 2 but not a power of 2",
			squareSize:  18,
			expectedErr: fmt.Errorf("unsupported squareSize: 18"),
		},
		{
			name:       "when squareSize is greater than 2 and a power of 2",
			squareSize: 16,
		},
	}

	for _, tc := range tests {
		s.Run(tc.name, func() {
			_, actualErr := s.cctx.FillBlock(tc.squareSize, s.accounts, flags.BroadcastAsync)
			s.Equal(tc.expectedErr, actualErr)
		})
	}
}

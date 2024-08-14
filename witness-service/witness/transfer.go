package witness

import (
	"math/big"
	"time"

	solcommon "github.com/blocto/solana-go-sdk/common"
	soltypes "github.com/blocto/solana-go-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/iotexproject/ioTube/witness-service/grpc/types"
	"github.com/iotexproject/ioTube/witness-service/util"
)

// Transfer defines a record
type Transfer struct {
	cashier      common.Address
	token        common.Address
	coToken      util.Address
	index        uint64
	sender       common.Address
	recipient    util.Address
	ataOwner     util.Address // ata owner on solana
	amount       *big.Int
	decimalRound int
	fee          *big.Int
	id           common.Hash
	status       TransferStatus
	blockHeight  uint64
	txHash       common.Hash
	timestamp    time.Time
	gas          uint64
	gasPrice     *big.Int
	txSender     common.Address
	payload      []byte
}

func (t *Transfer) Cashier() util.Address {
	return util.ETHAddressToAddress(t.cashier)
}

func (t *Transfer) Token() util.Address {
	return util.ETHAddressToAddress(t.token)
}

func (t *Transfer) CoToken() util.Address {
	return t.coToken
}

func (t *Transfer) Index() *big.Int {
	return new(big.Int).SetUint64(t.index)
}

func (t *Transfer) ID() ([]byte, error) {
	for i := range t.id[:] {
		if t.id[i] != 0 {
			return t.id[:], nil
		}
	}
	return nil, errors.New("t id is empty")
}

func (t *Transfer) SetID(id common.Hash) {
	t.id = id
}

func (t *Transfer) BlockHeight() uint64 {
	return t.blockHeight
}

func (t *Transfer) ToTypesTransfer() *types.Transfer {
	gasPrice := "0"
	if t.gasPrice != nil {
		gasPrice = t.gasPrice.String()
	}
	ataOwnerBytes := []byte{}
	if t.ataOwner != nil {
		ataOwnerBytes = t.ataOwner.Bytes()
	}
	return &types.Transfer{
		Cashier:      t.cashier.Bytes(),
		Token:        t.coToken.Bytes(),
		Index:        int64(t.index),
		Sender:       t.sender.Bytes(),
		Recipient:    t.recipient.Bytes(),
		Amount:       roundAmount(t.amount, t.decimalRound).String(),
		Timestamp:    timestamppb.New(t.timestamp),
		Fee:          t.fee.String(),
		TxSender:     t.txSender.Bytes(),
		Gas:          t.gas,
		GasPrice:     gasPrice,
		Payload:      t.payload,
		SourceTxHash: t.txHash.Bytes(),
		AtaOwner:     ataOwnerBytes,
	}
}

func (t *Transfer) Sender() util.Address {
	return util.ETHAddressToAddress(t.sender)
}

func (t *Transfer) Recipient() util.Address {
	return t.recipient
}

func (t *Transfer) Payload() []byte {
	return t.payload
}

func (t *Transfer) Amount() *big.Int {
	return roundAmount(t.amount, t.decimalRound)
}

func roundAmount(raw *big.Int, cnt int) *big.Int {
	ret := big.NewInt(0).Set(raw)
	switch {
	case cnt == 0:
	case cnt < 0:
		cnt = -cnt
		for i := 0; i < cnt; i++ {
			ret.Div(ret, big.NewInt(10))
		}
	case cnt > 0:
		for i := 0; i < cnt; i++ {
			ret.Mul(ret, big.NewInt(10))
		}
	}
	return ret
}

func (t *Transfer) Status() TransferStatus {
	return t.status
}

// solTransfer defines a SOL record
type solTransfer struct {
	cashier      solcommon.PublicKey
	token        solcommon.PublicKey
	coToken      util.Address
	index        uint64
	sender       solcommon.PublicKey
	recipient    util.Address
	amount       *big.Int
	decimalRound int
	fee          *big.Int
	id           common.Hash
	status       TransferStatus
	blockHeight  uint64
	txSignature  soltypes.Signature
	txPayer      solcommon.PublicKey
	payload      []byte
}

func (s *solTransfer) Cashier() util.Address {
	return util.SOLAddressToAddress(s.cashier)
}

func (s *solTransfer) Token() util.Address {
	return util.SOLAddressToAddress(s.token)
}

func (s *solTransfer) CoToken() util.Address {
	return s.coToken
}

func (s *solTransfer) Index() *big.Int {
	return new(big.Int).SetUint64(s.index)
}

func (s *solTransfer) ID() ([]byte, error) {
	for i := range s.id[:] {
		if s.id[i] != 0 {
			return s.id[:], nil
		}
	}
	return nil, errors.New("t id is empty")
}

func (s *solTransfer) SetID(id common.Hash) {
	s.id = id
}

func (s *solTransfer) BlockHeight() uint64 {
	return s.blockHeight
}

func (s *solTransfer) Payload() []byte {
	return s.payload
}

func (s *solTransfer) Amount() *big.Int {
	return roundAmount(s.amount, s.decimalRound)
}

func (s *solTransfer) Sender() util.Address {
	return util.SOLAddressToAddress(s.sender)
}

func (s *solTransfer) Recipient() util.Address {
	return s.recipient
}

func (s *solTransfer) Status() TransferStatus {
	return s.status
}

func (s *solTransfer) ToTypesTransfer() *types.Transfer {
	return &types.Transfer{
		Cashier:      s.cashier.Bytes(),
		Token:        s.coToken.Bytes(),
		Index:        int64(s.index),
		Sender:       s.sender.Bytes(),
		Recipient:    s.recipient.Bytes(),
		Amount:       roundAmount(s.amount, s.decimalRound).String(),
		Timestamp:    timestamppb.Now(),
		Fee:          s.fee.String(),
		TxSender:     s.txPayer.Bytes(),
		Payload:      s.payload,
		SourceTxHash: s.txSignature,
	}
}

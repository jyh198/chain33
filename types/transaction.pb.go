// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// assert transfer struct
type AssetsGenesis struct {
	Amount        int64  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	ReturnAddress string `protobuf:"bytes,3,opt,name=returnAddress" json:"returnAddress,omitempty"`
}

func (m *AssetsGenesis) Reset()                    { *m = AssetsGenesis{} }
func (m *AssetsGenesis) String() string            { return proto.CompactTextString(m) }
func (*AssetsGenesis) ProtoMessage()               {}
func (*AssetsGenesis) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *AssetsGenesis) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *AssetsGenesis) GetReturnAddress() string {
	if m != nil {
		return m.ReturnAddress
	}
	return ""
}

type AssetsTransferToExec struct {
	Cointoken string `protobuf:"bytes,1,opt,name=cointoken" json:"cointoken,omitempty"`
	Amount    int64  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	Note      string `protobuf:"bytes,3,opt,name=note" json:"note,omitempty"`
	ExecName  string `protobuf:"bytes,4,opt,name=execName" json:"execName,omitempty"`
	To        string `protobuf:"bytes,5,opt,name=to" json:"to,omitempty"`
}

func (m *AssetsTransferToExec) Reset()                    { *m = AssetsTransferToExec{} }
func (m *AssetsTransferToExec) String() string            { return proto.CompactTextString(m) }
func (*AssetsTransferToExec) ProtoMessage()               {}
func (*AssetsTransferToExec) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *AssetsTransferToExec) GetCointoken() string {
	if m != nil {
		return m.Cointoken
	}
	return ""
}

func (m *AssetsTransferToExec) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *AssetsTransferToExec) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *AssetsTransferToExec) GetExecName() string {
	if m != nil {
		return m.ExecName
	}
	return ""
}

func (m *AssetsTransferToExec) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type AssetsWithdraw struct {
	Cointoken string `protobuf:"bytes,1,opt,name=cointoken" json:"cointoken,omitempty"`
	Amount    int64  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	Note      string `protobuf:"bytes,3,opt,name=note" json:"note,omitempty"`
	ExecName  string `protobuf:"bytes,4,opt,name=execName" json:"execName,omitempty"`
	To        string `protobuf:"bytes,5,opt,name=to" json:"to,omitempty"`
}

func (m *AssetsWithdraw) Reset()                    { *m = AssetsWithdraw{} }
func (m *AssetsWithdraw) String() string            { return proto.CompactTextString(m) }
func (*AssetsWithdraw) ProtoMessage()               {}
func (*AssetsWithdraw) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

func (m *AssetsWithdraw) GetCointoken() string {
	if m != nil {
		return m.Cointoken
	}
	return ""
}

func (m *AssetsWithdraw) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *AssetsWithdraw) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *AssetsWithdraw) GetExecName() string {
	if m != nil {
		return m.ExecName
	}
	return ""
}

func (m *AssetsWithdraw) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type AssetsTransfer struct {
	Cointoken string `protobuf:"bytes,1,opt,name=cointoken" json:"cointoken,omitempty"`
	Amount    int64  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	Note      string `protobuf:"bytes,3,opt,name=note" json:"note,omitempty"`
	To        string `protobuf:"bytes,4,opt,name=to" json:"to,omitempty"`
}

func (m *AssetsTransfer) Reset()                    { *m = AssetsTransfer{} }
func (m *AssetsTransfer) String() string            { return proto.CompactTextString(m) }
func (*AssetsTransfer) ProtoMessage()               {}
func (*AssetsTransfer) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{3} }

func (m *AssetsTransfer) GetCointoken() string {
	if m != nil {
		return m.Cointoken
	}
	return ""
}

func (m *AssetsTransfer) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *AssetsTransfer) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *AssetsTransfer) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type CreateTx struct {
	To          string `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
	Amount      int64  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	Fee         int64  `protobuf:"varint,3,opt,name=fee" json:"fee,omitempty"`
	Note        string `protobuf:"bytes,4,opt,name=note" json:"note,omitempty"`
	IsWithdraw  bool   `protobuf:"varint,5,opt,name=isWithdraw" json:"isWithdraw,omitempty"`
	IsToken     bool   `protobuf:"varint,6,opt,name=isToken" json:"isToken,omitempty"`
	TokenSymbol string `protobuf:"bytes,7,opt,name=tokenSymbol" json:"tokenSymbol,omitempty"`
	ExecName    string `protobuf:"bytes,8,opt,name=execName" json:"execName,omitempty"`
}

func (m *CreateTx) Reset()                    { *m = CreateTx{} }
func (m *CreateTx) String() string            { return proto.CompactTextString(m) }
func (*CreateTx) ProtoMessage()               {}
func (*CreateTx) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{4} }

func (m *CreateTx) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *CreateTx) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *CreateTx) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *CreateTx) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *CreateTx) GetIsWithdraw() bool {
	if m != nil {
		return m.IsWithdraw
	}
	return false
}

func (m *CreateTx) GetIsToken() bool {
	if m != nil {
		return m.IsToken
	}
	return false
}

func (m *CreateTx) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

func (m *CreateTx) GetExecName() string {
	if m != nil {
		return m.ExecName
	}
	return ""
}

type CreateTransactionGroup struct {
	Txs []string `protobuf:"bytes,1,rep,name=txs" json:"txs,omitempty"`
}

func (m *CreateTransactionGroup) Reset()                    { *m = CreateTransactionGroup{} }
func (m *CreateTransactionGroup) String() string            { return proto.CompactTextString(m) }
func (*CreateTransactionGroup) ProtoMessage()               {}
func (*CreateTransactionGroup) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{5} }

func (m *CreateTransactionGroup) GetTxs() []string {
	if m != nil {
		return m.Txs
	}
	return nil
}

type UnsignTx struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *UnsignTx) Reset()                    { *m = UnsignTx{} }
func (m *UnsignTx) String() string            { return proto.CompactTextString(m) }
func (*UnsignTx) ProtoMessage()               {}
func (*UnsignTx) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{6} }

func (m *UnsignTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// payAddr 可以支持 1. 地址 2. 私钥
type NoBalanceTx struct {
	TxHex   string `protobuf:"bytes,1,opt,name=txHex" json:"txHex,omitempty"`
	PayAddr string `protobuf:"bytes,2,opt,name=payAddr" json:"payAddr,omitempty"`
	Privkey string `protobuf:"bytes,3,opt,name=privkey" json:"privkey,omitempty"`
	Expire  string `protobuf:"bytes,4,opt,name=expire" json:"expire,omitempty"`
}

func (m *NoBalanceTx) Reset()                    { *m = NoBalanceTx{} }
func (m *NoBalanceTx) String() string            { return proto.CompactTextString(m) }
func (*NoBalanceTx) ProtoMessage()               {}
func (*NoBalanceTx) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{7} }

func (m *NoBalanceTx) GetTxHex() string {
	if m != nil {
		return m.TxHex
	}
	return ""
}

func (m *NoBalanceTx) GetPayAddr() string {
	if m != nil {
		return m.PayAddr
	}
	return ""
}

func (m *NoBalanceTx) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

func (m *NoBalanceTx) GetExpire() string {
	if m != nil {
		return m.Expire
	}
	return ""
}

type SignedTx struct {
	Unsign []byte `protobuf:"bytes,1,opt,name=unsign,proto3" json:"unsign,omitempty"`
	Sign   []byte `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty"`
	Pubkey []byte `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Ty     int32  `protobuf:"varint,4,opt,name=ty" json:"ty,omitempty"`
}

func (m *SignedTx) Reset()                    { *m = SignedTx{} }
func (m *SignedTx) String() string            { return proto.CompactTextString(m) }
func (*SignedTx) ProtoMessage()               {}
func (*SignedTx) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{8} }

func (m *SignedTx) GetUnsign() []byte {
	if m != nil {
		return m.Unsign
	}
	return nil
}

func (m *SignedTx) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *SignedTx) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *SignedTx) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

type Transaction struct {
	Execer    []byte     `protobuf:"bytes,1,opt,name=execer,proto3" json:"execer,omitempty"`
	Payload   []byte     `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature *Signature `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
	Fee       int64      `protobuf:"varint,4,opt,name=fee" json:"fee,omitempty"`
	Expire    int64      `protobuf:"varint,5,opt,name=expire" json:"expire,omitempty"`
	// 随机ID，可以防止payload 相同的时候，交易重复
	Nonce int64 `protobuf:"varint,6,opt,name=nonce" json:"nonce,omitempty"`
	// 对方地址，如果没有对方地址，可以为空
	To         string `protobuf:"bytes,7,opt,name=to" json:"to,omitempty"`
	GroupCount int32  `protobuf:"varint,8,opt,name=groupCount" json:"groupCount,omitempty"`
	Header     []byte `protobuf:"bytes,9,opt,name=header,proto3" json:"header,omitempty"`
	Next       []byte `protobuf:"bytes,10,opt,name=next,proto3" json:"next,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{9} }

func (m *Transaction) GetExecer() []byte {
	if m != nil {
		return m.Execer
	}
	return nil
}

func (m *Transaction) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Transaction) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Transaction) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Transaction) GetExpire() int64 {
	if m != nil {
		return m.Expire
	}
	return 0
}

func (m *Transaction) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Transaction) GetGroupCount() int32 {
	if m != nil {
		return m.GroupCount
	}
	return 0
}

func (m *Transaction) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Transaction) GetNext() []byte {
	if m != nil {
		return m.Next
	}
	return nil
}

type Transactions struct {
	Txs []*Transaction `protobuf:"bytes,1,rep,name=txs" json:"txs,omitempty"`
}

func (m *Transactions) Reset()                    { *m = Transactions{} }
func (m *Transactions) String() string            { return proto.CompactTextString(m) }
func (*Transactions) ProtoMessage()               {}
func (*Transactions) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{10} }

func (m *Transactions) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

// 环签名类型时，签名字段存储的环签名信息
type RingSignature struct {
	Items []*RingSignatureItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *RingSignature) Reset()                    { *m = RingSignature{} }
func (m *RingSignature) String() string            { return proto.CompactTextString(m) }
func (*RingSignature) ProtoMessage()               {}
func (*RingSignature) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{11} }

func (m *RingSignature) GetItems() []*RingSignatureItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// 环签名中的一组签名数据
type RingSignatureItem struct {
	Pubkey    [][]byte `protobuf:"bytes,1,rep,name=pubkey,proto3" json:"pubkey,omitempty"`
	Signature [][]byte `protobuf:"bytes,2,rep,name=signature,proto3" json:"signature,omitempty"`
}

func (m *RingSignatureItem) Reset()                    { *m = RingSignatureItem{} }
func (m *RingSignatureItem) String() string            { return proto.CompactTextString(m) }
func (*RingSignatureItem) ProtoMessage()               {}
func (*RingSignatureItem) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{12} }

func (m *RingSignatureItem) GetPubkey() [][]byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *RingSignatureItem) GetSignature() [][]byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// 对于一个交易组中的交易，要么全部成功，要么全部失败
// 这个要好好设计一下
// 最好交易构成一个链条[prevhash].独立的交易构成链条
// 只要这个组中有一个执行是出错的，那么就执行不成功
// 三种签名支持
// ty = 1 -> secp256k1
// ty = 2 -> ed25519
// ty = 3 -> sm2
// ty = 4 -> OnetimeED25519
// ty = 5 -> RingBaseonED25519
type Signature struct {
	Ty     int32  `protobuf:"varint,1,opt,name=ty" json:"ty,omitempty"`
	Pubkey []byte `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	// 当ty为5时，格式应该用RingSignature去解析
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{13} }

func (m *Signature) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

func (m *Signature) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type AddrOverview struct {
	Reciver int64 `protobuf:"varint,1,opt,name=reciver" json:"reciver,omitempty"`
	Balance int64 `protobuf:"varint,2,opt,name=balance" json:"balance,omitempty"`
	TxCount int64 `protobuf:"varint,3,opt,name=txCount" json:"txCount,omitempty"`
}

func (m *AddrOverview) Reset()                    { *m = AddrOverview{} }
func (m *AddrOverview) String() string            { return proto.CompactTextString(m) }
func (*AddrOverview) ProtoMessage()               {}
func (*AddrOverview) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{14} }

func (m *AddrOverview) GetReciver() int64 {
	if m != nil {
		return m.Reciver
	}
	return 0
}

func (m *AddrOverview) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *AddrOverview) GetTxCount() int64 {
	if m != nil {
		return m.TxCount
	}
	return 0
}

type ReqAddr struct {
	Addr string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	// 表示取所有/from/to/其他的hash列表
	Flag      int32 `protobuf:"varint,2,opt,name=flag" json:"flag,omitempty"`
	Count     int32 `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Direction int32 `protobuf:"varint,4,opt,name=direction" json:"direction,omitempty"`
	Height    int64 `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	Index     int64 `protobuf:"varint,6,opt,name=index" json:"index,omitempty"`
}

func (m *ReqAddr) Reset()                    { *m = ReqAddr{} }
func (m *ReqAddr) String() string            { return proto.CompactTextString(m) }
func (*ReqAddr) ProtoMessage()               {}
func (*ReqAddr) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{15} }

func (m *ReqAddr) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReqAddr) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *ReqAddr) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqAddr) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *ReqAddr) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReqAddr) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type ReqPrivacy struct {
	Count     int32 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Direction int32 `protobuf:"varint,2,opt,name=direction" json:"direction,omitempty"`
	Height    int64 `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
}

func (m *ReqPrivacy) Reset()                    { *m = ReqPrivacy{} }
func (m *ReqPrivacy) String() string            { return proto.CompactTextString(m) }
func (*ReqPrivacy) ProtoMessage()               {}
func (*ReqPrivacy) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{16} }

func (m *ReqPrivacy) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqPrivacy) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *ReqPrivacy) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type HexTx struct {
	Tx string `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
}

func (m *HexTx) Reset()                    { *m = HexTx{} }
func (m *HexTx) String() string            { return proto.CompactTextString(m) }
func (*HexTx) ProtoMessage()               {}
func (*HexTx) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{17} }

func (m *HexTx) GetTx() string {
	if m != nil {
		return m.Tx
	}
	return ""
}

type ReplyTxInfo struct {
	Hash   []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Height int64  `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	Index  int64  `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
}

func (m *ReplyTxInfo) Reset()                    { *m = ReplyTxInfo{} }
func (m *ReplyTxInfo) String() string            { return proto.CompactTextString(m) }
func (*ReplyTxInfo) ProtoMessage()               {}
func (*ReplyTxInfo) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{18} }

func (m *ReplyTxInfo) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ReplyTxInfo) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReplyTxInfo) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type ReqTxList struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *ReqTxList) Reset()                    { *m = ReqTxList{} }
func (m *ReqTxList) String() string            { return proto.CompactTextString(m) }
func (*ReqTxList) ProtoMessage()               {}
func (*ReqTxList) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{19} }

func (m *ReqTxList) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ReplyTxList struct {
	Txs []*Transaction `protobuf:"bytes,1,rep,name=txs" json:"txs,omitempty"`
}

func (m *ReplyTxList) Reset()                    { *m = ReplyTxList{} }
func (m *ReplyTxList) String() string            { return proto.CompactTextString(m) }
func (*ReplyTxList) ProtoMessage()               {}
func (*ReplyTxList) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{20} }

func (m *ReplyTxList) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type TxHashList struct {
	Hashes [][]byte `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
	Count  int64    `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Expire []int64  `protobuf:"varint,3,rep,packed,name=expire" json:"expire,omitempty"`
}

func (m *TxHashList) Reset()                    { *m = TxHashList{} }
func (m *TxHashList) String() string            { return proto.CompactTextString(m) }
func (*TxHashList) ProtoMessage()               {}
func (*TxHashList) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{21} }

func (m *TxHashList) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

func (m *TxHashList) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *TxHashList) GetExpire() []int64 {
	if m != nil {
		return m.Expire
	}
	return nil
}

type ReplyTxInfos struct {
	TxInfos []*ReplyTxInfo `protobuf:"bytes,1,rep,name=txInfos" json:"txInfos,omitempty"`
}

func (m *ReplyTxInfos) Reset()                    { *m = ReplyTxInfos{} }
func (m *ReplyTxInfos) String() string            { return proto.CompactTextString(m) }
func (*ReplyTxInfos) ProtoMessage()               {}
func (*ReplyTxInfos) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{22} }

func (m *ReplyTxInfos) GetTxInfos() []*ReplyTxInfo {
	if m != nil {
		return m.TxInfos
	}
	return nil
}

type ReceiptLog struct {
	Ty  int32  `protobuf:"varint,1,opt,name=ty" json:"ty,omitempty"`
	Log []byte `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
}

func (m *ReceiptLog) Reset()                    { *m = ReceiptLog{} }
func (m *ReceiptLog) String() string            { return proto.CompactTextString(m) }
func (*ReceiptLog) ProtoMessage()               {}
func (*ReceiptLog) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{23} }

func (m *ReceiptLog) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

func (m *ReceiptLog) GetLog() []byte {
	if m != nil {
		return m.Log
	}
	return nil
}

// ty = 0 -> error Receipt
// ty = 1 -> CutFee //cut fee ,bug exec not ok
// ty = 2 -> exec ok
type Receipt struct {
	Ty   int32         `protobuf:"varint,1,opt,name=ty" json:"ty,omitempty"`
	KV   []*KeyValue   `protobuf:"bytes,2,rep,name=KV" json:"KV,omitempty"`
	Logs []*ReceiptLog `protobuf:"bytes,3,rep,name=logs" json:"logs,omitempty"`
}

func (m *Receipt) Reset()                    { *m = Receipt{} }
func (m *Receipt) String() string            { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()               {}
func (*Receipt) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{24} }

func (m *Receipt) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

func (m *Receipt) GetKV() []*KeyValue {
	if m != nil {
		return m.KV
	}
	return nil
}

func (m *Receipt) GetLogs() []*ReceiptLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

type ReceiptData struct {
	Ty   int32         `protobuf:"varint,1,opt,name=ty" json:"ty,omitempty"`
	Logs []*ReceiptLog `protobuf:"bytes,3,rep,name=logs" json:"logs,omitempty"`
}

func (m *ReceiptData) Reset()                    { *m = ReceiptData{} }
func (m *ReceiptData) String() string            { return proto.CompactTextString(m) }
func (*ReceiptData) ProtoMessage()               {}
func (*ReceiptData) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{25} }

func (m *ReceiptData) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

func (m *ReceiptData) GetLogs() []*ReceiptLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

type TxResult struct {
	Height      int64        `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	Index       int32        `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	Tx          *Transaction `protobuf:"bytes,3,opt,name=tx" json:"tx,omitempty"`
	Receiptdate *ReceiptData `protobuf:"bytes,4,opt,name=receiptdate" json:"receiptdate,omitempty"`
	Blocktime   int64        `protobuf:"varint,5,opt,name=blocktime" json:"blocktime,omitempty"`
	ActionName  string       `protobuf:"bytes,6,opt,name=actionName" json:"actionName,omitempty"`
}

func (m *TxResult) Reset()                    { *m = TxResult{} }
func (m *TxResult) String() string            { return proto.CompactTextString(m) }
func (*TxResult) ProtoMessage()               {}
func (*TxResult) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{26} }

func (m *TxResult) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TxResult) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TxResult) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *TxResult) GetReceiptdate() *ReceiptData {
	if m != nil {
		return m.Receiptdate
	}
	return nil
}

func (m *TxResult) GetBlocktime() int64 {
	if m != nil {
		return m.Blocktime
	}
	return 0
}

func (m *TxResult) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

type TransactionDetail struct {
	Tx         *Transaction `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
	Receipt    *ReceiptData `protobuf:"bytes,2,opt,name=receipt" json:"receipt,omitempty"`
	Proofs     [][]byte     `protobuf:"bytes,3,rep,name=proofs,proto3" json:"proofs,omitempty"`
	Height     int64        `protobuf:"varint,4,opt,name=height" json:"height,omitempty"`
	Index      int64        `protobuf:"varint,5,opt,name=index" json:"index,omitempty"`
	Blocktime  int64        `protobuf:"varint,6,opt,name=blocktime" json:"blocktime,omitempty"`
	Amount     int64        `protobuf:"varint,7,opt,name=amount" json:"amount,omitempty"`
	Fromaddr   string       `protobuf:"bytes,8,opt,name=fromaddr" json:"fromaddr,omitempty"`
	ActionName string       `protobuf:"bytes,9,opt,name=actionName" json:"actionName,omitempty"`
}

func (m *TransactionDetail) Reset()                    { *m = TransactionDetail{} }
func (m *TransactionDetail) String() string            { return proto.CompactTextString(m) }
func (*TransactionDetail) ProtoMessage()               {}
func (*TransactionDetail) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{27} }

func (m *TransactionDetail) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *TransactionDetail) GetReceipt() *ReceiptData {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func (m *TransactionDetail) GetProofs() [][]byte {
	if m != nil {
		return m.Proofs
	}
	return nil
}

func (m *TransactionDetail) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TransactionDetail) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TransactionDetail) GetBlocktime() int64 {
	if m != nil {
		return m.Blocktime
	}
	return 0
}

func (m *TransactionDetail) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TransactionDetail) GetFromaddr() string {
	if m != nil {
		return m.Fromaddr
	}
	return ""
}

func (m *TransactionDetail) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

type TransactionDetails struct {
	Txs []*TransactionDetail `protobuf:"bytes,1,rep,name=txs" json:"txs,omitempty"`
}

func (m *TransactionDetails) Reset()                    { *m = TransactionDetails{} }
func (m *TransactionDetails) String() string            { return proto.CompactTextString(m) }
func (*TransactionDetails) ProtoMessage()               {}
func (*TransactionDetails) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{28} }

func (m *TransactionDetails) GetTxs() []*TransactionDetail {
	if m != nil {
		return m.Txs
	}
	return nil
}

type ReqAddrs struct {
	Addrs []string `protobuf:"bytes,1,rep,name=addrs" json:"addrs,omitempty"`
}

func (m *ReqAddrs) Reset()                    { *m = ReqAddrs{} }
func (m *ReqAddrs) String() string            { return proto.CompactTextString(m) }
func (*ReqAddrs) ProtoMessage()               {}
func (*ReqAddrs) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{29} }

func (m *ReqAddrs) GetAddrs() []string {
	if m != nil {
		return m.Addrs
	}
	return nil
}

type ReqDecodeRawTransaction struct {
	TxHex string `protobuf:"bytes,1,opt,name=txHex" json:"txHex,omitempty"`
}

func (m *ReqDecodeRawTransaction) Reset()                    { *m = ReqDecodeRawTransaction{} }
func (m *ReqDecodeRawTransaction) String() string            { return proto.CompactTextString(m) }
func (*ReqDecodeRawTransaction) ProtoMessage()               {}
func (*ReqDecodeRawTransaction) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{30} }

func (m *ReqDecodeRawTransaction) GetTxHex() string {
	if m != nil {
		return m.TxHex
	}
	return ""
}

type ReqTokenTx struct {
	Symbol string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	// 表示取所有/from/to/其他的hash列表
	Flag      int32  `protobuf:"varint,2,opt,name=flag" json:"flag,omitempty"`
	Count     int32  `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Direction int32  `protobuf:"varint,4,opt,name=direction" json:"direction,omitempty"`
	Height    int64  `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	Index     int64  `protobuf:"varint,6,opt,name=index" json:"index,omitempty"`
	Addr      string `protobuf:"bytes,7,opt,name=addr" json:"addr,omitempty"`
}

func (m *ReqTokenTx) Reset()                    { *m = ReqTokenTx{} }
func (m *ReqTokenTx) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenTx) ProtoMessage()               {}
func (*ReqTokenTx) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{31} }

func (m *ReqTokenTx) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqTokenTx) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *ReqTokenTx) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqTokenTx) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *ReqTokenTx) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReqTokenTx) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ReqTokenTx) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func init() {
	proto.RegisterType((*AssetsGenesis)(nil), "types.AssetsGenesis")
	proto.RegisterType((*AssetsTransferToExec)(nil), "types.AssetsTransferToExec")
	proto.RegisterType((*AssetsWithdraw)(nil), "types.AssetsWithdraw")
	proto.RegisterType((*AssetsTransfer)(nil), "types.AssetsTransfer")
	proto.RegisterType((*CreateTx)(nil), "types.CreateTx")
	proto.RegisterType((*CreateTransactionGroup)(nil), "types.CreateTransactionGroup")
	proto.RegisterType((*UnsignTx)(nil), "types.UnsignTx")
	proto.RegisterType((*NoBalanceTx)(nil), "types.NoBalanceTx")
	proto.RegisterType((*SignedTx)(nil), "types.SignedTx")
	proto.RegisterType((*Transaction)(nil), "types.Transaction")
	proto.RegisterType((*Transactions)(nil), "types.Transactions")
	proto.RegisterType((*RingSignature)(nil), "types.RingSignature")
	proto.RegisterType((*RingSignatureItem)(nil), "types.RingSignatureItem")
	proto.RegisterType((*Signature)(nil), "types.Signature")
	proto.RegisterType((*AddrOverview)(nil), "types.AddrOverview")
	proto.RegisterType((*ReqAddr)(nil), "types.ReqAddr")
	proto.RegisterType((*ReqPrivacy)(nil), "types.ReqPrivacy")
	proto.RegisterType((*HexTx)(nil), "types.HexTx")
	proto.RegisterType((*ReplyTxInfo)(nil), "types.ReplyTxInfo")
	proto.RegisterType((*ReqTxList)(nil), "types.ReqTxList")
	proto.RegisterType((*ReplyTxList)(nil), "types.ReplyTxList")
	proto.RegisterType((*TxHashList)(nil), "types.TxHashList")
	proto.RegisterType((*ReplyTxInfos)(nil), "types.ReplyTxInfos")
	proto.RegisterType((*ReceiptLog)(nil), "types.ReceiptLog")
	proto.RegisterType((*Receipt)(nil), "types.Receipt")
	proto.RegisterType((*ReceiptData)(nil), "types.ReceiptData")
	proto.RegisterType((*TxResult)(nil), "types.TxResult")
	proto.RegisterType((*TransactionDetail)(nil), "types.TransactionDetail")
	proto.RegisterType((*TransactionDetails)(nil), "types.TransactionDetails")
	proto.RegisterType((*ReqAddrs)(nil), "types.ReqAddrs")
	proto.RegisterType((*ReqDecodeRawTransaction)(nil), "types.ReqDecodeRawTransaction")
	proto.RegisterType((*ReqTokenTx)(nil), "types.ReqTokenTx")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 1210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0xdd, 0x8e, 0x13, 0x37,
	0x14, 0xd6, 0xcc, 0x24, 0xbb, 0xc9, 0x49, 0xa0, 0xec, 0x08, 0xc1, 0x08, 0x21, 0x9a, 0x5a, 0x54,
	0x42, 0x08, 0x05, 0x69, 0xc3, 0x65, 0xa5, 0x16, 0xd8, 0x0a, 0x10, 0x14, 0x5a, 0x93, 0xd2, 0xaa,
	0xaa, 0x2a, 0x39, 0x13, 0x27, 0x71, 0x77, 0x32, 0xce, 0xce, 0x38, 0xcb, 0xe4, 0x05, 0x7a, 0xd3,
	0xde, 0xf5, 0xba, 0xaf, 0xd0, 0x67, 0xe9, 0x23, 0x55, 0x3e, 0xb6, 0x67, 0x9c, 0x4d, 0xb6, 0xe2,
	0xa2, 0x12, 0x57, 0xeb, 0xef, 0xf8, 0xc4, 0xdf, 0xf9, 0xf9, 0x7c, 0xc6, 0x0b, 0x47, 0xaa, 0x60,
	0x79, 0xc9, 0x52, 0x25, 0x64, 0x3e, 0x5c, 0x15, 0x52, 0xc9, 0xb8, 0xad, 0x36, 0x2b, 0x5e, 0xde,
	0xea, 0xa7, 0x72, 0xb9, 0x74, 0x46, 0xf2, 0x0d, 0x5c, 0x79, 0x5c, 0x96, 0x5c, 0x95, 0xcf, 0x78,
	0xce, 0x4b, 0x51, 0xc6, 0x37, 0xe0, 0x80, 0x2d, 0xe5, 0x3a, 0x57, 0x49, 0x38, 0x08, 0xee, 0x45,
	0xd4, 0xa2, 0xf8, 0x2e, 0x5c, 0x29, 0xb8, 0x5a, 0x17, 0xf9, 0xe3, 0xe9, 0xb4, 0xe0, 0x65, 0x99,
	0x44, 0x83, 0xe0, 0x5e, 0x97, 0x6e, 0x1b, 0xc9, 0x1f, 0x01, 0x5c, 0x37, 0xe7, 0x8d, 0x35, 0xff,
	0x8c, 0x17, 0x63, 0xf9, 0x75, 0xc5, 0xd3, 0xf8, 0x36, 0x74, 0x53, 0x29, 0x72, 0x25, 0x4f, 0x79,
	0x9e, 0x04, 0xf8, 0xd3, 0xc6, 0x70, 0x29, 0x69, 0x0c, 0xad, 0x5c, 0x2a, 0x6e, 0xb9, 0x70, 0x1d,
	0xdf, 0x82, 0x0e, 0xaf, 0x78, 0xfa, 0x9a, 0x2d, 0x79, 0xd2, 0x42, 0x7b, 0x8d, 0xe3, 0xab, 0x10,
	0x2a, 0x99, 0xb4, 0xd1, 0x1a, 0x2a, 0x49, 0x7e, 0x0b, 0xe0, 0xaa, 0x09, 0xe7, 0x07, 0xa1, 0x16,
	0xd3, 0x82, 0xbd, 0xff, 0x48, 0x81, 0xfc, 0xea, 0xe2, 0x70, 0x65, 0xf9, 0x1f, 0xe3, 0x30, 0x5c,
	0xad, 0x9a, 0xeb, 0x9f, 0x00, 0x3a, 0x4f, 0x0b, 0xce, 0x14, 0x1f, 0x57, 0x76, 0x33, 0x70, 0x9b,
	0x97, 0x1e, 0x7c, 0x0d, 0xa2, 0x19, 0x37, 0xe7, 0x46, 0x54, 0x2f, 0x6b, 0xaa, 0x96, 0x47, 0x75,
	0x07, 0x40, 0xd4, 0xa5, 0xc4, 0xf4, 0x3a, 0xd4, 0xb3, 0xc4, 0x09, 0x1c, 0x8a, 0x72, 0x8c, 0x29,
	0x1d, 0xe0, 0xa6, 0x83, 0xf1, 0x00, 0x7a, 0x98, 0xd9, 0xdb, 0xcd, 0x72, 0x22, 0xb3, 0xe4, 0x10,
	0x0f, 0xf5, 0x4d, 0x5b, 0xe5, 0xec, 0x6c, 0x97, 0x93, 0xdc, 0x87, 0x1b, 0x36, 0xa3, 0x46, 0xd5,
	0xcf, 0x0a, 0xb9, 0x5e, 0xe9, 0xb8, 0x55, 0x55, 0x26, 0xc1, 0x20, 0xba, 0xd7, 0xa5, 0x7a, 0x49,
	0xee, 0x40, 0xe7, 0xfb, 0xbc, 0x14, 0xf3, 0x7c, 0x5c, 0xe9, 0x1c, 0xa6, 0x4c, 0x31, 0xcc, 0xbf,
	0x4f, 0x71, 0x4d, 0x24, 0xf4, 0x5e, 0xcb, 0x27, 0x2c, 0x63, 0x79, 0xaa, 0x0b, 0x74, 0x1d, 0xda,
	0xaa, 0x7a, 0xce, 0x2b, 0x5b, 0x23, 0x03, 0x74, 0x22, 0x2b, 0xb6, 0xd1, 0xaa, 0xc6, 0x3a, 0x75,
	0xa9, 0x83, 0xb8, 0x53, 0x88, 0xf3, 0x53, 0xbe, 0xb1, 0x4d, 0x70, 0x50, 0x97, 0x96, 0x57, 0x2b,
	0x51, 0xb8, 0x92, 0x59, 0x44, 0x7e, 0x81, 0xce, 0x5b, 0x31, 0xcf, 0xf9, 0x74, 0x5c, 0x69, 0x9f,
	0x35, 0x06, 0x67, 0x43, 0xb2, 0x48, 0x07, 0x8a, 0xd6, 0xd0, 0x04, 0x8a, 0xb6, 0x1b, 0x70, 0xb0,
	0x5a, 0x4f, 0x1c, 0x51, 0x9f, 0x5a, 0x84, 0x2d, 0xdd, 0x20, 0x47, 0x9b, 0x86, 0x6a, 0x43, 0x7e,
	0x0f, 0xa1, 0xe7, 0xd5, 0xc5, 0xc4, 0xc1, 0x53, 0x5e, 0x38, 0x0e, 0x83, 0x6c, 0x4e, 0x99, 0x64,
	0x53, 0x4b, 0xe3, 0x60, 0x3c, 0x84, 0xae, 0x66, 0x64, 0x6a, 0x5d, 0x18, 0x09, 0xf4, 0x8e, 0xaf,
	0x0d, 0x71, 0x5a, 0x0c, 0xdf, 0x3a, 0x3b, 0x6d, 0x5c, 0x9c, 0x58, 0x5a, 0x8d, 0x58, 0x9a, 0xdc,
	0xdb, 0x46, 0x56, 0x06, 0xe9, 0xea, 0xe6, 0x32, 0x4f, 0x39, 0xca, 0x21, 0xa2, 0x06, 0x58, 0x51,
	0x1e, 0xd6, 0xa2, 0xbc, 0x03, 0x30, 0xd7, 0xdd, 0x7c, 0x8a, 0xc2, 0xec, 0x60, 0x66, 0x9e, 0x45,
	0x9f, 0xbe, 0xe0, 0x6c, 0xca, 0x8b, 0xa4, 0x6b, 0x32, 0x32, 0x08, 0x25, 0xca, 0x2b, 0x95, 0x80,
	0xa9, 0x9a, 0x5e, 0x93, 0x47, 0xd0, 0xf7, 0x8a, 0x51, 0xc6, 0x77, 0x1b, 0x81, 0xf4, 0x8e, 0x63,
	0x9b, 0x95, 0xe7, 0x61, 0x44, 0xf3, 0x25, 0x5c, 0xa1, 0x22, 0x9f, 0xd7, 0xd9, 0xc6, 0x43, 0x68,
	0x0b, 0xc5, 0x97, 0xee, 0x87, 0x89, 0xfd, 0xe1, 0x96, 0xd3, 0x0b, 0xc5, 0x97, 0xd4, 0xb8, 0x91,
	0x17, 0x70, 0xb4, 0xb3, 0xe7, 0x75, 0x50, 0x9f, 0xd2, 0x74, 0xf0, 0xb6, 0x5f, 0xef, 0x10, 0xb7,
	0x1a, 0x03, 0xf9, 0x0e, 0xba, 0x4d, 0x1c, 0xa6, 0xd9, 0x81, 0x6b, 0xb6, 0x77, 0x64, 0xb8, 0x25,
	0x8a, 0xdb, 0x17, 0x5b, 0xb8, 0x75, 0xe4, 0xcf, 0xd0, 0xd7, 0xe2, 0x7d, 0x73, 0xce, 0x8b, 0x73,
	0xc1, 0xf1, 0x9e, 0x16, 0x3c, 0x15, 0xe7, 0x56, 0x23, 0x11, 0x75, 0x50, 0xef, 0x4c, 0xcc, 0xdd,
	0xb0, 0x03, 0xc2, 0x41, 0xbd, 0xa3, 0x2a, 0xd3, 0x21, 0x33, 0x25, 0x1c, 0x24, 0x7f, 0x06, 0x70,
	0x48, 0xf9, 0x19, 0x5e, 0x8f, 0x18, 0x5a, 0x4c, 0xdf, 0x1a, 0x73, 0x9b, 0x70, 0xad, 0x6d, 0xb3,
	0x8c, 0xcd, 0xf1, 0xc0, 0x36, 0xc5, 0xb5, 0x16, 0x46, 0x5a, 0x9f, 0xd5, 0xa6, 0x06, 0xe8, 0x2c,
	0xa6, 0xa2, 0xe0, 0xd8, 0x18, 0xab, 0xf0, 0xc6, 0x60, 0x64, 0x20, 0xe6, 0x0b, 0xe5, 0x44, 0x66,
	0x90, 0x3e, 0x4b, 0xe4, 0x53, 0x5e, 0x39, 0x91, 0x21, 0x20, 0x3f, 0x02, 0x50, 0x7e, 0xf6, 0x6d,
	0x21, 0xce, 0x59, 0xba, 0x69, 0xf8, 0x82, 0x4b, 0xf9, 0xc2, 0xcb, 0xf9, 0x22, 0x9f, 0x8f, 0xdc,
	0x84, 0xf6, 0x73, 0x5e, 0xd9, 0xe1, 0x5a, 0xd5, 0xc3, 0xb5, 0x22, 0x6f, 0xa0, 0x47, 0xf9, 0x2a,
	0xdb, 0x8c, 0xab, 0x17, 0xf9, 0x4c, 0xea, 0xbc, 0x17, 0xac, 0x5c, 0xb8, 0xe9, 0xa3, 0xd7, 0xde,
	0x99, 0xe1, 0xfe, 0x1c, 0x22, 0x3f, 0x87, 0xcf, 0xa0, 0x4b, 0xf9, 0xd9, 0xb8, 0x7a, 0x25, 0x4a,
	0xb5, 0x9d, 0x42, 0x64, 0x53, 0x20, 0xa3, 0x9a, 0x13, 0x9d, 0x3e, 0x4c, 0xee, 0x14, 0x60, 0x5c,
	0x3d, 0x67, 0xe5, 0x02, 0x7f, 0xa3, 0x63, 0x62, 0xe5, 0x82, 0x97, 0x4e, 0xa6, 0x06, 0x35, 0x84,
	0xa1, 0x47, 0xe8, 0x5d, 0xf5, 0x68, 0x10, 0x35, 0x57, 0x9d, 0x7c, 0x01, 0x7d, 0x2f, 0xf9, 0x32,
	0x7e, 0xa0, 0xf5, 0x82, 0xcb, 0x0b, 0xd1, 0x78, 0x5e, 0xd4, 0xb9, 0x90, 0xa1, 0xee, 0x56, 0xca,
	0xc5, 0x4a, 0xbd, 0x92, 0xf3, 0x1d, 0xd5, 0x5f, 0x83, 0x28, 0x93, 0x73, 0x2b, 0x79, 0xbd, 0x24,
	0x4c, 0x4b, 0x0e, 0xfd, 0x77, 0x9c, 0x3f, 0x85, 0xf0, 0xe5, 0x3b, 0xbc, 0x56, 0xbd, 0xe3, 0x4f,
	0x2c, 0xe7, 0x4b, 0xbe, 0x79, 0xc7, 0xb2, 0x35, 0xa7, 0xe1, 0xcb, 0x77, 0xf1, 0xe7, 0xd0, 0xca,
	0xe4, 0xbc, 0xc4, 0xf8, 0x7b, 0xc7, 0x47, 0x75, 0x58, 0x8e, 0x9e, 0xe2, 0x36, 0x39, 0xd1, 0x95,
	0x45, 0xdb, 0x09, 0x53, 0x6c, 0x87, 0xe6, 0x03, 0x4f, 0xd1, 0x5f, 0xe3, 0x71, 0x45, 0x79, 0xb9,
	0xce, 0x94, 0xd7, 0xfd, 0x60, 0x7f, 0xf7, 0x8d, 0x06, 0x0d, 0x88, 0x09, 0xca, 0xcb, 0xcc, 0xe3,
	0x7d, 0xad, 0x0c, 0x55, 0x15, 0x3f, 0x82, 0x5e, 0x61, 0x28, 0xa7, 0xcc, 0x7e, 0xac, 0xfd, 0x4a,
	0xd7, 0xe1, 0x53, 0xdf, 0x4d, 0xeb, 0x7e, 0x92, 0xc9, 0xf4, 0x54, 0x89, 0xa5, 0x9b, 0xd8, 0x8d,
	0x41, 0x8f, 0x63, 0xc3, 0x80, 0xdf, 0xe2, 0x03, 0x94, 0xb7, 0x67, 0x21, 0x7f, 0x85, 0x70, 0xe4,
	0xc5, 0x71, 0xc2, 0x15, 0x13, 0x99, 0x8d, 0x36, 0xf8, 0xcf, 0x68, 0x1f, 0xe0, 0xdc, 0xd1, 0x61,
	0x60, 0xa6, 0xfb, 0x23, 0x75, 0x2e, 0x38, 0xeb, 0x0a, 0x29, 0x67, 0xa6, 0xc6, 0x7a, 0xd6, 0x21,
	0xf2, 0xaa, 0xd8, 0xda, 0x5f, 0xc5, 0xb6, 0x77, 0x87, 0xb6, 0x73, 0x3d, 0xb8, 0x98, 0x6b, 0xf3,
	0x1e, 0x3a, 0xdc, 0x7a, 0x0f, 0xdd, 0x82, 0xce, 0xac, 0x90, 0x4b, 0x9c, 0x65, 0xf6, 0x35, 0xe2,
	0xf0, 0x85, 0xfa, 0x74, 0x77, 0xea, 0xf3, 0x15, 0xc4, 0x3b, 0xe5, 0x29, 0xe3, 0xfb, 0xfe, 0xcd,
	0x4c, 0x76, 0x0b, 0x64, 0xfc, 0xcc, 0xfd, 0x1c, 0x40, 0xc7, 0x0e, 0x54, 0xbc, 0x85, 0x9a, 0xd5,
	0xbd, 0x71, 0x0c, 0x20, 0x0f, 0xe1, 0x26, 0xe5, 0x67, 0x27, 0x3c, 0x95, 0x53, 0x4e, 0xd9, 0x7b,
	0xff, 0xfb, 0xbf, 0xf7, 0x45, 0x43, 0xfe, 0x0e, 0x70, 0x1e, 0xe2, 0x6b, 0xcc, 0x3c, 0x44, 0x4a,
	0xf3, 0x14, 0x33, 0x5e, 0x16, 0x7d, 0x9c, 0x59, 0x5d, 0x7f, 0x35, 0x0e, 0x9b, 0xaf, 0xc6, 0x93,
	0xbb, 0x3f, 0x91, 0xb9, 0x50, 0x19, 0x9b, 0x0c, 0x47, 0xa3, 0x61, 0x9a, 0x3f, 0x4c, 0x17, 0x4c,
	0xe4, 0xa3, 0x51, 0xfd, 0x17, 0x0b, 0x38, 0x39, 0xc0, 0x7f, 0x63, 0x46, 0xff, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x9e, 0x95, 0x84, 0xfb, 0xf0, 0x0c, 0x00, 0x00,
}

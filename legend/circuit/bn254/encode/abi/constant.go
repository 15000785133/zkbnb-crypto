package abi

import (
	"encoding/hex"
	"math/big"
)

type AbiId uint64

func (abiId AbiId) ToBigIntRegular(res *big.Int) *big.Int {
	return new(big.Int).SetUint64(uint64(abiId))
}

const (
	Transfer         = "Transfer"
	Withdraw         = "Withdraw"
	AddLiquidity     = "AddLiquidity"
	RemoveLiquidity  = "RemoveLiquidity"
	Swap             = "Swap"
	CreateCollection = "CreateCollection"
	MintNft          = "MintNft"
	TransferNft      = "TransferNft"
	WithdrawNft      = "WithdrawNft"
	CancelOffer      = "CancelOffer"
	AtomicMatch      = "AtomicMatch"
)

// HexEIP712MessageTypeHash Metamask implemented eip712 as while encodeData,
// the type(Transfer e.g.) defined in the sign request
// will first be called by join of its subTypes then the join string be put in keccak,
// then the keccak results of it will be used as the first part (type bytes32) of abi.encode,
// so the hash constants below is the type related keccak results, and will be used as the first value in our encodes.
// here we record them as static value in order to reduce usage of dynamic keccak.
// Be cared: once the type properties change, the related hash will change accordingly.
var HexEIP712MessageTypeHash map[string]string

// GetEIP712MessageTypeHashBytes32 map the name of transaction type to type hash results,
// as first values argument passed for abi.encode
func GetEIP712MessageTypeHashBytes32(name string) [32]byte {
	messageTypeBytes, err := hex.DecodeString(HexEIP712MessageTypeHash[name])
	if err != nil {
		// the string should always be decoded.
		panic(err)
	}
	var messageTypeBytes32 [32]byte
	copy(messageTypeBytes32[:], messageTypeBytes)
	return messageTypeBytes32
}

func init() {
	HexEIP712MessageTypeHash = make(map[string]string)
	HexEIP712MessageTypeHash[Transfer] = "96695797a85b65c62a1eb8e28852fc7d5a34b668e127752d9a132d6d5e2d3717"
	HexEIP712MessageTypeHash[Withdraw] = "09ed82aece8178a11d49bfa4432352b5f19495bc410e77e6411666015bca2d5a"
	HexEIP712MessageTypeHash[AddLiquidity] = "1d0c49cf49fe0e24dac1467598ed87c0f11e2e6297f0047878db9d14bc0f4666"
	HexEIP712MessageTypeHash[RemoveLiquidity] = "f8756b5312d5c28d3105707462dc74e73e0f69e885b2f74c0673d502e93907c4"
	HexEIP712MessageTypeHash[Swap] = "8c8e829f99313ec0b7163eee2c779f3b643d6d9a842d595385632a461e78176f"
	HexEIP712MessageTypeHash[CreateCollection] = "0d4c4b6153bea091b919a39d7915a2155673c15089b7bb09cedff8109f478009"
	HexEIP712MessageTypeHash[MintNft] = "cacc7ff451f0d4db55b4b995712f56456bed3dcf1c99450a83822e2c10867892"
	HexEIP712MessageTypeHash[TransferNft] = "a44319586642f197aa0151a82e6390a26ecf7a045202885c76acc9d993ebdc07"
	HexEIP712MessageTypeHash[WithdrawNft] = "c1cfe1ac4b0b18691abc9f738129b173f7ab0cf80eea90a78d52399a0fb9dd6f"
	HexEIP712MessageTypeHash[CancelOffer] = "2afdc1852648f0134e5e99ce96465f52ea04a0747b3c143f8d003cfac71b01a0"
	HexEIP712MessageTypeHash[AtomicMatch] = "c40d8c33405fddd01ba170c2d918f34e298d9c16b2718078fddb3520431e2ee1"
}

// HexPrefixAndEip712DomainKeccakHash Metamask implemented eip712 as while keccak,
// it will not only keccak the hashStruct of message, but also a prefix and the hashStruct of eip712Domain
// the hash is keccak(b'1901' + hashStruct(eip712Domain) + hashStruct(message))
// here we pretend the prefix and hash eip712Domain to be static for the sake of reduce usage of keccak.
var HexPrefixAndEip712DomainKeccakHash = "1901c8404fcb774064f8f6661c4fcaef7ed215866d423188610ece3533d068413a06"

// PairingABI is the input ABI used to generate the binding from.
const PairingABI = "[]"
const GeneralABIJSON = "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"FromAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"PairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint40\",\"name\":\"packedAAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"packedBAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"AddLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"AccountIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"OfferType\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"OfferId\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"AccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"NftIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"packedAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint64\",\"name\":\"OfferListedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"OfferExpiredAt\",\"type\":\"uint64\"},{\"internalType\":\"bytes16\",\"name\":\"SigRx\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"SigRy\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32\",\"name\":\"SigS\",\"type\":\"bytes32\"}],\"internalType\":\"struct Storage.Offer\",\"name\":\"BuyerOffer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"OfferType\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"OfferId\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"AccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"NftIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"packedAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint64\",\"name\":\"OfferListedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"OfferExpiredAt\",\"type\":\"uint64\"},{\"internalType\":\"bytes16\",\"name\":\"SigRx\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"SigRy\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32\",\"name\":\"SigS\",\"type\":\"bytes32\"}],\"internalType\":\"struct Storage.Offer\",\"name\":\"SellerOffer\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"AtomicMatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"AccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"OfferId\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"CancelOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"AccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"CreateCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"CreatorAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ToAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"ToAccountNameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"NftContentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"CreatorTreasureRate\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"NftCollectionId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"MintNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"FromAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"PairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint40\",\"name\":\"packedAAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"packedBAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"lpAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"RemoveLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"FromAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"PairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint40\",\"name\":\"packedAAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"packedBAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"Swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EIP712MessageType\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"FromAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ToAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"ToAccountNameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"AssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint40\",\"name\":\"packedAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"CallDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"FromAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ToAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"ToAccountNameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"NftIndex\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"CallDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"TransferNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"FromAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"AssetId\",\"type\":\"uint16\"},{\"internalType\":\"bytes16\",\"name\":\"AssetAmount\",\"type\":\"bytes16\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"bytes20\",\"name\":\"ToAddress\",\"type\":\"bytes20\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"AccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"NftIndex\",\"type\":\"uint40\"},{\"internalType\":\"bytes20\",\"name\":\"ToAddress\",\"type\":\"bytes20\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"WithdrawNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
const TransferABIJSON = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EIP712MessageType\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"FromAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ToAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"ToAccountNameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"AssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint40\",\"name\":\"packedAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"CallDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
const WithdrawABIJSON = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EIP712MessageType\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"FromAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"AssetId\",\"type\":\"uint16\"},{\"internalType\":\"bytes16\",\"name\":\"AssetAmount\",\"type\":\"bytes16\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"bytes20\",\"name\":\"ToAddress\",\"type\":\"bytes20\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
const AddLiquidityABIJSON = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EIP712MessageType\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"FromAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"PairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint40\",\"name\":\"packedAAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"packedBAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
const RemoveLiquidityABIJSON = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EIP712MessageType\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"FromAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"PairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint40\",\"name\":\"packedAAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"packedBAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"lpAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
const SwapABIJSON = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EIP712MessageType\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"FromAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"PairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint40\",\"name\":\"packedAAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"packedBAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
const CreateCollectionABIJSON = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EIP712MessageType\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"AccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
const WithdrawNftABIJSON = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EIP712MessageType\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"AccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"NftIndex\",\"type\":\"uint40\"},{\"internalType\":\"bytes20\",\"name\":\"ToAddress\",\"type\":\"bytes20\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
const TransferNftABIJSON = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EIP712MessageType\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"FromAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ToAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"ToAccountNameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"NftIndex\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"CallDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
const MintNftABIJSON = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EIP712MessageType\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"CreatorAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ToAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"ToAccountNameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"NftContentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"CreatorTreasureRate\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"NftCollectionId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
const CancelOfferABIJSON = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EIP712MessageType\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"AccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"OfferId\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"GasAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"GasFeeAssetId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"packedFee\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"ExpireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"ChainId\",\"type\":\"uint32\"}],\"name\":\"\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
const AtomicMatchABIJSON = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EIP712MessageType\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sellerAccountIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellerNftIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellerOfferId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellerType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellerAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellerListedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellerExpiredAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellerTreasureRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sellerSigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sellerSigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"buyerAccountIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyerNftIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyerOfferId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyerType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyerAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyerListedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyerExpiredAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyerTreasureRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"buyerSigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"buyerSigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"Nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ChainId\",\"type\":\"uint256\"}],\"name\":\"\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
const ABI_ENCODE_EMPTY_BYTE = 0xffff

const (
	DefaultAbi          AbiId = 0
	TransferAbi               = 6
	WithdrawAbi               = 10
	AddLiquidityAbi           = 8
	RemoveLiquidityAbi        = 9
	SwapAbi                   = 7
	CreateCollectionAbi       = 11
	WithdrawNftAbi            = 16
	TransferNftAbi            = 13
	MintNftAbi                = 12
	AtomicMatchAbi            = 14
	CancelOfferAbi            = 15
)

const StaticArgsOutput = 1024

func init() {
}

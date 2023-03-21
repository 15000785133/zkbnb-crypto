package txtypes

import "fmt"

var (
	ErrAccountIndexTooLow       = fmt.Errorf("AccountIndex should not be less than %d", minAccountIndex)
	ErrAccountIndexTooHigh      = fmt.Errorf("AccountIndex should not be larger than %d", maxAccountIndex)
	ErrGasAccountIndexTooLow    = fmt.Errorf("GasAccountIndex should not be less than %d", minAccountIndex)
	ErrGasAccountIndexTooHigh   = fmt.Errorf("GasAccountIndex should not be larger than %d", maxAccountIndex)
	ErrGasFeeAssetIdTooLow      = fmt.Errorf("GasFeeAssetId should not be less than %d", minAssetId)
	ErrGasFeeAssetIdTooHigh     = fmt.Errorf("GasFeeAssetId should not be larger than %d", maxAssetId)
	ErrGasFeeAssetAmountTooLow  = fmt.Errorf("GasFeeAssetAmount should not be less than %s", minPackedFeeAmount.String())
	ErrGasFeeAssetAmountTooHigh = fmt.Errorf("GasFeeAssetAmount should not be larger than %s", maxPackedFeeAmount.String())
	ErrNonceTooLow              = fmt.Errorf("Nonce should not be less than %d", minNonce)
	ErrOfferTypeInvalid         = fmt.Errorf("Type should only be buy(%d) and sell(%d)", BuyOfferType, SellOfferType)
	ErrOfferIdTooLow            = fmt.Errorf("OfferId should not be less than 0")
	ErrNftIndexTooLow           = fmt.Errorf("NftIndex should not be less than %d", minNftIndex)
	ErrNftIndexTooHigh          = fmt.Errorf("NftIndex should not be larger than %d", maxNftIndex)
	ErrAssetIdTooLow            = fmt.Errorf("AssetId should not be less than %d", minAssetId)
	ErrAssetIdTooHigh           = fmt.Errorf("AssetId should not be larger than %d", maxAssetId)
	ErrAssetAmountTooLow        = fmt.Errorf("AssetAmount should be larger than %s", minAssetAmount.String())
	ErrAssetAmountTooHigh       = fmt.Errorf("AssetAmount should not be larger than %s", maxAssetAmount.String())
	ErrPlatformFeeTooLow        = fmt.Errorf("PlatformFee should be larger than %s", minAssetAmount.String())
	ErrPlatformFeeTooHigh       = fmt.Errorf("PlatformFee should not be larger than %s", maxAssetAmount.String())
	ErrListedAtTooLow           = fmt.Errorf("ListedAt should be larger than 0")
	ErrChanelRateTooLow         = fmt.Errorf("ChanelRate should  not be less than %d", minRate)
	ErrChanelRateTooHigh        = fmt.Errorf("ChanelRate should  not be larger than %d", maxRate)
	ErrPlatformFeeRateTooLow    = fmt.Errorf("PlatformFeeRate should  not be less than %d", minRate)
	ErrPlatformFeeRateTooHigh   = fmt.Errorf("PlatformFeeRate should  not be larger than %d", maxRate)
	ErrCollectionNameTooShort   = fmt.Errorf("length of Name should not be less than %d", minCollectionNameLength)
	ErrCollectionNameTooLong    = fmt.Errorf("length of Name should not be larger than %d", maxCollectionNameLength)
	ErrIntroductionTooLong      = fmt.Errorf("length of Introduction should not be larger than %d", maxCollectionIntroductionLength)
	ErrNftContentHashInvalid    = fmt.Errorf("NftContentHash is invalid")
	ErrNftCollectionIdTooLow    = fmt.Errorf("NftCollectionId should not be less than %d", minCollectionId)
	ErrNftCollectionIdTooHigh   = fmt.Errorf("NftCollectionId should not be larger than %d", maxCollectionId)
	ErrCallDataHashInvalid      = fmt.Errorf("CallDataHash is invalid")
	ErrPubKeyXYInvalid          = fmt.Errorf("PubKeyX or PubKeyY is invalid")

	ErrCreatorAccountIndexTooLow  = fmt.Errorf("CreatorAccountIndex should not be less than %d", minAccountIndex)
	ErrCreatorAccountIndexTooHigh = fmt.Errorf("CreatorAccountIndex should not be larger than %d", maxAccountIndex)
	ErrToAccountIndexTooLow       = fmt.Errorf("ToAccountIndex should not be less than %d", minAccountIndex)
	ErrToAccountIndexTooHigh      = fmt.Errorf("ToAccountIndex should not be larger than %d", maxAccountIndex)
	ErrToL1AddressInvalid         = fmt.Errorf("ToL1Address is invalid")
	ErrCreatorTreasuryRateTooLow  = fmt.Errorf("CreatorTreasuryRate should  not be less than %d", minRate)
	ErrCreatorTreasuryRateTooHigh = fmt.Errorf("CreatorTreasuryRate should not be larger than %d", maxRate)
	ErrFromAccountIndexTooLow     = fmt.Errorf("FromAccountIndex should not be less than %d", minAccountIndex)
	ErrFromAccountIndexTooHigh    = fmt.Errorf("FromAccountIndex should not be larger than %d", maxAccountIndex)
	ErrToAddressInvalid           = fmt.Errorf("ToAddress is invalid")
	ErrBuyOfferInvalid            = fmt.Errorf("BuyOffer is invalid")
	ErrSellOfferInvalid           = fmt.Errorf("SellOffer is invalid")
)

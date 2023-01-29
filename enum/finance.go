package enum

const (
	GuilinBankIntrabankTransferServiceCode                  = "CBE020100"
	GuilinBankOutofbankTransferServiceCode                  = "CBE020102"
	GuilinBankTransactionDetailServiceCode                  = "CBE010201"
	GuilinBankTransactionDetailElectronicReceiptServiceCode = "CBE070123"
	GuilinBankTransferResultServiceCode                     = "CBE020105"
	GuilinBankQueryAccountBalanceServiceCode                = "CBE010101"
	GuilinBankBatchTransferServiceCode                      = "CBE020110"
	GuilinBankBatchTransferDetailServiceCode                = "CBE020111"
	GuilinBankBatchTransferResultServiceCode                = "CBE020112"

	GuilinBankIntrabankTransferBusinessCode    = "020100"
	GuilinBankOutofbankTransferBusinessCode    = "020102"
	GuilinBankTransferRequestType              = "0"
	GuilinBankCollectionNoticeType             = "1"
	GuilinBankToPublicFlag                     = "对公"
	GuilinBankToPrivateFlag                    = "对私"
	GuilinBankTransactionDetailCashFlag        = "现钞"
	GuilinBankTransactionDetailTransferFlag    = "现汇"
	GuilinBankTransferSmallAndLargeChannelFlag = "大小额渠道"
	GuilinBankTransferSupernetChannelFlag      = "超网渠道"
	GuilinBankTransferCommonRmtType            = "普通"
	GuilinBankTransferRealtimeRmtType          = "实时"

	DefaultRecommend = "银企直联转账"
	CurrencyTypeCNY  = "CNY"

	GuilinBankSuccessRetCode = "000000"

	GuilinBankTransferRejectResult  = "71"
	GuilinBankTransferCreditResult  = "75"
	GuilinBankTransferRevokeResult  = "91"
	GuilinBankTransferDeleteResult  = "92"
	GuilinBankTransferSuccessResult = "90"
	GuilinBankTransferFailResult    = "99"
	GuilinBankTransferRefuseResult  = "101"

	GuilinBankTransactionDetailPayType = "pay"
	GuilinBankTransactionDetailRecType = "rec"
)

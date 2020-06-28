package action

/**
@ app 结构体
@ SDK结构体
@
0　assetList  基础资产ID
1　　currentAssetID 　要获取哪个资产的子集
2　　updateTime　更新操作时间
*/

type Cdatagcae struct {
	AssetID        string `json:"assetId"`
	CurrentAssetID string `json:"currentAssetId"`
	UpdateTime     string `json:"updateTime"`

	OrgName       string `json:"orgName"`
	OrgID         string `json:"orgId"`
	Peer          string `json:"peer"`
	Anchor        string `json:"anchor"`
	UserName      string `json:"userName"`
	UserID        string `json:"userId"`
	AffiliationId string `json:"affiliationId"`
}

/*
	assetID 基础资产ID
	updateTime time
	assetType
	assetKey
*/
// 资产基本字段
type BaseAsset struct {
	// IsDeleted  bool   `json:"isDeleted"`
	UpdateDate string `json:"updateDate"`
	// 仅用在底层资产中、其Parent有可能为产品、保理
	ParentType string `json:"parentType"`
	ParentID   string `json:"parentID"`
}

// 发票信息
type Invoice struct {
	BaseAsset
	InvoiceType          string `json:"invoiceType"`
	InvoiceCode          string `json:"invoiceCode"`
	InvoiceOperationType string `json:"InvoiceOperationType"`
	// 唯一标识
	InvoiceNum            string `json:"invoiceNum" pkey:""`
	IssueDate             string `json:"issueDate"`
	TotalAmount           string `json:"totalAmount"`
	PreTaxAmount          string `json:"preTaxAmount"`
	AmountUnit            string `json:"amountUnit"`
	CheckCodes            string `json:"checkCodes"`
	BuyerTaxNum           string `json:"buyerTaxNum"`
	SellerTaxNum          string `json:"sellerTaxNum"`
	Verified              string `json:"verified"`
	InvoiceAttachmentList []struct {
		IpfsHash          string `json:"ipfsHash"`
		InvoiceAttachName string `json:"invoiceAttachName"`
		InvoiceAttachType string `json:"invoiceAttachType"`
		InvoiceAttachNo   string `json:"invoiceAttachNo"`
		InvoiceAttachMD   string `json:"invoiceAttachMD"`
		InvoiceAttachAddr string `json:"invoiceAttachAddr"`
	} `json:"invoiceAttachmentList"`
}

// 合同信息
type TradingContr struct {
	BaseAsset
	// 唯一标识
	TradingContrNum              string `json:"tradingContrNum" pkey:""`
	TradingContrOperationType    string `json:"tradingContrOperationType"`
	TradingContrName             string `json:"tradingContrName"`
	TradingContrAmount           string `json:"tradingContrAmount"`
	TradingContrAmountUnit       string `json:"tradingContrAmountUnit"`
	TradingContrCurrency         string `json:"tradingContrCurrency"`
	TradingContrPayerCertificate string `json:"tradingContrPayerCertificate"`
	TradingContrPayerName        string `json:"tradingContrPayerName"`
	TradingContrPayeeCertificate string `json:"tradingContrPayeeCertificate"`
	TradingContrPayeeName        string `json:"tradingContrPayeeName"`
	TradingContrSigningDate      string `json:"tradingContrSigningDate"`
	TradingContrFrequency        string `json:"tradingContrFrequency"`
	TradingContrAccountPeriod    string `json:"tradingContrAccountPeriod"`
	TradingContrAttachList       []struct {
		TradingContrAttachName string `json:"tradingContrAttachName"`
		TradingContrAttachTyp  string `json:"tradingContrAttachTyp"`
		TradingContrAttachNo   string `json:"tradingContrAttachNo"`
		TradingContrAttachMD   string `json:"tradingContrAttachMD"`
		TradingContrAttachAddr string `json:"tradingContrAttachAddr"`
		IpfsHash               string `json:"ipfsHash"`
	} `json:"tradingContrAttachList"`
}

// 其他附件
type ExtraInfoAttach struct {
	BaseAsset           `json:"base"`
	IpfsHash            string `json:"ipfsHash"`
	ExtraInfoFileName   string `json:"extraInfoFileName"`
	ExtraInfoAttachType string `json:"extraInfoAttachType"`
	// 唯一标识
	ExtraInfoAttachNo   string `json:"extraInfoAttachNo" pkey:""`
	ExtraInfoAttachMd   string `json:"extraInfoAttachMd"`
	ExtraInfoAttachAddr string `json:"extraInfoAttachAddr"`
}

//
type User struct {
	//
	Datatime      string `json:"dataTime"`
	UUID          string `json:"uuid"`
	Peer          string `json:"peer"`
	Anchor        string `json:"anchor"`
	OrgName       string `json:"orgName"`
	OrgID         string `json:"orgId"`
	UserName      string `json:"userName"`
	UserID        string `json:"userId"`
	AffiliationId string `json:"affiliationId"`
	BycByKey      string `json:"bycByKey"`
	BycByfMaster  string `json:"bycByfMaster"`
	Span          string `json:"span"`
}

type CDdata struct {
	Datastate string `json:"dataState"`
	DataTime  string `json:"dataTime"`
	User      User
}

type CDdata1 struct {
	Datastate string `json:"datastate"`
	Parmas    string `json:"parmas"`
	Datatime  string `json:"datatime"`
}

//底层资产
type ArAsset struct {
	BaseAsset
	// 唯一标识
	ArNum            string   `json:"arNum" pkey:""`
	ArPayer          string   `json:"arPayer"`
	ArPayee          string   `json:"arPayee"`
	ArTotalAmount    string   `json:"arTotalAmount"`
	ArNet            string   `json:"arNet"`
	ArDueDate        string   `json:"arDueDate"`
	AssetStatus      string   `json:"assetStatus"`
	AssetAttribute   string   `json:"assetAttribute"`
	TradingContrList []string `json:"tradeContractList"`
	InvoiceAmount    string   `json:"invoiceAmount"`
	InvoiceList      []string `json:"invoiceList"`
	ExtraInfoList    []string `json:"extraInfoList"`
}

// 保理信息
type Factoring struct {
	BaseAsset
	// SendID             string `json:"sendId"`
	// SenderName         string `json:"senderName"`
	// UserName           string `json:"userName"`
	// Token              string `json:"token"`
	// TargetName         string `json:"targetName"`
	FOperationType string `json:"fOperationType"`
	// FStatus            string `json:"fStatus"`
	FAttribute       string `json:"fAttribute"`
	FirstHandAsset   string `json:"firstHandAsset"`
	FPossessor       string `json:"fPossessor"`
	FPossessorTaxNum string `json:"fPossessorTaxNum"`
	DisclosedF       string `json:"disclosedF"`
	RecourseF        string `json:"recourseF"`
	FDuration        string `json:"fDuration"`
	ArTotalAmount    string `json:"arTotalAmount"`
	// 唯一标识
	FContrNum          string `json:"fContrNum" pkey:""`
	FContrName         string `json:"fContrName"`
	FContrAttachAmount string `json:"fContrAttachAmount"`
	FContrAttachList   []struct {
		IpfsHash         string `json:"ipfsHash"`
		FContrAttachName string `json:"fContrAttachName"`
		FContrAttachType string `json:"fContrAttachType"`
		FContrAttachNo   string `json:"fContrAttachNo"`
		FContrAttachMD   string `json:"fContrAttachMD"`
		FContrAttachAddr string `json:"fContrAttachAddr"`
	} `json:"fContrAttachList"`
	FContrTotalAmount            string   `json:"fContrTotalAmount"`
	FinancingPrincipal           string   `json:"financingPrincipal"`
	FinancingInterestTotalAmount string   `json:"financingInterestTotalAmount"`
	FinancingInterestRate        string   `json:"financingInterestRate"`
	FinancingDueDate             string   `json:"financingDueDate"`
	InterestPaymentFrequency     string   `json:"interestPaymentFrequency"`
	ManagementFee                string   `json:"managementFee"`
	FinancingDuration            string   `json:"financingDuration"`
	FFinancingRatio              string   `json:"fFinancingRatio"`
	ArAssetAmount                string   `json:"arAssetAmount"`
	ArAssetList                  []string `json:"arAssetList"`
}

//　产品信息
type Product struct {
	BaseAsset
	SendID     string `json:"sendId"`
	SenderName string `json:"senderName"`
	UserName   string `json:"userName"`
	// TransUUID                string   `json:"transUuid"`
	// MsgUUID                  string   `json:"msgUuid"`
	/** 唯一值使用后台生成的的编号 */
	ProNo     string `json:"proNo" pkey:""`
	PreRegNo  string `json:"preRegNo"`
	RegNo     string `json:"regNo"`
	PreProdNo string `json:"preProdNo"`
	ProdNo    string `json:"prodNo"`
	Status    string `json:"status"`
	AssetUUID string `json:"assetUuid"`

	// 0：保理资产
	// 1：底层资产
	AssetType string `json:"assetType"`

	AssetAmount              string   `json:"assetAmount"`
	AssetScope               string   `json:"assetScope"`
	ProductUUID              string   `json:"productUuid"`
	ProductName              string   `json:"productName"`
	Financier                string   `json:"financier"`
	TotalAmount              string   `json:"totalAmount"`
	PriorAmount              string   `json:"priorAmount"`
	PosteriorAmount          string   `json:"posteriorAmount"`
	Unit                     string   `json:"unit"`
	ProductDuration          string   `json:"productDuration"`
	ProductDenomination      string   `json:"productDenomination"`
	ListingMethod            string   `json:"listingMethod"`
	ListingDate              string   `json:"listingDate"`
	DatedDate                string   `json:"datedDate"`
	ProductInterestRate      string   `json:"productInterestRate"`
	DebtServiceMethod        string   `json:"debtServiceMethod"`
	InterestPaymentFrequency string   `json:"interestPaymentFrequency"`
	RedemptionPrice          string   `json:"redemptionPrice"`
	RedemptionDate           string   `json:"redemptionDate"`
	CallProvision            string   `json:"callProvision"`
	CreditRatingResult       string   `json:"creditRatingResult"`
	CollectionObject         string   `json:"collectionObject"`
	ListingPrice             string   `json:"listingPrice"`
	TrustPlanName            string   `json:"trustPlanName"`
	TrustPlanNum             string   `json:"trustPlanNum"`
	ProductCategory          string   `json:"productCategory"`
	BasicAssetPen            string   `json:"basicAssetPen"`
	CreditorNumber           string   `json:"creditorNumber"`
	ObligorNumber            string   `json:"obligorNumber"`
	Originator               string   `json:"originator"`
	IssuingCarrier           string   `json:"issuingCarrier"`
	RatingAgency             string   `json:"ratingAgency"`
	AccountingFirm           string   `json:"accountingFirm"`
	LawFirm                  string   `json:"lawFirm"`
	MainUnderwriter          string   `json:"mainUnderwriter"`
	AssetList                []string `json:"assetList"`
}

// ------------------------历史追溯 子资产的结构体-----------------------------

//历史追溯、针对产品的子资产　 (有保理 无保理)
//type SubAssetsResult struct {
//	FactoringTxID       string         `json:"factoringTxID"`
//	Factoring           Factoring      `json:"factoring"`
//	FactoringIsModified bool           `json:"factoringIsModified"`
//	ArAssets            []ArAsset      `json:"arAssets"`
//	ArAsset             ArAsset        `json:"arAsset"`
//	ArAssetIsModified   bool           `json:"arAssetIsModified"`
//	Invoice             []Invoice      `json:"invoices"`
//	TradingContr        []TradingContr `json:"tradingContrs"`
//}
type ChaincodegetforArAsset struct {
	AssetID        string          `json:"assetID"`
	SubAssetResult SubAssetsResult `json:"subAssetResult"`
	UpdateTime     string          `json:"updateTime"`
}

//子资产的key以及txID信息
type SubAssetKeyAndTxIDResult struct {
	TxID string `json:"txID"`
	Key  string `json:"key"`
}

const (
	// 发票
	assetType_INVOICE string = "1"
	// 合同
	assetType_CONTRACT string = "2"
	// 其他附件
	assetType_ATTACHMENT string = "3"
	// 底层资产
	assetType_UNDERLYING string = "4"
	// 保理资产
	assetType_FACTORING string = "5"
	// 基础资产
	assetType_PRODUCT string = "6"
)

type Userid struct {
	UserID string `json:"userid"`
}

//MSP
// RegistrationRequest defines the attributes required to register a user with the CA
type RegistrationRequest struct {
	// Name is the unique name of the identity
	Name string
	// Type of identity being registered (e.g. "peer, app, user")
	Type string
	// MaxEnrollments is the number of times the secret can  be reused to enroll.
	// if omitted, this defaults to max_enrollments configured on the server
	MaxEnrollments int
	// The identity's affiliation e.g. org1.department1
	Affiliation string
	// Optional attributes associated with this identity
	Attributes []Attribute
	// CAName is the name of the CA to connect to
	CAName string
	// Secret is an optional password.  If not specified,
	// a random secret is generated.  In both cases, the secret
	// is returned from registration.
	Secret string
}

type Attribute struct {
	Name  string
	Value string
	ECert bool
}

//  new channel
type ChannelConfig struct {
	ChannelID     string `json:"channelid"`
	ChannelConfig string `json:"channelconfig"`
	Datatime      string `json:"datatime"`
	UUID          string `json:"uuid"`
	ArAssetUUID   string `json:"arassetuuid"`
	OrgName       string `json:"orgname"`
	Peer          string `json:"peer"`
	Anchor        string `json:"anchor"`
	UserName      string `json:"username"`
	UserPass      string `json:"userpass"`
}

// Func_InstallChainCode
type ChainCodeConfig struct {
	ChannelID        string `json:"ChannelID"`
	ChainCodeID      string `json:"ChainCodeID"`
	ChainCodeConfig  string `json:"ChainCodeConfig"`
	ChaincodeVersion string `json:"ChaincodeVersion"`
	Args             string `json:"Args"`
	Datatime         string `json:"datatime"`
	UUID             string `json:"uuid"`
	ArAssetUUID      string `json:"arassetuuid"`
	OrgName          string `json:"orgname"`
	Peer             string `json:"peer"`
	Anchor           string `json:"anchor"`
	UserName         string `json:"username"`
	UserPass         string `json:"userpass"`
}

type AssetsT struct {
	BaseAsset
	BaseType                  `json:"-"`
	BaseUser                  `json:"-"`
	ProUUID                   string   `json:"prouuid"`
	ProNo                     string   `json:"proNo" pkey:""`
	ProPlatForm               string   `json:"proPlatForm"`
	ProName                   string   `json:"proName"`
	ProAddress                string   `json:"proAddress"`
	ProOwnerName              string   `json:"proOwnerName"`
	ProOwnerId                string   `json:"proOwnerId"`
	ProDisposalName           string   `json:"proDisposalName"`
	ProDisposalId             string   `json:"proDisposalId"`
	ProType                   string   `json:"proType"`
	ProSendTime               string   `json:"proSendTime"`
	ProFinaningNum            string   `json:"proFinaningNum"`
	ProFinancingUnit          string   `json:"proFinancingUnit"`
	ProFinancingEndLine       string   `json:"proFinancingEndLine"`
	ProFinancingEndUnit       string   `json:"proFinancingEndUnit"`
	ProFinancingWay           string   `json:"proFinancingWay"`
	ProFinancingPurpose       string   `json:"proFinancingPurpose"`
	ProEcreditWay             string   `json:"proEcreditWay"`
	ProEcreditAssure          string   `json:"proEcreditAssure"`
	ProEcreditGuaranteeName   string   `json:"proEcreditGuaranteeName"`
	ProEcreditGuaranteeId     string   `json:"proEcreditGuaranteeId"`
	ProEcreditGuaranteeType   string   `json:"proEcreditGuaranteeType"`
	ProReportDuediligenceHash string   `json:"proReportDuediligenceHash"`
	ProReportDuediligenceId   string   `json:"proReportDuediligenceId"`
	ProReportThirdHash        string   `json:"proReportThirdHash"`
	ProReportThirdId          string   `json:"proReportThirdId"`
	ProState                  string   `json:"proState"`
	ProNote                   []string `json:"proNote"`
	ProAttachmentlist         []struct {
		IpfsHash      string `json:"ipfsHash"`
		ProAttachName string `json:"proAttachName"`
		ProAttachType string `json:"proAttachType"`
		ProAttachNo   string `json:"proAttachNo" pkey:""`
		ProAttachMD   string `json:"proAttachMD"`
		ProAttachAddr string `json:"proAttachAddr"`
	}
}

// config
type Blockconfig struct {
	ChannelID     string `json:"channelId"`
	Datatime      string `json:"dataTime"`
	OrgBlockName  string `json:"orgBlockName"`
	OrgName       string `json:"orgName"`
	OrgID         string `json:"orgId"`
	Peer          string `json:"peer"`
	Anchor        string `json:"anchor"`
	UserName      string `json:"userName"`
	UserID        string `json:"userId"`
	AffiliationId string `json:"affiliationId"`
}

type ChannelBlockConfig struct {
	Datatime string `json:"dataTime"`
	User     User   `json:"user"`
}

type Blockconfigca struct {
	CaName        string `json:"caName"`
	Datatime      string `json:"datatime"`
	OrgName       string `json:"orgName"`
	OrgID         string `json:"orgId"`
	UserName      string `json:"userName"`
	UserID        string `json:"userId"`
	Peer          string `json:"peer"`
	Anchor        string `json:"anchor"`
	AffiliationId string `json:"affiliationId"`
}

// config
type BlockconfigNum struct {
	ChannelID    string `json:"channelid"`
	Datatime     string `json:"datatime"`
	OrgName      string `json:"orgname"`
	OrgBlockName string `json:"orgblockname"`
	OrgID        string `json:"orgid"`
	Peer         string `json:"peer"`
	Anchor       string `json:"anchor"`
	UserName     string `json:"username"`
	UserId       string `json:"userid"`
	BlockNum     string `json:"blocknum"`
}

// NewAffiliation
type NewAffiliation struct {
	Datatime        string `json:"dataTime"`
	OrgName         string `json:"orgName"`
	OrgID           string `json:"orgId"`
	AffiliationName string `json:"affiliationName"`
	AffiliationID   string `json:"affiliationId"`
	Peer            string `json:"peer"`
	Anchor          string `json:"anchor"`
	UserName        string `json:"userName"`
	UserId          string `json:"userId"`
}

// 其他附件
type ProAttachmentlist struct {
	BaseAsset
	BaseType
	BaseUser
	BaseFetter
	IpfsHash      string `json:"ipfsHash"`
	ProAttachName string `json:"proAttachName"`
	ProAttachType string `json:"proAttachType"`
	// 唯一标识
	ProAttachNo   string `json:"proAttachNo" pkey:""`
	ProAttachMD   string `json:"proAttachMD"`
	ProAttachAddr string `json:"proAttachAddr"`

	//所属资产信息
	FMasterID string `json:fMasterID"`
}

// 其他附件
type ProAttachmentlistT struct {
	BaseAsset
	BaseType      `json:"-"`
	BaseUser      `json:"-"`
	BaseFetter    `json:"-"`
	IpfsHash      string `json:"ipfsHash"`
	ProAttachName string `json:"proAttachName"`
	ProAttachType string `json:"proAttachType"`
	// 唯一标识
	ProAttachNo   string `json:"proAttachNo" pkey:""`
	ProAttachMD   string `json:"proAttachMD"`
	ProAttachAddr string `json:"proAttachAddr"`

	//所属资产信息
	FMasterID string `json:"fMasterID"`
}

//历史追溯、针对产品的子资产
type SubAssetsResult struct {
	ProAttachmentlist []ProAttachmentlist `json:"proAttachmentlist"`
	BaseSurvey        []BaseSurvey        `json:"baseSurvey"`
	BaseReport        []BaseReport        `json:"baseReport"`
}

// config
type BlockQueryTxid struct {
	QueTxid       string `json:"queTxid"`
	Datatime      string `json:"dataTime"`
	OrgBlockName  string `json:"orgBlockName"`
	OrgName       string `json:"orgName"`
	OrgID         string `json:"orgId"`
	Peer          string `json:"peer"`
	Anchor        string `json:"anchor"`
	UserName      string `json:"userName"`
	UserID        string `json:"userId"`
	AffiliationId string `json:"affiliationId"`
}

//Query the txid corresponding to the key
type NewQueryTxID struct {
	Txid string `json:"txID"`
	Key  string `json:"key"`
}

// 云图配置结构体
type YunConfig struct {
	ConfigHeight         string `json:"configHeight"`
	ConfigOrgNum         string `json:"configOrgNum"`
	ConfigNodeNum        string `json:"configNodeNum"`
	ConfigTransactionNum string `json:"configTransactionNum"`
	ConfigChannelNum     string `json:"configChannelNum"`
	ConfigChaincodeNum   string `json:"configChaincodeNum"`
	ConfigTime           string `json:"configTime"`
}

// config
type BlockQuery struct {
	QueTxid  string `json:"queTxid"`
	Datatime string `json:"dataTime"`
	User     User   `json:"user"`
}

type ChannelConfigRe struct {
	ChannelID string `json:"channelId"`
	Datatime  string `json:"dataTime"`
	User      User   `json:"user"`
}

//  根据区块高度查询后续关联的五个区块信息
type BlockFive struct {
	//
	ChannelID string `json:"channelId"`
	BlockNum  string `json:"blockNum"`
	BlockTxID string `json:"blockTxid"`
	User      User   `json:"user"`
}

type BlockWordList struct {
	Datatime string `json:"dataTime"`
	User     User   `json:"user"`
}

type BlockWordback struct {
	Height         uint64 `json:"height"`
	BlockHash      string `json:"blockHash"`
	PreHash        string `json:"preHash"`
	TransactionNum string `json:"transactionNum"`
	BlockTime      string `json:"blockTime"`
	ChannelID      string `json:"channelId"`
}

// 区块配置信息 没有通道
type BlockWordbacknoChannel struct {
	Height         uint64 `json:"height"`
	BlockHash      string `json:"blockHash"`
	PreHash        string `json:"preHash"`
	TransactionNum string `json:"transactionNum"`
	BlockTime      string `json:"blockTime"`
	ChannelID      string `json:"channelId"`
}

//
type QueryBlockCompound struct {
	Id   string `json:"id"`
	User User   `json:"user"`
}

type BlockTransaction struct {
	BlockNum string `json:"blockNum"`
	Datatime string `json:"dataTime"`
	User     User   `json:"user"`
}

// 资产查询
type BlockAssetQuery struct {
	No       string `json:"no"`
	Type     string `json:"type"`
	Datatime string `json:"dataTime"`
	User     User   `json:"user"`
}

type BlockTransactionBack struct {
	Height        int64  `json:"height"`
	BlockHash     string `json:"blockHash"`
	PreHash       string `json:"preHash"`
	TransactionID string `json:"transactionID"`
	BlockTime     string `json:"blockTime"`
	AssetName     string `json:"assetName"`
	AssetUUID     string `json:"assetUuid"`
}

// 尽调结果 集合 上链 T
type BaseSurveyT struct {
	BaseAsset
	BaseType   `json:"-"`
	BaseUser   `json:"-"`
	BaseFetter `json:"-"`
	IpfsHash   string `json:"ipfsHash"`
	// 唯一标识
	SurveyNo string `json:"surveyNo" pkey:""`
	//所属资产信息
	FMasterID string `json:"fMasterID"`
}

// 尽调报告 集合 上链
type BaseReportT struct {
	BaseAsset
	BaseType   `json:"-"`
	BaseUser   `json:"-"`
	BaseFetter `json:"-"`
	IpfsHash   string `json:"ipfsHash"`
	ReportName string `json:"reportName"`
	ReportType string `json:"reportType"`
	// 唯一标识
	ReportNo   string `json:"reportNo" pkey:""`
	ReportTx   string `json:"reportTx"`
	ReportMD   string `json:"reportMD"`
	ReportAddr string `json:"reportAddr"`

	//所属资产信息
	FMasterID string `json:"fMasterID"`
}

// 历史记录表
type HistoryLog struct {
	BaseAsset
	// 数据信息
	HisUUID        string `json:"uuid" pkey:""`
	HisCurrentTime string `json:"hisCurrentTime"`
	HisCurrentTx   string `json:"hisCurrentTx"`
	HisPreTime     string `json:"hisPreTime"`
	HisPreTx       string `json:"hisPreTx"`
	HisRemark      string `json:"hisRemark"`

	HisType   string `json:"hisType"`
	HisDataID string `json:"hisDataId"`

	// 操作人信息
	HisOrgID      string `json:"hisOrgId"`
	HisOrgName    string `json:"hisOrgName"`
	HisUserName   string `json:"hisUserName"`
	HisUserID     string `json:"hisUserId"`
	HisUserRemark string `json:"hisUserRemark"`
	HisCreateTime string `json:"hisCreateTime"`
	HisModify     string `json:"hisModify"`
	HisState      string `json:"hisState"`
}

// 查询历史记录  web

type History struct {
	//
	HisName string `json:"hisName"`
	HisUUID string `json:"uuid"`
	HisNum  string `json:"hisNum"`
	HisType string `json:"hisType"`
}

type QueryOne struct {
	//
	QueName string `json:"queName"`
	QueUUID string `json:"uuid"`
	QueNum  string `json:"queNum"`
	QueType string `json:"queType"`
}

type HistoryData struct {
	History History `json:"history"`
	User    User    `json:"user"`
}

type QueryData struct {
	Data QueryOne `json:"data"`
	User User     `json:"user"`
}

//子资产信息关联主资产信息
type BaseFetter struct {
	FetUUID  string `json:"fetUUID"`
	FetTxID  string `json:"fetTxID"`
	FetTime  string `json:"fetTime"`
	FetTpye  string `json:"fetTpye"`
	FetBz    string `json:"fetBz"`
	FetState string `json:"fetState"`
}

// 资产 子资产集合列表
type BaseAssetsList struct {
	//
	BaseAsset
	SubsetNo    string       `json:"subsetNo"`
	SubsetUUID  string       `json:"subsetUUID" pkey:""`
	SubsetName  string       `json:"subsetName"`
	SubsetType  string       `json:"subsetType"`
	SubsetState string       `json:"subsetState"`
	SubsetList  []BaseAssets `json:"subsetList"`
}

// 子集信息
type BaseAssets struct {
	BstName       string   `json:"bstName"`
	BstType       string   `json:"bstName"`
	BstUUID       string   `json:"bstUUID"`
	BstParentUUID string   `json:"bstParentUUID"`
	BstUpDateTime string   `json:"bstUpDateTime"`
	BstState      string   `json:"bstState"`
	BaseUser      BaseUser `json:"baseUser"`
}

// user 信息  ，数据上传的信息
type BaseUser struct {
	//
	BurUseName    string `json:"burUseName"`
	BurUseID      string `json:"burUseID"`
	BurUseOrgName string `json:"burUseOrgName"`
	BurUseOrgID   string `json:"burUseOrgID"`
	BurUseType    string `json:"burUseType"`
}

// 类型 信息
type BaseType struct {
	BtType  string `json:"btType"`
	BtState string `json:"btState"`
}

//-------------------------------------------------------2019 08 16 新数据结构上链

/*
应收账款数据结构整理：
资产包：
		资产：（多个）
				1.合同
				2.发票
		融资信息：（单个）
		增信措施信息：（单个）
					1. 担保
					2. 质押
					3. 抵押
		发布人信息：（单个）

*/

//	发布人信息
type AstSendInfo struct {
	AstSenduuid      string `json:"astSenduuid" pigkey:""`
	AstSendfMasterID string `json:"astSendfMasterID"`
	AstSendName      string `json:"astSendName"`
	AstSendid        string `json:"astSendid"`
	AstSendContact   string `json:"astSendContact"`
	AstSendTime      string `json:"astSendTime"`
}

//
//	发布人信息 JIU
type AstSendInfoJiu struct {
	AstSenduuid      string `json:"astSenduuid" pigkey:""`
	AstSendfMasterID string `json:"AstSendfMasterID"`
	AstSendName      string `json:"AstSendName"`
	AstSendid        string `json:"AstSendid"`
	AstSendContact   string `json:"AstSendContact"`
	AstSendTime      string `json:"AstSendTime"`
}

//	质押品
type AstCreEnsureList struct {
	AstCreEnsuuid      string `json:"astCreEnsuuid piokey:""` // 质押品uuid
	AstCreEnsfMasterID string `json:"astCreEnsfMasterID"`     //质押品fMasterID
	AstCreEnsType      string `json:"astCreEnsType"`          //质押品类型
	AstCreEnsName      string `json:"astCreEnsName"`          // 质押品名称
	AstCreEnsOwner     string `json:"astCreEnsOwner"`         //质押品所有人
}

type AstCreEnsureListJiu struct {
	AstCreEnsuuid      string `json:"astCreEnsuuid piokey:""` // 质押品uuid
	AstCreEnsfMasterID string `json:"AstCreEnsfMasterID"`     //质押品fMasterID
	AstCreEnsType      string `json:"AstCreEnsType"`          //质押品类型
	AstCreEnsName      string `json:"AstCreEnsName"`          // 质押品名称
	AstCreEnsOwner     string `json:"AstCreEnsOwner"`         //质押品所有人
}

//	抵押品
type AstCrePledgeList struct {
	AstCrePleuuid      string `json:"astCrePleuuid pitkey:""` // 抵押品uuid
	AstCrePlefMasterID string `json:"astCrePlefMasterID"`     // 抵押品fMasterID
	AstCrePleType      string `json:"astCrePleType"`          // 抵押品类型
	AstCrePleName      string `json:"astCrePleName"`          // 抵押品名称
	AstCrePleOwner     string `json:"astCrePleOwner"`         // 抵押品所有人
}

//	抵押品
type AstCrePledgeListJiu struct {
	AstCrePleuuid      string `json:"astCrePleuuid pitkey:""` // 抵押品uuid
	AstCrePlefMasterID string `json:"AstCrePlefMasterID"`     // 抵押品fMasterID
	AstCrePleType      string `json:"AstCrePleType"`          // 抵押品类型
	AstCrePleName      string `json:"AstCrePleName"`          // 抵押品名称
	AstCrePleOwner     string `json:"AstCrePleOwner"`         // 抵押品所有人
}

// 担保 信息
type AstCreGuarantyList struct {
	AstCreGuauuid       string `json:"astCreGuauuid piakey:""` // 担保uuid
	AstCreGuafMasterID  string `json:"astCreGuafMasterID"`     // 担保fMasterID
	AstCreGuaType       string `json:"astCreGuaType"`          // 增信措施类型
	AstCreGuaName       string `json:"astCreGuaName"`          // 保证人名称
	AstCreGuaManner     string `json:"astCreGuaManner"`        // 担保方式
	AstCreGuaEnsureId   string `json:"astCreGuaEnsureId"`      // 担保人ID
	AstCreGuaEnsureName string `json:"astCreGuaEnsureName"`    // 担保人名称
}

// 担保 信息
type AstCreGuarantyListJiu struct {
	AstCreGuauuid       string `json:"astCreGuauuid piakey:""` // 担保uuid
	AstCreGuafMasterID  string `json:"AstCreGuafMasterID"`     // 担保fMasterID
	AstCreGuaType       string `json:"AstCreGuaType"`          // 增信措施类型
	AstCreGuaName       string `json:"AstCreGuaName"`          // 保证人名称
	AstCreGuaManner     string `json:"AstCreGuaManner"`        // 担保方式
	AstCreGuaEnsureId   string `json:"AstCreGuaEnsureId"`      // 担保人ID
	AstCreGuaEnsureName string `json:"AstCreGuaEnsureName"`    // 担保人名称
}

// 担保列表
type AstCreditInfo struct {
	AstCreGuarantyList []AstCreGuarantyList `json:"astCreGuarantyList"` // 担保
	AstCrePledgeList   []AstCrePledgeList   `json:"astCrePledgeList"`   // 抵押品
	AstCreEnsureList   []AstCreEnsureList   `json:"astCreEnsureList"`   // 质押品
}

//
type AstCreditInfoJiu struct {
	AstCreGuarantyList []AstCreGuarantyListJiu `json:"AstCreGuarantyList"` // 担保
	AstCrePledgeList   []AstCrePledgeListJiu   `json:"AstCrePledgeList"`   // 抵押品
	AstCreEnsureList   []AstCreEnsureListJiu   `json:"AstCreEnsureList"`   // 质押品
}

// 融资 信息
type AstFinancingInfo struct {
	AstFinuuid        string `json:"astFinuuid" pivkey:""` // 融资uuid
	AstFinfMasterID   string `json:"astFinfMasterID"`      // 融资fMasterID
	AstFinPrice       string `json:"astFinPrice"`          // 融资金额
	AstCreGuaName     string `json:"astFinLimit"`          // 融资期限
	AstCreGuaManner   string `json:"astFinType"`           // 融资方式
	AstCreGuaEnsureId string `json:"astFinUsefor"`         // 资金用途
}

// 融资 信息 JIU
type AstFinancingInfoJiu struct {
	AstFinuuid        string `json:"astFinuuid" pivkey:""` // 融资uuid
	AstFinfMasterID   string `json:"AstFinfMasterID"`      // 融资fMasterID
	AstFinPrice       string `json:"AstFinPrice"`          // 融资金额
	AstCreGuaName     string `json:"AstCreGuaName"`        // 融资期限
	AstCreGuaManner   string `json:"AstCreGuaManner"`      // 融资方式
	AstCreGuaEnsureId string `json:"AstCreGuaEnsureId"`    // 资金用途
}

// 发票  信息
type AstInv struct {
	AstInvuuid           string           `json:"astInvuuid" pkey:""`   // 发票uuid
	AstInvfMasterID      string           `json:"astInvfMasterID"`      // 发票fMasterID
	AstInvType           string           `json:"astInvType"`           // 发票类型
	AstInvNum            string           `json:"astInvNum"`            // 发票号码
	AstInvCode           string           `json:"astInvCode"`           // 发票代码
	AstInvChecksum       string           `json:"astInvChecksum"`       // 校验码
	AstInvPrice          string           `json:"astInvPrice"`          // 发票金额
	AstInvUnit           string           `json:"astInvUnit"`           // 金额单位
	AstInvTime           string           `json:"astInvTime"`           // 发票开具日期
	AstInvBuyerTaxNum    string           `json:"astInvBuyerTaxNum"`    // 购买方纳税识别号
	AstInvSellerTaxNum   string           `json:"astInvSellerTaxNum"`   // 销售方纳税识别号
	AstInvCheckResult    string           `json:"astInvCheckResult"`    // 发票验真结果
	AstInvAttachmentList []AttachmentList `json:"astInvAttachmentList"` // 发票列表
}

type AstInvJiu struct {
	AstInvuuid           string           `json:"astInvuuid" pkey:""`   // 发票uuid
	AstInvfMasterID      string           `json:"AstInvfMasterID"`      // 发票fMasterID
	AstInvType           string           `json:"AstInvType"`           // 发票类型
	AstInvNum            string           `json:"AstInvNum"`            // 发票号码
	AstInvCode           string           `json:"AstInvCode"`           // 发票代码
	AstInvChecksum       string           `json:"AstInvChecksum"`       // 校验码
	AstInvPrice          string           `json:"AstInvPrice"`          // 发票金额
	AstInvUnit           string           `json:"AstInvUnit"`           // 金额单位
	AstInvTime           string           `json:"AstInvTime"`           // 发票开具日期
	AstInvBuyerTaxNum    string           `json:"AstInvBuyerTaxNum"`    // 购买方纳税识别号
	AstInvSellerTaxNum   string           `json:"AstInvSellerTaxNum"`   // 销售方纳税识别号
	AstInvCheckResult    string           `json:"AstInvCheckResult"`    // 发票验真结果
	AstInvAttachmentList []AttachmentList `json:"AstInvAttachmentList"` // 发票列表
}

// 发票  信息
type AstCon struct {
	AstConuuid           string           `json:"astConuuid" pkey:""`   // 合同uuid
	AstConfMasterID      string           `json:"astConfMasterID"`      // 合同fMasterID
	AstConNo             string           `json:"astConNo"`             // 合同编号
	AstConName           string           `json:"astConName"`           // 合同名称
	AstConPrice          string           `json:"astConPrice"`          // 合同金额
	AstConPriceUnit      string           `json:"astConPriceUnit"`      // 金额单位
	AstConType           string           `json:"astConType"`           // 币种
	AstConPayerTaxNum    string           `json:"astConPayerTaxNum"`    // 付款方纳税识别号
	AstConPayerName      string           `json:"astConPayerName"`      // 付款方名称
	AstConPayeeTaxNum    string           `json:"astConPayeeTaxNum"`    // 收款方纳税识别号
	AstConPayeeName      string           `json:"astConPayeeName"`      // 收款方名称
	AstConTime           string           `json:"astConTime"`           // 合同签署时间
	AstConCount          string           `json:"astConCount"`          // 单次/多次
	AstConDays           string           `json:"astConDays"`           // 账期
	AstConAttachmentList []AttachmentList `json:"astConAttachmentList"` // 附件列表
}

// 合同 动产质押
type AstConByPle struct {
	AstConuuid           string           `json:"astConuuid" pkey:""`   // 合同uuid
	AstConfMasterID      string           `json:"astConfMasterID"`      // 合同fMasterID
	AstConNo             string           `json:"astConNo"`             // 合同编号
	AstConName           string           `json:"astConName"`           // 合同名称
	AstConPrice          string           `json:"astConPrice"`          // 合同金额
	AstConPriceUnit      string           `json:"astConPriceUnit"`      // 金额单位
	AstConType           string           `json:"astConType"`           // 币种
	AstConPayerTaxNum    string           `json:"astConPayerTaxNum"`    // 付款方纳税识别号
	AstConPayerName      string           `json:"astConPayerName"`      // 付款方名称
	AstConPayeeTaxNum    string           `json:"astConPayeeTaxNum"`    // 收款方纳税识别号
	AstConPayeeName      string           `json:"astConPayeeName"`      // 收款方名称
	AstConTime           string           `json:"astConTime"`           // 合同签署时间
	AstConCount          string           `json:"astConCount"`          // 单次/多次
	AstConDays           string           `json:"astConDays"`           // 账期
	AstConAttachmentList []AttachmentList `json:"astConAttachmentList"` // 附件列表
}

//
// 发票  信息
type AstConJiu struct {
	AstConuuid           string              `json:"astConuuid" pkey:""`   // 合同uuid
	AstConfMasterID      string              `json:"AstConfMasterID"`      // 合同fMasterID
	AstConNo             string              `json:"AstConNo"`             // 合同编号
	AstConName           string              `json:"AstConName"`           // 合同名称
	AstConPrice          string              `json:"AstConPrice"`          // 合同金额
	AstConPriceUnit      string              `json:"AstConPriceUnit"`      // 金额单位
	AstConType           string              `json:"AstConType"`           // 币种
	AstConPayerTaxNum    string              `json:"AstConPayerTaxNum"`    // 付款方纳税识别号
	AstConPayerName      string              `json:"AstConPayerName"`      // 付款方名称
	AstConPayeeTaxNum    string              `json:"AstConPayeeTaxNum"`    // 收款方纳税识别号
	AstConPayeeName      string              `json:"AstConPayeeName"`      // 收款方名称
	AstConTime           string              `json:"AstConTime"`           // 合同签署时间
	AstConCount          string              `json:"AstConCount"`          // 单次/多次
	AstConDays           string              `json:"AstConDays"`           // 账期
	AstConAttachmentList []AttachmentListJiu `json:"AstConAttachmentList"` // 附件列表
}

//	附件 信息
type AttachmentList struct {
	AstAttachNo        string `json:"astAttachNo" pkey:""` //附件附加编号
	IpfsHash           string `json:"ipfsHash"`            //IPFS-Hash
	AstAttachfMasterID string `json:"astAttachfMasterID"`  //附件附加fMasterID
	AstAttachName      string `json:"astAttachName"`       //附件附加名称
	AstAttachType      string `json:"astAttachType"`       //附件附加类型
	AstAttachMD        string `json:"astAttachMD"`         //附件附MD
	AstAttachAddr      string `json:"astAttachAddr"`       //附件附加地址
}

//	附件 信息
type AttachmentListJiu struct {
	AstAttachNo        string `json:"astAttachNo" pkey:""` //附件附加编号
	IpfsHash           string `json:"IfsHash"`             //IPFS-Hash
	AstAttachfMasterID string `json:"AstAttachfMasterID"`  //附件附加fMasterID
	AstAttachName      string `json:"AstAttachName"`       //附件附加名称
	AstAttachType      string `json:"AstAttachType"`       //附件附加类型
	AstAttachMD        string `json:"AstAttachMD"`         //附件附MD
	AstAttachAddr      string `json:"AstAttachAddr"`       //附件附加地址
}

// 资产信息  应收账款信息
type AstAssetsList struct {
	//
	AstAssetsuuid           string   `json:"astAssetsuuid" pkey:""`   //资产
	AstAssetsfMasterID      string   `json:"astAssetsfMasterID"`      //外键
	AstAssetsType           string   `json:"astAssetsType"`           //资产类型
	AstAssetsIntroduce      string   `json:"astAssetsIntroduce"`      //资产简介
	AstAssetsCreditorName   string   `json:"astAssetsCreditorName"`   //债权人名称
	AstAssetsCreditorTaxNum string   `json:"astAssetsCreditorTaxNum"` //债权人纳税识别号
	AstAssetsDebtorName     string   `json:"astAssetsDebtorName"`     //债务人名称
	AstAssetsDebtorTaxNum   string   `json:"astAssetsDebtorTaxNum"`   //债务人纳税识别号
	AstAssetsValuation      string   `json:"astAssetsValuation"`      //资产估值
	AstAssetsHonour         string   `json:"astAssetsHonour"`         //资产兑付日
	AstAssetsState          string   `json:"astAssetsState"`          //当前状态
	AstAssetsPrimeval       string   `json:"astAssetsPrimeval"`       //是否为原始资产
	AstContractInfoList     []AstCon `json:"astContractInfoList"`     //合同列表
	AstInvoiceInfoList      []AstInv `json:"astInvoiceInfoList"`      //发票列表
}

// 资产信息  动产融资信息
type AssetsListByPle struct {
	//
	//Id           string   `json:"uuid" pkey:""`   //资产
	Id                  string        `json:"id" pkey:""`          //主键
	Type                string        `json:"type"`                //资产类型
	OwnerEnterName      string        `json:"ownerEnterName"`      //ownerEnterName
	OwnerEnterTaxNum    string        `json:"ownerEnterTaxNum"`    //货主企业纳税识别号
	StoreInfo           StoreInfo     `json:"storeInfo"`           //仓库信息
	AstContractInfoList []AstConByPle `json:"astContractInfoList"` //合同列表
}

// storeInfo --
type StoreInfo struct {
	Id        string      `json:"Id" pkey:""` //仓库id
	Org       string      `json:"org"`        //所属机构
	StoreName string      `json:"storeName"`  //仓库名称
	Attr      string      `json:"attr"`       //归属地
	Address   string      `json:"address"`    //具体地址
	GoodsInfo []GoodsInfo `json:"goodsInfo"`  //资产类型
}

//
type GoodsInfo struct {
	GoodsId     string     `json:"id" pkey:""`  //仓储物名称
	GoodsName   string     `json:"goodsName"`   //仓储物名称
	GoodsMate   string     `json:"goodsMate"`   //仓储物材质
	GoodsOrigin string     `json:"goodsOrigin"` //仓储物产地
	GoodsSum    string     `json:"goodsSum"`    //总数量
	GoodsWeight string     `json:"goodsWeight"` //总重量
	AttrInfo    []AttrInfo `json:"attrInfo"`    //仓储物明细
}

//
type AttrInfo struct {
	AttrName    string `json:"attrName"`    // 仓储物名称
	AttrCode    string `json:"attrCode"`    // 仓储物身份编码
	AttrSize    string `json:"attrSize"`    // 规格
	AttrMate    string `json:"attrMate"`    // 材质
	AttrOrigin  string `json:"attrOrigin"`  // 产地
	AttrArea    string `json:"attrArea"`    // 库区
	AttrPlace   string `json:"attrPlace"`   // 库位
	AttrCeng    string `json:"attrCeng"`    // 层数
	AttrNum     string `json:"attrNum"`     // 数量
	AttrWeight  string `json:"attrWeight"`  // 重量
	AttrQuality string `json:"attrQuality"` // 质量等级
}

//
type AstAssetsListJiu struct {
	//
	AstAssetsuuid           string      `json:"astAssetsuuid" pkey:""`   //资产
	AstAssetsfMasterID      string      `json:"AstAssetsfMasterID"`      //外键
	AstAssetsType           string      `json:"AstAssetsType"`           //资产类型
	AstAssetsIntroduce      string      `json:"AstAssetsIntroduce"`      //资产简介
	AstAssetsCreditorName   string      `json:"AstAssetsCreditorName"`   //债权人名称
	AstAssetsCreditorTaxNum string      `json:"AstAssetsCreditorTaxNum"` //债权人纳税识别号
	AstAssetsDebtorName     string      `json:"AstAssetsDebtorName"`     //债务人名称
	AstAssetsDebtorTaxNum   string      `json:"AstAssetsDebtorTaxNum"`   //债务人纳税识别号
	AstAssetsValuation      string      `json:"AstAssetsValuation"`      //资产估值
	AstAssetsHonour         string      `json:"AstAssetsHonour"`         //资产兑付日
	AstAssetsState          string      `json:"AstAssetsState"`          //当前状态
	AstAssetsPrimeval       string      `json:"AstAssetsPrimeval"`       //是否为原始资产
	AstContractInfoList     []AstConJiu `json:"AstContractInfoList"`     //合同列表
	AstInvoiceInfoList      []AstInvJiu `json:"AstInvoiceInfoList"`      //发票列表
}

//  资产包 信息
type AstAssetsInfo struct {
	//
	BaseAsset            `json:"-"`
	AstPackageuuid       string           `json:"astPackageuuid" pidkey:""` //资产包UUID
	AstPackageNo         string           `json:"astPackageNo"`             //资产包编号
	AstPackagePlatForm   string           `json:"astPackagePlatForm"`       //平台信息-鼎链名称
	AstPackagePlatFormID string           `json:"astPackagePlatFormID"`     //平台信息-鼎链ID
	AstPackageName       string           `json:"astPackageName"`           //资产包名称
	AstPackageNumber     string           `json:"astPackageNumber"`         //资产数量
	AstPackageOwnerName  string           `json:"astPackageOwnerName"`      //资产所有人名称
	AstPackageOwnerId    string           `json:"astPackageOwnerId"`        //资产所有人ID
	AstPackageSplit      string           `json:"astPackageSplit"`          //是否允许拆分
	AstPackageEvaluation string           `json:"astPackageEvaluation"`     //综合评价
	AstAssetsList        []AstAssetsList  `json:"astAssetsList"`            //合同发票列表
	AstFinancingInfo     AstFinancingInfo `json:"astFinancingInfo"`         //融资信息
	AstCreditInfo        AstCreditInfo    `json:"astCreditInfo"`            //增信措施信息
	AstSendInfo          AstSendInfo      `json:"astSendInfo"`              //发布人信息
}

// 动产质押
type AstAssetsInfoByPle struct {
	//
	BaseAsset            `json:"-"`
	AstPackageuuid       string            `json:"astPackageuuid" pidkey:""` //资产包UUID
	AstPackageNo         string            `json:"astPackageNo"`             //资产包编号
	AstPackagePlatForm   string            `json:"astPackagePlatForm"`       //平台信息-鼎链名称
	AstPackagePlatFormID string            `json:"astPackagePlatFormID"`     //平台信息-鼎链ID
	AstPackageName       string            `json:"astPackageName"`           //资产包名称
	AstPackageNumber     string            `json:"astPackageNumber"`         //资产数量
	AstPackageOwnerName  string            `json:"astPackageOwnerName"`      //资产所有人名称
	AstPackageOwnerId    string            `json:"astPackageOwnerId"`        //资产所有人ID
	AstPackageSplit      string            `json:"astPackageSplit"`          //是否允许拆分
	AstPackageEvaluation string            `json:"astPackageEvaluation"`     //综合评价
	AstAssetsListByPle   []AssetsListByPle `json:"astAssetsList"`            //合同发票列表 || 资产发票信息
	AstFinancingInfo     AstFinancingInfo  `json:"astFinancingInfo"`         //融资信息
	AstCreditInfo        AstCreditInfo     `json:"astCreditInfo"`            //增信措施信息
	AstSendInfo          AstSendInfo       `json:"astSendInfo"`              //发布人信息
}

type AstAssetsInfoj struct {
	//
	BaseAsset            `json:"-"`
	AstPackageuuid       string              `json:"astPackageuuid" pidkey:""` //资产包UUID
	AstPackageNo         string              `json:"AstPackageNo"`             //资产包编号
	AstPackagePlatForm   string              `json:"AstPackagePlatForm"`       //平台信息-鼎链名称
	AstPackagePlatFormID string              `json:"AstPackagePlatFormID"`     //平台信息-鼎链ID
	AstPackageName       string              `json:"AstPackageName"`           //资产包名称
	AstPackageNumber     string              `json:"AstPackageNumber"`         //资产数量
	AstPackageOwnerName  string              `json:"AstPackageOwnerName"`      //资产所有人名称
	AstPackageOwnerId    string              `json:"AstPackageOwnerId"`        //资产所有人ID
	AstPackageSplit      string              `json:"AstPackageSplit"`          //是否允许拆分
	AstPackageEvaluation string              `json:"AstPackageEvaluation"`     //综合评价
	AstAssetsList        []AstAssetsListJiu  `json:"AstAssetsList"`            //合同发票列表
	AstFinancingInfo     AstFinancingInfoJiu `json:"AstFinancingInfo"`         //融资信息
	AstCreditInfo        AstCreditInfoJiu    `json:"AstCreditInfo"`            //增信措施信息
	AstSendInfo          AstSendInfoJiu      `json:"AstSendInfo"`              //发布人信息
}

// 应收账款   主体包
type AstInfo struct {
	FBase         Base          `json:"base"` //标记关系
	FUse          Use           `json:"use"`  //标记 用户信息
	AstAssetsInfo AstAssetsInfo `json:"astAssetsInfo"`
	//AstAssetsInfo AstAssetsInfo `json:"astAssetsInfo"`
	// 单个应收账款的UUID
	AstAssetsInfoUUID string `json:"uuid" pkey:""` //资产包UUID
}

// 应收账款   主体包
type AstInfoWebJiu struct {
	FBase Base `json:"-"` //标记关系
	FUse  Use  `json:"-"` //标记 用户信息
	//AstAssetsInfo AstAssetsInfo `json:"astAssetsInfo"`
	AstAssetsInfo AstAssetsInfoj `json:"AstAssetsInfo"`
	// 单个应收账款的UUID
	AstAssetsInfoUUID string `json:"uuid" pkey:""` //资产包UUID
}

// 应收账款   主体包
type AstInfoWeb struct {
	FBase Base `json:"-"` //标记关系
	FUse  Use  `json:"-"` //标记 用户信息
	//AstAssetsInfo AstAssetsInfo `json:"astAssetsInfo"`
	AstAssetsInfo AstAssetsInfo `json:"astAssetsInfo"`
	// 单个应收账款的UUID
	AstAssetsInfoUUID string `json:"uuid" pkey:""` //资产包UUID
}

/**----------------------------------------------------常量--------------------------------------------------------**/

//***********************************************************************************************************************************//
//******************************************房地产********************************//
//***********************************************************************************************************************************//
// 数据结构=>编码表
const (
	Label string = "szbm"

	LabelRequesRealty       = Label + RequestRealty      //编码||房地产标识类型请求
	LabelRequestReceivables = Label + RequestReceivables //编码||应收账款标识类型请求
)

const (
	Label_12           string = "scbm"
	RequestRealty             = "scbn01" //房地产类型请求
	RequestReceivables        = "scbn02" //应收账款类型请求
	Realty                    = "scbm01" //房地产标识
	Receivables               = "scbm02" //应收账款
	accessory                 = "scbm03" //附件标识
	survey                    = "scbm04" //尽调结果标识
	report                    = "scbm05" //报告标识

	//web

	WebIndex string = "webindex_1" //web 首页展示
)

// 请求类型
const (
	FuncNew    = "new"    //新增
	FuncUpdate = "update" //更新
	FuncQuery  = "query"  //查询
)

// 数据类型标识 || 房地产
const (
	DataDueProjectAll   string = "25"     // 主表集合
	DataDueProject      string = "20"     //主表标识
	DataDueProjectFdc   string = "21"     //房地产数据标识
	DataDueProjectCemGt string = "22"     //保证人
	DataDueProjectCemMg string = "23"     //抵押品
	DataDueProjectCemPg string = "24"     //质押品
	Test                string = "10019"  //测试
	Test12039           string = "100312" //测试
)

//	 2019年11月5日09:52:09  || 数据统一标识
const (
	// 房地产
	Pro_Rea_All string = "25" //	房地产主表(整体)
	Pro_Rea_One string = "30" //	主表标识
	Pro_Rea_Fdc string = "31" //	房地产标识
	Pro_Rea_Cgt string = "32" //	保证人标识
	Pro_Rea_Cmg string = "33" //	抵押品标识
	Pro_Rea_Cpg string = "34" //	质押品标识

	// 应收账款
	Pro_Rec_Ysk string = "20" // 应收账款标识
	Pro_Rec_Yjg string = "9"  // 尽调结果标识
	Pro_Rec_Jdb string = "10" // 尽调报告标识
	// 动产质押
	Pro_Rec_JCZY string = "26" // 动产质押

	// 微分格    --  创建 质押 解压
	Pro_Rec_CR  string = "41" // 创建
	Pro_Rec_ZH  string = "42" // 质押
	Pro_Rec_JI  string = "43" // 解压
	Pro_Rec_WI  string = "49" // 微分格标识
	Pro_Rec_Jiu string = "51" // 微分格标识 旧数据

	// 金茂钢包
)

// 微分格
const (
	Wei_Rea_01 string = "41" //
	Wei_Rea_02 string = "42" //
	Wei_Rea_03 string = "43" //
	Wei_Rea_04 string = "44" //

)

//	查询类型	||
const (
	QueryDataListByID string = "queryDataListByID" // 根据主键查询数据关联数据
	QueryDataByID     string = "queryDataByID"     // 根据主键查询元数据
	QueryDataTimeByID string = "queryDataTimeByID" // 根据主键和时间查询元数据
)

// 链码函数名称标识
const (
	Sql_Channel string = "assetpublish" // sql -  通道名称

	Sql_ChainCodeName = "AssetToChain_realty34" // sql -- 链码名称
	Sql_ChainCodoFunc = "uploadAsset"           // sql -- 链码名称
	//UploadAsset string = "uploadAsset" //上链.
	ChainUploadAsset string = "uploadAssetKey" //上链
	ChainUpDataAsset string = "updateAsset"    //数据更新
	//UploadAsset1     string = "uploadAssetKey" //上链
	//ChainUploadAsset string = "uploadAsset" //上链
	//--- 统一链码
	TYUPLOADASSET   string = "uploadAsset" //上链
	TYUPUPDATEASSET string = "updateAsset" //上链 更新

	//-- select  By
	GetAssetList  string = "getAssetList" //查询 链码 By id list
	GETASSEETBYID string = "getAssetByID" //	根据主键查询元数据
	getHistory    string = "getHistory"   //	根据主键查询元数据
	TYUPloadAsset
)

// 房地产主体包
type RealtyInfo struct {
	FBase       Base       `json:"base"` //标记关系
	FUse        Use        `json:"use"`  //标记 用户信息
	FDueProject DueProject `json:"dueProject"`
}

// 返回结构
type RealtyInfoWeb struct {
	FBase       Base       `json:"-"` //标记关系
	FUse        Use        `json:"-"` //标记 用户信息
	FDueProject DueProject `json:"dueProject"`
}

// 动产质押
//TODO
type PleInfo struct {
	FBase             Base               `json:"base"` //标记关系
	FUse              Use                `json:"use"`  //标记 用户信息
	AstAssetsInfo     AstAssetsInfoByPle `json:"astAssetsInfo"`
	AstAssetsInfoUUID string             `json:"uuid" pkey:""` //资产包UUID
}

type PleInfoWeb struct {
	FBase             Base               `json:"-"` //标记关系
	FUse              Use                `json:"-"` //标记 用户信息
	AstAssetsInfo     AstAssetsInfoByPle `json:"astAssetsInfo"`
	AstAssetsInfoUUID string             `json:"uuid" pkey:""` //资产包UUID
}

//主表信息
type DueProject struct {
	FBase Base `json:"-"` //标记关系
	FUse  Use  `json:"-"` //标记 用户信息

	FID                 string `json:"Id" pkey:""`         // uuid 主键
	FSource             string `json:"source"`             //资产来源
	FNumber             string `json:"number"`             //资产包编号
	FName               string `json:"name"`               //资产包名称
	FProjectNum         string `json:"projectNum"`         //资产数量
	FProjectOwner       string `json:"projectOwner"`       //资产所有人名称
	FSplit              string `json:"split"`              //是否允许拆分
	FRemark             string `json:"remark"`             //资产描述
	FIntro              string `json:"intro"`              //资产简介
	FCity               string `json:"city"`               //城市
	FCardNo             string `json:"cardNo"`             //资产方三证合一码
	FType               string `json:"type"`               //标的物类型
	FOrgan              string `json:"organ"`              //所属机构
	FState              string `json:"state"`              //状态
	FLogo               string `json:"logo"`               //封面图
	FPushOrg            string `json:"pushOrg"`            //推送组织
	FFinanceAmount      string `json:"financeAmount"`      //融资金额
	FFnancePeriod       string `json:"financePeriod"`      //融资期限
	FFinanceType        string `json:"financeType"`        //融资方式
	FFundUse            string `json:"fundUse"`            //资金用途
	FMeasure            string `json:"measure"`            //增信措施
	FProjectSourceOrgan string `json:"projectSourceOrgan"` //发布人/机构
	FTelephone          string `json:"telePhone"`          //联系电话
	FProjectSourceTime  string `json:"projectSourceTime"`  //发布时间
	FCreateUser         string `json:"createUser"`         //创建人
	FCreateTime         string `json:"createTime"`         //创建时间
	FUpdateTime         string `json:"updateTime"`         //更新时间

	FastRealtyList   []DueProjectFdc  `json:"astRealtyList"`    //房地产信息
	AstFinancingInfo AstFinancingInfo `json:"astFinancingInfo"` //融资信息
	AstCreditInfo    AstCreditInfo    `json:"astCreditInfo"`    //增信措施信息
	AstSendInfo      AstSendInfo      `json:"astSendInfo"`      //发布人信息
}

//
type DueProjectWeb struct {
	FBase Base `json:"--"` //标记关系
	FUse  Use  `json:"--"` //标记 用户信息

	FID                 string `json:"Id" pkey:""`         // uuid 主键
	FSource             string `json:"source"`             //资产来源
	FNumber             string `json:"number"`             //资产包编号
	FName               string `json:"name"`               //资产包名称
	FProjectNum         string `json:"projectNum"`         //资产数量
	FProjectOwner       string `json:"projectOwner"`       //资产所有人名称
	FSplit              string `json:"split"`              //是否允许拆分
	FRemark             string `json:"remark"`             //资产描述
	FIntro              string `json:"intro"`              //资产简介
	FCity               string `json:"city"`               //城市
	FCardNo             string `json:"cardNo"`             //资产方三证合一码
	FType               string `json:"type"`               //标的物类型
	FOrgan              string `json:"organ"`              //所属机构
	FState              int64  `json:"state"`              //状态
	FLogo               string `json:"logo"`               //封面图
	FPushOrg            string `json:"pushOrg"`            //推送组织
	FFinanceAmount      string `json:"financeAmount"`      //融资金额
	FFnancePeriod       string `json:"financePeriod"`      //融资期限
	FFinanceType        string `json:"financeType"`        //融资方式
	FFundUse            string `json:"fundUse"`            //资金用途
	FMeasure            string `json:"measure"`            //增信措施
	FProjectSourceOrgan string `json:"projectSourceOrgan"` //发布人/机构
	FTelephone          string `json:"telePhone"`          //联系电话
	FProjectSourceTime  string `json:"projectSourceTime"`  //发布时间
	FCreateUser         string `json:"createUser"`         //创建人
	FCreateTime         string `json:"createTime"`         //创建时间
	FUpdateTime         string `json:"updateTime"`         //更新时间

	FastRealtyList   []DueProjectFdcWeb `json:"astRealtyList"`    //房地产信息
	AstFinancingInfo AstFinancingInfo   `json:"astFinancingInfo"` //融资信息
	AstCreditInfo    AstCreditInfo      `json:"astCreditInfo"`    //增信措施信息
	AstSendInfo      AstSendInfo        `json:"astSendInfo"`      //发布人信息
}

type DueProjectKey struct {
	FBase Base `json:"base"` //标记关系
	FUse  Use  `json:"use"`  //标记 用户信息

	FID                 string `json:"Id" pkey:""`         // uuid 主键
	FSource             string `json:"source"`             //资产来源
	FNumber             string `json:"number"`             //资产包编号
	FName               string `json:"name"`               //资产包名称
	FProjectNum         string `json:"projectNum"`         //资产数量
	FProjectOwner       string `json:"projectOwner"`       //资产所有人名称
	FSplit              string `json:"split"`              //是否允许拆分
	FRemark             string `json:"remark"`             //资产描述
	FIntro              string `json:"intro"`              //资产简介
	FCity               string `json:"city"`               //城市
	FCardNo             string `json:"cardNo"`             //资产方三证合一码
	FType               string `json:"type"`               //标的物类型
	FOrgan              string `json:"organ"`              //所属机构
	FState              int64  `json:"state"`              //状态
	FLogo               string `json:"logo"`               //封面图
	FPushOrg            string `json:"pushOrg"`            //推送组织
	FFinanceAmount      string `json:"financeAmount"`      //融资金额
	FFnancePeriod       string `json:"financePeriod"`      //融资期限
	FFinanceType        string `json:"financeType"`        //融资方式
	FFundUse            string `json:"fundUse"`            //资金用途
	FMeasure            string `json:"measure"`            //增信措施
	FProjectSourceOrgan string `json:"projectSourceOrgan"` //发布人/机构
	FTelephone          string `json:"telePhone"`          //联系电话
	FProjectSourceTime  string `json:"projectSourceTime"`  //发布时间
	FCreateUser         string `json:"createUser"`         //创建人
	FCreateTime         string `json:"createTime"`         //创建时间
	FUpdateTime         string `json:"updateTime"`         //更新时间
}

/**
@ 房地产标识
@ 2019年10月18日15:15:35
@ lidongsun
*/
type DueProjectFdc struct {
	FBase Base `json:"-"` //标记关系
	FUse  Use  `json:"-"` //标记 用户信息

	FID             string `json:"Id" `            // uuid 主键 房地产id
	FProjectId      string `json:"projectId"`      //资产id
	FProjectNum     string `json:"projectNum"`     //资产编号
	FType           string `json:"type"`           //资产类型
	FTypeTwo        string `json:"typeTwo"`        //二级分类
	FFdcNum         string `json:"fdcNum"`         //房产证编号
	FAddress        string `json:"address"`        //地址
	FProjectComment string `json:"projectComment"` //资产简介
	FTotalLandArea  string `json:"totalLandArea"`  //总用地面积
	FTotalBuildArea string `json:"totalBuildArea"` //总建筑面积
	FProjectValua   string `json:"projectValua"`   //资产估值
	FProjress       string `json:"projress"`       //建设进度
	FCurrentType    string `json:"currentType"`    //当前状态
	FFile           string `json:"file"`           //附件
	FCreateTime     string `json:"createTime"`     //创建时间
}

type DueProjectFdcWeb struct {
	FBase Base `json:"--"` //标记关系
	FUse  Use  `json:"--"` //标记 用户信息

	FID             string `json:"Id" `            // uuid 主键 房地产id
	FProjectId      string `json:"projectId"`      //资产id
	FProjectNum     string `json:"projectNum"`     //资产编号
	FType           string `json:"type"`           //资产类型
	FTypeTwo        string `json:"typeTwo"`        //二级分类
	FFdcNum         string `json:"fdcNum"`         //房产证编号
	FAddress        string `json:"address"`        //地址
	FProjectComment string `json:"projectComment"` //资产简介
	FTotalLandArea  string `json:"totalLandArea"`  //总用地面积
	FTotalBuildArea string `json:"totalBuildArea"` //总建筑面积
	FProjectValua   string `json:"projectValua"`   //资产估值
	FProjress       string `json:"projress"`       //建设进度
	FCurrentType    string `json:"currentType"`    //当前状态
	FFile           string `json:"file"`           //附件
	FCreateTime     string `json:"createTime"`     //创建时间
}

type DueProjectFdcKey struct {
	FBase Base `json:"base"` //标记关系
	FUse  Use  `json:"use"`  //标记 用户信息

	FID             string `json:"Id" pkey:""`     // uuid 主键 房地产id
	FProjectId      string `json:"projectId"`      //资产id
	FProjectNum     string `json:"projectNum"`     //资产编号
	FType           string `json:"type"`           //资产类型
	FTypeTwo        string `json:"typeTwo"`        //二级分类
	FFdcNum         string `json:"fdcNum"`         //房产证编号
	FAddress        string `json:"address"`        //地址
	FProjectComment string `json:"projectComment"` //资产简介
	FTotalLandArea  string `json:"totalLandArea"`  //总用地面积
	FTotalBuildArea string `json:"totalBuildArea"` //总建筑面积
	FProjectValua   string `json:"projectValua"`   //资产估值
	FProjress       string `json:"projress"`       //建设进度
	FCurrentType    string `json:"currentType"`    //当前状态
	FFile           string `json:"file"`           //附件
	FCreateTime     string `json:"createTime"`     //创建时间
}

/**
@ 保证人 DueProjectCemGt
@ 2019年10月18日15:21:54
@
*/
type DueProjectCemGt struct {
	FBase Base `json:"base"` //标记关系
	FUse  Use  `json:"use"`  //标记 用户信息

	FID            string `json:"id" pkey:""`    // uuid 主键 房地产id
	FProjectId     string `json:"projectId"`     //资产id
	FGuaranteeMode string `json:"guaranteeMode"` //担保方式
	FGuaranteeUser string `json:"guaranteeUser"` //保证人名称
	FCreateTime    string `json:"createTime"`    //创建时间
}

/**
@ 抵押品 DueProjectCemGt
@ 2019年10月18日15:21:54
@
*/
type DueProjectCemMg struct {
	FBase Base `json:"base"` //标记关系
	FUse  Use  `json:"use"`  //标记 用户信息

	FID              string `json:"Id" pkey:""`      // uuid 主键 房地产id
	FProjectId       string `json:"projectId"`       //资产id
	FCollateralType  string `json:"collateralType"`  //抵押品类型
	FCollateralName  string `json:"collateralName"`  //抵押品名称
	FCollateralOwner string `json:"collateralOwner"` //抵押品所有人
	FCreateTime      string `json:"createTime"`      //创建时间
}

/**
@ 质押 DueProjectCemPg
@ 2019年10月18日15:21:54
@
*/
type DueProjectCemPg struct {
	FBase Base `json:"base"` //标记关系
	FUse  Use  `json:"use"`  //标记 用户信息

	FID          string `json:"Id" pkey:""`  // uuid 主键 房地产id
	FProjectId   string `json:"projectId"`   //资产id
	FPledge      string `json:"pledge"`      //质押品类型
	FPledegName  string `json:"pledegName"`  //质押品名称
	FPledegOwner string `json:"pledegOwner"` //质押品所有人
	FCreateTime  string `json:"createTime"`  //创建时间
}

//***********************************************************************************************************************************//
//******************************************房地产********************************//
//***********************************************************************************************************************************//
// 基础信息
// 资产基本字段
type Base struct {
	// IsDeleted  bool   `json:"isDeleted"`
	UpdateDate string `json:"updateDate"`
	// 仅用在底层资产中、其Parent有可能为产品、保理
	ParentType string `json:"parentType"`
	ParentID   string `json:"parentID"`
	Ptype      string `json:"Ptype"` // 字段类型
}

//状态码 标识 |||
const (
	StatusContinue           = 100 // RFC 7231, 6.2.1
	StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1

	StatusOK                   = 200  // RFC 7231, 6.3.1
	StatusFailed               = 4111 // RFC 7231, 6.3.1
	StatusCreated              = 201  // RFC 7231, 6.3.2
	StatusAccepted             = 202  // RFC 7231, 6.3.3
	StatusNonAuthoritativeInfo = 203  // RFC 7231, 6.3.4
	StatusNoContent            = 204  // RFC 7231, 6.3.5
	StatusResetContent         = 205  // RFC 7231, 6.3.6
	StatusPartialContent       = 206  // RFC 7233, 4.1
	StatusMultiStatus          = 207  // RFC 4918, 11.1
	StatusAlreadyReported      = 208  // RFC 5842, 7.1
	StatusIMUsed               = 226  // RFC 3229, 10.4.1

	StatusMultipleChoices   = 300 // RFC 7231, 6.4.1
	StatusMovedPermanently  = 301 // RFC 7231, 6.4.2
	StatusFound             = 302 // RFC 7231, 6.4.3
	StatusSeeOther          = 303 // RFC 7231, 6.4.4
	StatusNotModified       = 304 // RFC 7232, 4.1
	StatusUseProxy          = 305 // RFC 7231, 6.4.5
	_                       = 306 // RFC 7231, 6.4.6 (Unused)
	StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7
	StatusPermanentRedirect = 308 // RFC 7538, 3

	StatusBadRequest                   = 400 // RFC 7231, 6.5.1
	StatusUnauthorized                 = 401 // RFC 7235, 3.1
	StatusPaymentRequired              = 402 // RFC 7231, 6.5.2
	StatusForbidden                    = 403 // RFC 7231, 6.5.3
	StatusNotFound                     = 404 // RFC 7231, 6.5.4
	StatusMethodNotAllowed             = 405 // RFC 7231, 6.5.5
	StatusNotAcceptable                = 406 // RFC 7231, 6.5.6
	StatusProxyAuthRequired            = 407 // RFC 7235, 3.2
	StatusRequestTimeout               = 408 // RFC 7231, 6.5.7
	StatusConflict                     = 409 // RFC 7231, 6.5.8
	StatusGone                         = 410 // RFC 7231, 6.5.9
	StatusLengthRequired               = 411 // RFC 7231, 6.5.10
	StatusPreconditionFailed           = 412 // RFC 7232, 4.2
	StatusRequestEntityTooLarge        = 413 // RFC 7231, 6.5.11
	StatusRequestURITooLong            = 414 // RFC 7231, 6.5.12
	StatusUnsupportedMediaType         = 415 // RFC 7231, 6.5.13
	StatusRequestedRangeNotSatisfiable = 416 // RFC 7233, 4.4
	StatusExpectationFailed            = 417 // RFC 7231, 6.5.14
	StatusTeapot                       = 418 // RFC 7168, 2.3.3
	StatusMisdirectedRequest           = 421 // RFC 7540, 9.1.2
	StatusUnprocessableEntity          = 422 // RFC 4918, 11.2
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusTooEarly                     = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              = 426 // RFC 7231, 6.5.15
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 7231, 6.6.1
	StatusNotImplemented                = 501 // RFC 7231, 6.6.2
	StatusBadGateway                    = 502 // RFC 7231, 6.6.3
	StatusServiceUnavailable            = 503 // RFC 7231, 6.6.4
	StatusGatewayTimeout                = 504 // RFC 7231, 6.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 7231, 6.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)

var statusText = map[int64]string{
	StatusContinue:             "调用链码处理错误",
	StatusSwitchingProtocols:   "查询交易出错",
	StatusProcessing:           "Processing",
	StatusOK:                   "success",
	StatusFailed:               "failed",
	StatusCreated:              "Created",
	StatusAccepted:             "Accepted",
	StatusNonAuthoritativeInfo: "Non-Authoritative Information",
	StatusNoContent:            "No Content",
	StatusResetContent:         "Reset Content",
	StatusPartialContent:       "Partial Content",
	StatusMultiStatus:          "Multi-Status",
	StatusAlreadyReported:      "Already Reported",
	StatusIMUsed:               "IM Used",

	StatusMultipleChoices:   "Multiple Choices",
	StatusMovedPermanently:  "Moved Permanently",
	StatusFound:             "Found",
	StatusSeeOther:          "See Other",
	StatusNotModified:       "Not Modified",
	StatusUseProxy:          "Use Proxy",
	StatusTemporaryRedirect: "Temporary Redirect",
	StatusPermanentRedirect: "Permanent Redirect",

	StatusBadRequest:                   "Bad Request",
	StatusUnauthorized:                 "Unauthorized",
	StatusPaymentRequired:              "Payment Required",
	StatusForbidden:                    "Forbidden",
	StatusNotFound:                     "Not Found",
	StatusMethodNotAllowed:             "Method Not Allowed",
	StatusNotAcceptable:                "Not Acceptable",
	StatusProxyAuthRequired:            "Proxy Authentication Required",
	StatusRequestTimeout:               "Request Timeout",
	StatusConflict:                     "Conflict",
	StatusGone:                         "Gone",
	StatusLengthRequired:               "Length Required",
	StatusPreconditionFailed:           "Precondition Failed",
	StatusRequestEntityTooLarge:        "Request Entity Too Large",
	StatusRequestURITooLong:            "Request URI Too Long",
	StatusUnsupportedMediaType:         "Unsupported Media Type",
	StatusRequestedRangeNotSatisfiable: "Requested Range Not Satisfiable",
	StatusExpectationFailed:            "Expectation Failed",
	StatusTeapot:                       "I'm a teapot",
	StatusMisdirectedRequest:           "Misdirected Request",
	StatusUnprocessableEntity:          "Unprocessable Entity",
	StatusLocked:                       "Locked",
	StatusFailedDependency:             "Failed Dependency",
	StatusTooEarly:                     "Too Early",
	StatusUpgradeRequired:              "Upgrade Required",
	StatusPreconditionRequired:         "Precondition Required",
	StatusTooManyRequests:              "Too Many Requests",
	StatusRequestHeaderFieldsTooLarge:  "Request Header Fields Too Large",
	StatusUnavailableForLegalReasons:   "Unavailable For Legal Reasons",

	StatusInternalServerError:           "Internal Server Error",
	StatusNotImplemented:                "Not Implemented",
	StatusBadGateway:                    "Bad Gateway",
	StatusServiceUnavailable:            "Service Unavailable",
	StatusGatewayTimeout:                "Gateway Timeout",
	StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
	StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	StatusInsufficientStorage:           "Insufficient Storage",
	StatusLoopDetected:                  "Loop Detected",
	StatusNotExtended:                   "Not Extended",
	StatusNetworkAuthenticationRequired: "Network Authentication Required",
}

func StatusText(code int64) string {
	return statusText[code]
}

//---------------------------------------  数据上链结构体

// Kyc 数据结构
type Kyc struct {
	// file 	KycID     string `json:"id" pkey:""` // 主键
	KycID     string `json:"id" pkey:""` // 主键
	KycType   string `json:"kycType"`    // 数据类型
	KycTime   string `json:"kycTime"`    // 上链时间
	KycString string `json:"kycString"`  // 内容密文
	SignKey   string `json:"signKey"`    // 签名key
	SignPower string `json:"signPower"`  // 访问权限  || 访问权限    1 2 3 4
	SignUser  Use    `json:"signUser"`   //签名用户信息
	//SignUser string `json:"signUser"` //签名用户信息
}

// test  Kyc
type Kyctest struct {
	// file 	KycID     string `json:"id" pkey:""` // 主键
	KycID     string `json:"id" pkey:""` // 主键
	KycType   string `json:"kycType"`    // 数据类型
	KycTime   string `json:"kycTime"`    // 上链时间
	KycString string `json:"kycString"`  // 内容密文
	SignKey   string `json:"signKey"`    // 签名key
	SignPower string `json:"signPower"`  // 访问权限  || 访问权限    1 2 3 4
	//SignUser  Use    `json:"signUser"`   //签名用户信息
	SignUser string `json:"signUser"` //签名用户信息
}

//	Byc 数据结构
type Byc struct {
	BycID        string     `json:"id" pkey:""`   //  uuid
	BycUser      Use        `json:"bycUser"`      // 用户信息
	BycMes       []BycMes   `json:"bycMes"`       // 数据关联信息
	BycSign      Sign       `json:"sign"`         // 权限信息
	BycSignList  []SignList `json:"signList"`     // 权限信息列表
	BycByKey     string     `json:"bycByKey"`     // 宿主ID
	BycByfMaster string     `json:"bycByfMaster"` // 关联宿主
	BycBySpan    string     `json:"span"`         // 是否跨链码
}

// Byc 数据关联结构
type BycMes struct {
	BycID  string `json:"bycId"`  //外键
	BycKey string `json:"bycKey"` //秘钥

	ParBycID  string `json:"parBycID"`  //外键
	ParBycKey string `json:"parBycKey"` //秘钥

	BycRelation string `json:"bycRelation"` //外键关系  parent  sun
	BySol       string `json:"bySol"`       //开关
}

// Byc 权限结构
type Sign struct {
	OrgKey  string `json:"orgKey"`  // 组织秘钥
	UserKey string `json:"userKey"` // 用户秘钥
	Public  string `json:"public"`  // 公共秘钥
	Self    string `json:"self"`    // 本数据秘钥
}

// Byc 访问权限列表
type SignList struct {
	RelID   string `json:"relID"`   // 访问ID
	RelType string `json:"relType"` // 用户 || 数据
	RelTime string `json:"relTime"` // 访问最后时间
	RelNum  string `json:"relNum"`  // 访问数量
	RelSol  string `json:"relSol"`  // 是否开启
}

// Kyc||Byc|| 用户信息记录
type Use struct {
	UseName    string `json:"useName"`    // 用户名称
	UseID      string `json:"useId"`      // 用户ID
	UseOrgName string `json:"useOrgName"` // 组织名称
	UseOrgID   string `json:"useOrgId"`   // 组织ID
	UseType    string `json:"useType"`    // 用户类型
	UseCa      string `json:"useCa"`      // 用户ca名称 ||
}

//	数据查看权限
const (
	signPower01 string = "SPXY01" // 一级|| 个人级别 仅自己看
	signPower02 string = "SPXY02" // 二级|| 个人级别 可授权
	signPower03 string = "SPXY03" // 三级|| 公司级别
	signPower04 string = "SPXY04" // 四级|| 公开
	signPower05 string = "SPXY05" // 五级|| 备用
)

//	状态开关
const (
	//
	State_Yes string = "yes"
	State_No  string = "no"
	State_Zan string = "zan"
)

const (
	KYC = "KYC"
	BYC = "BYC"
)

// 	TXID  查询返回
type ChainTransactionConfig struct {
	Height      int64       `json:"height"`
	Timestamp   int64       `json:"timestamp"`
	CreatedFlag bool        `json:"createdFlag"`
	TxArgs      [][]byte    `json:"-"`
	Hash        string      `json:"hash"`
	Time        string      `json:"time"`
	PreHash     string      `json:"preHash"`
	Type        string      `json:"type"`
	DataTxid    string      `json:"dataTxid"`
	TxId        string      `json:"txid"`
	Version     string      `json:"version"`
	Chaincode   string      `json:"chaincode"`
	Method      string      `json:"method"`
	ChannelId   string      `json:"channelId"`
	PledgeType  string      `json:"pledgeType"`
	UserId      string      `json:"userId"`
	Incident    string      `json:"incident"`
	Timestampes interface{} `json:"timeshs"`
}

type TxIDrenData struct {
	//
	ChainTransactionConfig ChainTransactionConfig `json:"config"`
	Data                   interface{}            `json:"info"`
	DataType               string                 `json:"dType"`
}

//
type TxIDrenDataJiu struct {
	//
	ChainTransactionConfig ChainTransactionConfig `json:"config"`
	Data                   interface{}            `json:"AstInfo"`
	DataType               string                 `json:"dType"`
}

// 担保列表
type TxIDren struct {
	InfoList           interface{} `json:"infoList"`           // 担保
	AstCreGuarantyList interface{} `json:"astCreGuarantyList"` // 担保
	AstCrePledgeList   interface{} `json:"astCrePledgeList"`   // 抵押品
	AstCreEnsureList   interface{} `json:"astCreEnsureList"`   // 质押品
	AstSendInfo        interface{} `json:"astSendInfo"`        //发布人信息
}

// V3 查询ByID List
type UseDataById struct {
	ById   string `json:"id"`
	ByTime string `json:"time"`
}

//  父子 标签
const (
	Title_Parent string = "p"
	Title_Sun    string = "s"
	Title_Tong   string = "t"
)

//	关系标签 结构
type LinkedData struct {
	//
	Rel  string      `json:"rel"`
	Data interface{} `json:"data"`
}

// V3 ---------------------  升级

//  注册组织
type UnionWeb struct {
	OrgName         string `json:"orgName"`
	OrgID           string `json:"orgId"`
	UserName        string `json:"userName"`
	UserId          string `json:"userId"`
	AffiliationName string `json:"affiliationName"`
	AffiliationID   string `json:"affiliationId"`
	Peer            string `json:"peer"`
	Anchor          string `json:"anchor"`
	Datatime        string `json:"dataTime"`
}

// 注册组织用户
type Registered struct {
	//必备 字段
	UserName           string `json:"userName"`         // 验证 用户名称
	UserID             string `json:"userId"`           // 验证 用户主键
	RegOrgName         string `json:"regOrgName"`       // 验证 组织名称
	RegOrgID           string `json:"regOrgId"`         // 验证 组织ID
	RegAffiliationName string `json:"regAffiliationId"` //验证  联盟名称
	RegAffiliationId   string `json:"regAffiliationId"` //验证 联盟ID
	RegUserName        string `json:"regUserName"`      // 注册用户名称
	RegUserID          string `json:"regUserId"`        // 注册用户ID
	// sf
	RegID    string `json:"regID"`  // 注册数据ID
	RegState string `json:"state"`  // 注册状态
	Peer     string `json:"peer"`   // Peer 节点
	Anchor   string `json:"anchor"` // 锚节点
}

// 上链请求数据
type Info struct {
	Datatype string      `json:"dataType"`
	Data     interface{} `json:"data"`
	User     User        `json:"user"`
	Datafunc string      `json:"datafunc"`
	//Chain    Chain       `json:"chain"`
}

// 区块链信息
type Chain struct {
	ChannelName   string `json:"channelName"`
	ChainCodeName string `json:"chainCodeName"`
	FuncName      string `json:"funcName"`
	UUID          string `json:"uuid"`
}

// 尽调结果 集合 上链
type BaseSurvey struct {
	IpfsHash string `json:"ipfsHash"` // ipfs hash
}

type BaseSurveyByJiu struct {
	IpfsHash string `json:"ipfsHash"` // ipfs hash
}

// 尽调报告 集合 上链
type BaseReport struct {
	IpfsHash   string `json:"ipfsHash"`         // ipfs hash
	ReportName string `json:"reportName"`       // 报告名称
	ReportType string `json:"reportType"`       //报告类型
	ReportNo   string `json:"reportNo" pkey:""` // 唯一标识  TODO 报告主键
	ReportTx   string `json:"reportTx"`         // 报告txid
	ReportMD   string `json:"reportMD"`         // 报告MD5
	ReportAddr string `json:"reportAddr"`       //报告地址
	FMasterID  string `json:"fMasterID"`        //所属资产信息 外键
}

// 尽调报告 集合 上链
type BaseReportByJiu struct {
	IpfsHash   string `json:"ipfsHash"`         // ipfs hash
	ReportName string `json:"reportName"`       // 报告名称
	ReportType string `json:"reportType"`       //报告类型
	ReportNo   string `json:"reportNo" pkey:""` // 唯一标识  TODO 报告主键
	ReportTx   string `json:"reportTx"`         // 报告txid
	ReportMD   string `json:"reportMD"`         // 报告MD5
	ReportAddr string `json:"reportAddr"`       //报告地址
	FMasterID  string `json:"fMasterID"`        //所属资产信息 外键
}

//	 微分格测试数据结构

type Box struct {
	XbId      string `json:"ipfsHash" pkey:""` // XbId
	XbNname   string `json:"ipfsHash"`         // XbNname
	XbAge     string `json:"ipfsHash"`         // XbAge
	XbAdrress string `json:"ipfsHash"`         // XbAdrress
	Xbtag     string `json:"ipfsHash"`         // Xbtag
}

// 微分格
// 上链请求数据
type InfoWei struct {
	Data  []string `json:"data"`   // 数据集
	Chain Chain    `json:"config"` // 链上信息
	User  User     `json:"user"`   // 用户信息
}

//
type OrgWeb struct {
	Affiliations string      `json:"Affiliations"`
	CAName       interface{} `json:"CAName"`
	Identities   interface{} `json:"Identities"`
	Name         string      `json:"Name"`
}

// ----------------------------------------- 微分格 上链数据定义

// ByCr
type WeiByCr struct {
	// 基础数据类型
	WhrId                     interface{} `json:"whrId"`                     // 仓单编号
	WarehouserName            interface{} `json:"warehouserName"`            // 仓储方名称
	WhName                    interface{} `json:"whName"`                    // 仓库地址
	Fullname                  interface{} `json:"fullname"`                  // 贸易方
	WhrLotno                  interface{} `json:"whrLotno"`                  // 仓单编号
	OwnerName                 interface{} `json:"ownerName"`                 // 当前持有人
	Remarks                   interface{} `json:"remarks"`                   // 备注
	WhFeeRate                 interface{} `json:"whFeeRate"`                 // 仓储费
	TradeConfirmTime          interface{} `json:"tradeConfirmTime"`          // 仓单生效时间
	WhrExpireTime             interface{} `json:"whrExpireTime"`             // 仓单到期时间
	WhrPreparedBy             interface{} `json:"whrPreparedBy"`             // 仓单制单人名称
	CreateTime                interface{} `json:"createTime"`                // 仓单制单时间
	WhrReviewTheBiller        interface{} `json:"whrReviewTheBiller"`        // 复核记账人
	ReviewTime                interface{} `json:"reviewTime"`                // 复合记账时间
	WhrTradePartyConfirmation interface{} `json:"whrTradePartyConfirmation"` //贸易方确认
	//TradeConfirmTime          interface{} `json:"tradeConfirmTime"`          // 贸易方确认时间
	RemainValue   interface{} `json:"remainValue"`   // 货物合计
	WarehouseName interface{} `json:"warehouseName"` // 仓储名称
	// 默认
	FromOwnerName interface{} `json:"fromOwnerName"`
	SignTime      interface{} `json:"signTime"`
	ToOwnerName   interface{} `json:"toOwnerName"`
	TradeId       interface{} `json:"tradeId"`

	// 列表
	GoodsList                    []WeiByCrForGoodsList `json:"goodsList"`                    // 货物列表
	InboundAndOutboundRecords    []WeiByCrForInbo      `json:"inboundAndOutboundRecords"`    // 货物出入库记录数据
	QualityInspectionInformation []WeiByCrForQual      `json:"qualityInspectionInformation"` // 仓单质检信息
	Responsibility               []string              `json:"responsibility"`               // 责任补偿
	Warranty                     []string              `json:"warranty"`                     // 担保补偿
	//InsureCert                   []InsureCert          `json:"insureCert"`                   // 保证书
	InsureCert      []string      `json:"insureCert"`      // 保证书
	OwnershipChange []interface{} `json:"ownershipChange"` // 备用
}

type WeiByCrbf struct {

	// 基础数据类型
	CreateTime                interface{} `json:"createTime"`
	Fullname                  interface{} `json:"fullname"`
	OwnerName                 interface{} `json:"ownerName"`
	RemainValue               interface{} `json:"remainValue"`
	Remarks                   interface{} `json:"remarks"`
	ReviewTime                interface{} `json:"reviewTime"`
	WarehouserName            interface{} `json:"warehouserName"`
	WhrExpireTime             interface{} `json:"whrExpireTime"`
	WhrId                     interface{} `json:"whrId"`
	WhrLotno                  interface{} `json:"whrLotno"`
	WhrPreparedBy             interface{} `json:"whrPreparedBy"`
	WhrReviewTheBiller        interface{} `json:"whrReviewTheBiller"`
	WhrTradePartyConfirmation interface{} `json:"whrTradePartyConfirmation"`

	// 默认
	FromOwnerName interface{} `json:"fromOwnerName"`
	SignTime      interface{} `json:"signTime"`
	ToOwnerName   interface{} `json:"toOwnerName"`
	TradeId       interface{} `json:"tradeId"`

	// 列表
	GoodsList                    []WeiByCrForGoodsListbf `json:"goodsList"`                    // 货物列表
	InboundAndOutboundRecords    []WeiByCrForInbobf      `json:"inboundAndOutboundRecords"`    // 货物出入库记录数据
	QualityInspectionInformation []WeiByCrForQualbf      `json:"qualityInspectionInformation"` // 仓单质检信息
	Responsibility               []Responsibility        `json:"responsibility"`               // 责任补偿
	Warranty                     []Warranty              `json:"warranty"`                     // 担保补偿
	InsureCert                   []InsureCert            `json:"insureCert"`                   // 保证书
	OwnershipChange              []interface{}           `json:"ownershipChange"`
}

//  保证书
type InsureCert struct {
	Id           interface{} `json:"id"`           // 记录编号
	CertId       interface{} `json:"certId"`       // 凭证编号
	CertVer      interface{} `json:"certVer"`      //凭证版本
	FileId       interface{} `json:"fileId"`       // 文件编号
	FilePath     interface{} `json:"filePath"`     //文件路径
	FileName     interface{} `json:"fileName"`     //文件名
	FileState    interface{} `json:"fileState"`    //文件状态
	FileVer      interface{} `json:"fileVer"`      //文件版本
	FileMd5      interface{} `json:"fileMd5"`      //文件md5值
	BlockId      interface{} `json:"blockId"`      //区块号
	ChainNo      interface{} `json:"chainNo"`      //上链编号
	Remarks      interface{} `json:"remarks"`      //备注
	CreateUser   interface{} `json:"createUser"`   //创建ID
	CreateTime   interface{} `json:"createTime"`   //创建时间
	UpdateUser   interface{} `json:"updateUser"`   //修改人ID
	UpdateTime   interface{} `json:"updateTime"`   //修改时间
	IsDeleted    interface{} `json:"isDeleted"`    //删除标识
	RelativePath interface{} `json:"relativePath"` //相对路径
}

// 担保补偿
type Warranty struct {
	Id           interface{} `json:"id"`           // 记录编号
	CertId       interface{} `json:"certId"`       // 凭证编号
	CertVer      interface{} `json:"certVer"`      //凭证版本
	FileId       interface{} `json:"fileId"`       // 文件编号
	FilePath     interface{} `json:"filePath"`     //文件路径
	FileName     interface{} `json:"fileName"`     //文件名
	FileState    interface{} `json:"fileState"`    //文件状态
	FileVer      interface{} `json:"fileVer"`      //文件版本
	FileMd5      interface{} `json:"fileMd5"`      //文件md5值
	BlockId      interface{} `json:"blockId"`      //区块号
	ChainNo      interface{} `json:"chainNo"`      //上链编号
	Remarks      interface{} `json:"remarks"`      //备注
	CreateUser   interface{} `json:"createUser"`   //创建ID
	CreateTime   interface{} `json:"createTime"`   //创建时间
	UpdateUser   interface{} `json:"updateUser"`   //修改人ID
	UpdateTime   interface{} `json:"updateTime"`   //修改时间
	IsDeleted    interface{} `json:"isDeleted"`    //删除标识
	RelativePath interface{} `json:"relativePath"` //相对路径
}

//
type Responsibility struct {
	Id           interface{} `json:"id"`           // 记录编号
	CertId       interface{} `json:"certId"`       // 凭证编号
	CertVer      interface{} `json:"certVer"`      //凭证版本
	FileId       interface{} `json:"fileId"`       // 文件编号
	FilePath     interface{} `json:"filePath"`     //文件路径
	FileName     interface{} `json:"fileName"`     //文件名
	FileState    interface{} `json:"fileState"`    //文件状态
	FileVer      interface{} `json:"fileVer"`      //文件版本
	FileMd5      interface{} `json:"fileMd5"`      //文件md5值
	BlockId      interface{} `json:"blockId"`      //区块号
	ChainNo      interface{} `json:"chainNo"`      //上链编号
	Remarks      interface{} `json:"remarks"`      //备注
	CreateUser   interface{} `json:"createUser"`   //创建ID
	CreateTime   interface{} `json:"createTime"`   //创建时间
	UpdateUser   interface{} `json:"updateUser"`   //修改人ID
	UpdateTime   interface{} `json:"updateTime"`   //修改时间
	IsDeleted    interface{} `json:"isDeleted"`    //删除标识
	RelativePath interface{} `json:"relativePath"` //相对路径
}

//
type WeiByCrForGoodsList struct {
	GoodsId         interface{} `json:"goodsId"`         // 货物编号
	GoodsName       interface{} `json:"goodsName"`       // 货物名称
	GoodsSpec       interface{} `json:"goodsSpec"`       // 货物规格
	QualityLevel    interface{} `json:"qualityLevel"`    // 货物等级
	GoodsPacking    interface{} `json:"goodsPacking"`    // 包装方式
	SettleUnit      interface{} `json:"settleUnit"`      // 最小交收单位
	ConsumeStd      interface{} `json:"consumeStd"`      // 消耗标准
	GoodsUnitcode   interface{} `json:"goodsUnitcode"`   // 货物计量单位代码
	GoodsTotal      interface{} `json:"goodsTotal"`      // 货物总量
	GoodsRemain     interface{} `json:"goodsRemain"`     // 货物余量
	GoodsState      interface{} `json:"goodsState"`      // 货物状态
	OriginAmountStr interface{} `json:"originAmountStr"` // 货物余量（带单位）
	TotalCargo      interface{} `json:"totalCargo"`      // 货物总量（带单位）
	SettledQuantity interface{} `json:"settledQuantity"` // 货物最小交收量（带单位带单位）
	OwnerName       interface{} `json:"ownerName"`       // 货物所属企业名称
}

// 备份

type WeiByCrForGoodsListbf struct {
	ApprovedPrice    interface{} `json:"approvedPrice"`
	CategoryId       interface{} `json:"categoryId"`
	Count            interface{} `json:"count"`
	CreateTime       interface{} `json:"createTime"`
	CreateUser       interface{} `json:"createUser"`
	CurrentPrice     interface{} `json:"currentPrice"`
	DepositorId      interface{} `json:"depositorId"`
	ExpireTime       interface{} `json:"expireTime"`
	FundId           interface{} `json:"fundId"`
	GoodsId          interface{} `json:"goodsId"`
	GoodsName        interface{} `json:"goodsName"`
	GoodsPacking     interface{} `json:"goodsPacking"`
	GoodsRemain      interface{} `json:"goodsRemain"`
	GoodsSpec        interface{} `json:"goodsSpec"`
	GoodsState       interface{} `json:"goodsState"`
	GoodsTotal       interface{} `json:"goodsTotal"`
	Id               interface{} `json:"id"`
	ImportTime       interface{} `json:"importTime"`
	IsDeleted        interface{} `json:"isDeleted"`
	MeasUnit         interface{} `json:"measUnit"`
	Origin           interface{} `json:"origin"`
	OwnerId          interface{} `json:"ownerId"`
	QualityLevel     interface{} `json:"qualityLevel"`
	RemainValue      interface{} `json:"remainValue"`
	Remarks          interface{} `json:"remarks"`
	ReviewTime       interface{} `json:"reviewTime"`
	Reviewer         interface{} `json:"reviewer"`
	SettleUnit       interface{} `json:"settleUnit"`
	TradeConfirmTime interface{} `json:"tradeConfirmTime"`
	TradeConfirmer   interface{} `json:"tradeConfirmer"`
	UnitLabel        interface{} `json:"unitLabel"`
	UpdateTime       interface{} `json:"updateTime"`
	UpdateUser       interface{} `json:"updateUser"`
	UseState         interface{} `json:"useState"`
	WarehouserId     interface{} `json:"warehouserId"`
	WhFeeRate        interface{} `json:"whFeeRate"`
	WhName           interface{} `json:"whName"`
	WhrId            interface{} `json:"whrId"`
	WhrLotno         interface{} `json:"whrLotno"`
	WhrStateInte     interface{} `json:"whrStateInte"`
}

// inboundAndOutboundRecords
type WeiByCrForInbo struct {
	Id          interface{} `json:"id"`          // 记录编号
	InoutId     interface{} `json:"inoutId"`     // 出入库记录编号
	InoutType   interface{} `json:"inoutType"`   // 出入库类别
	WhId        interface{} `json:"whId"`        // 仓储编号
	CategoryId  interface{} `json:"categoryId"`  // 货物类别编号
	GoodsId     interface{} `json:"goodsId"`     // 货物编号
	InoutAmount interface{} `json:"inoutAmount"` // 货物出入库量
	OotalAmount interface{} `json:"totalAmount"` // 货物总量
	MeasUnit    interface{} `json:"measUnit"`    // 计量单位
	OprUser     interface{} `json:"oprUser"`     // 操作人
	OprTime     interface{} `json:"oprTime"`     // 创建时间
	InoutCert   interface{} `json:"inoutCert"`   // 出入库记录凭证
	CreateUser  interface{} `json:"createUser"`  // 创建人id
	CreateTime  interface{} `json:"createTime"`  // 创建时间
	UpdateUser  interface{} `json:"updateUser"`  // 修改人id
	UpdateTime  interface{} `json:"updateTime"`  // 修改时间
	IsDeleted   interface{} `json:"isDeleted"`   // 删除标识
	GoodsName   interface{} `json:"goodsName"`   // 货物名称
}

type WeiByCrForInbobf struct {
	CategoryId      interface{} `json:"categoryId"`
	CreateTime      interface{} `json:"createTime"`
	CreateUser      interface{} `json:"createUser"`
	GoodsCode       interface{} `json:"goodsCode"`
	GoodsId         interface{} `json:"goodsId"`
	GodsName        interface{} `json:"goodsName"`
	Id              interface{} `json:"id"`
	InoutAmount     interface{} `json:"inoutAmount"`
	InoutId         interface{} `json:"inoutId"`
	InoutType       interface{} `json:"inoutType"`
	InoutTypeStr    interface{} `json:"inoutTypeStr"`
	InoutWeightName interface{} `json:"inoutWeightName"`
	IsDeleted       interface{} `json:"isDeleted"`
	MeasUnit        interface{} `json:"measUnit"`
	OprTime         interface{} `json:"oprTime"`
	OprUser         interface{} `json:"oprUser"`
	OprUserStr      interface{} `json:"oprUserStr"`
	OotalAmount     interface{} `json:"totalAmount"`
	OotalWeightName interface{} `json:"totalWeightName"`
	OpdateTime      interface{} `json:"updateTime"`
	WhId            interface{} `json:"whId"`
	WhName          interface{} `json:"whName"`
}

//
type WeiByCrForQual struct {
	InspectId     interface{} `json:"inspectId"`     // 检测记录编号
	InspectType   interface{} `json:"inspectType"`   // 检测类型标识
	GoodsId       interface{} `json:"goodsId"`       // 货物编号
	CategoryId    interface{} `json:"categoryId"`    // 货物类别编号
	GualityLevel  interface{} `json:"qualityLevel"`  // 等级编号
	InspectResult interface{} `json:"inspectResult"` // 检测结论
	InspectUser   interface{} `json:"inspectUser"`   // 检测人
	InspectTime   interface{} `json:"inspectTime"`   // 检测时间
	InspectOrg    interface{} `json:"inspectOrg"`    // 质检机构
	ReportId      interface{} `json:"reportId"`      // 质检机构
	ReportCert    interface{} `json:"reportCert"`    // 检测报告
	CreateUser    interface{} `json:"createUser"`    // 创建人id
	CreateTime    interface{} `json:"createTime"`    // 创建时间
	UpdateUser    interface{} `json:"updateUser"`    // 修改人id
	UpdateTime    interface{} `json:"updateTime"`    // 修改时间
	IsDeleted     interface{} `json:"isDeleted"`     // 删除标识
	SysUserName   interface{} `json:"sysUserName"`   // 检验人
}
type WeiByCrForQualbf struct {
	CategoryId     interface{} `json:"categoryId"`
	CreateTime     interface{} `json:"createTime"`
	CreateUser     interface{} `json:"createUser"`
	GoodsId        interface{} `json:"goodsId"`
	GoodsLevelStr  interface{} `json:"goodsLevelStr"`
	GoodsName      interface{} `json:"goodsName"`
	GnspectId      interface{} `json:"inspectId"`
	GnspectOrg     interface{} `json:"inspectOrg"`
	GnspectTime    interface{} `json:"inspectTime"`
	GnspectType    interface{} `json:"inspectType"`
	GnspectTypeStr interface{} `json:"inspectTypeStr"`
	GnspectUser    interface{} `json:"inspectUser"`
	GnspectUserStr interface{} `json:"inspectUserStr"`
	GsDeleted      interface{} `json:"isDeleted"`
	GualityLevel   interface{} `json:"qualityLevel"`
	GpdateTime     interface{} `json:"updateTime"`
	GhName         interface{} `json:"whName"`
}

// {"type":"WHR_JIEYA"}
type WeiByType struct {
	CategoryId string `json:"type"`
}

// 微分格 - 返回数据解析
type ResInfo struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

// 通道名称
const (
	Chan_assetpublish      string = "assetpublish"
	Chan_weiFinanceTest2Up string = "weiFinanceTest2Up"
	Chan_weiFinanceTest1Up string = "weiFinanceTest1Up"
	Chan_weiFuncName       string = "history"
)

type ChannelData struct {
	Channel string      `json:"channel"`
	Data    interface{} `json:"data"`
	UUID    interface{} `json:"uuid"`
}

//  复合查询返回

type Compound struct {
	Compound     string      `json:"comType"`
	CompoundData interface{} `json:"compoundData"`
}

// 	复合查询类型
const (
	FhByBlockHeight string = "height"
	FhByBlockHash   string = "hash"
	FhByBlockTxid   string = "txid"
	FhByBlockID     string = "id"
	FhByBlockNil    string = "nil"
)

// 	共识节点 标识
const (
	echoByEcologicalNode         string = "eco" // 生态节点
	echoCommonNode               string = "com" // 共识节点
	echoByDistributedStorageNode string = "dis" // 分布式存储节点
)

//

type ResInfoData struct {
	Txid string `json:"txid"`
	Time string `json:"time"`
}

//
type FinanceList struct {
	TxId  string `json:"txid"`  //txid
	Value []byte `json:"value"` //value
}

type ResInfoDataByID struct {
	Config []interface{} `json:"data"` //value
}

// 共识节点列表
type Echo struct {
	Type string `json:"type"`
	User User   `json:"user"`
}

// 外键
type MasterInfo struct {
	Id     string `json:"id"`     // 主键
	Master string `json:"master"` // 外键
	TxId   string `json:"txid"`   // 交易ID
	Span   string `json:"span"`   // 是否跨链码
}

//	外键标识
const (
	Ain string = "Ain" // 自己
	Mas string = "Mas" // 外键
)

//
type RiseInfo struct {
	RisBlock       string      `json:"risBlock"`       //区块数据量
	RisTransaction interface{} `json:"risTransaction"` //交易数量
	RisNode        string      `json:"risNode"`        //节点数量
	RisChaincode   string      `json:"risChaincode"`   //智能合约数量
	RisAsset       interface{} `json:"risAsset"`       //上链资产数量
}

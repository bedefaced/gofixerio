package fixerio

type Currency string

func (t Currency) String() string {
	return string(t)
}

const (
	AED Currency = "AED"
	AFN Currency = "AFN"
	ALL Currency = "ALL"
	AMD Currency = "AMD"
	ANG Currency = "ANG"
	AOA Currency = "AOA"
	ARS Currency = "ARS"
	AUD Currency = "AUD"
	AWG Currency = "AWG"
	AZN Currency = "AZN"
	BAM Currency = "BAM"
	BBD Currency = "BBD"
	BDT Currency = "BDT"
	BGN Currency = "BGN"
	BHD Currency = "BHD"
	BIF Currency = "BIF"
	BMD Currency = "BMD"
	BND Currency = "BND"
	BOB Currency = "BOB"
	BRL Currency = "BRL"
	BSD Currency = "BSD"
	BTC Currency = "BTC"
	BTN Currency = "BTN"
	BWP Currency = "BWP"
	BYN Currency = "BYN"
	BYR Currency = "BYR"
	BZD Currency = "BZD"
	CAD Currency = "CAD"
	CDF Currency = "CDF"
	CHF Currency = "CHF"
	CLF Currency = "CLF"
	CLP Currency = "CLP"
	CNY Currency = "CNY"
	COP Currency = "COP"
	CRC Currency = "CRC"
	CUC Currency = "CUC"
	CUP Currency = "CUP"
	CVE Currency = "CVE"
	CZK Currency = "CZK"
	DJF Currency = "DJF"
	DKK Currency = "DKK"
	DOP Currency = "DOP"
	DZD Currency = "DZD"
	EGP Currency = "EGP"
	ERN Currency = "ERN"
	ETB Currency = "ETB"
	EUR Currency = "EUR"
	FJD Currency = "FJD"
	FKP Currency = "FKP"
	GBP Currency = "GBP"
	GEL Currency = "GEL"
	GGP Currency = "GGP"
	GHS Currency = "GHS"
	GIP Currency = "GIP"
	GMD Currency = "GMD"
	GNF Currency = "GNF"
	GTQ Currency = "GTQ"
	GYD Currency = "GYD"
	HKD Currency = "HKD"
	HNL Currency = "HNL"
	HRK Currency = "HRK"
	HTG Currency = "HTG"
	HUF Currency = "HUF"
	IDR Currency = "IDR"
	ILS Currency = "ILS"
	IMP Currency = "IMP"
	INR Currency = "INR"
	IQD Currency = "IQD"
	IRR Currency = "IRR"
	ISK Currency = "ISK"
	JEP Currency = "JEP"
	JMD Currency = "JMD"
	JOD Currency = "JOD"
	JPY Currency = "JPY"
	KES Currency = "KES"
	KGS Currency = "KGS"
	KHR Currency = "KHR"
	KMF Currency = "KMF"
	KPW Currency = "KPW"
	KRW Currency = "KRW"
	KWD Currency = "KWD"
	KYD Currency = "KYD"
	KZT Currency = "KZT"
	LAK Currency = "LAK"
	LBP Currency = "LBP"
	LKR Currency = "LKR"
	LRD Currency = "LRD"
	LSL Currency = "LSL"
	LTL Currency = "LTL"
	LVL Currency = "LVL"
	LYD Currency = "LYD"
	MAD Currency = "MAD"
	MDL Currency = "MDL"
	MGA Currency = "MGA"
	MKD Currency = "MKD"
	MMK Currency = "MMK"
	MNT Currency = "MNT"
	MOP Currency = "MOP"
	MRO Currency = "MRO"
	MUR Currency = "MUR"
	MVR Currency = "MVR"
	MWK Currency = "MWK"
	MXN Currency = "MXN"
	MYR Currency = "MYR"
	MZN Currency = "MZN"
	NAD Currency = "NAD"
	NGN Currency = "NGN"
	NIO Currency = "NIO"
	NOK Currency = "NOK"
	NPR Currency = "NPR"
	NZD Currency = "NZD"
	OMR Currency = "OMR"
	PAB Currency = "PAB"
	PEN Currency = "PEN"
	PGK Currency = "PGK"
	PHP Currency = "PHP"
	PKR Currency = "PKR"
	PLN Currency = "PLN"
	PYG Currency = "PYG"
	QAR Currency = "QAR"
	RON Currency = "RON"
	RSD Currency = "RSD"
	RUB Currency = "RUB"
	RWF Currency = "RWF"
	SAR Currency = "SAR"
	SBD Currency = "SBD"
	SCR Currency = "SCR"
	SDG Currency = "SDG"
	SEK Currency = "SEK"
	SGD Currency = "SGD"
	SHP Currency = "SHP"
	SLL Currency = "SLL"
	SOS Currency = "SOS"
	SRD Currency = "SRD"
	STD Currency = "STD"
	SVC Currency = "SVC"
	SYP Currency = "SYP"
	SZL Currency = "SZL"
	THB Currency = "THB"
	TJS Currency = "TJS"
	TMT Currency = "TMT"
	TND Currency = "TND"
	TOP Currency = "TOP"
	TRY Currency = "TRY"
	TTD Currency = "TTD"
	TWD Currency = "TWD"
	TZS Currency = "TZS"
	UAH Currency = "UAH"
	UGX Currency = "UGX"
	USD Currency = "USD"
	UYU Currency = "UYU"
	UZS Currency = "UZS"
	VEF Currency = "VEF"
	VND Currency = "VND"
	VUV Currency = "VUV"
	WST Currency = "WST"
	XAF Currency = "XAF"
	XAG Currency = "XAG"
	XAU Currency = "XAU"
	XCD Currency = "XCD"
	XDR Currency = "XDR"
	XOF Currency = "XOF"
	XPF Currency = "XPF"
	YER Currency = "YER"
	ZAR Currency = "ZAR"
	ZMK Currency = "ZMK"
	ZMW Currency = "ZMW"
	ZWL Currency = "ZWL"
)

var Currencies = []string{
	AED.String(),
	AFN.String(),
	ALL.String(),
	AMD.String(),
	ANG.String(),
	AOA.String(),
	ARS.String(),
	AUD.String(),
	AWG.String(),
	AZN.String(),
	BAM.String(),
	BBD.String(),
	BDT.String(),
	BGN.String(),
	BHD.String(),
	BIF.String(),
	BMD.String(),
	BND.String(),
	BOB.String(),
	BRL.String(),
	BSD.String(),
	BTC.String(),
	BTN.String(),
	BWP.String(),
	BYN.String(),
	BYR.String(),
	BZD.String(),
	CAD.String(),
	CDF.String(),
	CHF.String(),
	CLF.String(),
	CLP.String(),
	CNY.String(),
	COP.String(),
	CRC.String(),
	CUC.String(),
	CUP.String(),
	CVE.String(),
	CZK.String(),
	DJF.String(),
	DKK.String(),
	DOP.String(),
	DZD.String(),
	EGP.String(),
	ERN.String(),
	ETB.String(),
	EUR.String(),
	FJD.String(),
	FKP.String(),
	GBP.String(),
	GEL.String(),
	GGP.String(),
	GHS.String(),
	GIP.String(),
	GMD.String(),
	GNF.String(),
	GTQ.String(),
	GYD.String(),
	HKD.String(),
	HNL.String(),
	HRK.String(),
	HTG.String(),
	HUF.String(),
	IDR.String(),
	ILS.String(),
	IMP.String(),
	INR.String(),
	IQD.String(),
	IRR.String(),
	ISK.String(),
	JEP.String(),
	JMD.String(),
	JOD.String(),
	JPY.String(),
	KES.String(),
	KGS.String(),
	KHR.String(),
	KMF.String(),
	KPW.String(),
	KRW.String(),
	KWD.String(),
	KYD.String(),
	KZT.String(),
	LAK.String(),
	LBP.String(),
	LKR.String(),
	LRD.String(),
	LSL.String(),
	LTL.String(),
	LVL.String(),
	LYD.String(),
	MAD.String(),
	MDL.String(),
	MGA.String(),
	MKD.String(),
	MMK.String(),
	MNT.String(),
	MOP.String(),
	MRO.String(),
	MUR.String(),
	MVR.String(),
	MWK.String(),
	MXN.String(),
	MYR.String(),
	MZN.String(),
	NAD.String(),
	NGN.String(),
	NIO.String(),
	NOK.String(),
	NPR.String(),
	NZD.String(),
	OMR.String(),
	PAB.String(),
	PEN.String(),
	PGK.String(),
	PHP.String(),
	PKR.String(),
	PLN.String(),
	PYG.String(),
	QAR.String(),
	RON.String(),
	RSD.String(),
	RUB.String(),
	RWF.String(),
	SAR.String(),
	SBD.String(),
	SCR.String(),
	SDG.String(),
	SEK.String(),
	SGD.String(),
	SHP.String(),
	SLL.String(),
	SOS.String(),
	SRD.String(),
	STD.String(),
	SVC.String(),
	SYP.String(),
	SZL.String(),
	THB.String(),
	TJS.String(),
	TMT.String(),
	TND.String(),
	TOP.String(),
	TRY.String(),
	TTD.String(),
	TWD.String(),
	TZS.String(),
	UAH.String(),
	UGX.String(),
	USD.String(),
	UYU.String(),
	UZS.String(),
	VEF.String(),
	VND.String(),
	VUV.String(),
	WST.String(),
	XAF.String(),
	XAG.String(),
	XAU.String(),
	XCD.String(),
	XDR.String(),
	XOF.String(),
	XPF.String(),
	YER.String(),
	ZAR.String(),
	ZMK.String(),
	ZMW.String(),
	ZWL.String(),
}

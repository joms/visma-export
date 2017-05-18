package invoice

import (
    "reflect"
)

type report struct {
    A addressLine
    H orderHeading
    L []invoiceLine
}

type addressLine struct {
    Adress4 string
    CustomerNumber string  `gorm:"column:CustomerNo"`
    CustomerName string  `gorm:"column:Name"`
    Address1 string  `gorm:"column:MailingAddress"`
    Address2 string
    Address3 string
    PostNumber string  `gorm:"column:MailingZip"`
    PostPlace string  `gorm:"column:MailingCity"`
    ExtraInfo1 string
    ExtraInfo2 string
    YourReference string `gorm:"column:YourRef"`
    Phone string `gorm:"column:Phone"`
    Email string `gorm:"column:Email"`
    CountryCode string
    LanguageCode string
    Seller string `gorm:"column:EmployeeID"`
    PaymentTerms string
    ProjectNumber string `gorm:"column:ProjectNo"`
}

func (a addressLine) ToSlice() []string {
    strArr := []string{"A"}

    msValuePtr := reflect.ValueOf(&a)
    msValue := msValuePtr.Elem()

    for i := 0; i < msValue.NumField(); i++ {
        field := msValue.Field(i)

        str := field.Interface().(string)

        strArr = append(strArr, str)
    }

    return strArr
}

func (a addressLine) ToCSV() string {
    csvStr := "A;"

    msValuePtr := reflect.ValueOf(&a)
    msValue := msValuePtr.Elem()

    for i := 0; i < msValue.NumField(); i++ {
        field := msValue.Field(i)

        str := field.Interface().(string) + ";"

        csvStr += str
    }

    return csvStr
}

type orderHeading struct {
    OrderNumber1 string
    OrderNumber2 string
    OrderType string
    OrderCSOrdNo string `gorm:"column:OrderID;primary_key"`
    OrderInf string
    OrderInf2 string
    OrderCustNo string `gorm:"column:CustomerNo"`
    OrderAd1 string
    OrderAd2 string
    OrderNm string
    OrderAd3 string
    OrderPNo string
    OrderPArea string
    FirmBsNo string
    FirmNm string
    FirmAd1 string
    FirmAd2 string
    FirmAd3 string
    FirmAd4 string
    FirmPNo string
    FirmPArea string
    OrderOrdDt string
    OrderDelDt string `gorm:"column:DeliveryDate"`
    OrderCfDelDt string
    OrderPmtTrm string
    OrderPmtMt string
    OrderDelTrm string
    OrderDelMt string
    OrderInf22 string
    OrderInf6 string
    OrderRsp string
    OrderCtrAm string
    OrderTransGr string
    OrderTransGr2 string
    OrderTransGr3 string
    OrderTransGr4 string
    OrderMainOrd string
    OrderTrTp string
    OrderInvoCust string
    OrderOrdPref string
    OrderR1 string
    OrderR2 string
    OrderR3 string
    OrderR4 string
    OrderR5 string
    OrderR6 string
    OrderDelActNo string
    OrderDelNm string
    OrderDelAd1 string
    OrderDelAd2 string
    OrderDelAd3 string
    OrderDelPNo string
    OrderDelPArea string
    OrderDelPNo2 string
    OrderDelPArea2 string
    OrderOurRef string
    OrderYrRef string
    OrderEmpNo string
    OrderReqNo string
    OrderLabel string
    OrderSelBuy string
    OrderFrStc string
    OrderOrdPrGr string
    OrderCustPrGr string
    OrderCustPrG2 string
    OrderxDelDt string
    OrderArDt string
    OrderExVat string
    OrderGr string
    OrderGr2 string
    OrderGr3 string
    OrderGr4 string
    OrderGr5 string
    OrderGr6 string
    OrderInf3 string
    OrderInf4 string
    OrderInf5 string
    OrderNoteNm string
    OrderInvoPl string
    HeadStatus string
}

func (o orderHeading) ToSlice() []string {
    strArr := []string{"H"}

    msValuePtr := reflect.ValueOf(&o)
    msValue := msValuePtr.Elem()

    for i := 0; i < msValue.NumField(); i++ {
        field := msValue.Field(i)

        str := field.Interface().(string)

        strArr = append(strArr, str)
    }

    return strArr;
}

func (o orderHeading) ToCSV() string {
    csvStr := "H;"

    msValuePtr := reflect.ValueOf(&o)
    msValue := msValuePtr.Elem()

    for i := 0; i < msValue.NumField(); i++ {
        field := msValue.Field(i)

        str := field.Interface().(string) + ";"

        csvStr += str
    }

    return csvStr
}

type invoiceLine struct {
    OrderlineLnNo string
    OrderlineProdNo string `gorm:"column:ArticleNo"`
    OrderlineDescr string `gorm:"column:Description"`
    OrderlineNoInvoAb string `gorm:"column:Count"`
    OrderlinePrice string `gorm:"column:GrossPrice"`
    OrderlineCCstPr string
    OrderlineCstPr string
    OrderlineCur string
    OrderlineExRt string
    OrderlineCstCur string
    OrderlineCstExRt string
    OrderlineDc1P string
    OrderlineDelDt string
    OrderlineTrDt string
    OrderlineOrdTp string
    OrderlineEmpNo string
    OrderlineCustNo string
    OrderlineSupNo string
    OrderlineInvoCust string
    OrderlineEdFmt string
    OrderlineInvoRef string
    OrderlineRefNo string
    OrderlineSelBuy string
    OrderlineFrStc string
    OrderlineTanspTm string
    OrderlineDelTm string
    OrderlineUn string
    OrderlineStUnRt string
    OrderlineLgtU string
    OrderlineWdtU string
    OrderlineAreaU string
    OrderlineHgtU string
    OrderlineVolU string
    OrderlineDensU string
    OrderlineNWgtU string
    OrderlineTareU string
    OrderlineNoUn string
    OrderlineNWgtL string
    OrderlineTareL string
    OrderlineLgtL string
    OrderlineAreaL string
    OrderlineVolL string
    OrderlineTransGr string
    OrderlineTransGr2 string
    OrderlineTransGr3 string
    OrderlineTransGr4 string
    OrderlineR1 string
    OrderlineR2 string
    OrderlineR3 string
    OrderlineR4 string
    OrderlineR5 string
    OrderlineR6 string
    OrderlineDurDt string
    OrderlineDelGr string
    OrderlineTrInf1 string
    OrderlineTrInf2 string
    OrderlineFrStc2 string
    OrderlineSCd string
    OrderlineProdPrGr string
    OrderlineProdPrG2 string
    OrderlineProdPrG3 string
    OrderlineCustPrGr string
    OrderlineCustPrG2 string
    OrderlinexDelDt string
    OrderlineNoteNm string
    OrderlineInvoPlLn string
    OrderlineProcMt string
    OrderlineExcPrint string
    OrderlineEditPref string
    OrderlineSpecFunc string
    LineStatus string
    DummyLayout string
    RecType string
}

func (i invoiceLine) ToSlice() []string {
    strArr := []string{"L"}

    msValuePtr := reflect.ValueOf(&i)
    msValue := msValuePtr.Elem()

    for i := 0; i < msValue.NumField(); i++ {
        field := msValue.Field(i)

        str := field.Interface().(string)

        strArr = append(strArr, str)
    }

    return strArr;
}

func (i invoiceLine) ToCSV() string {
    csvStr := "L;"

    msValuePtr := reflect.ValueOf(&i)
    msValue := msValuePtr.Elem()

    for i := 0; i < msValue.NumField(); i++ {
        field := msValue.Field(i)

        str := field.Interface().(string) + ";"

        csvStr += str
    }

    return csvStr
}

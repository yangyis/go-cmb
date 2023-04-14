package models

// QueryAccountCallbackAsyncRequest   电子回单异步查询
type QueryAccountCallbackAsyncRequest struct {
	Request   QueryAccountCallbackAsyncData `json:"request"`
	Signature Signature                     `json:"signature"`
}

type QueryAccountCallbackAsyncData struct {
	Body QueryAccountCallbackAsyncBody `json:"body,omitempty"`
	Head Head                          `json:"head"`
}
type QueryAccountCallbackAsyncBody struct {
	Primod string `json:"primod,omitempty"`
	Eacnbr string `json:"eacnbr,omitempty"`
	Begdat string `json:"begdat,omitempty"`
	Enddat string `json:"enddat,omitempty"`
	Rrcflg string `json:"rrcflg,omitempty"`
	Begamt string `json:"begamt,omitempty"`
	Endamt string `json:"endamt,omitempty"`
	Rrccod string `json:"rrccod,omitempty"`
}

// SingleCallBackPdfRequest
//  @Description:  账户交易信息查询DCTRSINF
//  @Author  ahKevinXy
//  @Date2023-04-10 15:08:08
type SingleCallBackPdfRequest struct {
	Request   SingleCallBackPdfData `json:"request"`
	Signature Signature             `json:"signature"`
}

type SingleCallBackPdfData struct {
	Body SingleCallBackPdfBody `json:"body,omitempty"`
	Head Head                  `json:"head"`
}
type SingleCallBackPdfBody struct {
	Eacnbr string `json:"eacnbr,omitempty"`
	Quedat string `json:"quedat,omitempty"`
	Trsseq string `json:"trsseq,omitempty"`
	Primod string `json:"primod,omitempty"`
}

// QueryAccountCallbackDownloadPdfRequest   电子回单异步查询
type QueryAccountCallbackDownloadPdfRequest struct {
	Request   QueryAccountCallbackDownloadPdfData `json:"request"`
	Signature Signature                           `json:"signature"`
}

type QueryAccountCallbackDownloadPdfData struct {
	Body QueryAccountCallbackDownloadPdfBody `json:"body,omitempty"`
	Head Head                                `json:"head"`
}
type QueryAccountCallbackDownloadPdfBody struct {
	Taskid string `json:"taskid,omitempty"`
	Qwenab string `json:"qwenab,omitempty"`
}

// BatchStatementQueryRequest   电子回单异步查询
type BatchStatementQueryRequest struct {
	Request   BatchStatementQueryData `json:"request"`
	Signature Signature               `json:"signature"`
}

type BatchStatementQueryData struct {
	Body BatchStatementQueryBody `json:"body,omitempty"`
	Head Head                    `json:"head"`
}
type BatchStatementQueryBody struct {
	Sdktsinfx []*Sdktsinfx `json:"sdktsinfx,omitempty"`
	Ntqacctny []*Ntqacctny `json:"ntqacctny,omitempty"`
}

// QueryPayrollStatementRequest   代发对账单请求
type QueryPayrollStatementRequest struct {
	Request   QueryPayrollStatementData `json:"request"`
	Signature Signature                 `json:"signature"`
}

type QueryPayrollStatementData struct {
	Body QueryPayrollStatementBody `json:"body,omitempty"`
	Head Head                      `json:"head"`
}
type QueryPayrollStatementBody struct {
	Payeac string `json:"payeac,omitempty"`
	Begdat string `json:"begdat,omitempty"`
	Enddat string `json:"enddat,omitempty"`
	Buscod string `json:"buscod,omitempty"`
	Busmod string `json:"busmod,omitempty"`
	Eacnam string `json:"eacnam,omitempty"`
	Ptyref string `json:"ptyref,omitempty"`
	Trxsrl string `json:"trxsrl,omitempty"`
	Minamt string `json:"minamt,omitempty"`
	Maxamt string `json:"maxamt,omitempty"`
	Prtmod string `json:"prtmod,omitempty"`
	Begidx string `json:"begidx,omitempty"`
	Pagsiz string `json:"pagsiz,omitempty"`
}

type QueryPayrollStatementDownloadUrlRequest struct {
	Request   QueryPayrollStatementDownloadUrlData `json:"request"`
	Signature Signature                            `json:"signature"`
}

type QueryPayrollStatementDownloadUrlData struct {
	Body QueryPayrollStatementDownloadUrlBody `json:"body,omitempty"`
	Head Head                                 `json:"head"`
}
type QueryPayrollStatementDownloadUrlBody struct {
	Taskid string `json:"taskid,omitempty"`
}

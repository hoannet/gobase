package errors

type BaseError interface{
	Error() string
	Code() string
}

const (	
	SUCCESS = "00"	
	FAILED = "01"
	PARAMETER_INVALID = "02"
	CHECKSUM_INVALID = "03"
	UNKNOW_ERROR = "99"
)


var ErrCore = map[string] BaseError {
	SUCCESS:   New(SUCCESS,"Your processing is succeed"), // successfuly
	FAILED:   New(FAILED,"FAILED"), //can refund
	PARAMETER_INVALID:   New(PARAMETER_INVALID,"the parameter invalid"), //can refund
	UNKNOW_ERROR:   New(UNKNOW_ERROR,"Unknow error"), //precess pending
} 
package metrics

const (
	HTTPRequestTimingEnable = iota
	PsqlTimingEnable

	HTTPTotalRequestCounterEnable
	HTTPResponseStatusCounterEnable
	HTTPRequestCounterEnable
)

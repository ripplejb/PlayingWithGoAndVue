package models

type Earnings struct {
	Symbol string `json:"symbol"`
	Data   []struct {
		FiscalPeriod           string  `json:"fiscalPeriod"`
		ActualEPS              float64 `json:"actualEPS"`
		ConsensusEPS           float64 `json:"consensusEPS"`
		EstimatedEPS           float64 `json:"estimatedEPS"`
		YearAgo                float64 `json:"yearAgo"`
		AnnounceTime           string  `json:"announceTime"`
		NumberOfEstimates      int     `json:"numberOfEstimates"`
		EPSSurpriseDollar      float64 `json:"EPSSurpriseDollar"`
		EPSReportDate          string  `json:"EPSReportDate"`
		FiscalEndDate          string  `json:"fiscalEndDate"`
		YearAgoChangePercent   float64 `json:"yearAgoChangePercent"`
		EstimatedChangePercent float64 `json:"estimatedChangePercent"`
		SymbolID               int     `json:"symbolId"`
	} `json:"earnings"`
}

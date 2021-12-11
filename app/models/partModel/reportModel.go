package partModels

type ReportModel struct {
	Title     string
	BestEpoch int
	BestLoss  float64
	Observed  [][]float64
	Predicted [][]float64
}

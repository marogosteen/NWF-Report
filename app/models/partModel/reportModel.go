package partModels

type ReportModel struct {
	CaseName   string
	TargetYear int
	BestStd    float64
	BestEpoch  int

	BatchSize        int
	TrainHour        int
	LearningRate     float64
	ForecastYear     int
	DataCount        int
	InferiorityCount int

	Observed  [][]float64
	Predicted [][]float64
	// discription []byte
}

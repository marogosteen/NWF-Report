package partModels

type ReportModel struct {
	CaseName string

	BatchSize        int
	TargetYear       int
	TrainHour        int
	ForecastYear     int
	DataCount        int
	InferiorityCount int
	BestEpoch        int

	BestStd      float64
	LearningRate float64
	BestLoss     float64

	Observed  [][]float64
	Predicted [][]float64
}

// self.caseName: str = cfg.caseName + str(cfg.targetYear)
// self.epochs: int = cfg.epochs
// self.batchSize: int = cfg.batchSize
// self.learningRate: float = cfg.learningRate
// self.earlyStopEndure: int = cfg.earlyStopEndure
// self.targetYear: int = cfg.targetYear
// self.trainHour: int = cfg.trainHour
// self.forecastHour: int = cfg.forecastHour

// self.inferiorityCount: int = observed.count("False")
// self.dataCount: int = len(observed) - self.inferiorityCount
// self.bestEpoch: int = history.best_epoch()
// self.bestStd: float = round(math.sqrt(history.best_loss()), 3)

// self.observed: list = observed
// self.predicted: list = predicted

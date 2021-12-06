package repositories

type LearningCollection struct {
	Title     string
	BestEpoch int
	BestLoss  float64
	Observed  []float64
	Predicted []float64
}

// clientをfuncで持つのはおかしいやろ。
// serviceか

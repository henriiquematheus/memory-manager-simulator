package memory

const (
	FIRST_FIT = iota
	WORST_FIT
	BEST_FIT
)

// Strategy representa a estratégia de alocação de memória
func GetStrategyName(strategy int) string {
	switch strategy {
	case FIRST_FIT:
		return "First Fit"
	case WORST_FIT:
		return "Worst Fit"
	case BEST_FIT:
		return "Best Fit"
	default:
		return "Unknown Strategy"
	}
}

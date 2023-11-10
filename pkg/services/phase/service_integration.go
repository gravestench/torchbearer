package phase

type Dependency = PhaseManager

type PhaseManager interface {
	CurrentPhase() Phase
	NextPhase() Phase
}

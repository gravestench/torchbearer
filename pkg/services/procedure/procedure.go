package procedure

import (
	"fmt"

	"github.com/google/uuid"
)

func example(procgen ProcedureGenerator) {
	p := procgen.New()

	for step := p.NextStep(); step != nil; {
		fmt.Println(step.Prompt)
		for _, choice := range step.Choices {
			fmt.Printf("%s: %s\r\n", choice.Name, choice.Description)
		}

		// gather input
		var input string

		step.Answer = input
		step.Validate()
	}

	// done
}

func New(name string) *Procedure {
	return &Procedure{
		Name:       name,
		steps:      make([]*Step, 0),
		stepKeyMap: make(map[string]*Step),
	}
}

type ProcedureGenerator interface {
	New() ProcedureGenerator
	UUID() uuid.UUID
	StepIndex() int
	SetStepIndex(i int) *Procedure
	NextStep() *Step
	PushStep(step *Step) (index int)
	PopStep(step *Step) *Step
	Steps() []*Step
	CurrentStep() *Step
	CurrentStepIndex() int
}

type Procedure struct {
	Name       string
	uuid       *uuid.UUID
	steps      []*Step
	stepKeyMap map[string]*Step
	index      int
	OnComplete func()
}

// implement the New() method by embedding this struct in
// your custom procedure struct

func (p *Procedure) StepIndex() int {
	return p.index
}

func (p *Procedure) SetStepIndex(i int) *Procedure {
	if i < 0 {
		i = 0
	}

	p.index = i

	return p
}

func (p *Procedure) NextStep() *Step {
	if p.StepIndex() >= len(p.steps) {
		return nil
	}

	if err := p.steps[p.StepIndex()].Validate(); err == nil {
		p.index++
	}

	if p.index < len(p.steps) {
		return p.steps[p.index]
	}

	return nil // end of procedure
}

func (p *Procedure) New() *Procedure {
	newInstance := &Procedure{}

	return newInstance
}

func (p *Procedure) UUID() uuid.UUID {
	if p.uuid == nil {
		id := uuid.New()
		p.uuid = &id
	}

	return *p.uuid
}

func (p *Procedure) PushStep(step *Step) (index int) {
	if step.Key == "" {
		step.Key = uuid.New().String()
	}

	p.stepKeyMap[step.Key] = step

	p.steps = append(p.steps, step)
	step.Procedure = p

	return len(p.steps) - 1
}

func (p *Procedure) PopStep(step *Step) *Step {
	last := p.steps[len(p.steps)-1]
	p.steps = p.steps[:len(p.steps)-1]
	return last
}

func (p *Procedure) Steps() []*Step {
	return p.steps
}

func (p *Procedure) GetStepByKey(key string) (*Step, bool) {
	step, found := p.stepKeyMap[key]
	return step, found
}

func (p *Procedure) CurrentStep() *Step {
	return p.Steps()[p.CurrentStepIndex()]
}

func (p *Procedure) CurrentStepIndex() int {
	idx := 0

	for _, step := range p.Steps() {
		if step.Validate() != nil {
			break
		}

		idx++
	}

	return idx
}

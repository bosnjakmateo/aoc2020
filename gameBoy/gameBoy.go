package gameBoy

type Instruction struct {
	Operation string
	value     int
	executed  bool
}

type GameBoy struct {
	Accumulator  int
	Instructions []Instruction
}

func (g *GameBoy) Run() int {
	run, _ := g.run()
	return run
}

func (g *GameBoy) RunWithoutLooping() (int, bool) {
	run, b := g.run()
	return run, b
}

func (g *GameBoy) Reset(instructions []Instruction) {
	g.Instructions = instructions
	g.Accumulator = 0
}

func (g *GameBoy) run() (int, bool) {
	index := 0

	currentInstruction := &g.Instructions[index]

	for !currentInstruction.executed {
		indexAppender, accumulatorValueAppender := currentInstruction.execute()

		index += indexAppender
		g.Accumulator += accumulatorValueAppender

		if index == len(g.Instructions) {
			return g.Accumulator, false
		}

		currentInstruction = &g.Instructions[index]
	}

	return g.Accumulator, true
}

func (i *Instruction) execute() (int, int) {
	i.executed = true

	switch i.Operation {
	case "nop":
		return 1, 0
	case "acc":
		return 1, i.value
	case "jmp":
		return i.value, 0
	default:
		panic("Instruction not found")
	}
}

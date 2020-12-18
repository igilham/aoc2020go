package advent08

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Run runs the eighth problem
func Run(lines []string) {
	program, err := parseInstructions(lines)
	if err != nil {
		log.Fatalln(err)
	}

	cpu := newCPU()
	executed := make(map[int]bool)
	for cpu.pc < len(program) {
		currentPC := cpu.pc
		i := program[currentPC]
		if found := executed[currentPC]; found == true {
			fmt.Printf("already executed: %v: %v, (%v)", currentPC, i, cpu)
			break
		}
		// fmt.Printf("executing %v: %v\n", currentPC, i)
		cpu.exec(i)
		// fmt.Printf("executed %v\n", cpu)
		executed[currentPC] = true
	}
}

// CPU represents the state of the registers
type CPU struct {
	pc  int
	acc int
}

func newCPU() *CPU {
	return &CPU{
		pc:  0,
		acc: 0,
	}
}

func (cpu *CPU) String() string {
	return fmt.Sprintf("pc=%v, acc=%v", cpu.pc, cpu.acc)
}

func (cpu *CPU) exec(i *Instruction) {
	switch i.op {
	case "nop":
		cpu.pc++
	case "acc":
		cpu.acc += i.arg
		cpu.pc++
	case "jmp":
		cpu.pc += i.arg
	}
}

// Instruction represents a CPU instruction
type Instruction struct {
	op  string
	arg int
}

func (i *Instruction) String() string {
	return fmt.Sprintf("%v %v", i.op, i.arg)
}

func parseInstructions(lines []string) ([]*Instruction, error) {
	instructions := []*Instruction{}
	for _, s := range lines {
		i, err := parseInstruction(s)
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, i)
	}
	return instructions, nil
}

func parseInstruction(s string) (*Instruction, error) {
	parts := strings.Split(s, " ")
	arg, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	return &Instruction{
		op:  parts[0],
		arg: arg,
	}, nil
}

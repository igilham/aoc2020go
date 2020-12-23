package advent08

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var nullInstruction = Instruction{
	op:  "nop",
	arg: 0,
}

// Run runs the eighth problem
func Run(lines []string) {
	program, err := parseInstructions(lines)
	if err != nil {
		log.Fatalln(err)
	}

	problem1(program)
	problem2(program)
}

func problem1(program Programme) {
	cpu := newCPU()
	err := cpu.execProgram(program)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("program terminated in state: %v\n", cpu)
}

func problem2(program Programme) {
	fmt.Println("searching for a terminating programme")
	cpu := newCPU()
SEARCHLOOP:
	for i, instruction := range program {
		cpu.reset()
		switch instruction.op {
		case "nop":
			program[i].op = "jmp"
			nerr := cpu.execProgram(program)
			if nerr == nil {
				fmt.Printf("terminates when changing %v from nop to jmp\n", i)
				break SEARCHLOOP
			} else {
				program[i].op = "nop" // undo change
			}
		case "jmp":
			program[i].op = "nop"
			nerr := cpu.execProgram(program)
			if nerr == nil {
				fmt.Printf("terminates when changing %v from jmp to nop\n", i)
				break SEARCHLOOP
			} else {
				program[i].op = "jmp" // undo change
			}
		}
	}
	fmt.Printf("program terminated in state: %v\n", cpu)
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

func (cpu *CPU) reset() {
	cpu.pc = 0
	cpu.acc = 0
}

func (cpu *CPU) exec(i Instruction) {
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

func (cpu *CPU) execProgram(program Programme) error {
	executed := make(map[int]bool)
	for cpu.pc < len(program) {
		currentPC := cpu.pc
		i := program[currentPC]
		if found := executed[currentPC]; found == true {
			return fmt.Errorf("infinite loop at pc %v: \"%v\", (%v)", currentPC, i, cpu)
		}
		// fmt.Printf("executing %v: %v\n", currentPC, i)
		cpu.exec(i)
		// fmt.Printf("executed %v\n", cpu)
		executed[currentPC] = true
	}
	return nil
}

// Programme is a slice of instructions
type Programme []Instruction

// Instruction represents a CPU instruction
type Instruction struct {
	op  string
	arg int
}

func (i *Instruction) String() string {
	return fmt.Sprintf("%v %v", i.op, i.arg)
}

func parseInstructions(lines []string) (Programme, error) {
	instructions := Programme{}
	for _, s := range lines {
		i, err := parseInstruction(s)
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, i)
	}
	return instructions, nil
}

func parseInstruction(s string) (Instruction, error) {
	parts := strings.Split(s, " ")
	arg, err := strconv.Atoi(parts[1])
	if err != nil {
		return nullInstruction, err
	}

	return Instruction{
		op:  parts[0],
		arg: arg,
	}, nil
}

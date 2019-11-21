package symboltable

type Address int

type SymbolTable struct {
	table map[string]Address
}

func NewSymbolTable() *SymbolTable {
	table := map[string]Address{
		"SP": 0, "LCL": 1, "ARG": 2, "THIS": 3, "THAT": 4,
		"R0": 0, "R1": 1, "R2": 2, "R3": 3, "R4": 4, "R5": 5, "R6": 6, "R7": 7,
		"R8": 8, "R9": 9, "R10": 10, "R11": 11, "R12": 12, "R13": 13, "R14": 14, "R15": 15,
		"SCREEN": 0x4000, "KBD": 0x6000,
	}
	return &SymbolTable{table: table}
}

func (s *SymbolTable) AddEntry(symbol string, address Address) {
	s.table[symbol] = address
}

func (s *SymbolTable) GetAddress(symbol string) (Address, bool) {
	address, has := s.table[symbol]
	return address, has
}

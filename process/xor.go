package process

// Xor function determines the turte table
func Xor(table []int) bool {
	nodes := []Node{
		Node{value: table[0]},
		Node{value: table[1]},
		Node{weight: 1.5, isThreShold: true},
		Node{weight: 0.5, isThreShold: true},
		Node{weight: 0.5, isThreShold: true}, // Last node convergenty
	}

	connections := []Connection{
		Connection{from: 0, to: 2, value: 1},
		Connection{from: 0, to: 3, value: 1},
		Connection{from: 1, to: 2, value: 1},
		Connection{from: 1, to: 3, value: 1},
		Connection{from: 2, to: 4, value: -1},
		Connection{from: 3, to: 4, value: 1},
	}

	return ResolveTable(&nodes, &connections)

}

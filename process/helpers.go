package process

// Node struct
type Node struct {
	value       int
	weight      float32
	isThreShold bool
}

// Connection struct
type Connection struct {
	from  uint
	to    uint
	value int
}

// ResolveTable resolve the table and return the result, this method is prepared for anything operations
func ResolveTable(nodes *[]Node, connections *[]Connection) bool {
	for index, node := range *nodes {
		if node.isThreShold {
			var weightInside int
			for _, connection := range *connections {
				if index == int(connection.to) {
					weightInside += (*nodes)[connection.from].value * (connection.value)

				}
			}
			if index == len(*nodes)-1 {
				if float32(weightInside) > node.weight {
					return true
				}
				break
			}
			if float32(weightInside) > node.weight {
				(*nodes)[index].value = 1
			} else {
				(*nodes)[index].value = 0
			}

		}
	}
	return false
}

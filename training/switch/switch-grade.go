package _switch

func GradeLevel(score int) {
	level := score / 10

	switch level {
	case 9, 10:
		println("A")
	case 8:
		println("B")
	case 7:
		println("C")
	case 6:
		println("D")
	default:
		println("E")
	}

	println(level)
}

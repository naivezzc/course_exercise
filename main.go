package main

import "fmt"

type Board struct {
	tokens [9]int
}

func (b *Board) put(x, y int, u string) {
	if u == "o" {
		b.tokens[x+3*y] = 1
	} else {
		b.tokens[x+3*y] = 2
	}
}

func (b *Board) get(x, y int) string {
	if b.tokens[x+3*y] == 1 {
		return "o"
	} else {
		b.tokens[x+3*y] = 2
		return "x"
	}
}

func (b *Board) win(i int) {
    
    if   b.tokens[0] == b.tokens[1] == b.tokens[2]  && (b.tokens[0] !=0)  ||
        b.tokens[3] == b.tokens[4] == b.tokens[5]  && (b.tokens[3] !=0)  ||
        b.tokens[6] == b.tokens[7] == b.tokens[8]  && (b.tokens[6] !=0)  ||
        b.tokens[0] == b.tokens[3] == b.tokens[6]  && (b.tokens[0] !=0)  ||
        b.tokens[1] == b.tokens[4] == b.tokens[7]  && (b.tokens[1] !=0)  ||
        b.tokens[2] == b.tokens[5] == b.tokens[8]  && (b.tokens[2] !=0)  ||
        b.tokens[0] == b.tokens[4] == b.tokens[8]  && (b.tokens[0] !=0)  ||
        b.tokens[2] == b.tokens[4] == b.tokens[6]  && (b.tokens[2] !=0)  {
            if i % 2 != 0{
                    fmt.Println("Player 1 won")
            }
            if i % 2 == 0{
                    fmt.Println("Player 2 won")
            }
    } else {
        fmt.Println("Continue")
    }
}



func main() {
	board := Board{
		tokens: [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Println("board before put", board.tokens)
	board.put(1, 2, "o")
	board.put(0, 0, "x")
	fmt.Println("board after put", board.tokens)
	fmt.Println("get(1, 2)", board.get(1, 2))
}

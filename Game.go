package main

import (
	
	"fmt"
)
const UP_RIGHT = "UP_RIGHT";
const UP_LEFT = "UP_LEFT";
const LEFT = "LEFT";
const RIGHT = "RIGHT";
const DOWN_RIGHT ="DOWN_RIGHT";
const DOWN_LEFT = "DOWN_LEFT"; 



// Sobhan said a man died when he give up !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!



type Point struct {
	i int ;
	j int ;
}

type Game struct {
	board [][]int;	
	out_1 int;
	out_2 int;
	
}



func get_next_pos(i int ,j int , direction string) Point {
	var return_val Point ;
	switch direction {
	case UP_RIGHT: 
		if i < 5 { 
			return_val.i = i-1;
			return_val.j = j; 
		} else {
			return_val.i = i-1;
			return_val.j = j+1;
		}
	case UP_LEFT:
		if i < 5 { 
			return_val.i = i-1;
			return_val.j = j-1; 
		} else {
			return_val.i = i-1;
			return_val.j = j;
		}
	case LEFT:
		return_val.i = i;
		return_val.j=j-1;
	case RIGHT:	
		return_val.i = i;
		return_val.j=j+1;
	case DOWN_RIGHT:
		if i < 4 { 
			return_val.i = i+1;
			return_val.j = j+1; 
		} else {
			return_val.i = i+1;
			return_val.j = j;
		}
	case DOWN_LEFT:
		if i < 4 { 
			return_val.i = i+1;
			return_val.j = j; 
		} else {
			return_val.i = i+1;
			return_val.j = j-1;
		}
	}
	return return_val;

}

func is_in_board(p Point) bool{
	ret := false;
	switch p.i {
	case 0: if p.j > -1 && p.j < 6{ ret = true}
	case 1:	if p.j > -1 && p.j < 7{ ret = true}
	case 2:if p.j > -1 && p.j < 8{ ret = true}
	case 3:if p.j > -1 && p.j < 9{ ret = true}
	case 4:if p.j > -1 && p.j < 10{ ret = true}
	case 5:if p.j > -1 && p.j < 9{ ret = true}
	case 6:if p.j > -1 && p.j < 8{ ret = true}
	case 7:if p.j > -1 && p.j < 7{ ret = true}
	case 8:if p.j > -1 && p.j < 6{ ret = true}
		
	}
	return ret
}


//0			2 2 2 2 2
//1		   2 2 2 2 2 2
//2		  0 0 2 2 2 0 0	
//3		 0 0 0 0 0 0 0 0
//4		0 0 0 0 0 0 0 0 0
//5		 0 0 0 0 0 0 0 0 
//6		  0 0 1 1 1 0 0 
//7		   1 1 1 1 1 1
//8			1 1 1 1 1


func (this *Game) init(){
	this.board = make([][]int ,  9);
	this.board[0] = make([]int, 5);
	this.board[1] = make([]int, 6);
	this.board[2] = make([]int, 7);
	this.board[3] = make([]int, 8);
	this.board[4] = make([]int, 9);
	this.board[5] = make([]int, 8);
	this.board[6] = make([]int, 7);
	this.board[7] = make([]int, 6);
	this.board[8] = make([]int, 5);


	this.board[0] = []int {2,2,2,2,2};
	this.board[1] = []int {2,2,2,2,2,2};
	this.board[8] = []int {1,1,1,1,1};
	this.board[7] = []int {1,1,1,1,1,1};
	this.board[2] = []int {0,0,2,2,2 ,0,0};
	this.board[6] = []int {0,0,1,1,1,0,0};
	this.board[3] = []int {0,0,0,0,0,0,0,0};
	this.board[5] = []int {0,0,0,0,0,0,0,0};
	this.board[4] = []int {0,0,0,0,0,0,0,0,0};

	
	this.out_1 =0;
	this.out_2= 0;

}


func (this *Game)   print_board(){
	print("in print func \n")
	for i:=0 ; i < 9 ; i++{
		for k:=0 ; k < 15 - len(this.board[i]);k++{
			fmt.Printf(" ");
		}
		
		for j := range this.board[i]{
			fmt.Printf("%d " , this.board[i][j]);
		}
		fmt.Print("\n")

	}
	println("end in print_board");
}




func (this *Game) move(player int , row int , col int , direction string) (int ){
	print("in move \n")

	if !is_in_board(Point{row , col}){
		return -1 //, errors.New("bad request");
	}
	if player != this.board[row][col]{
		return -1 //, errors.New("access denied!");
	}
	var power int =1;
	free_cell :=0;
	enemy :=0;
	next := get_next_pos(row , col , direction);
	for power <4 &&  enemy < 4 && is_in_board(next) {
		if this.board[next.i][next.j] == 0{
			free_cell++;
			break;
		}else if this.board[next.i][next.j] == player{
			power++;
		}else {
			enemy++;
		}

		next = get_next_pos(next.i , next.j , direction);

	}
	if power > 3 || enemy >= power || (free_cell==0 && enemy == 0){
		return -1 //, errors.New("you not can do this move!");
	}
	
	element := player;
	next = get_next_pos(row , col , direction);
	for element != 0 && is_in_board(next){
		next_element := this.board[next.i][next.j];
		this.board[next.i][next.j] = element;
		element = next_element;
		next = get_next_pos(next.i , next.j , direction);
	}
	this.board[row][col] = 0;
	if free_cell == 0 {
		if enemy==2{
			this.out_2++;
		}else {
			this.out_1++;
		}
	}

	ret :=0;
	if this.out_1 > 5 {
		ret =2;
	}else if this.out_2 > 5{
		ret =1;
	}
	return ret //, nil;

}


// func main(){
// 	var game Game ;
// 	game.init();
// 	game.print_board();

// 	//var ls [5]int= [5]int{ 3 , 3, 1 , 4, 9};
	
	

	
// 	//ls = []int {9 , 2, 1  , 1 , 4};
// 	// fmt.Println(ls);
// 	player := 2;
// 	winner := 0;
// 	for winner < 1{
// 		var s string;
// 		var i,j int;
// 		//println("in for")
// 		fmt.Scanf("%d %d %s\n" , &i , &j , &s);
// 		print(i , j , s)
// 		winner = game.move(player , i , j , s);
// 		game.print_board();
// 		if (winner == -1){
// 			continue;
// 		}
// 		if player == 1{
// 			player=2;
// 		}else{
// 			player=1
// 		}
		
// 	}

// }





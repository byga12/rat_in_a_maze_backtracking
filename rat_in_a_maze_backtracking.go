package main

import (
	"errors"
	"fmt"
)

type Coord struct {
    x int
    y int
}

func main(){
    labyrinth := [][]int{
        {0,0,0,0,0,0,0,0,0,1},
        {0,0,0,0,0,0,0,1,1,1},
        {0,0,0,0,0,0,0,1,0,0},
        {0,1,1,1,1,1,1,1,0,0},
        {1,1,0,0,0,0,1,0,0,0},
        {0,0,0,0,0,0,1,0,0,0},
    }
    start := Coord{4,0}
    end := Coord{5,6}
    if !isValidCoord(labyrinth,start){
        panic(errors.New("Índice fuera de rango"))
    }
        if !isValidCoord(labyrinth,end){
        panic(errors.New("Índice fuera de rango"))
    }
    backtrackLabyrinth(labyrinth,start,end)
}

func printMatrix(matrix [][]int){
    // rows := len(matrix)
    // cols := len(matrix[0])
    // fmt.Printf("Filas: %d\nColumnas: %d\n", rows, cols)
    for _,rows:=range(matrix){
        fmt.Printf("%v\n",rows)
    }
    fmt.Println()
}

func backtrackLabyrinth(labyrinth [][]int, currCoord Coord, endCoord Coord){
    labyrinth[currCoord.x][currCoord.y] = 2
    printMatrix(labyrinth)

    if currCoord==endCoord {
        fmt.Println("Finish!!!")
        currCoord=endCoord
        return
    }

    // Look around

    // Up
    if isValidCoord(labyrinth,Coord{currCoord.x-1,currCoord.y}) && labyrinth[currCoord.x-1][currCoord.y]==1{
        fmt.Println("Up")
        newCoord:=Coord{currCoord.x-1,currCoord.y}
        backtrackLabyrinth(labyrinth,newCoord,endCoord)
    }
    // Right
    if isValidCoord(labyrinth,Coord{currCoord.x,currCoord.y+1}) && labyrinth[currCoord.x][currCoord.y+1]==1{
        fmt.Println("Right")
        newCoord:=Coord{currCoord.x,currCoord.y+1}
        backtrackLabyrinth(labyrinth,newCoord,endCoord)
    }
    // Down
    if isValidCoord(labyrinth,Coord{currCoord.x+1,currCoord.y}) && labyrinth[currCoord.x+1][currCoord.y]==1{
        fmt.Println("Down")
        newCoord:=Coord{currCoord.x+1,currCoord.y}
        backtrackLabyrinth(labyrinth,newCoord,endCoord)
    }
    // Left
    if isValidCoord(labyrinth,Coord{currCoord.x,currCoord.y-1}) && labyrinth[currCoord.x][currCoord.y-1]==1{
        fmt.Println("Left")
        newCoord:=Coord{currCoord.x,currCoord.y-1}
        backtrackLabyrinth(labyrinth,newCoord,endCoord)
    }
    labyrinth[currCoord.x][currCoord.y]=1
    return
}

func isValidCoord(matrix [][]int, coord Coord) bool{
    rowsQ := len(matrix)
    colsQ := len(matrix[0])

    if coord.x < 0 || coord.x >= rowsQ || coord.y < 0 || coord.y >= colsQ {
        return false
    } 
    
    return true
}
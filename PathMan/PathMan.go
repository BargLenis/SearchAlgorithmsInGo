package main

import (
    "fmt"
    "Containers"
)

type Vector struct {
    X int
    Y int
    Z int
}

func (n *Vector) String() string {
	return fmt.Sprint("{", n.X, ",", n.Y, "}")
}

func (v Vector) Add(n Vector) Vector {
    return Vector{v.X+n.X,  v.Y+n.Y, v.Z+n.Z} //The v.z addition is used to add alayer using go right/left/etc.
}

const (
    PATHMAN = 80
    PATH = 45
    WALL = 37
    FOOD = 46
)
type GameBoard struct {
    Board []string
    Visited [][]int
    Player, Path, Wall, Food uint8
    MapSize Vector
}

func (gb *GameBoard) IsPlayer(v Vector) bool{
    return gb.IsVal(v, PATHMAN)
}
func (gb *GameBoard) IsFood(v Vector) bool{
    return gb.IsVal(v, FOOD)
}
func (gb *GameBoard) IsWall(v Vector) bool{
    return gb.IsVal(v, WALL)
}
func (gb *GameBoard) IsPath(v Vector) bool{
    return gb.IsVal(v, PATH)
}
func (gb *GameBoard) IsVal(v Vector, i uint8) bool {
    if gb.Board[v.X][v.Y] == i  {
        return true
    }
    return false
}
func (gb *GameBoard) MarkVisited(v Vector){
    gb.Visited[v.X][v.Y] = v.Z
}
func (gb *GameBoard) IsNew(v Vector) bool{
    if gb.Visited[v.X][v.Y] == 0 {
        return true
    }
    return false
}
func (gb *GameBoard) MarkWall(v Vector){
    gb.Visited[v.X][v.Y] = 77
}
func (gb *GameBoard) SetVisited(v Vector){

    gb.Visited  = make([][]int, v.X)
    for i := 0; i < v.X; i++ {
        gb.Visited[i]  = make([]int, v.Y)
    }
}
func (gb *GameBoard) PrintBoard(){
    for i := 0; i < gb.MapSize.X; i++ {
        fmt.Println(gb.Visited[i])
    }
}

type Player struct {
    Position Vector
    Up, Down, Left, Right Vector
    b bool
    count int
    hit, tv Containers.Container
}

func (p *Player) GoRight(gb GameBoard, pos Vector) (Vector) {
    return pos.Add(p.Right)
}

func (p *Player) GoLeft(gb GameBoard, pos Vector) (Vector)  {
    return pos.Add(p.Left)
}

func (p *Player) GoUp(gb GameBoard, pos Vector) (Vector)  {
    return pos.Add(p.Up)
}

func (p *Player) GoDown(gb GameBoard, pos Vector) (Vector)  {
    return pos.Add(p.Down)
}

func (p *Player) IsValidMove(gb GameBoard, mover Vector) (bool)  {
    if gb.IsFood(mover) {
        p.b = true
        return true
    } else if gb.IsPath(mover) && gb.IsNew(mover) {
        return true //p.DepthFirstSearch(gb)
    }
    return false
}

func (p *Player) DepthFirstSearch(gb GameBoard, start Vector, traversed Containers.Container) {
    traversed.Push(start)
    p.hit.Push(start)

    gb.MarkVisited(start)
    p.count ++
    var move Vector
    if p.b {

        p.tv = traversed
    }

    if !p.b {
        move = p.GoDown(gb, start)
        if p.IsValidMove(gb, move) {
            p.DepthFirstSearch(gb, move, traversed)
        }
    }
    if !p.b {
        move = p.GoRight(gb, start)
        if p.IsValidMove(gb, move) {
            p.DepthFirstSearch(gb, move, traversed)
        }
    }
    if !p.b {
        move = p.GoLeft(gb, start)
        if p.IsValidMove(gb, move) {
            p.DepthFirstSearch(gb, move, traversed)
        }
    }
    if !p.b {
        move = p.GoUp(gb, start)
        if p.IsValidMove(gb, move) {
            p.DepthFirstSearch(gb, move, traversed)
        }
    }
}

func (p *Player) BreadthFirstSearch(gb GameBoard, nextMoves *Containers.Container, traversed Containers.Container) {

    if nextMoves.PeekTop() != nil {
        if !p.b {
            start := nextMoves.PopTop().(Vector)
            traversed.Push(start)
            p.hit.Push(start)

            gb.MarkVisited(start)
            p.count ++

            var move Vector
            move = p.GoDown(gb, start)
            if p.IsValidMove(gb, move) {

            fmt.Println("down")
                nextMoves.Push(move)
            }

            move = p.GoRight(gb, start)
            if p.IsValidMove(gb, move) {

            fmt.Println("right")
                nextMoves.Push(move)
            }


            move = p.GoLeft(gb, start)
            if p.IsValidMove(gb, move) {
                nextMoves.Push(move)
            }


            move = p.GoUp(gb, start)
            if p.IsValidMove(gb, move) {
                nextMoves.Push(move)
            }

            p.BreadthFirstSearch(gb, nextMoves, traversed)

        }
    }

    if p.b {
        p.tv = traversed
    }

}


func main() {

    var p Player
    p.count = 0

    Pac := Vector{3,9, 0}

    p.Position = Pac
    p.b = false

    p.Up = Vector{-1,0,1}      // UP
    p.Left = Vector{0,-1,1}    // LEFT is inserted second
    p.Right = Vector{0,1,1}   // RIGHT is inserted third
    p.Down = Vector{1,0,1}    // DOWN is inserted fourth (on top)

    var gb GameBoard
    gb.Board = []string{"%%%%%%%%%%%%%%%%%%%%", "%--------------%---%",  "%-%%-%%-%%-%%-%%-%-%", "%--------P-------%-%", "%%%%%%%%%%%%%%%%%%-%", "%.-----------------%", "%%%%%%%%%%%%%%%%%%%%"}
    gb.Player = PATHMAN
    gb.Wall = WALL
    gb.Path = PATH
    gb.Food = FOOD


    gb.MapSize = Vector{7,20,0}

    gb.SetVisited(gb.MapSize)
    for i := 0; i < gb.MapSize.X; i++ {
        for j := 0; j < gb.MapSize.Y; j++ {
            if gb.IsWall(Vector{i, j, 0}) {
                gb.MarkWall(Vector{i,j, 0})
            }
        }
    }
    var zz Containers.Container
    posplus := new(Containers.Container)
    posplus.Push(p.Position)
    gb.MarkVisited(Pac)
    p.BreadthFirstSearch(gb, posplus, zz)
    //
    //p.DepthFirstSearch(gb, p.Position, zz)
    gb.PrintBoard()

    fmt.Println(p.hit.Len())
    x := p.hit.PopTop()
    for x != nil {
        fmt.Println(x)
        x = p.hit.PopTop()

    }

    fmt.Println(p.tv.Len())

    x = p.tv.PopTop()
    for x != nil {
        fmt.Println(x)
        x = p.tv.PopTop()

    }


}

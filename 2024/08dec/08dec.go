package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)
type Pair struct{
	Row int
	Column int
}

func checkValid(p Pair, m,n int) bool{
	r:=p.Row
	c:=p.Column
	return (r>=0 && c >=0 && r<m && c<n)
}
func part1(store [][]string){
	m := len(store)
	n := len(store[0])

	mapStorePreviousDigits := make(map[string][]Pair)
	mapStoreAntinodes := make(map[string]int)

	for j:=0;j<n;j++{
		if store[0][j]!="."{
			mapStorePreviousDigits[store[0][j]] = append(mapStorePreviousDigits[store[0][j]], Pair{Row:0, Column: j})
		}	
	}

	for i:=1;i<m;i++{
		for j:=0;j<n;j++{
			if store[i][j]!="."{
				str := store[i][j]
				pairs, exists := mapStorePreviousDigits[str]
				if exists{
					for _, pair:= range pairs{
						rowdiff := i-pair.Row
						coldiff:=j-pair.Column
						pairtop := Pair{Row:pair.Row-rowdiff, Column: pair.Column-coldiff}
						pairbottom := Pair{Row:i+rowdiff, Column: j+coldiff}
						
						if checkValid(pairtop,m,n){
							mapStoreAntinodes[fmt.Sprintf("%d,%d", pairtop.Row, pairtop.Column)] =1
						}

						if checkValid(pairbottom,m,n){
							mapStoreAntinodes[fmt.Sprintf("%d,%d", pairbottom.Row, pairbottom.Column)] =1
						}

					}
				}
				
			
				mapStorePreviousDigits[str] = append(mapStorePreviousDigits[str], Pair{Row: i, Column: j})
			
			}
		}
	}


	fmt.Println(len(mapStoreAntinodes))
}

func part2(store [][]string){
	m := len(store)
	n := len(store[0])



	mapStorePreviousDigits := make(map[string][]Pair)
	mapStoreAntinodes := make(map[string]int)

	for j:=0;j<n;j++{
		if store[0][j]!="."{
			mapStorePreviousDigits[store[0][j]] = append(mapStorePreviousDigits[store[0][j]], Pair{Row:0, Column: j})
			mapStoreAntinodes[fmt.Sprintf("%d,%d",0,j)]=1
		}	
	}

	for i:=1;i<m;i++{
		for j:=0;j<n;j++{
			if store[i][j]!="."{
				str := store[i][j]
				pairs, exists := mapStorePreviousDigits[str]
				if exists{
					for _, pair:= range pairs{
						rowdiff := i-pair.Row
						coldiff:=j-pair.Column
						
						ptc := pair.Column
						for ptr:=pair.Row-rowdiff;ptr>=0 && ptr<m;ptr-=rowdiff{

							ptc-=coldiff
							pt := Pair{Row: ptr, Column: ptc}
							if checkValid(pt,m,n){
								
								mapStoreAntinodes[fmt.Sprintf("%d,%d", pt.Row, pt.Column)] =1
							}

						}
						
						pbc:=j
						for pbr:=i+rowdiff;pbr>=0 && pbr<m;pbr+=rowdiff{
							pbc+=coldiff
							pb := Pair{Row: pbr, Column: pbc}
							if checkValid(pb,m,n){
								mapStoreAntinodes[fmt.Sprintf("%d,%d", pb.Row, pb.Column)] =1
							}
						}

					}
				}
				
				mapStoreAntinodes[fmt.Sprintf("%d,%d",i,j)]=1
				mapStorePreviousDigits[str] = append(mapStorePreviousDigits[str], Pair{Row: i, Column: j})
			
			}
		}
	}

	fmt.Println(len(mapStoreAntinodes))
}

func main(){

	if len(os.Args)<1 {	
		fmt.Println("Error while gettin the file")
		return 
	}

	file, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println("error while opening the file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var store [][]string
	for scanner.Scan(){
		str := scanner.Text()
		parts := strings.Split(str, "")
		
		store = append(store, parts)

	}

	//fmt.Println(store)
	start1 := time.Now()
	part1(store)
	end1 := time.Now() 

	start2:=time.Now()
	part2(store)
	end2:=time.Now()

	fmt.Println("Time for part1 :", end1.Sub(start1))
	fmt.Println("Time for part2 :", end2.Sub(start2))
}	
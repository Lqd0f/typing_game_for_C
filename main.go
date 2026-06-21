package main

import(

	"os"
	"os/exec"
	"fmt"
	"math/rand/v2"
	"bufio"
)

const ROUND = 1

func main(){
	
	game()
}

func initi()(int, error){

	cnt := 0

	files, err := os.ReadDir("./problems")
	if err != nil {
		fmt.Println("ERROR WITH READING A PROBLEM DIRECTORY", err)
		return -1, err
	}

	for _, f := range files {
		if f.IsDir() {
			cnt++
		}
	}

	return cnt, nil
}

func game()error{

	dirCnt, err := initi()

	if err != nil {
		return err
	}

	order := make([]int, dirCnt)
	for i := range dirCnt {
		order[i] = i
	}

	shfl(order);

	for i := range ROUND {
		query(order[i])
	}

	return nil
}

func shfl(ar []int){

	var sz int = len(ar)
	for i := range sz {
		fir, sec := i, rand.N(sz-i)
		ar[fir], ar[sec] = ar[sec], ar[fir]
	}
}

func query(id int){

	dirNm := fmt.Sprintf("./problems/p%d/", id)

	txtFi, err := os.Open(fmt.Sprintf("%starg.txt", dirNm))
	if err != nil {
		fmt.Println("ERROR WITH OPENING A TXT FILE", err)
		return
	}
	defer txtFi.Close()

	var targ []string
	targScnr := bufio.NewScanner(txtFi)
	for targScnr.Scan() {
		targ = append(targ, targScnr.Text())
	}

	fmt.Println()
	fmt.Println("====== GOAL OF OUTPUT ======")
	for i := range len(targ) {
		fmt.Println(targ[i])
	}
	fmt.Println("============================\n")

	if err := targScnr.Err(); err != nil {
		fmt.Println("ERROR WITH  READING A TXT FILE", err)
		return
	}

	srcFi, err := os.Open(fmt.Sprintf("%sfull.c", dirNm))
	if err != nil {
		fmt.Println("ERROR WITH OPENING A SOURCE FILE", err)
		return
	}
	defer srcFi.Close()

	var src []string
	srcScnr := bufio.NewScanner(srcFi)
	for srcScnr.Scan() {
		src = append(src, srcScnr.Text())
	}

	if err := srcScnr.Err(); err != nil {
		fmt.Println("ERROR WITH  READING A SOURCE FILE", err)
		return
	}

	/*
	include
	\n
	main
	\n	
		
	\n
	ret
	}
	*/
	hole := rand.N(len(src)-7)+4 

	for i := range src {
		for _, ch := range src[i] {
			if ch == '\t' {
				fmt.Printf("  ")
			} else {
				break
			}
		}

		if i != hole {
			for _, ch := range src[i] {
				if ch != '\t' {
					fmt.Printf("%c", ch)
				}
			}
			fmt.Println()
			continue
		}

		fmt.Println("/* ======[ WHAT PROGRAM DOES FIT IN HERE ? ]====== */")
	}
	fmt.Printf("\n\n")

	var inpt string

	stdScnr := bufio.NewScanner(os.Stdin)
	isCrr := true 
	for{
		os.Remove("./generated/gene.c")
		geneFi, err := os.OpenFile(
			"./generated/gene.c",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644,
		)
		if err != nil {
			fmt.Println("ERROR WITH OPENING A FILE:", err)
			return
		}
		defer geneFi.Close() 

		fmt.Println("enter the correct sentence or syboml")
		fmt.Printf("> ")
		stdScnr.Scan()
		inpt = stdScnr.Text()
		for i := range src {
			var line string
			if i == hole {
				line = fmt.Sprintf("%s\n", inpt)
			} else {
				line = fmt.Sprintf("%s\n", src[i])
			}

			_, err = geneFi.Write([]byte(line))
			if err != nil {
				fmt.Println("ERROR WITH WRITING TO FILE:", err)
				return 
			}
		}

		cpil := exec.Command(
			"gcc",
			"./generated/gene.c",
			"-o",
			"./generated/EXE",
		)

		if err := cpil.Run(); err != nil {
			continue
		}

		resFi, err := os.OpenFile(
			"./generated/rslt.txt",
			os.O_RDWR|os.O_CREATE|os.O_TRUNC,
			0644,
		)
		if err != nil {
			fmt.Println("ERROR WITH CREATING A FILE:", err)
			return
		}
		defer resFi.Close()

		xqt := exec.Command("./generated/EXE")
		xqt.Stdout = resFi

		if err := xqt.Run(); err != nil {
			fmt.Println("ERROR ON 209", err)
			continue
		}

		resFi.Seek(0, 0)
		var rslt []string
		rsltScnr := bufio.NewScanner(resFi)
		for rsltScnr.Scan() {
			rslt = append(rslt, rsltScnr.Text())
		}

		if len(rslt) != len(targ) {
			fmt.Println("ERROR ON 222", len(rslt), len(targ))
			continue
		}

		for i := range len(targ) {
			if rslt[i] != targ[i] {
				isCrr = false
			}
		}

		if isCrr {
			fmt.Println("CORRECT!", len(rslt), len(targ))
			break
		}
	}
}

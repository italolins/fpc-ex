package main

import (
	"fmt";
	"time"
)

type prod struct {
	id int
	value int
}


func produtor (id int, produced chan prod, doneProd chan int, sizeBuffer int ) {
	for i := 0; i < sizeBuffer * 10 ; i++{
		produ := prod{id,i}
		//fmt.Println("P: ", produ)
		produced <- produ
	} 
	doneProd <- 1
}

func consumidor (id int, produced chan prod, endCons chan int, sleepTime time.Duration){
	for x := range(produced){
		fmt.Println("C",id,":  ", x)
		time.Sleep( sleepTime * time.Millisecond)
	}
	close(endCons)
}

func produtor_consumidor( qtdProd int, qtdCons int, sizeBuffer int, endf chan int){
	produced := make(chan prod, sizeBuffer)
	endCons := make( chan int )
	doneP := make(chan int )
	
	for i := 1; i <= qtdProd; i++ {
		go produtor(i, produced, doneP, sizeBuffer)
	}
	for i := 1; i <= qtdCons; i++ {
		go consumidor(i, produced, endCons, time.Duration( 5))
	}
	
	auxi := 0
	
	for{
		<- doneP
		auxi++
		if(auxi == qtdProd){
			close(produced)
			<- endCons
			close(endf)
		}
	}
	fmt.Println("passou aki")
	
}

func main() {
	endf := make(chan int)
	
	qtdeProdutor := 2
	qtdeConsumidor := 3
	sizeBuffer := 3
	
	go produtor_consumidor(qtdeProdutor,qtdeConsumidor,sizeBuffer,endf)
	<- endf
	fmt.Println("FIM!!")
}

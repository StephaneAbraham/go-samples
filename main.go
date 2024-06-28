package main

import "fmt"

func main() {
	for {
		const tickets int =250
		var remTickets =250
		
	    if(remTickets==0){
		    fmt.Println("Sorry tickets are sold out")
	    } else{

	        fmt.Printf("\nGrab the opportunity to visit MUNNAR before the tickets are sold out\n ")
	    }
	    input()
	    options()
	
                
	}
	
}


func input(){
	var name string
	var number int
	var email string
	fmt.Printf("\nPlease enter your name : \n")
	fmt.Scanln(&name)
	fmt.Println("Please enter your mobile number : ")
	fmt.Scanln(&number)
	fmt.Println("Please enter your email id : ")
	fmt.Scanln(&email)
}
func options(){
	options := 1
	fmt.Printf("\nChoose the option you wish to visit \n 1. **Munnar hill Station**\n 2. **Eravikulam National park** \n 3. **Kolukkumalai Tea estate**\n 4. **Echo point munnar**\n 5. **Mattupetty Dam**\n\n")
	fmt.Scanln(&options)
	switch options{
	    case 1 :
            hill := 50
			var htickets int 
			fmt.Printf("\n\nThe place you have chosen to visit is Munnar hill station\n")
			fmt.Printf("\n\nThe cost of visiting the place is 2000 rupees per head\n")
			fmt.Printf("\n\nEnter the number of tickets required\n")
			fmt.Scanln(&htickets)
			remhill := hill- htickets 
			if(remhill==0) {
				fmt.Printf("\nTickets sold out for Munnar hill station\n")
			} 
			if(htickets>50){
				fmt.Println("\nNumber exceeded")
			} else if(htickets==0) {
				fmt.Println("\nTickets not booked")
			} else {
			fmt.Printf("\nThankyou for booking %v tickets and the tickets will be send to you via the mail provided\n",htickets)
			fmt.Printf("\nThe remainging tickets for Munnar hill station is % v \n", remhill)
			}
		case 2 :
			park := 50
			var ptickets int 
			fmt.Println("\nThe place you have chosen to visit is Eravikulam National Park")
			fmt.Println("\nThe cost of visiting the place is 2500 rupees per head")
			fmt.Println("\nEnter the number of tickets required")
			fmt.Scanln(&ptickets)
			var rempark int = park-ptickets
			if(rempark==0) {
				fmt.Println("\nTickets sold out for Eravikulam National Park")
			} 
			if(ptickets>50){
				fmt.Println("\nNumber exceeded")
			} else if(ptickets==0) {
				fmt.Println("\nTickets not booked")
			} else {

			fmt.Printf("\nThankyou for booking %v tickets and the tickets will be send to you via the mail provided\n",ptickets)
			fmt.Printf("\nThe remainging tickets for Eravikulam National Park is % v\n ", rempark)
			}
		case 3 :
		    estate := 50
			var etickets int 
			fmt.Printf("\nThe place you have chosen to visit is Kolukkumalai Tea Estate\n")
			fmt.Printf("\nThe cost of visiting the place is 3000 rupees per head\n")
			fmt.Printf("\nEnter the number of tickets required\n")
			fmt.Scanln(&etickets)
			remestate := estate-etickets
			if(remestate==0) {
				fmt.Printf("\nTickets sold out for Kolukkumalai Tea Estate\n")
			} 
			if(etickets>50){
				fmt.Printf("\nNumber exceeded")
			} else if(etickets==0) {
				fmt.Printf("\nTickets not booked")
			} else {
			fmt.Printf("\nThankyou for booking %v tickets and the tickets will be send to you via the mail provided\n",etickets)
			fmt.Printf("\nThe remainging tickets for Kolukkumalai Tea Estate is % v \n", remestate)
			}
	    case 4 :
			echo := 50
			var echotickets int 
			fmt.Printf("\nThe place you have chosen to visit is Echo point munnar\n")
			fmt.Printf("\nThe cost of visiting the place is 3500 rupees per head\n")
			fmt.Printf("\nEnter the number of tickets required\n")
			fmt.Scanln(&echotickets)
			remecho := echo-echotickets
			if(remecho==0) {
				fmt.Println("\nTickets sold out for Echo point munnar")
			} 
			if(echotickets>50){
				fmt.Println("\nNumber exceeded")
			} else if(echotickets==0) {
				fmt.Println("\nTickets not booked")
			} else {
			fmt.Printf("\nThankyou for booking %v tickets and the tickets will be send to you via the mail provided\n",echotickets)
			fmt.Printf("\nThe remainging tickets for Echo point munnar is % v \n", remecho)
			}
		case 5 :
			dam := 50
			var dtickets int 
			fmt.Printf("\nThe place you have chosen to visit is Mattupetty Dam\n")
			fmt.Printf("\nThe cost of visiting the place is 3000 rupees per head\n")
			fmt.Printf("\nEnter the number of tickets required\n")
			fmt.Scanln(&dtickets)
			remdam := dam-dtickets
			if(remdam==0) {
				fmt.Println("\nTickets sold out for EMattupetty Dam")
			} 
			if(dtickets>50){
				fmt.Println("\nNumber exceeded")
			} else if(dtickets==0) {
				fmt.Println("\nTickets not booked")
			} else {
			fmt.Printf("\nThankyou for booking %v tickets and the tickets will be send to you via the mail provided\n",dtickets)
			fmt.Printf("\nThe remainging tickets for Mattupetty Dam is % v \n", remdam)
			}
	}
}
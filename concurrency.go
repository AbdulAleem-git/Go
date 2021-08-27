package main

// This is how goroutine works
// func printalphabets() {
// 	for _, c := range "hello"{
// 		time.Sleep(time.Millisecond * 250)
// 		fmt.Printf("%c " , c)
// 	}
// }
// func printnumbers(){
// 	for _, i := range []int{1,2,3,4,5,6,5} {
// 		time.Sleep(time.Millisecond * 500)
// 		fmt.Printf("%d ",i)
// 	}
// }
// func main(){
// 	go printalphabets()
// 	go printnumbers()
// 	time.Sleep (time.Second * 4 )
// 	fmt.Println("main terminated")
// }

// Lets just try to understand channel
// Actually when some recieving and or inserting is happen it is blocking calls.

// func main(){
// 	var c chan int
// 	if c == nil {
// 		fmt.Println("channel is empty going to init the channel")
// 		c = make(chan int)
// 		fmt.Printf("%T",c)
// 	}
// 	c <- 3
// 	data := <- c
// 	fmt.Println(data)

// }

// Lets just try to see how channels are useful with goroutine

// func printalphabets(s chan rune) {
// 	for _, v := range "AbdulAleem" {
// 		time.Sleep(time.Millisecond * 100)
// 		s <- v
// 	}
// 	close(s)
// }
// func main() {
// 	var str chan rune
// 	str = make(chan rune)
// 	go printalphabets(str)

// 	// for {
// 	// 		v , ok := <-str
// 	// 		if !ok{
// 	// 			break
// 	// 		}
// 	// 		fmt.Printf("%c ",v)
// 	// }

// 	for v := range str {
// 		fmt.Printf("%c", v)
// 	}
// 	fmt.Println()

// }

// Here we will see on problem in wich we are trying to get the squares sum and cubes sum
//

// func digits(number int, digchannel chan int) {
// 	defer close(digchannel)
// 	for number != 0 {
// 		d := number % 10
// 		digchannel <- d
// 		number = number / 10
// 	}
// }

// func calcSquare(number int, sqch chan int) {
// 	defer close(sqch)
// 	sum := 0
// 	dch := make(chan int)
// 	go digits(number , sqch)
// 	for d := range dch {
// 		sum += d * d
// 	}
// 	sqch <- sum
// }

// func calccube(number int , cubech chan int){
// 	defer close(cubech)
// 	sum := 0
// 	cubch := make (chan int)
// 	go digits(number,cubch)
// 	for d := range cubch{
// 		sum += d * d *d
// 	}
// 	cubech <- sum
// }

// func main(){
// 	reader := bufio.NewReader(os.Stdin)
// 	number , _  := reader.ReadString('\n')
// 	num , _  := strconv.Atoi(number)

// 	if num == 0{
// 		num = 123
// 	}
// 	sqch := make(chan int)
// 	cubch := make(chan int)
// 	start := time.Now()
// 	go calcSquare(num,sqch)
// 	go calccube(num, cubch)
// 	squares , cubes := <- sqch , <-cubch

// 	fmt.Printf("squares: %d, cubes: %d\n",squares,cubes)
// 	fmt.Printf("%d\n",time.Since(start).Milliseconds())
// }

// Buffered Channel and Working pool

// 1-

// func main(){
// 	ch := make( chan int , 2)
// 	ch <- 12
// 	ch <- 13
// 	for c := range ch {
// 		fmt.Println(c)
// 	}
// 	close(ch)
// }

//2.
// func write ( ch chan int ) {
// 	for i := 0 ; i < 5 ; i ++{
// 		ch <- i
// 		fmt.Println("write done")
// 	}
// 	close(ch)
// }

// func main(){
// 	ch := make (chan int , 2)
// 	go write(ch)
// 	for v := range ch {
// 		fmt.Printf("%d \n", v)
// 		time.Sleep(1 * time.Second)
// 	}
// }

// 3- length and capicity

// func main(){
// 	ch := make(chan int, 10 )
// 	ch <- 1
// 	ch <- 2
// 	ch <- 3
// 	fmt.Printf("len:%d  cap:%d",len(ch), cap(ch))

// }

// 4- WaithGroup

// func print(wg *sync.WaitGroup ){
// 	for i := 0 ;  i < 10 ; i++{
// 		time.Sleep(time.Millisecond * 200)
// 		fmt.Printf("%d ", i)
// 	}
// 	fmt.Println()
// 	wg.Done()

// }
// func main(){
// 	 var wg sync.WaitGroup
// 	 for i := 0 ; i < 3 ; i++{
// 		 wg.Add(1)
// 		 go print(&wg)
// 	 }
// 	 wg.Wait()
// }

// 5. Working Pools and WaitGroup

// type Job struct {
// 	id       int
// 	randomno int
// }
// type Result struct {
// 	job          Job
// 	sumofdigits int
// }
// var jobs = make (chan Job , 10 )
// var results = make (chan Result , 10 )

// func digits (number int) int {
// 	sum := 0
// 	num := number
// 	for num != 0 {
// 		d := num % 10
// 		num = num / 10
// 		sum += d
// 	}
// 	time.Sleep(time.Second * 2 )
// 	return sum
// }

// func worker (wg *sync.WaitGroup){
// 	for job := range jobs {
// 		output := Result{job, digits(job.randomno)}
// 		results <- output
// 	}
// 	wg.Done()
// }
// func createWorkerPool(noOfWorkers int){
// 	var wg sync.WaitGroup
// 	for i := 0 ; i < noOfWorkers ; i++{
// 		wg.Add(1)
// 		go worker(&wg)
// 	}
// 	wg.Wait()
// 	close(results)
// }
// func allocate(noOfJobs int){
// 	for i := 0 ; i < noOfJobs ; i++{
// 		randomno := rand.Intn(999)
// 		job := Job{i, randomno}
// 		jobs <- job
// 	}
// 	close(jobs)
// }

// func result (done chan bool){
// 	for result := range results{
// 		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
// 	}
// 	done <- true
// }

// func main(){
// 	startTime := time.Now()

// 	noOfJobs := 10
// 	go allocate(noOfJobs)
// 	done := make (chan bool)
// 	go result(done)
// 	noOfWorkers := 5
// 	createWorkerPool(noOfWorkers)
// 	<- done
// 	endTime := time.Now()
//     diff := endTime.Sub(startTime)
//     fmt.Println("total time taken ", diff.Seconds(), "seconds")
// }

// Select
//  Real world scenerio where Select is usful

// func server1(ch chan string){
// 	time.Sleep(time.Millisecond * 2)
// 	ch <- "from server1"
// }

// func serve
// r2(ch chan string){
// 	time.Sleep(time.Millisecond * 3)
// 	ch <- "from server2"
// }

// func main(){
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)
// 	go server1(ch1)
// 	go server2(ch2)
// 	select{
// 	case s1 := <-ch1:
// 		fmt.Println(s1)
// 	case s2 := <-ch2:
// 		fmt.Println(s2)
// 	}
// }

// func server1 (ch chan string){
// 	time.Sleep(time.Second * 2)
// 	ch <- "from server1"

// }
// func server2 (ch chan string){
// 	time.Sleep(time.Second * 1)
// 	ch <- "from server2"
// }

// func main(){

// 	ch1  := make (chan string)
// 	ch2  := make (chan string)
// 	go server1 (ch1)
// 	go server2 (ch2)
// 	for {
// 		time.Sleep(time.Millisecond * 250)
// 		select {
// 		case r1 :=  <- ch1:
// 			fmt.Println(r1)
// 			return
// 		case r2  := <- ch2:
// 			fmt.Println(r2)
// 			return
// 		default:
// 			fmt.Println("default case")
// 		}
// 	}
// }

func main(){

	select{}
}
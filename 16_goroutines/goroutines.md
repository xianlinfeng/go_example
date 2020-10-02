怎么终止一个goroutine：  
# 1. First Response  
创建一个channel（with buffer）,当有数据发送至该channel时，终止任务。  
# 2. Once task 
使用系统包 sync.Once 执行一次性任务。  
```go 
var once sync.Once  
once.Do(  
    func() {fmt.Println("Create Obj")}  
    )  
```
# 3. until all done: 
## 3.1.   使用waitgroup:
```go 
var wg sync.WaitGroup 

wg.Add(1) // add a task 
go worker(i, &wg)

func worker(id int, wg *sync.WaitGroup) {
    wg.Done() // finish a task
}

wg.Wait() // wait until all task are done
```

## 3.2.   遍历该channel，直到所有内容被取出 或 channel被关闭。
## 3.3.   使用time.sleep 等待所有goroutine完成。
## 3.4  (sync)在main函数最后增加一行，（`done是一个buffer channel`），当任务完成后，再往done里面发送一个信息。
```go
var done = make(chan bool, 1)
func main(){

    <- done
}
```


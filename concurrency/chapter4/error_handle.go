package main

//func main() {
//
//	type Result struct {
//		Error    error
//		Response *http.Response
//	}
//
//	checkStatus := func(done <-chan interface{}, strings ...string) <-chan Result {
//		response := make(chan Result)
//		go func() {
//			defer close(response)
//			for _, url := range strings {
//				var result Result
//				resp, err := http.Get(url)
//				result = Result{Error: err, Response: resp}
//				select {
//				case <-done:
//					return
//				case response <- result:
//
//				}
//			}
//		}()
//
//		return response
//	}
//
//	done := make(chan interface{})
//	defer close(done)
//	urls := []string{"https://www.baidu.com", "https://www.163.com", "http:badhttp"}
//	for response := range checkStatus(done, urls...) {
//		if response.Error != nil {
//			fmt.Printf("error:%v", response.Error)
//			continue
//		}
//		fmt.Printf("Respone: %v\n", response.Response.Status)
//	}
//
//}

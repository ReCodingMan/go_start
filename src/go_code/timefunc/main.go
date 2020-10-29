package main
import (
	"fmt"
	"time"
)

func main() {
	
	now := time.Now()
	fmt.Printf("now=%v now Type=%T\n", now, now)

	//通过日期获取年月日，时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//格式化日期
	//方式一：
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(),
	now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Printf("dateStr=%v", dateStr)

	//第二种：
	fmt.Printf(now.Format("2006/01/02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006/01/02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	//每隔0.1s打印
	i := 0
	for {
		i++
		fmt.Println(i)
		//休眠
		time.Sleep(time.Millisecond * 100)
		if i==100 {
			break
		}
	}

	//Unix和UnixNano纳秒
	fmt.Printf("unix时间戳=%v unixnano时间戳=%v", now.Unix(), now.UnixNano() )
}
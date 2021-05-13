# go-rule34-scraper
Simple but efficient rule34 scraper using the r34.app API

# Supported domains
- rule34.xxx
- rule34.paheal.net
- danbooru.donmai.us
- gelbooru.com
- e621.net
- safebooru.org
- e926.net

# Instructions
1. Download the repository 
2. CD into the folder
3. run `go get` to install needed depencies
4. Run `go run main.go <domain> <limit> <tags>` to start downloading
    - domain can be anything from the supported domains above.
    - limit can be anything.
    - tags is an array of images you want to scrape. Example: `nude sonic_the_hedgehog`
4 (optional). Run `go build main.go` to compile it into a executable 
5. Wait till every image is downloaded.
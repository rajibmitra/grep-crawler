# grep-crawler
golang grep crawler 

We start by defining the initial URL to begin crawling from.

We maintain a map visited to keep track of visited URLs and a channel queue to manage the URLs to be crawled.

We initiate the crawling process by sending the initial URL to the queue.

We use a goroutine to process URLs from the queue. For each URL, we fetch its content using http.Get, parse the HTML using html.NewTokenizer, and extract links.

We ensure that we only process links that haven't been visited before and add them back to the queue for further crawling.

This is a very basic web crawler 
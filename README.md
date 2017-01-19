hn
==

A CLI(Command Line Interface) tool for hacker news[0].


Setup
=====

First of all, ensure that your system has Go development environment[1].  Then,
just `$ go get github.com/sjp38/hn` from your shell.


Usage
=====

Basic usage is `hn [options]`.  With no option, it will print out top ten
stories of hacker news in simple form like below:
```
$ hn
[1] U.S. sues Oracle, alleges salary and hiring discrimination (247)
[2] What I Wish I'd Known About Equity Before Joining a Unicorn (933)
[3] Too much sitting, too little exercise accelerate biological aging (12)
[4] PyTorch – Tensors and Dynamic neural networks in Python (313)
[5] Welcoming Fabric to Google (248)
[6] Caching at Reddit (147)
[7] Rust vs. Go (303)
[8] How the Human Brain Decides What Is Important and What’s Not (32)
[9] Go vs. Swift [pdf] (50)
[10] Iceland knows how to stop teen substance abuse (67)
```

Number inside parantheses is the HN score of the story.


Options
-------

It also provides few command line options.  You can show description about that
using `-help` option as below:
```
$ hn -help
Usage of hn:
  -category string
        Category of items to show.  It should be (top|new|best) (default "top")
  -nrItems int
        Number of items to print out (default 10)
  -showCommentURL
        Show URL for HN comments
  -showOrigURL
        Show URL for the story
  -showTitle
        Show Title
  -showURLs
        Show URLs for the story and HN comments
exit status 2
```


ShowURLs option
---------------

While default execution output has only title and score of each stories, you
can also show more detailed information including URLs for each story using
`-showURLs` option as below:
```
$ go run hn.go -showURLs
[1] U.S. sues Oracle, alleges salary and hiring discrimination (247)
(http://www.reuters.com/article/us-oracle-usa-labor-idUSKBN1522O6?il=0)
(https://news.ycombinator.com/item?id=13430222)

[2] What I Wish I'd Known About Equity Before Joining a Unicorn (934)
(https://gist.github.com/yossorion/4965df74fd6da6cdc280ec57e83a202d)
(https://news.ycombinator.com/item?id=13426494)

[3] Too much sitting, too little exercise accelerate biological aging (13)
(http://sciencebulletin.org/archives/9448.html)
(https://news.ycombinator.com/item?id=13431806)

[4] PyTorch – Tensors and Dynamic neural networks in Python (313)
(http://pytorch.org/)
(https://news.ycombinator.com/item?id=13428098)

[5] Welcoming Fabric to Google (249)
(https://firebase.googleblog.com/2017/01/FabricJoinsGoogle17.html)
(https://news.ycombinator.com/item?id=13428595)

[6] Caching at Reddit (147)
(https://redditblog.com/2017/1/17/caching-at-reddit/)
(https://news.ycombinator.com/item?id=13429314)

[7] Rust vs. Go (303)
(https://blog.ntpsec.org/2017/01/18/rust-vs-go.html)
(https://news.ycombinator.com/item?id=13430108)

[8] How the Human Brain Decides What Is Important and What’s Not (32)
(http://neurosciencenews.com/importance-neuroscience-decisions-5967/)
(https://news.ycombinator.com/item?id=13430892)

[9] Go vs. Swift [pdf] (50)
(https://github.com/jakerockland/go-vs-swift/blob/master/go-vs-swift.pdf)
(https://news.ycombinator.com/item?id=13430640)

[10] Iceland knows how to stop teen substance abuse (67)
(https://mosaicscience.com/story/iceland-prevent-teen-substance-abuse)
(https://news.ycombinator.com/item?id=13430547)
```


License
=======

GPL v3


Author
======

SeongJae Park (sj38.park@gmail.com)


References
==========

[0] https://news.ycombinator.com/

[1] https://golang.org/doc/install

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
[1] Three.js editor (139)
[2] What Could Have Entered the Public Domain on January 1, 2017 (416)
[3] Ask HN: Excluding WordPress, what is your favorite for blogs or small stores? (56)
[4] Why do traders in investment banks feel their jobs are immune from AI, etc? (113)
[5] Recreating 3D renderings in real life (166)
[6] IR is better than assembly (2013) (129)
[7] Lenovo ThinkPad T460 – A Good Linux Laptop for Development (125)
[8] The moving sofa problem (526)
[9] Ask HN: What is your side project for 2017? (19)
[10] How I Write Tests (255)
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
# 10 top stories

[1] Three.js editor (143)
[https://threejs.org/editor/]
[https://news.ycombinator.com/item?id=13299479]

[2] What Could Have Entered the Public Domain on January 1, 2017 (418)
[https://web.law.duke.edu/cspd/publicdomainday/2017/pre-1976]
[https://news.ycombinator.com/item?id=13297792]

[3] Ask HN: Excluding WordPress, what is your favorite for blogs or small stores? (57)
[]
[https://news.ycombinator.com/item?id=13300023]

[4] Why do traders in investment banks feel their jobs are immune from AI, etc? (114)
[https://www.quora.com/Why-do-traders-in-investment-banks-feel-their-jobs-are-immune-from-AI-automation-and-deep-learning]
[https://news.ycombinator.com/item?id=13299311]

[5] Recreating 3D renderings in real life (168)
[http://skrekkogle.com/still-file/]
[https://news.ycombinator.com/item?id=13298004]

[6] IR is better than assembly (2013) (129)
[https://idea.popcount.org/2013-07-24-ir-is-better-than-assembly/]
[https://news.ycombinator.com/item?id=13297424]

[7] Lenovo ThinkPad T460 – A Good Linux Laptop for Development (127)
[https://karussell.wordpress.com/2017/01/02/lenovo-thinkpad-t460-a-good-linux-laptop-for-development/]
[https://news.ycombinator.com/item?id=13299585]

[8] The moving sofa problem (527)
[https://www.math.ucdavis.edu/~romik/movingsofa/]
[https://news.ycombinator.com/item?id=13296502]

[9] Ask HN: What is your side project for 2017? (21)
[]
[https://news.ycombinator.com/item?id=13300178]

[10] Swiss say goodbye to banking secrecy (47)
[http://www.swissinfo.ch/eng/business/tax-evasion_swiss-say-goodbye-to-banking-secrecy-/42799134]
[https://news.ycombinator.com/item?id=13298021]
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

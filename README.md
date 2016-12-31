hn
==

Command Line Interface tool for hacker news[0].


Usage
=====

Because this program is written in Go, you need Go development environment[1]
setup.  If your system has Go development environment, just `$ go run hn.go`.

It will print out top ten stories of hacker news in simple form like below:
```
$ go run hn.go
[1] Be Careful with Python's New-Style String Format (162)
[http://lucumr.pocoo.org/2016/12/29/careful-with-str-format/]
[https://news.ycombinator.com/item?id=13292350]

[2] Deep Learning 2016: The Year in Review (117)
[http://www.deeplearningweekly.com/blog/deep-learning-2016-the-year-in-review]
[https://news.ycombinator.com/item?id=13291640]

[3] Silicon Valley's obscure unicorns could boost 2017 IPO market (34)
[http://www.reuters.com/article/us-technology-ipos-idUSKBN14J0GY]
[https://news.ycombinator.com/item?id=13283624]

[4] Implementing Function Spreadsheets (2008) [pdf] (12)
[https://pdfs.semanticscholar.org/ab87/31cd70495b715acd33ba683c94c47e88ea14.pdf]
[https://news.ycombinator.com/item?id=13282127]

[5] Recursive Anonymous Functions in Elixir: Combinators and Macros (20)
[https://github.com/jisaacstone/ex_rfn]
[https://news.ycombinator.com/item?id=13291798]

[6] A Paper Algorithm Notation (44)
[http://canonical.org/~kragen/sw/dev3/paperalgo]
[https://news.ycombinator.com/item?id=13286503]

[7] Russia Hysteria Infects WashPost: False Story About Hacking U.S. Electric Grid (144)
[https://theintercept.com/2016/12/31/russia-hysteria-infects-washpost-again-false-story-about-hacking-u-s-electric-grid/]
[https://news.ycombinator.com/item?id=13292607]

[8] Interactive Map: The Flow of International Trade (8)
[http://www.visualcapitalist.com/interactive-mapping-flow-international-trade/]
[https://news.ycombinator.com/item?id=13293008]

[9] To Make the World Better, Think Small (29)
[http://www.nytimes.com/2016/12/30/opinion/to-make-the-world-better-think-small.html?action=click&pgtype=Homepage&clickSource=story-heading&module=opinion-c-col-left-region&region=opinion-c-col-left-region&WT.nav=opinion-c-col-left-region&_r=0]
[https://news.ycombinator.com/item?id=13292238]

[10] The IT security culture, hackers vs. industry consortia (55)
[http://laforge.gnumonks.org/blog/20161206-it_security_culture_telecoms/]
[https://news.ycombinator.com/item?id=13291594]
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

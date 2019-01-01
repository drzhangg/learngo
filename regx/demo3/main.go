package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := `</head>
<body>
	<header class="navbar navbar-default navbar-fixed-top" role="navigation">
		<div class="container">
			
			<div class="navbar-header">
				<a href="/" class="navbar-brand" title="Go语言中文网"><img alt="Go语言中文网" src="https://static.studygolang.com/img/logo1.png" style="margin-top: -7px; height: 45px;"></a>
				<button class="navbar-toggle" type="button" data-toggle="collapse" data-target="#navbar-main">
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
				</button>
			</div>
			<div class="navbar-collapse collapse" id="navbar-main">
				<ul class="nav navbar-nav">
					<li class="">
						<a href="/topics">主题</a>
					</li>
					<li class="">
						<a href="/articles">文章</a>
					</li>
					<li class="">
						<a href="/projects">项目</a>
					</li>
					<li class="">
						<a href="/resources">资源</a>
					</li>
					<li class="">
						<a href="/books">图书</a>
					</li>
					<div>测试</div>
					<div>你好
						hahahahah
						大神
					</div>
					<div>我号</div>
					<div>他也好</div>
					<li class="dropdown ">
						
						<a class="dropdown-toggle" data-toggle
		`

	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	if reg == nil {
		fmt.Println("reg error ")
		return
	}

	result := reg.FindAllStringSubmatch(buf,-1)

	for _,text := range result{
		fmt.Println("text[0] = ",text[1])
		//fmt.Println("text[0] = ",text)
	}
	//fmt.Println(result)

}

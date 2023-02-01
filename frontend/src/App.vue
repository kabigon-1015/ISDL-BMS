<template>
<head>
<meta charset="UTF-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<title>企業・ビジネスサイト向け 無料ホームページテンプレート tp_biz52</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="description" content="ここにサイト説明を入れます">
<meta name="keywords" content="キーワード１,キーワード２,キーワード３,キーワード４,キーワード５">
<!--[if lt IE 9]>
<script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
<script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
<style>
.menu1 a {background-position: -10px -10px;}
.menu2 a {background-position: -10px -130px;}
.menu3 a {background-position: -10px -250px;}
.menu4 a {background-position: -10px -370px;}
.menu5 a {background-position: -10px -490px;}
</style>
<![endif]-->
<!--[if lt IE 10]>
<style>
.slide0,.slide1,.slide2,.slide3 {background: url(images/1.jpg) no-repeat center center;}
</style>
<![endif]-->
</head>

<body class="home">
<div id="container">
<!--PC用（801px以上端末）で表示させるブロック-->
<header class="pc">

	<h1 class="logo"><a href="index.html"><img src="./assets/images/logo.png" alt="SAMPLE SITE"></a></h1>

	<!--PC用（801px以上端末)メニュー-->
	<nav id="menubar">
	<ul>
	<li class="menuimg menu1 current"><router-link to="/"><span>Home</span></router-link></li>
	<li class="menuimg menu2"><router-link to="/about"><span>Rent</span></router-link></li>
	<li class="menuimg menu3"><a href="works.html"><span>Works</span></a></li>
	<li class="menuimg menu4"><a href="link.html"><span>Links</span></a></li>
	<li class="menuimg menu5"><a href="contact.html"><span>Contact</span></a></li>
	</ul>
	</nav>
	<ul class="icon">
	<li><a href="#"><img src="./assets/images/icon_facebook.png" alt="Facebook"></a></li>
	<li><a href="#"><img src="./assets/images/icon_twitter.png" alt="Twitter"></a></li>
	<li><a href="#"><img src="./assets/images/icon_instagram.png" alt="Instagram"></a></li>
	<li><a href="#"><img src="./assets/images/icon_youtube.png" alt="YouTube"></a></li>
	</ul>

</header>
<!--/.pc-->

<!--小さな端末用（800px以下端末）で表示させるブロック-->
<header class="sh">

	<h1 class="logo"><a href="index.html"><img src="./assets/images/logo.png" alt="SAMPLE SITE"></a></h1>

	<!--小さな端末用（800px以下端末）メニュー-->
	<div id="menubar-s">
	<nav>
	<ul>
	<li class="menuimg menu1 current"><a href="index.html"><span>Home</span></a></li>
	<li class="menuimg menu2"><a href="company.html"><span>Company</span></a></li>
	<li class="menuimg menu3"><a href="works.html"><span>Works</span></a></li>
	<li class="menuimg menu4"><a href="link.html"><span>Links</span></a></li>
	<li class="menuimg menu5"><a href="contact.html"><span>Contact</span></a></li>
	</ul>
	</nav>
	<ul class="icon">
	<li><a href="#"><img src="./assets/images/icon_facebook.png" alt="Facebook"></a></li>
	<li><a href="#"><img src="./assets/images/icon_twitter.png" alt="Twitter"></a></li>
	<li><a href="#"><img src="./assets/images/icon_instagram.png" alt="Instagram"></a></li>
	<li><a href="#"><img src="./assets/images/icon_youtube.png" alt="YouTube"></a></li>
	</ul>
	
	</div>
	<!--/#menubar-s-->

</header>
<!--/.sh-->

<!--スライドショー-->
<aside id="mainimg">
<div class="slide0">slide0</div>
<div class="slide1">slide1</div>
<div class="slide2">slide2</div>
<div class="slide3">slide3</div>
</aside>

</div>
<!--/#container-->

<!--ページの上部に戻る「↑」ボタン-->
<p class="nav-fix-pos-pagetop"><a href="#pagetop">↑</a></p>

<!--メニュー開閉ボタン-->
<div id="menubar_hdr" class="close"></div>

<!--メニューの開閉処理条件設定800px以下-->

</body>
  <router-view/>
</template>

<style>
/* #app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
} */
</style>

<script>
export default {
  head: {
    title() {
      return {
        inner: 'App',
        separator: '|',
        complement: 'page',
      }
    },
    meta: [
      { name: 'description', content: 'My description', id: 'desc' },
      { charset: 'utf-8' }
    ],
    link: [
		{ rel: 'stylesheet', href: "../assets/css/style.css", async: true, body: true}
    ],
	script:[
		{ type: 'text/javascript', src: "../assets/js/openclose.js", async: true, body: true}
	]
  }
}
</script>
<template>
<div>
<header class="pc">

	<h1 class="logo"><a href="index.html"><img src="../assets/images/ISDL-BMS.png" alt="SAMPLE SITE"></a></h1>

	<!--PC用（801px以上端末)メニュー-->
	<nav id="menubar">
	<ul>
	<li class="menuimg menu1 current"><router-link to="/"><span>Home</span></router-link></li>
	<li class="menuimg menu2"><router-link :to="{name:'rental',params:{show:1}}"><span>Rental</span></router-link></li>
	<li class="menuimg menu3"><router-link to="/return"><span>Return</span></router-link></li>
	<li class="menuimg menu4"><a href="link.html"><span>Links</span></a></li>
	<li class="menuimg menu5"><router-link to="/booklist"><span>Book List</span></router-link></li>
	</ul>
	</nav>
	<!-- <ul class="icon">
	<li><a href="#"><img src="../assets/images/icon_facebook.png" alt="Facebook"></a></li>
	<li><a href="#"><img src="../assets/images/icon_twitter.png" alt="Twitter"></a></li>
	<li><a href="#"><img src="../assets/images/icon_instagram.png" alt="Instagram"></a></li>
	<li><a href="#"><img src="../assets/images/icon_youtube.png" alt="YouTube"></a></li>
	</ul> -->

</header>
<!--/.pc-->

<!--小さな端末用（800px以下端末）で表示させるブロック-->
<header class="sh">

	<h1 class="logo"><a href="index.html"><img src="../assets/images/ISDL-BMS.png" alt="SAMPLE SITE"></a></h1>

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
	<li><a href="#"><img src="../assets/images/icon_facebook.png" alt="Facebook"></a></li>
	<li><a href="#"><img src="../assets/images/icon_twitter.png" alt="Twitter"></a></li>
	<li><a href="#"><img src="../assets/images/icon_instagram.png" alt="Instagram"></a></li>
	<li><a href="#"><img src="../assets/images/icon_youtube.png" alt="YouTube"></a></li>
	</ul>
	</div>
</header>
</div>
</template>

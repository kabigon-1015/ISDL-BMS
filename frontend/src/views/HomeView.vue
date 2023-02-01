<template>
<div id="contents">

<div id="main">

<span id="pagetop"></span>

<section id="new">

<h2>What's New</h2>
<dl>
<dt>2019/11/09</dt>
<dd>幅800px以下の端末でトップページの「What's New」を見た場合にスクロールできなかった問題点を修正。利用中のお客様は、style.cssの800px以下に#new dl {overflow: auto;}を追加して下さい。<span class="newicon">NEW</span></dd>
<dt>2019/08/27</dt>
<dd>印刷した際に１ページ目しか印刷できない問題を修正しました。</dd>
<dt>2019/08/17</dt>
<dd>ページ右下の「↑」アイコンをクリックしても上部に戻らない不具合を修正しました。</dd>
<dt>2019/07/19</dt>
<dd>tp_biz52公開。</dd>
<dt>20XX/00/00</dt>
<dd>サンプルテキスト。サンプルテキスト。サンプルテキスト。サンプルテキスト。サンプルテキスト。サンプルテキスト。サンプルテキスト。サンプルテキスト。サンプルテキスト。サンプルテキスト。サンプルテキスト。サンプルテキスト。サンプルテキスト。</dd>
<dt>20XX/00/00</dt>
<dd>サンプルテキスト。サンプルテキスト。サンプルテキスト。</dd>
<dt>20XX/00/00</dt>
<dd>サンプルテキスト。サンプルテキスト。サンプルテキスト。</dd>
<dt>20XX/00/00</dt>
<dd>サンプルテキスト。サンプルテキスト。サンプルテキスト。</dd>
<dt>20XX/00/00</dt>
<dd>サンプルテキスト。サンプルテキスト。サンプルテキスト。</dd>
<dt>20XX/00/00</dt>
<dd>サンプルテキスト。サンプルテキスト。サンプルテキスト。</dd>
<dt>20XX/00/00</dt>
<dd>サンプルテキスト。サンプルテキスト。サンプルテキスト。</dd>
</dl>

</section>
<!--/#new-->

<footer>
<small>Copyright&copy; <a href="index.html">SAMPLE SITE</a> All Rights Reserved.</small>
<span class="pr"><a href="https://template-party.com/" target="_blank">《Web Design:Template-Party》</a></span>
</footer>

</div>
<!--/#main-->

</div>
<!--/#contents-->
</template>

<style>
/* ファイルパスに従って@importntの後にファイルパスを書き込む */
@import "../assets/css/style.css";
@import "../assets/css/print.css";
</style>

<script>
// @ is an alias to /src
//import HelloWorld from '@/components/HelloWorld.vue'
export default {
  head: {
    title() {
      return {
        inner: 'HomeView',
        separator: '|',
        complement: 'HelloWorld',
      }
    },
    meta: [
		{ name: 'viewport', content: 'width=device-width, initial-scale=1'},
		{ name: 'description', content: 'ここにサイト説明を入れます'},
		{ name: 'keywords', content: 'キーワード１,キーワード２,キーワード３,キーワード４,キーワード５'}
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

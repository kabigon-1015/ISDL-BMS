<template>
<div class="home">
  <div id="contents">
    <div id="main">
      <span id="pagetop"></span>
      <section id="new">
        <h2>What's New</h2>
        <dl>
          <dt>2023/1/25</dt>
          <dd>
            書籍管理システム作成開始<span class="newicon"
              >NEW</span
            >
          </dd>
        </dl>
      </section>
      <!--/#new-->
    </div>
  </div>
  <!--/#contents-->
</div>
</template>

<style scoped>
#contents {
  height: 100%;
	width: 100%;
	position: relative;
	width: 100%;
	top: 550px;	
	right: 0%;
}
</style>

<script>
export default {
  components: {
  },
  data(){
    return{
    }
  },
  methods:{
  },
  head: {
    title() {
      return {
        inner: "HomeView",
        separator: "|",
        complement: "HelloWorld",
      };
    },
    meta: [
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { name: "description", content: "ここにサイト説明を入れます" },
      {
        name: "keywords",
        content:
          "キーワード１,キーワード２,キーワード３,キーワード４,キーワード５",
      },
    ],
  },
};
</script>

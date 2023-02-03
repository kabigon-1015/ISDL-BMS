<template>
  <div id="contents">
    <div id="main">
      <section id="pagetop">
        <section class="box">
          <h2 class="title">Rental<span>貸出</span></h2>
          <!-- <div v-if="show==1"> -->
          <div class="list">
            <!-- <button v-on:click="showpage(2)">
              <figure><img src="../assets/images/sample1.jpg" alt="" /></figure>
              <h4>バーコードリーダーから検索</h4>
            </button> -->
            <router-link to="/rental/BarcodeReader">
            <figure><img src="../assets/images/sample1.jpg" alt="" /></figure>
              <h4>バーコードリーダーから検索</h4>
            </router-link>
          </div>

          <div class="list">
            <router-link to="/rental/BarcodeReader">
              <figure><img src="../assets/images/sample2.jpg" alt="" /></figure>
              <h4>書籍検索</h4>
            </router-link>
          </div>

          <div class="list">
            <router-link to="/rental/BarcodeReader">
              <figure><img src="../assets/images/sample3.jpg" alt="" /></figure>
              <h4>書籍一覧から探す</h4>
            </router-link>
          </div>
          <!-- </div> -->
          <!-- <div v-else-if="show==2">
            <router-view />
          </div>
          <div v-else-if="show==3">
            <BookSearch />
          </div>
          <div v-else-if="show==4">
          </div> -->
        </section>
      </section>
    </div>
    <!--/#main-->
  </div>
  <!--/#contents-->
</template>

<style scoped>
/* .BookRental{
  position: fixed;
	width: 100%;
	height: 100%;
  background: url(../assets/images/2.jpg) no-repeat center center;	
  background: url(../assets/images/2.jpg) no-repeat center center / cover;	
} */
</style>
<script>
// import BarcodeReader from "../components/BarcodeReader.vue";
// import BookSearch from "../components/BookSearch.vue";

export default {
  // components: {
  //   // BarcodeReader,
  //   BookSearch
  // },
  data:function(){
    return{
      show:1,
    };
  },
  methods:{
    showpage:function(num){
      this.show=num
    }
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


<template>
  <div id="container">
    {{ this.name }}
    <HomeHeader />
	<router-view />
	<!-- <HomeFooter /> -->
	<!--スライドショー-->
    <aside id="mainimg">
      <div class="slide0">slide0</div>
      <div class="slide1">slide1</div>
      <div class="slide2">slide2</div>
      <div class="slide3">slide3</div>
    </aside>
  </div>
</template>

<style>
@import "./assets/css/style.css";
@import "./assets/css/print.css";
</style>

<script>
import HomeHeader from "./components/HomeHeader.vue";
// import HomeFooter from "./components/HomeFooter.vue";
export default {
  components: {
    HomeHeader
	// HomeFooter
  },
  head: {
    title() {
      return {
        inner: "App",
        separator: "|",
        complement: "page",
      };
    },
    meta: [
      { name: "description", content: "My description", id: "desc" },
      { charset: "utf-8" },
    ],
  },
};
</script>

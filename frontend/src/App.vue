<template>
  <div id="container">
    <LoginIcon />
    <HomeHeader />
    <router-view />
	
	<!--スライドショー-->
    <aside id="mainimg">
      <div class="slide0">slide0</div>
      <div class="slide1">slide1</div>
      <div class="slide2">slide2</div>
      <div class="slide3">slide3</div>
    </aside>
  </div>
</template>

<script>
import HomeHeader from "./components/HomeHeader.vue";
import LoginIcon from "./components/LoginIcon.vue";

export default {
  components: {
    HomeHeader,
    LoginIcon
  },
  data() {
    return {

    }
  },
  mounted(){
    this.$store.dispatch("getcurrentpass", '/');
  },
  methods: {
    
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

<style>
@import "./assets/css/style.css";
@import "./assets/css/print.css";
</style>


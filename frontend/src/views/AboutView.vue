<template>
  <div class="about">
    <h1>This is an about page</h1>
    <HelloWorld msg="aaa"/>
  </div>
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
  name: 'HelloWorld',
  components:{
    HelloWorld
  },
  head: {
    title() {
      return {
        inner: 'About',
        separator: '|',
        complement: 'Page',
      }
    },
    meta: [
      { name: 'description', content: 'My description', id: 'desc' }
    ]
  }
}
</script>
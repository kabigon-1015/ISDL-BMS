<template>
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
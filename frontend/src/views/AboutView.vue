<template>
  <div class="about">
    <h1>This is an about page</h1>
  </div>
</template>

<script>
// @ is an alias to /src
//import HelloWorld from '@/components/HelloWorld.vue'

export default {
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
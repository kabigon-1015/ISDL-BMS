<template>
  <div class="user">
    <img src="../assets/images/usericon.png" />
  <div v-if="this.$store.state.username=='ログイン'">
    <router-link to="/userlogin"
      ><p>{{ this.$store.state.username }}</p>
    </router-link>
  </div>
  <div v-else>
    <router-link to="/userinfo"
      ><p>{{ this.$store.state.username }}</p>
    </router-link>
  </div>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
  methods: {},
};
</script>

<style scoped>
.user {
  position: relative;
  height: 10px;
  width: 150px;
  margin: 0 0 0 auto;
  padding: 10px;
  
}
.user img {
  display: inline-block;
  position: absolute;
  height: 30px;
  width: 30px;
  left: 0px;
}
.user p {
  font-size: 16px;
  position: absolute;
  display: inline-block;
  left: 40px;
  top: 12px;
}
</style>

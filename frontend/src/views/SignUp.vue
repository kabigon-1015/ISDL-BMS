<template>
  <div id="contents">
    <div id="main">
      <section id="pagetop">
        <section class="box">
          <h2 class="title">Sign Up<span>ユーザ登録</span></h2>
          <form v-on:submit.prevent="doLogin">
        <label>e-mail address</label>
        <input type="text" v-model="email" />
        <label>User ID</label>
        <input type="text" v-model="id"  />
        <label>name</label>
        <input type="text" v-model="id"  />
        <label>password</label>
        <input type="password" v-model="id"  />
        <label>password（confirm）</label>
        <input type="password" v-model="id"  />
        <button v-on:click="subscribe" class="btn btn--yellow btn--cubic">
              Done
        </button>
        </form>
        </section>
      </section>
    </div>
    <!--/#main-->
  </div>
</template>

<script>
//import axios from 'axios';

//const uri = "http://localhost:8080/logincheck";

export default {
    name:'UserLogin',
         data(){
    return {
        id:'',
      password:'',
      student_name:'',
      myresponse:''
    };
  },
 
  methods: {
    subscribe: function(){
      this.$router.push({ path: '/' });
    },
getresponse() {

          return this.id
},


  //ログインボタン
  //get_responseからresponseを抜き取れない
    doLogin() {
     
    console.log(this.getresponse())
    this.$store.dispatch("auth", {userId: this.getresponse(),userToken: 'dummy token'});     
    console.log(this.$store.state.userId )
     
      this.$router.push('home');
    }
  }
  
};
</script>

<style scoped>
.btn,
a.btn,
button.btn {
  border: none;
  font-size: 1rem;
  font-weight: 200;
  line-height: 0.7;
  position: relative;
  display: inline-block;
  padding: 1rem 2rem;
  cursor: pointer;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
  -webkit-transition: all 0.3s;
  transition: all 0.3s;
  text-align: center;
  vertical-align: middle;
  text-decoration: none;
  letter-spacing: 0.1em;
  color: #212529;
  border-radius: 0.5rem;
}

button.btn--yellow {
  color: #000;
  background-color: #00d0ff;
  border-bottom: 5px solid #0096cc;
}

button.btn--yellow:hover {
  margin-top: 3px;
  color: #000;
  background: #00d0ff;
  border-bottom: 2px solid #0096cc;
}
</style>
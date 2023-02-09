<template>
  <div id="contents">
    <div id="main">
      <section id="pagetop">
        <section class="box">
          <h2 class="title">Login<span>ログイン</span></h2>
          <form v-on:submit.prevent="doLogin">
        <label>User ID</label>
        <input type="text" v-model="id" />
        <label>Password</label>
        <input type="password" v-model="password"  />
        <button v-on:click="login" class="btn btn--yellow btn--cubic">
              Login
        </button>
        <button v-on:click="signup" class="btn btn--yellow btn--cubic">
              Sign Up
        </button>
        <button v-on:click="forgetpassword" class="btn btn--yellow btn--cubic">
              Forget Password
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
    login: function(){
      this.$router.push({ path: '/' });
    },
    signup: function(){
      this.$router.push({ path: '/signup' });
    },
    forgetpassword: function(){
      this.$router.push({ path: '/resetpassword' });
      //this.$router.replace({ path: '/resetpassword' });
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
<template>
  <div id="contents">
    <div id="main">
      <section id="pagetop">
        <section class="box">
          <h2 class="title">Change User Info<span>ユーザ情報変更</span></h2>
          <h2>基本情報をご入力ください</h2>
          <div class="cp_iptxt">
            <!-- <label>e-mail address</label> -->
              <input
                class="ef"
                type="text"
                v-model="oldpass"
                placeholder="old password"
              />
              <label>old password</label>
              <span class="focus_line"><i></i></span>
          </div>
          <div class="cp_iptxt">
              <input
                class="ef"
                type="text"
                v-model="password"
                placeholder="new password"
              />
              <label>new password</label>
              <span class="focus_line"><i></i></span>
          </div>
          <div class="cp_iptxt">
              <input
                class="ef"
                type="text"
                v-model="username"
                placeholder="new username"
              />
              <label>new username</label>
              <span class="focus_line"><i></i></span>
          </div>
          <div class="cp_iptxt">
              <input
                class="ef"
                type="text"
                v-model="emailadress"
                placeholder="new email adress"
              />
              <label>new email</label>
              <span class="focus_line"><i></i></span>
          </div>
          <div class="loginbtn">
            <button v-on:click="changeuserinfo" class="btn btn--yellow btn--cubic">
              Change
            </button>
            </div>
        </section>
      </section>
    </div>
    <!--/#main-->
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name:'ChangeUserInfo',
  methods: {
    changeuserinfo:function(){
      var params = new FormData();
      params.append('old_id', this.$store.state.userid);
      params.append("old_name", this.$store.state.username);
      params.append("old_email", this.$store.state.useremail);
      params.append("old_password", this.oldpass);
      params.append("new_password", this.password);
      params.append("new_username", this.username);
      params.append("new_email", this.emailadress);
      axios
        .post("/changeuserinfo", params)
        .then((response) => {
          alert(response.data['message']);
          this.$store.dispatch("loginusername", response.data.name);
          this.$store.dispatch("getuseremail", response.data.email);
          if (response.data.success==1){
            this.$router.push({ path: '/userinfo' });
          }
        })
        .catch((error) => {
          alert("データを送信できませんでした．");
        });
    },
    
  }
  
};
</script>

<style scoped>
.loginbtn{
  position: relative;
  margin:  0 auto;
  margin-top: 40px;
  width: 150px;
}
.cp_iptxt {
  /* display: inline-block; */
  position: relative;
  margin:  0 auto;
  width: 40%;
  margin-top: 40px;
  background-color: #b2adad;
}
.cp_iptxt input[type="text"] {
  font: 15px/24px sans-serif;
  color: #000;
  box-sizing: border-box;
  width: 100%;
  letter-spacing: 1px;
  /* padding-left: 4em; */
}
.cp_iptxt input[type="text"]:focus {
  outline: none;
}
.ef {
  padding: 7px 14px;
  transition: 0.4s;
  border: 1px solid #1b2538;
  background: transparent;
}
.ef ~ .focus_line:before,
.ef ~ .focus_line:after {
  position: absolute;
  top: -1px;
  left: 50%;
  width: 0;
  height: 2px;
  content: "";
  transition: 0.4s;
  background-color: #da3c41;
}
.ef ~ .focus_line:after {
  top: auto;
  bottom: 0;
}
.ef ~ .focus_line i:before,
.ef ~ .focus_line i:after {
  position: absolute;
  top: 50%;
  left: 0;
  width: 2px;
  height: 0;
  content: "";
  transition: 0.6s;
  background-color: #da3c41;
}
.ef ~ .focus_line i:after {
  right: 0;
  left: auto;
}
.ef:focus ~ .focus_line:before,
.ef:focus ~ .focus_line:after,
.cp_iptxt.ef ~ .focus_line:before,
.cp_iptxt.ef ~ .focus_line:after {
  left: 0;
  width: 100%;
  transition: 0.4s;
}
.ef:focus ~ .focus_line i:before,
.ef:focus ~ .focus_line i:after,
.cp_iptxt.ef ~ .focus_line i:before,
.cp_iptxt.ef ~ .focus_line i:after {
  top: -1px;
  height: 100%;
  transition: 0.6s;
}
.ef ~ label {
  z-index: -1;
  opacity: 0;
  position: absolute;
  font-size: 20px;
  top: 4px;
  left: 14px;
  width: 50%;
  transition: 0.3s;
  letter-spacing: 0.5px;
  color: #b2adad;
}
.ef:focus ~ label,
.cp_iptxt.ef ~ label {
  z-index: 0;
  opacity: 1;
  font-size: 25px;
  top: -40px;
  left: 0;
  transition: 0.3s;
  color: #fc040c;
}

.btn,
button.btn {
  /* margin-top: 40px; */
  border: none;
  font-size: 1rem;
  font-weight: 200;
  line-height: 0.7;
  width: 150px;
  /* position: relative; */
  /* margin:  0 auto; */
  display: inline-block;
  padding: 1rem 2rem;
  cursor: pointer;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
  -webkit-transition: all 0.3s;
  transition: all 0.3s;
  /* text-align: center; */
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
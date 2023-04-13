<template>
  <div id="contents">
    <div id="main">
      <section id="pagetop">
        <section class="box">
            <h2 class="title">User Info<span>ユーザー情報</span></h2>
            <h3>
              <div class=ntitle>名前:</div>
              <div class=user>{{this.$store.state.username}}</div>
              <div class=ntitle>email:</div>
              <div class=user>{{this.$store.state.useremail}}</div>
            </h3>
      <div class=rent>
      <h3>
        <div class=rentedtitle>貸出中</div>
        <div class=hist>貸出履歴</div>
      </h3>
			<table class="rentedtable" align="left">
				<thead>
					<tr>
						<th class=title>タイトル</th>
						<th class=author>著者</th>
						<th class=publisher>出版社</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="rented_tableData in rented_tableData" :key="rented_tableData.name">
						<td class=title><router-link :to="{name:'detail',query:{title:rented_tableData.title}}" v-text="rented_tableData.title"></router-link></td>
						<td class=author v-text="rented_tableData.author"></td>
						<td class=publisher v-text="rented_tableData.publisher"></td>
					</tr>
				</tbody>
			</table>
          <!-- <div class=hist> -->
            <!-- <h3>貸出履歴</h3> -->
            <table class="histtable">
                <thead>
                    <tr>
                        <th class=title>タイトル</th>
                        <th class=author>著者</th>
                        <th class=publisher>出版社</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="hist_tableData in hist_tableData" :key="hist_tableData.name">
                        <td class=title><router-link :to="{name:'detail',query:{title:hist_tableData.title}}" v-text="hist_tableData.title"></router-link></td>
                        <td class=author v-text="hist_tableData.author"></td>
                        <td class=publisher v-text="hist_tableData.publisher"></td>
                    </tr>
                </tbody>
            </table>
          <!-- </div> -->
          </div>
          <br>
          <button v-on:click="changeinfo" class="btn btn--yellow btn--cubic">ユーザー情報変更</button>
        </section>
      </section>
    </div>
    <!--/#main-->
  </div>
</template>

<script>
import axios from 'axios'
const URL = '/'
export default {
	name: 'UserInfo',
		data() {
			return {
			checked:false,
			isbn_return:"",
			returnbooks:[],
			returnisbn:[],
			sendisbn:[],
			bookinfo:[{
				title:"aaa",
				author:"bbb",
				publisher:"ccc",
				state:'貸出中',
				isbn:1
			}
			],
			hist_getData:[],
			hist_tableData: [],
            rented_getData:[],
			rented_tableData: [],
			isbn:'',
			keyword:'',
		}
	},
	created: function() {
    var params = new FormData()
    params.append('user_id', this.$store.state.userid)
		axios.post(`${URL}userinfo`, params)
		.then(response => {
			// alert(response.status)
			for(var i in response.data.title){
          console.log(i)
          this.hist_getData = {
            title:response.data.title[i],
            author:response.data.author[i],
            publisher:response.data.publisher[i]
          }
          console.log(response.data)
          // if(this.getData.state == "null") {
          //     this.getData.state = '貸出可能'
          // }
          this.hist_tableData.push(this.hist_getData)
      }
      for(var j in response.data.rented_title){
          console.log(j)
          this.rented_getData = {
            title:response.data.rented_title[j],
            author:response.data.rented_author[j],
            publisher:response.data.rented_publisher[j]
          }
          console.log(response.data)
          // if(this.getData.state == "null") {
          //     this.getData.state = '貸出可能'
          // }
          this.rented_tableData.push(this.rented_getData)
      }
		})
		.catch(error => {
			alert('データを送信できませんでした．')
			// alert(error)
		})
	},
  methods:{
    changeinfo: function(){
      this.$router.push({ path: '/changeuserinfo' });
    }
  }
}
</script>

<style scoped>
.ntitle, .rentedtitle, .hist{
  /* margin-bottom: 1px;	見出しの下に空けるスペース */
	font-size: 24px;		/*文字サイズ*/
	letter-spacing: 0.2em;	/*文字間隔を少し広くとる設定*/
	color: #c1bc9e;
  float: left;
}
.user{
  /* margin-bottom:0%;	 見出しの下に空けるスペース */
	font-size: 24px;		/*文字サイズ*/
	letter-spacing: 0.2em;
  color: white;
  float: left;
  margin-right: 20%;
}
.hist{
    /* margin-top: 200px; */
   margin-left: 45%;
}
.histtable{
  float: left;
  margin-left: 38%;
}
table {
  position: absolute;
  /* table-layout: fixed; */
  /* top: 130px; */
  size: 100px;
  border-collapse: collapse;
  border-spacing: 5px;
  margin: 0 auto;
  padding: 0;
  width: 33%;
  margin-left: auto;
  margin-right: auto;
  margin-top: 20px;
}
table tr {
	height: 20px;
}
/* thead */
thead th {
	font-size: 10px;
	/* width: 300px; */
	text-align: center;
	padding: 6px;
	/* color */
	color: white;
	background-color:#b3ae92;
	border-top: 1px solid #a2a7af;
	border-left: 1px solid #a2a7af;
	border-right: 1px solid #a2a7af;
	border-bottom:3px solid #a2a7af;
}
a {
	color: #111;		/*リンクテキストの色*/
	transition: 0.4s;	/*マウスオン時の移り変わるまでの時間設定。0.4秒。*/
	text-decoration: none;
}
a:hover {
	color: #448db3;			/*マウスオン時の文字色*/
	text-decoration: none;	/*マウスオン時に下線を消す設定。残したいならこの１行削除。*/
}
/* tbody */
thead{
  display: block;
}
tbody {
  display: block;
  overflow-y: scroll;
  height: 250px;
}
.title{
  min-width: 200px;
}
.author{
  min-width: 100px;
}
.publisher{
  width: 99999px;
}
tbody td {
	/* width: 300px; */
	text-align: center;
	padding: 8px;
	/* color */
	color: black;
	background-color: #e0dada;
	border-bottom: 1px solid #a2a7af;
	border-left:1px solid #a2a7af;
	border-right:1px solid #a2a7af;
}

/*Checkboxes styles*/
input[type="checkbox"] { display: none; }

input[type="checkbox"] + label {
  display: block;
  position: relative;
  padding-left: 35px;
  padding-bottom:15px;
  margin-bottom: 25px;
  font: 14px/20px 'Open Sans', Arial, sans-serif;
  color: #ddd;
  cursor: pointer;
}

input[type="checkbox"] + label:last-child { margin-bottom: 0; }

input[type="checkbox"] + label:before {
  content: '';
  display: block;
  width: 20px;
  height: 20px;
  border: 1px solid #6cc0e5;
  position: absolute;
  left: 0;
  top: 0;
  opacity: .6;
  -webkit-transition: all .12s, border-color .08s;
  transition: all .12s, border-color .08s;
}

input[type="checkbox"]:checked + label:before {
  width: 10px;
  top: -5px;
  left: 5px;
  border-radius: 0;
  opacity: 1;
  border-top-color: transparent;
  border-left-color: transparent;
  -webkit-transform: rotate(45deg);
  transform: rotate(45deg);
}

.btn,
a.btn,
button.btn {
  border: none;
  font-size: 12px;
  font-weight: 1000;
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
  top: 55%;
  left: 20%
}
button.btn--yellow {
  color: white;
  background-color: #b3ae92;
  border-bottom: 5px solid #7f7b68;
}

button.btn--yellow:hover {
  margin-top: 3px;
  color: white;
  background: #b3ae92;
  border-bottom: 2px solid #7f7b68;
}
</style>
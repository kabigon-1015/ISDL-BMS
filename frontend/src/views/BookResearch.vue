<template>
  <div id="contents">
    <div id="main">
      <section id="pagetop">
        <section class="box">
          <h2 class="title">Rental<span>貸出</span></h2>
          <div>
            <h2>お探しの書籍名をご入力下さい</h2>

            <div class="cp_iptxt">
              <input class="ef" type="text" v-model="book" placeholder="" />
              <label>書籍名</label>
              <span class="focus_line"><i></i></span>
            </div>
            <button v-on:click="research" class="btn btn--yellow btn--cubic">
              検索
            </button>
            <h2>書籍に関連するタグを選んで下さい</h2>
            <div class="tagbtn">
              <div v-for="item in tags" :key="item.name">
                <input type="checkbox" v-model="item.checked" @click="clicked(item)">
                {{item.name}}
              </div>
            </div>
          </div>
      <table class="returntable">
				<thead>
					<tr>
						<th class=title>タイトル</th>
						<th class=author>著者</th>
						<th class=publisher>出版社</th>
						<th class=status>貸出状況</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="Data in tableData" :key="Data.name">
						<td class=title><router-link :to="{name:'detail',query:{title:Data.title}}" v-text="Data.title"></router-link></td>
						<td class=author v-text="Data.author"></td>
						<td class=publisher v-text="Data.publisher"></td>
						<td class=status v-if="Data.state=='貸出可能'"><button variant="secondary" class="btn btn--yellow btn--cubic" v-on:click="sendrental(Data.isbn)">借りる</button></td>
						<td class=status v-else v-text="Data.state"></td>
					</tr>
				</tbody>
			</table>
        </section>
      </section>
    </div>
    <!--/#main-->
  </div>
  <!--/#contents-->
</template>

<script>
import axios from "axios";
const books=[
  {
    title:"python89898989898989898989898989898",
    author:"oka",
    publisher:"hayato",
    status:"ok"
  },
  {
    title:"python",
    author:"oka",
    publisher:"hayato",
    status:"ok"
  },
  {
    title:"python",
    author:"oka",
    publisher:"hayato",
    status:"ok"
  },
  {
    title:"python",
    author:"oka",
    publisher:"hayato",
    status:"ok"
  },
  {
    title:"python",
    author:"oka",
    publisher:"hayato",
    status:"ok"
  },
  {
    title:"python",
    author:"oka",
    publisher:"hayato",
    status:"ok"
  },
  {
    title:"python",
    author:"oka",
    publisher:"hayato",
    status:"ok"
  },
];
const tags=["機械学習","GAN","LSTM","CNN","EEG","DE"]
export default {
  data: function () {
    return {
      book: "",
      books:books,
      tags:tags,
      tag:[],
      getData:[],
			tableData: [],
      tmp_data:[],
    };
  },
  created: function () {
    var params = new FormData();
    params.append("isbn", "this.isbn");
    this.tags = []
    console.log("response.data")
    const response = axios
        .post("/gettag", params)
        .then((response) => {
            for(var i in response.data.alltagname){
                this.tmp_data = {
                  name:response.data.alltagname[i],
                  checked:false
                }
                this.tags.push(this.tmp_data);
            }
            console.log(response.data);
        })
        .catch((error) => {
          alert("データを送信できませんでした");
        });
    },
  mounted(){
    this.$store.dispatch("getcurrentpass", '/rental/BookResearch');
  },
  methods: {
    research: function () {

      var params = new FormData();
      params.append("isbn", this.book);
      const response = axios
        .post("/research", params)
        .then((response) => {
          console.log(response.data);
        })
        .catch((error) => {
          alert("データを送信できませんでした．");
        });
    },
    sendrental(isbn){
			var params = new FormData()
			params.append('user_id', this.$store.state.userid)
			params.append('isbn', isbn)
			this.sendisbn = []
			console.log(params['isbn'])
			axios.post(`${URL}rental_register`, params)
			.then(response => {
				alert(response.data['message'])
				axios.post(`${URL}book_list`)
				.then(response => {
					// alert(response.status)
					// console.log(response.data)
					this.tableData=[]
					for(var i in response.data.title){
						console.log(i)
						this.getData = {
							title:response.data.title[i],
							isbn:response.data.isbn[i],
							author:response.data.author[i],
							publisher:response.data.publisher[i],
							state: String(response.data.rentaluser_name[i])
						}
						console.log(response.data)
						// if(this.getData.state == "null") {
						// 	this.getData.state = '貸出可能'
						// }
						this.tableData.push(this.getData)
					}	
				})
				.catch(error => {
					alert('データを送信できませんでした．')
					console.log(error)
					// alert(error)
				})
				console.log(response.data)
			})
			.catch(error => {
				alert('データを送信できませんでした．')
				alert(error)
			})
			this.returnbooks.splice(0)
		},
    clicked(item){
      this.tag = []
      this.tableData = []
      for (var t in this.tags){
        if (this.tags[t].name == item.name){
          if (this.tags[t].checked == true){
            this.tags[t].checked = false
            break
          }
          else{
            this.tags[t].checked = true
            break
          }
        }
      }
      for (var u in this.tags){
        if (this.tags[u].checked == true){
          this.tag.push(this.tags[u].name)
        }
      }
      console.log(this.tag)
      var params = new FormData();
      params.append('taglength', this.tag.length)
      this.tag.forEach((text, index) => {
        params.append('tag[' + index + ']', text);
      })
      const response = axios
        .post("/filterbook", params)
        .then((response) => {
          for(var i in response.data.title){
                console.log(i)
                this.getData = {
                    title:response.data.title[i],
                    author:response.data.author[i],
                    publisher:response.data.publisher[i],
                    state: String(response.data.rentaluser_name[i])
                }
                console.log(response.data)
                if(this.getData.state == "null") {
                    this.getData.state = '貸出可能'
                }
                this.tableData.push(this.getData)
                console.log(response.data)
            }
        })
        .catch((error) => {
          alert("データを送信できませんでした．");
        });
    },
    catchtag(tagdata){
      this.tableData = []
      var select = 'off'
      for (var t in this.tag){
        if (this.tag[t] == tagdata){
          select = 'on'
          this.tag.splice(t,1)
        }
      }
      if (select == 'off'){
        this.tag.push(tagdata)
      }
      var params = new FormData();
      params.append('taglength', this.tag.length)
      this.tag.forEach((text, index) => {
        params.append('tag[' + index + ']', text);
      })
      const response = axios
        .post("/filterbook", params)
        .then((response) => {
          for(var i in response.data.title){
                console.log(i)
                this.getData = {
                    title:response.data.title[i],
                    author:response.data.author[i],
                    publisher:response.data.publisher[i],
                    state: String(response.data.rentaluser_name[i])
                }
                console.log(response.data)
                if(this.getData.state == "null") {
                    this.getData.state = '貸出可能'
                }
                this.tableData.push(this.getData)
                console.log(response.data)
            }
        })
        .catch((error) => {
          alert("データを送信できませんでした．");
        });
    }
  },
};
</script>

<style scoped>
.tagbtn{
  width: 500px;
  height: 200px;
  flex-wrap:wrap;
  
}
.rental{
   text-align: center;
}
.rentalstatus{
  width: 60px;
  height: 5px;
  
  border: none;
  font-size: 1rem;
  font-weight: 200;
  line-height: 0.7;
  position: relative;
  display: inline-block;
  margin-right: 2%;
  margin-bottom:2%;
  padding: 1rem 0.5rem;
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
table {
  position: absolute;
  /* table-layout: fixed; */
  /* top: 130px; */
  size: 200px;
  border-collapse: collapse;
  border-spacing: 5px;
  margin: 0 auto;
  padding: 0;
  width: 70%;
  margin-left: auto;
  margin-right: auto;
}
table tr {
	height: 20px;
}
/* thead */
thead th {
	font-size: 12px;
	/* width: 300px; */
	text-align: center;
	padding: 8px;
	/* color */
	color: white;
	background-color:#b3ae92;
	border-top: 1px solid #a2a7af;
	border-left: 1px solid #a2a7af;
	border-right: 1px solid #a2a7af;
	border-bottom:3px solid #a2a7af;
}
/* tbody */
tbody {
	overflow-y: scroll;
}
.tbody::-webkit-scrollbar {  /* Chrome, Safari 対応 */
        display:none;
    }
tbody td {
	/* width: 30%; */
	text-align: center;
	padding: 8px;
	/* color */
	color: black;
	background-color: #e0dada;
	border-bottom: 1px solid #a2a7af;
	border-left:1px solid #a2a7af;
	border-right:1px solid #a2a7af;
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
.txt{
   text-align: left;
   font-size: 1.3em;
   color: #000;
}
.txt1{
   text-align: center;
   font-size: 1.3em;
   color: #000;
}
.price{
   text-align: right;
}
.book{
    width:200px;
    text-align: center;
}
.book1{
    width:100px;
    text-align: center;
}

.cp_iptxt {
  display: inline-block;
  position: relative;
  width: 30%;
  margin: 40px 3%;
  background-color: #b2adad;
}
.cp_iptxt input[type="text"] {
  font: 15px/24px sans-serif;
  color: #000;
  box-sizing: border-box;
  width: 100%;
  letter-spacing: 1px;
  padding-left: 4em;
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

*,
*:before,
*:after {
  -webkit-box-sizing: inherit;
  box-sizing: inherit;
}

html {
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  font-size: 62.5%;
}

.btn,
a.btn,
button.btn {
  border: none;
  font-size: 1rem;
  font-weight: 200;
  line-height: 0.7;
  position: relative;
  display: inline-block;
  margin-right: 2%;
  margin-bottom:2%;
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

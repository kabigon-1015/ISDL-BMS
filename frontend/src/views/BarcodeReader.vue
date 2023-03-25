<template>
  <div id="contents">
    <div id="main">
      <section id="pagetop">
        <section class="box">
          <h2 class="title">Rental<span>貸出</span></h2>
          <div>
            <h2>書籍のISBNをバーコードリーダーで読み取って下さい</h2>

            <div class="cp_iptxt">
              <input class="ef" type="text" v-model="isbn_return" @change="bookrental" placeholder="" />
              <label>ISBN</label>
              <span class="focus_line"><i></i></span>
            </div>
            <button variant="secondary" class="btn btn--yellow btn--cubic" v-on:click="sendrental">送信</button>
			<table class="returntable">
				<thead>
					<tr>
						<th>タイトル</th>
						<th>著者</th>
						<th>出版社</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="(book, index) in returnbooks" :key="index">
						<td>{{ book.title }}</td>
						<td>{{ book.author }}</td>
						<td>{{ book.publisher }}</td>
					</tr>
				</tbody>
			</table>
          </div>
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
	name: 'BookReturn',
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
			getData:[],
			tableData: [],
			isbn:'',
			keyword:'',
		}
	},
	methods: {
		bookrental:function () {
			var params = new FormData()
			params.append("isbn", this.isbn_return)
			this.isbn = this.isbn_return
			axios.post(`${URL}bookinfo`, params)
			.then(response => {
				this.getData = {
					title: response.data.title,
					author: response.data.author,
					publisher: response.data.publisher,
					isbn: this.isbn
				}
				// 表に追加
				this.returnbooks.push({
					title: String(this.getData.title),
					author:String(this.getData.author),
					publisher:String(this.getData.publisher),
				});
				this.sendisbn.push(this.getData.isbn)
			})
			this.isbn_return='';
		},
		sendrental:function(){
			var params = new FormData()
			params.append('user_id', this.$store.state.userid)
			params.append('isbn', this.sendisbn)
			this.sendisbn = []
			console.log(params['isbn'])
			axios.post(`${URL}rental_register`, params)
			.then(response => {
				alert(response.data['message'])
				console.log(response.data)
			})
			.catch(error => {
				alert('データを送信できませんでした．')
				alert(error)
			})
			this.returnbooks.splice(0)
		},
	}
}
</script>

<style scoped>
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
table {
  position: absolute;
  table-layout: fixed;
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
	/* width: 100px; */
	text-align: center;
	padding: 8px 0;
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
</style>
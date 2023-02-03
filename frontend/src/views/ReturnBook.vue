<template>
  <div class="return">
    <h1>This is an about page</h1>
  </div>
  <div class="return">
		<h1>返却</h1>
		<!-- スキャン -->
		<div class="sendform">
			<input class="sendinput sendelement" type="text" v-model="isbn_return" @change="bookreturn" placeholder="スキャンしてください" style="margin-bottom: 5px; height:50px;width:86%;"/>
			<button variant="secondary" class="btn sendelement btn-default sendbtn" v-on:click="sendreturn" style="padding: 0.7rem 1.0rem;margin-left:5px;height:50px;width:100px">送信</button>
		</div>
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
</template>

<script>
import axios from 'axios'
const URL = 'localhost/'
export default {
	name: 'ReturnBook',
		data() {
			return {
			checked:false,
			isbn_rental:"",
			isbn_return:"",
			rentalbooks:[],
			returnbooks:[],
			rentalisbn:[],
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
	created: function() {
		axios.post(`${URL}book_list`)
		.then(response => {
			// alert(response.status)
			// console.log(response.data)
			for(var i in response.data){
				console.log(response.data[i].rental_username)
				this.getData = {
					title:response.data[i].title,
					author:response.data[i].author,
					publisher:response.data[i].publisher,
					state: String(response.data[i].rentaluser_name)
				}
				console.log(response.data)
				if(this.getData.state == "null") {
					this.getData.state = '貸出可能'
				}
				this.tableData.push(this.getData)
			}	
		})
		.catch(error => {
			alert('データを送信できませんでした．')
			// alert(error)
		})
	},
	methods: {
		rental: function () {
			var params = new FormData()
			params.append("isbn", this.isbn_rental)
			this.isbn = this.isbn_rental
			axios.post(`${URL}bookinfo`, params)
			.then(response => {
				this.getData = {
					title:response.data.title,
					author:response.data.author,
					publisher:response.data.publisher,
					state: response.data.rental_username,
					isbn: this.isbn
				}
				console.log(this.getData.title)
				// 表に追加処理
				if(this.getData.state == undefined) {
					this.rentalbooks.push({
						title: String(this.getData.title),
						author:String(this.getData.author),
						publisher:String(this.getData.publisher),
					});
					this.sendisbn.push(this.getData.isbn);
					console.log(this.sendisbn)
				} else {
					alert("貸出中です．");
				}
				// this.tableData.push(this.getData)
				// console.log(this.tableData)	
			})
			.catch(error => {
				alert('データを送信できませんでした．')
				alert(error)
			})
			this.isbn_rental = "";	
		},
		bookreturn:function () {
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
			/*サーバにsendisbnをおくる*/
			var params = new FormData()
			params.append('user_id', this.$store.state.userID)
			params.append('isbn', this.sendisbn)
			this.sendisbn = []
			console.log(params)
			axios.post(`${URL}rental_register`, params)
			.then(response => {
				alert(response.data['message'])
				console.log(response.data)
				axios.post(`${URL}book_list`)
				.then(response => {
					// alert(response.status)
					// console.log(response.data)
					this.tableData=[]
					for(var i in response.data){
						console.log(response.data[i].rental_username)
						this.getData = {
							title:response.data[i].title,
							author:response.data[i].author,
							publisher:response.data[i].publisher,
							state: String(response.data[i].rentaluser_name)
						}
						console.log(response.data)
						if(this.getData.state == "null") {
							this.getData.state = '貸出可能'
						}
						this.tableData.push(this.getData)
					}	
				})
				.catch(error => {
					alert('データを送信できませんでした．')
					console.log(error)
					// alert(error)
				})
			})
			.catch(error => {
				alert('データを送信できませんでした．')
				alert(error)
			})
			this.rentalbooks.splice(0)
		},
		sendreturn:function(){
			var params = new FormData()
			params.append('user_id', this.$store.state.userID)
			params.append('isbn', this.sendisbn)
			this.sendisbn = []
			console.log(params['isbn'])
			axios.post(`${URL}return_register`, params)
			.then(response => {
				alert(response.data['message'])
				axios.post(`${URL}book_list`)
				.then(response => {
					// alert(response.status)
					// console.log(response.data)
					this.tableData=[]
					for(var i in response.data){
						console.log(response.data[i].rental_username)
						this.getData = {
							title:response.data[i].title,
							author:response.data[i].author,
							publisher:response.data[i].publisher,
							state: String(response.data[i].rentaluser_name)
						}
						console.log(response.data)
						if(this.getData.state == "null") {
							this.getData.state = '貸出可能'
						}
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
	},
	computed: {	
		filteredbooks: function() {
			var tableData = [];
			var books=[];
			var book;
			for(var j in this.tableData){
				book=this.tableData[j];
				if(this.checked==true){
					if(book.state=='貸出可能'){
						// book.state = '貸出可能'
						books.push(book);
						// console.log(book)
					}
				}
				else{
					books.push(book);
					// console.log(book)
				}
			}
			
			for(var i in books) {
				book = books[i];
				if(book.title.indexOf(this.keyword) !== -1 ) {
					tableData.push(book);
				}
			}
			return tableData;
		},
	}
}
</script>

<style>
.sendform{
	display:flex;
}
/* .formcontrol{
	
	width:86%;
	
} */
/* .sendform.btn{
	
	height:100%;
	padding: 0.7rem 1.0rem;
	margin-left:10px;
	
} */
.contents{
	position:relative;
}
.block1{
	float:left;
	width:45%;
	position:relative;
	padding:20px 60px 60px 60px;
	height:75vh;
}
.block2{
  float:left;
  width:45%;
  margin-top:20px;
  margin-left: 60px;
  margin-bottom: 60px;
  position:relative;
  height: 75vh;
  /* overflow-y: scroll; */
}
div{
  width:100%;
}
h1{
  font-style:normal;
}
/* table */
table {
  position: absolute;
  table-layout: fixed;
  /* top: 130px; */
  size: 500px;
  border-collapse: collapse;
  border-spacing: 5px;
  margin: 0 auto;
  padding: 0;
  width: 100%;
  margin-left: auto;
  margin-right: auto;
}
table tr {
	height: 50px;
}
/* thead */
thead th {
	font-size: 16px;
	/* width: 100px; */
	text-align: center;
	padding: 10px 0;
	/* color */
	color: white;
	background-color:#3373d3;
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
	padding: 10px;
	/* color */
	color: black;
	background-color: white;
	border-bottom: 1px solid #a2a7af;
	border-left:1px solid #a2a7af;
	border-right:1px solid #a2a7af;
}
/* rental_area */
.rental{
  height: 60%;
  position:relative;
}
/* return_area */
.return{
  height:40%;
  position: relative;
}
</style>
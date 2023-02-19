<template>
  <div id="contents">
    <div id="main">
      <section id="pagetop">
        <section class="box">
            <h2 class="title">Book List<span>書籍一覧</span></h2>
          <div>
            <div class="boxes">
          <input type="checkbox" id="box-1" v-model="checked" name="checkbox01">
          <label for="box-1">貸出可能</label>
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
					<tr v-for="tableData in filteredbooks" :key="tableData.name">
						<td class=title><router-link :to="{name:'detail',query:{title:tableData.title}}" v-text="tableData.title"></router-link></td>
						<td class=author v-text="tableData.author"></td>
						<td class=publisher v-text="tableData.publisher"></td>
						<td class=status v-text="tableData.state"></td>
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
	name: 'BookList',
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
	created: function() {
		axios.post(`${URL}book_list`)
		.then(response => {
			// alert(response.status)
			// console.log(response.data)
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
            }
		})
		.catch(error => {
			alert('データを送信できませんでした．')
			// alert(error)
		})
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

<style scoped>
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
  height: 380px;
}
.title{
  min-width: 250px;
}
.author{
  min-width: 250px;
}
.publisher{
  min-width: 250px;
}
.status{
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
</style>
<template>
  <div id="contents">
    <div id="main">
      <section id="pagetop">
        <section class="box">
            <h2 class="title">Book Info<span>書籍情報</span></h2>
            <div class="boxes">
                <img :src="getData.imageurl" class="image">
            </div>
            <div class="boxes">
                <ul class="cp_list">
                    <li>書籍名</li>
                    {{getData.title}}<br>
                    <li>著者</li>
                    {{getData.author}}<br>
                    <li>出版社</li>
                    {{getData.publisher}}<br>
                    <li>内容</li>
                    {{getData.item_caption}}
                </ul>
          </div>
          <!-- <router-link to="/booklist" class="btn btn--purple btn--radius btn--cubic"><i class="fas fa-angle-double-right fa-position-left"></i>戻る</router-link> -->
          <br>
          <button type="button" onclick="history.back()" class="btn btn--yellow btn--cubic">戻る</button>
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
            booktitle: this.$route.query.title,
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
    var params = new FormData()
    params.append("booktitle", this.booktitle)
    axios.post(`${URL}book_detail`, params)
    .then(response => {
        // alert(response.status)
        // console.log(response.data)
        this.getData = {
            title:response.data.title,
            author:response.data.author,
            publisher:response.data.publisher,
            item_caption: response.data.item_caption,
            imageurl: response.data.imageurl
        }
        console.log(response.data)
        if(this.getData.state == "null") {
            this.getData.state = '貸出可能'
        }
        this.tableData.push(this.getData)
    })
    .catch(error => {
        alert('データを送信できませんでした．')
        // alert(error)
    })
  },
  mounted(){
    console.log(this.title)
    console.log(this.$route.query)
    console.log("this.$route.params.title", this.$route.query.title)
  },
}
</script>

<style scoped>
ul.cp_list {
    font-size:15px;
    display: inline-block;
    width:45%;
    position: absolute;
    margin-right: 15%;
    top:30%;
    left:40%;
	padding: 0.5em 1em 0.01em 2.3em;
	list-style-type: none;
}
ul.cp_list li {
	position: relative;
	padding: 0.5em 1em 0.01em 2.3em;
	margin-bottom:1px;
	border-bottom: 1px solid #c1bc9e;
}
ul.cp_list li:after,
ul.cp_list li:before{
	content:'';
	position: absolute;
	border-radius: 50%;
}
ul.cp_list li:before {
	top: 50%;
	left: 0.2em;
	width: 17px;
	height: 17px;
	background: #ebe6ce;
	transform: translateY(-50%);
}
ul.cp_list li:after {
	top: 1.1em;
	left: 0.7em;
	width: 14px;
	height: 14px;
	background: #c1bc9e;
}
.image{
  width:23%;
  left:50px;
  bottom:10%;
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
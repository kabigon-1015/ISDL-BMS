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
						<th>タイトル</th>
						<th>著者</th>
						<th>出版社</th>
						<th>貸出状況</th>
					</tr>
				</thead>
				<tbody>
                    <td v-text="getData.title"></td>
                    <td v-text="getData.author"></td>
                    <td v-text="getData.publisher"></td>
                    <td v-text="getData.item_caption"></td>
                    <td v-text="getData.imageurl"></td>
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
  position: absolute;
  font-size: 20px;
  top: 4px;
  left: 14px;
  width: 5%;
  transition: 0.3s;
  letter-spacing: 0.5px;
  color: #040404;
}
.ef:focus ~ label,
.cp_iptxt.ef ~ label {
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
	background-color:#1b2538;
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
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
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
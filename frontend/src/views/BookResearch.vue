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

            <h2>書籍に関連するタグを選んでください</h2>
            <button v-on:click="research" class="btn btn--yellow btn--cubic">
              検索
            </button>
            <button v-on:click="research" class="btn btn--yellow btn--cubic">
              検索
            </button>
            <button v-on:click="research" class="btn btn--yellow btn--cubic">
              検索
            </button>
            <button v-on:click="research" class="btn btn--yellow btn--cubic">
              検索
            </button>
            <button v-on:click="research" class="btn btn--yellow btn--cubic">
              検索
            </button>
          </div>
          <table>
            <thead>
              <tr>
                <th class="txt book">書籍名</th>
                <th scope="col" class="txt book1">著者</th>
                <th scope="col" class="txt book1">出版社</th>
                <th scope="col" class="txt book1">貸出状況</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <th class="txt">はじめてのPython</th>
                <td data-label="著者" class="txt">
                  岡颯人
                </td>
                <td data-label="出版社" class="txt">ISDL</td>
                <td data-label="貸出状況" class="txt">¥3,000</td>
              </tr>
              <tr>
                <th class="txt">はじめてのPython</th>
                <td data-label="著者" class="txt">
                  岡颯人
                </td>
                <td data-label="出版社" class="txt">ISDL</td>
                <td data-label="貸出状況" class="txt">¥3,000</td>
              </tr>
              <tr>
                <th class="txt">はじめてのPython</th>
                <td data-label="著者" class="txt">
                  岡颯人
                </td>
                <td data-label="出版社" class="txt">ISDL</td>
                <td data-label="貸出状況" class="txt">¥3,000</td>
              </tr>
              <tr>
                <th class="txt">はじめてのPython</th>
                <td data-label="著者" class="txt">
                  岡颯人
                </td>
                <td data-label="出版社" class="txt">ISDL</td>
                <td data-label="貸出状況" class="txt">¥3,000</td>
              </tr>
              <tr>
                <th class="txt">はじめてのPython</th>
                <td data-label="著者" class="txt">
                  岡颯人
                </td>
                <td data-label="出版社" class="txt">ISDL</td>
                <td data-label="貸出状況" class="txt">¥3,000</td>
              </tr>
              <tr>
                <th class="txt">はじめてのPython</th>
                <td data-label="著者" class="txt">
                  岡颯人
                </td>
                <td data-label="出版社" class="txt">ISDL</td>
                <td data-label="貸出状況" class="txt">¥3,000</td>
              </tr>
            </tbody>
          </table>
        </section>
      </section>
      {{ this.book }}
    </div>
    <!--/#main-->
  </div>
  <!--/#contents-->
</template>

<script>
import axios from "axios";
const URL = "localhost/";

export default {
  data: function () {
    return {
      book: "",
    };
  },
  methods: {
    research: function () {
      var params = new FormData();
      params.append("isbn", this.book);
      axios
        .post("/research", params)
        .then((response) => {
          console.log(response.data);
        })
        .catch((error) => {
          alert("データを送信できませんでした．");
        });
    },
  },
};
</script>

<style scoped>
/* body {
  font-family: "Open Sans", sans-serif;
  line-height: 1.25;
} */

table {
  position: absolute;
  right:10%;
  top:35%;
  border-collapse: collapse;
  margin: 0 auto;
  padding: 0;
  width: 650px;
  table-layout: fixed;
}

table tr {
  background-color: #fff;
  border: 1px solid #bbb;
  padding: .35em;
}
table th,
table td {
  padding: 1em 10px 1em 1em;
  border-right: 1px solid #bbb;
}
table th {
  font-size: .85em;
}
table thead tr{
  background-color: #eee;
}
.txt{
   text-align: left;
   font-size: 1.3em;
   color: #000;
}
.price{
   text-align: right;
}
.book{
    width:200px;
}
.book1{
    width:100px;
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

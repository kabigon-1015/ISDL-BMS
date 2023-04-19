<template>
  <div id="contents">
    <div id="main">
      <section id="pagetop">
        <section class="box">
          <h2 class="title">Admin<span>管理者権限ページ</span></h2>
          <h2>タグ追加機能</h2>
            <router-link to="/addbooktag">
                <h4>タグ追加ページへ</h4>
            </router-link>
          <h2>書籍追加機能</h2>
            <router-link to="/addbook">
                <h4>書籍追加ページへ</h4>
            </router-link>
          <h2>書籍削除機能</h2>
            <router-link to="/deletebook">
                <h4>書籍削除ページへ</h4>
            </router-link>
        </section>
      </section>
    </div>
  </div>
</template>
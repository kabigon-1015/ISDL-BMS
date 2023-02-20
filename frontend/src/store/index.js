import { createStore } from 'vuex'

export default createStore({
  state: {
    username:'ログイン',
    userid:'null'
  },
  getters: {
  },
  mutations: {
    loginusername: function(state, value) {
      state.username = value
    },
    getuserid: function(state, value) {
      state.userid = value
    },
  },
  actions: {
    loginusername: function(context, value) {
      context.commit('loginusername', value)
    },
    getuserid: function(context, value) {
      context.commit('getuserid', value)
    },
  },
  modules: {
  }
})

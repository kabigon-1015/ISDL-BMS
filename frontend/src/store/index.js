import { createStore } from 'vuex'

export default createStore({
  state: {
    username:'ログイン',
    userid:'null',
    currentpass:''
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
    getcurrentpass: function(state, value) {
      state.currentpass = value
    },
  },
  actions: {
    loginusername: function(context, value) {
      context.commit('loginusername', value)
    },
    getuserid: function(context, value) {
      context.commit('getuserid', value)
    },
    getcurrentpass: function(context, value) {
      context.commit('getcurrentpass', value)
    },
  },
  modules: {
  }
})

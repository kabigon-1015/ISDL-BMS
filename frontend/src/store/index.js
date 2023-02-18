import { createStore } from 'vuex'

export default createStore({
  state: {
    username:'ログイン'
  },
  getters: {
  },
  mutations: {
    loginusername: function(state, value) {
      state.username = value
    },
  },
  actions: {
    loginusername: function(context, value) {
      context.commit('loginusername', value)
    },
  },
  modules: {
  }
})

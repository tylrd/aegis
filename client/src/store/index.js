import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

export const defaultState = {
  user: {},
  loggedIn: false,
  status: ''
}

export const storeBuilder = (state = defaultState) => {
  return new Vuex.Store({
    state,
    mutations: {
      auth_request (state) {
        state.status = 'loading'
      },
      auth_success (state, user) {
        state.status = 'success'
        state.user = user
        state.loggedIn = true
      },
      auth_error (state) {
        state.status = 'error'
      },
      logout (state) {
        state.status = ''
        state.user = {}
        state.loggedIn = false
      }
    },
    actions: {
      login ({commit}, user) {
        return new Promise((resolve, reject) => {
          commit('auth_request')
          axios({
            url: '/login',
            data: user,
            method: 'POST'
          })
            .then(resp => {
              const user = resp.data.user
              commit('auth_success', user)
              resolve(resp)
            })
            .catch(err => {
              commit('auth_error')
              reject(err)
            })
        })
      },
      logout ({commit}) {
        return new Promise((resolve, reject) => {
          axios({
            url: '/logout',
            method: 'POST'
          })
            .then(() => {
              commit('logout')
              resolve()
            })
            .catch(err => {
              reject(err)
            })
        })
      },
      getSession ({commit}) {
        return new Promise((resolve, reject) => {
          axios({
            url: '/session',
            method: 'GET'
          })
            .then(response => {
              const user = response.data.user
              commit('auth_success', user)
              resolve(user)
            })
            .catch(err => {
              commit('auth_error')
              reject(err)
            })
        })
      }
    },
    getters: {
      isLoggedIn: state => state.loggedIn
    }
  })
}

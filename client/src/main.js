// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './components/App'
import routerBuilder from './router'
import { storeBuilder } from './store'
import axios from 'axios'
import 'bulma/css/bulma.css'

Vue.config.productionTip = false

Vue.prototype.$http = axios
Vue.prototype.$http.defaults.withCredentials = true
Vue.prototype.$http.defaults.baseURL = process.env.BASE_URL

/* eslint-disable no-new */
const createVue = (initialState) => {
  const store = storeBuilder(initialState)
  const router = routerBuilder(store)
  new Vue({
    el: '#app',
    router,
    store,
    template: '<App/>',
    components: { App }
  })
}

axios.get('/session')
  .then(data => {
    const session = {
      user: data,
      loggedIn: true
    }
    createVue(session)
  })
  .catch(() => createVue())

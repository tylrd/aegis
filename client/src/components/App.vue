<template>
  <div id="app">
    <nav class="navbar is-light" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <div class="navbar-item">
          <b>AEGIS</b>
        </div>
        <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample" @click="showNav = !showNav" :class="{ 'is-active': showNav}">
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>

      <div id="navbarBasicExample" class="navbar-menu" :class="{ 'is-active': showNav}">
        <div class="navbar-start">
          <router-link class="navbar-item" to="/">Home</router-link>
          <router-link class="navbar-item" to="/">About</router-link>
          <router-link class="navbar-item" to="/posts" v-if="isLoggedIn">Posts</router-link>
        </div>
        <div class="navbar-end">
          <div class="navbar-item">
            <div class="buttons">
              <a @click="logout" v-if="isLoggedIn" class="button is-primary">Logout</a>
              <router-link class="button is-primary" to="/register">Register</router-link>
              <router-link class="button is-light" to="/login">Login</router-link>
            </div>
          </div>
        </div>
      </div>

    </nav>

    <section class="section">
      <router-view></router-view>
    </section>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default{
  name: 'App',
  data () {
    return {
      showNav: false
    }
  },
  computed: {
    ...mapGetters([
      'isLoggedIn'
    ])
  },
  methods: {
    logout: function () {
      this.$store.dispatch('logout').then(() => {
        this.$router.push('/login')
      })
    }
  }
}
</script>

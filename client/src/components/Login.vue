<template>
  <div class="container login">
    <h1 class="title">Sign in</h1>

    <form class="login" @submit.prevent="login">

      <div class="field">
        <label class="label">Email</label>
        <input required class="input" v-model="email" type="email" placeholder="Name">
      </div>

      <div class="field">
        <label class="label">Password</label>
        <div class="control">
          <input required class="input" v-model="password" type="password" placeholder="Password" @input="test">
        </div>
      </div>

      <button class="button is-primary" type="submit">Login</button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data () {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    login () {
      const email = this.email
      const password = this.password
      this.$store.dispatch('login', { email, password })
        .then(() => this.$router.push('/'))
        .catch(err => {
          console.log('Incorrect login info:' + err)
        })
    }
  }
}
</script>

<style scoped>

.login {
  max-width: 600px;
  margin: 0 auto;
}

</style>

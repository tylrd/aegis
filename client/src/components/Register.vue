<template>
  <div class="register" >
    <h1 class="title">Register</h1>
    <form @submit.prevent="register">

      <div class="field">
        <label class="label">Email</label>
        <div class="control">
          <input required class="input" type="email" v-model="email" placeholder="Email">
        </div>
      </div>

      <div class="field">
        <label class="label">Password</label>
        <div class="control">
          <input required
            class="input"
            type="password"
            v-model="password"
            placeholder="Password"
            @input="onPasswordChange"
            :class="passwordClass"/>
        </div>
      </div>

      <div class="field">
        <label class="label">Password Confirmation</label>
        <div class="control">
          <input required
            class="input"
            type="password"
            placeholder="Password Confirmation"
            v-model="passwordConfirmation"
            @input="onPasswordConfirmationChange"
            :class="confirmationClass"/>
        </div>
      </div>

      <button class="button is-primary" type="submit">Register</button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'Register',
  data () {
    return {
      email: '',
      password: '',
      passwordConfirmation: '',
      passwordMatch: null
    }
  },
  computed: {
    passwordClass: function () {
      return {
        'is-success': this.passwordMatch
      }
    },
    confirmationClass: function () {
      return {
        'is-success': this.passwordConfirmation !== '' && this.passwordMatch,
        'is-danger': this.passwordMatch != null && !this.passwordMatch
      }
    }
  },
  methods: {
    onPasswordConfirmationChange (e) {
      const value = e.target.value
      this.passwordMatch = value !== '' && value === this.password
    },
    onPasswordChange (e) {
      const value = e.target.value
      this.passwordMatch = value !== '' && value === this.passwordConfirmation
    },
    register () {
      console.log('registered!')
    }
  }
}
</script>

<style>
.register {
  max-width: 600px;
  margin: 0 auto;
}
</style>

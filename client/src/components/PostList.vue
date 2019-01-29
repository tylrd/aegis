<template>
  <div id="app-4">
    <ol>
      <li v-for="todo in todos" :key="todo.text">{{ todo.text }}</li>
    </ol>
  </div>
</template>

<script>
export default {
  name: 'PostList',
  data () {
    return {
      todos: [
        { text: 'Learn JavaScript' },
        { text: 'Learn Vue' },
        { text: 'Build something awesome' }
      ]
    }
  }
}
</script>

<style>
</style>

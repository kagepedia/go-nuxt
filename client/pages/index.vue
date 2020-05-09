<template>
  <v-layout column justify-center align-center>
    <div>
      <h1>task一覧</h1>
      <b>{{ data.Task }}</b>
    </div>
  </v-layout>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      data: {}
    }
  },
  mounted() {
    axios.get('/api/tasks').then((res) => (this.data = res.data))
  }
}
</script>

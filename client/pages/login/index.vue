<template>
  <v-layout column justify-center align-center>
    <h1>ログイン画面</h1>
    <table>
      <tr>
        <th>ID</th>
        <td>
          <input id="id" v-model="ID" type="text" />
        </td>
      </tr>
      <tr>
        <th>PW</th>
        <td>
          <input id="password" v-model="PW" type="password" />
        </td>
      </tr>
      <tr>
        <td colspan="1">
          <button @click="login()">送信</button>
        </td>
      </tr>
    </table>
  </v-layout>
</template>

<style scoped></style>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      ID: '',
      PW: ''
    }
  },
  methods: {
    login() {
      const params = new FormData()
      params.append('ID', this.ID)
      params.append('PW', this.PW)
      console.log(this.ID, this.PW)
      axios.post('/api/login', params).then((res) => {
        if (res.status === 200) {
          this.$store.dispatch('userLogin', true)
          console.log(res)
          this.$router.push('/')
        }
      })
    }
  }
}
</script>

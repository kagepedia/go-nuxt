<template>
  <v-layout column justify-center align-center>
    <h1>登録画面</h1>
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
        <th>メールアドレス</th>
        <td>
          <input id="email" v-model="MAIL" type="email" />
        </td>
      </tr>
      <tr>
        <th>メールアドレス(確認)</th>
        <td>
          <input id="email_conf" v-model="MAIL_CONF" type="email" />
        </td>
      </tr>
      <tr>
        <td colspan="1">
          <button @click="registration()">送信</button>
        </td>
      </tr>
    </table>
  </v-layout>
</template>

<style scoped></style>

<script>
// import axios from 'axios'

export default {
  data() {
    return {
      ID: '',
      PW: '',
      MAIL: '',
      MAIL_CONF: ''
    }
  },
  methods: {
    registration() {
      const params = new FormData()
      params.append('ID', this.ID)
      params.append('PW', this.PW)
      params.append('MAIL', this.MAIL)
      params.append('MAIL_CONF', this.MAIL_CONF)
      console.log(this.ID, this.PW, this.MAIL, this.MAIL_CONF)
      /*
      axios.post('/api/cretate', params).then((res) => {
        if (res.status === 200) {
        }
      })
      */
    }
  }
}
</script>

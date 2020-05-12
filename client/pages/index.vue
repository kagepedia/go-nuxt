<template>
  <v-layout column justify-center align-center>
    <div>
      <h1>task一覧</h1>
      <table class="test">
        <tr>
          <th>日付</th>
          <th>集合場所</th>
          <th>操作</th>
          <th>操作</th>
        </tr>
        <tr v-for="data of datas" :key="data.Id">
          <td>{{ data.Id }}</td>
          <td>{{ data.Task }}</td>
          <td><button @click="edit(data.Id)">編集</button></td>
          <td><button @click="del(data.Id)">削除</button></td>
        </tr>
      </table>
    </div>
  </v-layout>
</template>

<style scoped>
table.test {
  border-spacing: 20px 5px;
}
</style>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      datas: []
    }
  },
  mounted() {
    axios.get('/api/tasks').then((res) => (this.datas = res.data))
  },
  methods: {
    edit(id) {
      console.log(id)
    },
    del(id) {
      console.log(id)
    }
  }
}
</script>

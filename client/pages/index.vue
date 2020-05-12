<template>
  <v-layout column justify-center align-center>
    <div>
      <h2>新しいtaskの追加</h2>
      <form class="add-form" @submit.prevent="doAdd">
        タスク <input ref="comment" type="text" />
        <button type="submit">追加</button>
      </form>
      <h2>task一覧</h2>
      <table class="test">
        <tr>
          <th>id</th>
          <th>task</th>
          <th>作成時間</th>
          <th>更新時間</th>
          <th>操作</th>
          <th>操作</th>
        </tr>
        <tr v-for="data of datas" :key="data.Id">
          <td>{{ data.Id }}</td>
          <td>{{ data.Task }}</td>
          <td>{{ data.CreateAt }}</td>
          <td>{{ data.UpdateAt }}</td>
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
    taskAdd() {
      console.log('task add')
    },
    edit(id) {
      console.log(id)
    },
    del(id) {
      console.log(id)
      axios.post('/api/taskdelete', { id })
    }
  }
}
</script>

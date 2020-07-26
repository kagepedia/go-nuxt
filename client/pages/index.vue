<template>
  <v-layout column justify-center align-center>
    <div>
      <nuxt-link to="/Create">Create</nuxt-link>
      <h2>新しいtaskの追加</h2>
      <div>
        タスク
        <input v-model="newTask" type="text" />
        <input type="file" @change="onChange" />
        <button @click="taskAdd()">追加</button>
      </div>
      <h2>taskの編集</h2>
      <div>
        タスク
        <input v-model="editTask" type="text" />
        <button @click="edit(editId)">データを更新</button>
      </div>
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
          <td>
            <button @click="find(data.Id)">編集</button>
          </td>
          <td>
            <button @click="del(data.Id)">削除</button>
          </td>
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
      newTask: '',
      editTask: '',
      editId: '',
      file: null,
      datas: []
    }
  },
  mounted() {
    axios.get('/api/read').then((res) => (this.datas = res.data))
  },
  methods: {
    onChange(event) {
      this.file = event.target.files[0]
    },
    taskAdd() {
      const params = new FormData()
      params.append('task', this.newTask)
      params.append('file', this.file)
      axios
        .post('/api/cretate', params, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })
        .then((res) => {
          if (res.status === 200) {
            this.newTask = ''
            location.reload()
          }
        })
    },
    find(id) {
      const params = new URLSearchParams()
      params.append('id', id)
      axios.post('/api/find', params).then((res) => {
        if (res.status === 200) {
          this.editId = res.data.Id
          this.editTask = res.data.Task
        }
      })
    },
    edit(id) {
      const params = new URLSearchParams()
      params.append('id', id)
      params.append('task', this.editTask)
      axios.post('/api/update', params).then((res) => {
        if (res.status === 200) {
          location.reload()
        }
      })
    },
    del(id) {
      const params = new URLSearchParams()
      params.append('id', id)
      axios.post('/api/delete', params).then((res) => {
        if (res.status === 200) {
          location.reload()
        }
      })
    }
  }
}
</script>

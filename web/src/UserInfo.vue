<template>
  <div class="info">
    <h1>用户信息管理界面</h1>
    <el-row>
      <el-col :span="20" :push='2'>
        <div>
          <el-form :inline="true">
            <el-form-item style="float: left" label="查询用户信息:">
              <el-input v-model="keyUser" placeholder="查询所需要的内容......"></el-input>
            </el-form-item>
            <el-form-item style="float: right">
              <el-button type="primary" size="small" icon="el-icon-edit-outline" @click="hanldeAdd()">添加用户</el-button>
            </el-form-item>
            <el-form-item style="float: right">
              <el-button type="primary" size="small" @click="getUserList()">显示所有用户</el-button>
            </el-form-item>
          </el-form>
        </div>
        <div class="table">
          <el-table
            :data="userList"
            border
            style="width: 100%">
            <el-table-column
              label="ID"
              align="center"
              prop="UserId">

            </el-table-column>
            <el-table-column
              label="昵称"
              align="center"
              prop="NickName">

            </el-table-column>
            <el-table-column
              label="角色"
              align="center"
              prop="Role">

            </el-table-column>
            <el-table-column
              label="登入时间"
              align="center"
              prop="LoginTime">

            </el-table-column>
          </el-table>
        </div>
      </el-col>
    </el-row>
    <AddUser :dialogAdd="dialogAdd" @update="getUserInfo"></AddUser>
  </div>
</template>

<script>
import AddUser from './AddUser.vue'
export default {
  name: 'info',
  data() {
    return {
      userList: [],
      form: {    //编辑信息
        UserId: '',
        NickName: '',
        Role: '',
        LoginTime: ''
      },
    }
  },
  methods: {
    getUserList() {
      this.$axios.get('http://localhost:8000/api/get_user_list').then(res => {
        this.userList = res.data
        console.log(this.userList)
      })
    },
    components: {
      AddUser
    },
  },
  created() {
    this.getUserList()
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1{
  font-size: 30px;
  color: #333;
  text-align: center;
  margin: 0 auto;
  padding-bottom: 5px;
  border-bottom: 2px solid #409EFF;
  width: 300px
}
</style>

<template>
  <div class="password-conaitner">
    <el-form ref="form" :model="formData" :rules="rules" label-width="100px">
      <el-form-item label="用户名">
        <el-input v-model="userName" autocomplete="off" disabled />
      </el-form-item>
      <el-form-item label="当前密码" prop="password" required>
        <el-input v-model="formData.password" type="password" />
      </el-form-item>
      <el-form-item label="新密码" prop="newPassword" required>
        <el-input v-model="formData.newPassword" type="password" />
      </el-form-item>
      <el-form-item label="确认新密码" prop="newPassword1" required>
        <el-input v-model="formData.newPassword1" type="password" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">确定</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import * as userApi from '@/api/user'
import { mapGetters } from 'vuex'

export default {
  name: 'Password',
  data() {
    return {
      formData: {
        password: '',
        newPassword: '',
        newPassword1: ''
      },
      rules: {
        password: [
          { required: true, message: '请输入当前密码', trigger: 'blur' }
        ],
        newPassword: [
          { required: true, message: '请输入新密码', trigger: 'blur' }
        ],
        newPassword1: [
          { required: true, message: '请输入确认新密码', trigger: 'blur' }
        ]
      }
    }
  },
  computed: {
    ...mapGetters({
      userName: 'name',
      userId: 'userId'
    })
  },
  methods: {
    onSubmit() {
      if (this.formData.newPassword !== this.formData.newPassword1) {
        this.$message({
          message: '新密码和确认密码不一致',
          type: 'error'
        })
      }
      this.$refs['form'].validate((valid) => {
        if (valid) {
          userApi.updatePassword(this.userId, this.userName, this.formData.password, this.formData.newPassword).then(d => {
            this.$message({
              message: '修改密码成功',
              type: 'success'
            })
          })
        }
      })
    }
  }
}
</script>
<style scoped>
.password-conaitner {
  height: 100%;
  padding: 10px;
  width: 60%;
  margin-left: 20%;
}
</style>

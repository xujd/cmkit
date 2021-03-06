<template>
  <div class="user-conaitner">
    <el-card class="search-box">
      <div style="position:relative;">
        <el-form :inline="true" :model="formData">
          <el-form-item label="用户名">
            <el-input v-model="formData.name" clearable placeholder="用户名" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">查询</el-button>
          </el-form-item>
        </el-form>
        <el-button type="success" style="position: absolute; right: 0; top: 0;" @click="onAddNew">新增</el-button>
      </div>
    </el-card>
    <el-card class="result-box">
      <el-table height="300" :data="tableData" style="width: 100%">
        <el-table-column prop="id" width="80" label="编号" />
        <el-table-column prop="name" width="120" label="用户名" />
        <el-table-column prop="staffName" width="120" label="员工姓名" />
        <el-table-column prop="startTimeStr" width="180" label="生效开始时间" />
        <el-table-column prop="endTimeStr" width="180" label="生效结束时间" />
        <el-table-column prop="statusStr" label="状态" width="80">
          <template slot-scope="scope">
            <el-tag
              :type="scope.row.status === 0 ? 'success' : 'danger'"
              disable-transitions
            >{{ scope.row.statusStr }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="remark" width="250" label="备注" />
        <el-table-column fixed="right" label="操作" width="160">
          <template slot-scope="scope">
            <el-button
              v-show="scope.row.id !== 1"
              type="text"
              size="small"
              @click="handleEditClick(scope.row)"
            >编辑</el-button>
            <el-button
              v-show="scope.row.id !== 1"
              type="text"
              size="small"
              @click="handleDeleteClick(scope.row)"
            >删除</el-button>
            <el-button
              v-show="scope.row.id !== 1"
              type="text"
              size="small"
              @click="handleResetPwdClick(scope.row)"
            >重置</el-button>
            <el-button
              v-show="scope.row.id !== 1"
              type="text"
              size="small"
              @click="handleRoleClick(scope.row)"
            >角色</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        :current-page="curPageIndex"
        :page-sizes="[10, 20, 50, 100]"
        :page-size="curPageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="dataTotal"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </el-card>
    <el-dialog :title="userNewTitle" :visible.sync="isAddUserVisible" width="40%" top="10px">
      <UserNew v-if="isAddUserVisible" ref="userNew" :user="curUser" />
      <span slot="footer" class="dialog-footer">
        <el-button @click="isAddUserVisible = false">取 消</el-button>
        <el-button type="primary" @click="onAddUserOK">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog title="设置角色" :visible.sync="isSetRoleVisible" width="40%" top="10px">
      <UserRole ref="userRole" />
      <span slot="footer" class="dialog-footer">
        <el-button @click="isSetRoleVisible = false">取 消</el-button>
        <el-button type="primary" @click="onSetUserRoleOK">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import * as userApi from '@/api/user'
import { mapGetters } from 'vuex'
import UserNew from './components/UserNew'
import UserRole from './components/UserRole'
import dayjs from 'dayjs'

export default {
  name: 'User',
  components: {
    UserNew,
    UserRole
  },
  data() {
    return {
      userNewTitle: '添加用户',
      isAddUserVisible: false,
      curUser: null,
      isSetRoleVisible: false,
      formData: {
        name: ''
      },
      tableData: [],
      dataTotal: 0,
      curPageIndex: 1,
      curPageSize: 10
    }
  },
  computed: {
    ...mapGetters({
      userId: 'userId'
    })
  },
  mounted() {
    this.queryUsers()
  },
  methods: {
    onSubmit() {
      this.queryUsers()
    },
    onAddNew() {
      this.curUser = null
      this.userNewTitle = '添加用户'
      this.isAddUserVisible = true
    },
    onAddUserOK() {
      const data = this.$refs.userNew.getData()
      if (!data.id) {
        userApi.addUser(data).then(d => {
          this.$message({
            message: '添加成功',
            type: 'success'
          })
          this.isAddUserVisible = false
          this.queryUsers()
        })
      } else {
        userApi.updateUser(data).then(d => {
          this.$message({
            message: '修改成功',
            type: 'success'
          })
          this.isAddUserVisible = false
          this.queryUsers()
        })
      }
    },
    handleSizeChange(val) {
      this.curPageSize = val
      this.queryUsers()
    },
    handleCurrentChange(val) {
      this.curPageIndex = val
      this.queryUsers()
    },
    handleEditClick(data) {
      this.curUser = data
      this.userNewTitle = '修改用户'
      this.isAddUserVisible = true
    },
    handleDeleteClick(data) {
      if (this.userId === data.id) {
        this.$message({
          message: '当前登录用户无法删除',
          type: 'error'
        })
        return
      }
      this.$confirm('确认删除该用户？', '删除', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        userApi.deleteUser(data.id).then(d => {
          if (d.code === 20000) {
            this.$message({
              message: '删除成功',
              type: 'success'
            })
            // 当前页只有一条数据时，查询前一页数据
            if (this.tableData.length === 1 && this.curPageIndex > 1) {
              this.curPageIndex -= 1
            }
            this.queryUsers()
          }
        })
      })
    },
    handleResetPwdClick(data) {
      this.$confirm('确认重置用户密码？', '重置密码', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        userApi.resetPassword(data.id).then(d => {
          if (d.code === 20000) {
            this.$message({
              message: '重置成功',
              type: 'success'
            })
          }
        })
      })
    },
    handleRoleClick(data) {
      this.curUser = data
      this.isSetRoleVisible = true
      this.$nextTick(() => {
        this.$refs.userRole.resetView(data.id)
      })
    },
    onSetUserRoleOK() {
      userApi.setUserRole(this.curUser.id, this.$refs.userRole.getSelectedRoles()).then(d => {
        this.$message({
          message: '设置角色成功',
          type: 'success'
        })
        this.isSetRoleVisible = false
      })
    },
    queryUsers() {
      userApi.queryUsers(this.formData.name, this.curPageSize, this.curPageIndex).then(d => {
        d.data.list.forEach(item => {
          item.statusStr = item.status === 0 ? '有效' : '无效'
          item.startTimeStr = item.startTime || '-'
          item.endTimeStr = item.endTime || '-'
          // 已失效
          if (item.endTime && item.status === 0 && dayjs(item.endTime).isBefore(dayjs())) {
            item.statusStr = '已过期'
          }
        })
        this.tableData = d.data.list
        this.dataTotal = d.data.total
      })
    }
  }
}
</script>
<style scoped>
.user-conaitner {
  width: 100%;
  height: 100%;
  padding: 10px;
}

.search-box ::v-deep .el-form-item {
  margin-bottom: 0;
}

.result-box {
  margin-top: 10px;
}
</style>

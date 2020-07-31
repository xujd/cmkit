<template>
  <div class="role-conaitner">
    <el-card class="search-box">
      <div style="position:relative;">
        <el-form :inline="true" :model="formData">
          <el-form-item label="角色名称">
            <el-input v-model="formData.name" clearable placeholder="角色名称" />
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
        <el-table-column prop="id" label="编号" />
        <el-table-column prop="name" label="角色名称" />
        <el-table-column prop="statusStr" label="状态" />
        <el-table-column prop="remark" label="备注" />
        <el-table-column fixed="right" label="操作" width="100">
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
    <el-dialog :title="roleNewTitle" :visible.sync="isAddRoleVisible" width="40%" top="10px">
      <RoleNew v-if="isAddRoleVisible" ref="roleNew" :role="curRole" />
      <span slot="footer" class="dialog-footer">
        <el-button @click="isAddRoleVisible = false">取 消</el-button>
        <el-button type="primary" @click="onAddRoleOK">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import * as roleApi from '@/api/role'
import { MessageBox, Message } from 'element-ui'
import RoleNew from './components/RoleNew'

export default {
  name: 'Role',
  components: {
    RoleNew
  },
  data() {
    return {
      roleNewTitle: '添加角色',
      isAddRoleVisible: false,
      curRole: null,
      formData: {
        name: '',
        companyId: null,
        departmentId: null
      },
      tableData: [],
      dataTotal: 0,
      curPageIndex: 1,
      curPageSize: 10
    }
  },
  mounted() {
    this.queryRoles()
  },
  methods: {
    onSubmit() {
      this.queryRoles()
    },
    onAddNew() {
      this.curRole = null
      this.roleNewTitle = '添加角色'
      this.isAddRoleVisible = true
    },
    onAddRoleOK() {
      const data = this.$refs.roleNew.getData()
      if (!data.id) {
        roleApi.addRole(data).then(d => {
          Message({
            message: '添加成功',
            type: 'success'
          })
          this.isAddRoleVisible = false
          this.queryRoles()
        })
      } else {
        roleApi.updateRole(data).then(d => {
          Message({
            message: '修改成功',
            type: 'success'
          })
          this.isAddRoleVisible = false
          this.queryRoles()
        })
      }
    },
    handleSizeChange(val) {
      this.curPageSize = val
      this.queryRoles()
    },
    handleCurrentChange(val) {
      this.curPageIndex = val
      this.queryRoles()
    },
    handleEditClick(data) {
      this.curRole = data
      this.roleNewTitle = '修改角色'
      this.isAddRoleVisible = true
    },
    handleDeleteClick(data) {
      MessageBox.confirm('确认删除该角色？', '删除', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        roleApi.deleteRole(data.id).then(d => {
          if (d.code === 20000) {
            Message({
              message: '删除成功',
              type: 'success'
            })
            // 当前页只有一条数据时，查询前一页数据
            if (this.tableData.length === 1 && this.curPageIndex > 1) {
              this.curPageIndex -= 1
            }
            this.queryRoles()
          }
        })
      })
    },
    queryRoles() {
      roleApi.getRoles(this.formData.name, this.curPageSize, this.curPageIndex).then(d => {
        d.data.list.forEach(item => {
          item.statusStr = item.status === 0 ? '有效' : '无效'
        })
        this.tableData = d.data.list
        this.dataTotal = d.data.total
      })
    }
  }
}
</script>
<style scoped>
.role-conaitner {
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

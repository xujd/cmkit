<template>
  <el-table
    ref="table"
    height="300"
    :data="tableData"
    style="width: 100%"
    @selection-change="handleSelectionChange"
  >
    <el-table-column type="selection" width="55" />
    <el-table-column prop="id" label="编号" />
    <el-table-column prop="name" label="角色名称" />
    <el-table-column prop="statusStr" label="状态" />
    <el-table-column prop="remark" label="备注" />
  </el-table>
</template>
<script>
import * as roleApi from '@/api/role'
import * as userApi from '@/api/user'
import * as _ from 'lodash'

export default {
  name: 'UserRole',
  data() {
    return {
      tableData: [],
      selectedRoles: []
    }
  },
  mounted() {
    roleApi.getRoles('', 100, 1).then(d => {
      d.data.list.forEach(item => {
        item.statusStr = item.status === 0 ? '有效' : '无效'
      })
      this.tableData = d.data.list.filter(item => item.status === 0)
    })
  },
  methods: {
    resetView(userId) {
      this.$refs.table.clearSelection()
      userApi.getUserRole(userId).then(d => {
        d.data.forEach(item => {
          const row = this.tableData.find(role => role.id === item.roleId)
          if (row) {
            this.$refs.table.toggleRowSelection(row)
          }
        })
      })
    },
    handleSelectionChange(val) {
      this.selectedRoles = val
    },
    getSelectedRoles() {
      return _.map(this.selectedRoles, 'id')
    }
  }
}
</script>
<style scoped>
</style>

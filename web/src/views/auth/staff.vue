<template>
  <div class="staff-conaitner">
    <el-card class="search-box">
      <div style="position:relative;">
        <el-form :inline="true" :model="formData">
          <el-form-item label="员工姓名">
            <el-input v-model="formData.name" clearable placeholder="姓名" />
          </el-form-item>
          <el-form-item label="所在公司">
            <el-select
              v-model="formData.companyId"
              clearable
              placeholder="请选择"
              @change="onCompanyChange"
            >
              <el-option
                v-for="item in companys"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="所在部门">
            <el-select v-model="formData.departmentId" clearable placeholder="请选择">
              <el-option
                v-for="item in departments"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
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
        <el-table-column prop="id" label="编号" width="60" />
        <el-table-column prop="name" label="员工姓名" width="100" />
        <el-table-column prop="companyName" label="所在公司" />
        <el-table-column prop="departmentName" label="所在部门" />
        <el-table-column prop="birthdayStr" label="出生日期" />
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
    <el-dialog :title="staffNewTitle" :visible.sync="isAddStaffVisible" width="40%" top="10px">
      <StaffNew v-if="isAddStaffVisible" ref="staffNew" :staff="curStaff" />
      <span slot="footer" class="dialog-footer">
        <el-button @click="isAddStaffVisible = false">取 消</el-button>
        <el-button type="primary" @click="onAddStaffOK">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import * as staffApi from '@/api/staff'
import * as sysApi from '@/api/sys'
import StaffNew from './components/StaffNew'

export default {
  name: 'Staff',
  components: {
    StaffNew
  },
  data() {
    return {
      staffNewTitle: '添加员工',
      isAddStaffVisible: false,
      curStaff: null,
      formData: {
        name: '',
        companyId: null,
        departmentId: null
      },
      companys: [],
      departments: [],
      tableData: [],
      dataTotal: 0,
      curPageIndex: 1,
      curPageSize: 10
    }
  },
  mounted() {
    sysApi.queryCompanys('').then(d => {
      this.companys = d.data.list
    })

    this.queryStaffs()
  },
  methods: {
    onSubmit() {
      this.queryStaffs()
    },
    onCompanyChange() {
      if (!this.formData.companyId) {
        this.departments = []
      } else {
        sysApi.queryDepartments('', this.formData.companyId).then(d => {
          this.departments = d.data.list
        })
      }
    },
    onAddNew() {
      this.curStaff = null
      this.staffNewTitle = '添加员工'
      this.isAddStaffVisible = true
    },
    onAddStaffOK() {
      const data = this.$refs.staffNew.getData()
      if (!data.id) {
        staffApi.addStaff(data).then(d => {
          this.$message({
            message: '添加成功',
            type: 'success'
          })
          this.isAddStaffVisible = false
          this.queryStaffs()
        })
      } else {
        staffApi.updateStaff(data).then(d => {
          this.$message({
            message: '修改成功',
            type: 'success'
          })
          this.isAddStaffVisible = false
          this.queryStaffs()
        })
      }
    },
    handleSizeChange(val) {
      this.curPageSize = val
      this.queryStaffs()
    },
    handleCurrentChange(val) {
      this.curPageIndex = val
      this.queryStaffs()
    },
    handleEditClick(data) {
      this.curStaff = data
      this.staffNewTitle = '修改员工'
      this.isAddStaffVisible = true
    },
    handleDeleteClick(data) {
      this.$confirm('确认删除该员工？', '删除', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        staffApi.deleteStaff(data.id).then(d => {
          if (d.code === 20000) {
            this.$message({
              message: '删除成功',
              type: 'success'
            })
            // 当前页只有一条数据时，查询前一页数据
            if (this.tableData.length === 1 && this.curPageIndex > 1) {
              this.curPageIndex -= 1
            }
            this.queryStaffs()
          }
        })
      })
    },
    queryStaffs() {
      staffApi.queryStaffs(this.formData, this.curPageSize, this.curPageIndex).then(d => {
        d.data.list.forEach(item => {
          item.birthdayStr = item.birthday || '-'
          item.birthdayStr = item.birthdayStr.split(' ')[0]
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
.staff-conaitner {
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

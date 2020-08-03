<template>
  <div class="return-conaitner">
    <el-card class="search-box">
      <div style="position:relative;">
        <el-form :inline="true" :model="formData">
          <el-form-item label="吊索具名称">
            <el-input v-model="formData.name" clearable placeholder="吊索具名称" />
          </el-form-item>
          <el-form-item label="借用人">
            <el-select v-model="formData.staffId" filterable clearable placeholder="请选择">
              <el-option
                v-for="item in staffList"
                :key="item.id"
                :label="item.label"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="借用时间">
            <el-date-picker
              v-model="formData.takeTimes"
              type="datetimerange"
              :value-format="'yyyy-MM-dd HH:mm:ss'"
              range-separator="至"
              start-placeholder="开始日期"
              end-placeholder="结束日期"
            />
          </el-form-item>
        </el-form>
        <div class="action-div">
          <el-button type="primary" @click="onSubmit">查询</el-button>
        </div>
      </div>
    </el-card>
    <el-card class="result-box">
      <el-table height="300" :data="tableData" style="width: 100%">
        <el-table-column prop="id" label="序号" width="80" />
        <el-table-column prop="resName" label="吊索具名称" width="120" />
        <el-table-column prop="takeStaffName" label="借用人" width="80" />
        <el-table-column prop="takeTime" label="借用时间" width="180" />
        <el-table-column prop="returnPlanTime" label="预计归还时间" width="180" />
        <el-table-column prop="returnStaffName" label="归还人" width="100" />
        <el-table-column prop="returnTime" label="归还时间" width="180" />
        <el-table-column prop="remark" label="用途说明" width="180" />
        <el-table-column fixed="right" label="操作" width="100">
          <template slot-scope="scope">
            <el-button type="text" size="small" @click="handleEditClick(scope.row)">归还</el-button>
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
    <el-dialog :title="returnTitle" :visible.sync="isReturnVisible" width="50%" top="10px">
      <ReturnNew v-if="isReturnVisible" ref="returnNew" :use-log="curSling" />
      <span slot="footer" class="dialog-footer">
        <el-button @click="isReturnVisible = false">取 消</el-button>
        <el-button type="primary" @click="onReturnOK">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import * as applyApi from '@/api/apply'
import { queryStaffs } from '@/api/staff'
import ReturnNew from './components/ReturnNew'

export default {
  name: 'Return',
  components: {
    ReturnNew
  },
  data() {
    return {
      returnTitle: '归还吊索具',
      isReturnVisible: false,
      curSling: null,
      formData: {
        name: '',
        staffId: null,
        takeTimes: []
      },
      tableData: [],
      dataTotal: 0,
      curPageIndex: 1,
      curPageSize: 10,
      staffList: []
    }
  },
  computed: {
  },
  mounted() {
    this.queryUseLogs()
    queryStaffs({ name: '' }, 1000, 1).then(d => {
      d.data.list.forEach(item => {
        item.label = item.id + '|' + item.name
      })
      this.staffList = d.data.list
    })
  },
  methods: {
    onSubmit() {
      this.queryUseLogs()
    },
    onReturnOK() {
      const data = this.$refs.returnNew.getData()
      applyApi.returnSling(data).then(d => {
        this.$message({
          message: '归还成功',
          type: 'success'
        })
        this.queryUseLogs()
        this.isReturnVisible = false
      })
    },
    handleSizeChange(val) {
      this.curPageSize = val
      this.queryUseLogs()
    },
    handleCurrentChange(val) {
      this.curPageIndex = val
      this.queryUseLogs()
    },
    handleEditClick(data) {
      this.curSling = data
      this.isReturnVisible = true
    },
    queryUseLogs() {
      const query = {
        resName: this.formData.name,
        returnFlag: 2,
        takeStaff: this.formData.staffId
      }
      if (this.formData.takeTimes.length === 2) {
        query['takeStartTime'] = this.formData.takeTimes[0]
        query['takeEndTime'] = this.formData.takeTimes[1]
      }
      applyApi.getResUseLogs(query, this.curPageSize, this.curPageIndex).then(d => {
        this.tableData = d.data.list
        this.dataTotal = d.data.total
      })
    }
  }
}
</script>
<style scoped>
.return-conaitner {
  width: 100%;
  height: 100%;
  padding: 10px;
}

.search-box ::v-deep .el-form-item {
  margin-bottom: 10px;
}
.action-div {
  position: absolute;
  right: 0px;
  top: 0px;
  width: 60px;
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  justify-content: center;
}
.result-box {
  margin-top: 10px;
}
</style>

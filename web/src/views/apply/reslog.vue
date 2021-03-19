<template>
  <div class="reslog-conaitner">
    <el-card class="search-box">
      <div style="position:relative;">
        <el-form :inline="true" :model="formData">
          <el-form-item label="吊索具名称">
            <el-input v-model="formData.name" clearable placeholder="吊索具名称" />
          </el-form-item>
          <el-form-item label="借用人">
            <el-select v-model="formData.takeStaffId" filterable clearable placeholder="请选择">
              <el-option
                v-for="item in staffList"
                :key="item.id"
                :label="item.label"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="归还人">
            <el-select v-model="formData.returnStaffId" filterable clearable placeholder="请选择">
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
          <el-form-item label="归还状态">
            <el-select v-model="formData.returnFlag" filterable clearable placeholder="请选择">
              <el-option
                v-for="item in flagList"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
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
        <el-table-column prop="takeStaffName" label="借用人员" width="80" />
        <el-table-column prop="createAt" label="借用时间" width="180" />
        <el-table-column prop="takeTime" label="出柜时间" width="180" />
        <el-table-column prop="returnPlanTime" label="预计归还时间" width="180" />
        <el-table-column prop="returnStaffName" label="归还人员" width="100" />
        <el-table-column prop="returnTime" label="归还时间" width="180" />
        <el-table-column prop="duration" label="使用时长" width="180" />
        <el-table-column prop="remark" label="用途说明" width="180" />
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
  </div>
</template>
<script>
import * as applyApi from '@/api/apply'
import { queryStaffs } from '@/api/staff'
import dayjs from 'dayjs'
export default {
  name: 'Return',
  components: {
  },
  data() {
    return {
      formData: {
        name: '',
        takeStaffId: null,
        returnStaffId: null,
        takeTimes: [],
        returnFlag: 0
      },
      tableData: [],
      dataTotal: 0,
      curPageIndex: 1,
      curPageSize: 10,
      staffList: [],
      flagList: [
        { id: 0, name: '全部' },
        { id: 1, name: '已归还' },
        { id: 2, name: '未归还' }
      ]
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
    formatDuring(mss) {
      var days = parseInt(mss / (1000 * 60 * 60 * 24))
      var hours = parseInt((mss % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))
      var minutes = parseInt((mss % (1000 * 60 * 60)) / (1000 * 60))
      var seconds = (mss % (1000 * 60)) / 1000
      return days + ' 天 ' + hours + ' 小时 ' + minutes + ' 分钟 ' + seconds + ' 秒 '
    },
    queryUseLogs() {
      const query = {
        resName: this.formData.name,
        returnFlag: this.formData.returnFlag,
        takeStaff: this.formData.takeStaffId,
        returnStaff: this.formData.returnStaffId
      }
      if (this.formData.takeTimes && this.formData.takeTimes.length === 2) {
        query['takeStartTime'] = this.formData.takeTimes[0]
        query['takeEndTime'] = this.formData.takeTimes[1]
      }
      applyApi.getResUseLogs(query, this.curPageSize, this.curPageIndex).then(d => {
        d.data.list.forEach(item => {
          if (item.returnTime) {
            let dur = dayjs(item.returnTime).diff(dayjs(item.createAt))
            if (dur < 0) {
              dur = 0
            }
            item.duration = this.formatDuring(dur)
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
.reslog-conaitner {
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

<template>
  <div class="sling-conaitner">
    <el-card class="search-box">
      <div style="position:relative;">
        <el-form
          style="width:calc(100% - 60px);"
          :inline="true"
          :model="formData"
          label-width="100px"
        >
          <el-form-item label="吊索具名称">
            <el-input v-model="formData.name" clearable placeholder="吊索具名称" />
          </el-form-item>
          <el-form-item label="型号">
            <el-select v-model="formData.slingType" clearable placeholder="请选择">
              <el-option
                v-for="item in slingTypes"
                :key="item.key"
                :label="item.name"
                :value="item.key"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="最大吨位">
            <el-select v-model="formData.maxTonnage" clearable placeholder="请选择">
              <el-option
                v-for="item in slingTons"
                :key="item.key"
                :label="item.name"
                :value="item.key"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="使用状态">
            <el-select v-model="formData.useStatus" clearable placeholder="请选择">
              <el-option
                v-for="item in slingUseStatus"
                :key="item.key"
                :label="item.name"
                :value="item.key"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="点检状态">
            <el-select v-model="formData.inspectStatus" clearable placeholder="请选择">
              <el-option
                v-for="item in slingInspectStatus"
                :key="item.key"
                :label="item.name"
                :value="item.key"
              />
            </el-select>
          </el-form-item>
        </el-form>
        <div class="action-div">
          <el-button type="primary" @click="onSubmit">查询</el-button>
          <el-button type="success" style="margin-left: 0; margin-top: 10px;" @click="onAddNew">新增</el-button>
        </div>
      </div>
    </el-card>
    <el-card class="result-box">
      <el-table height="300" :data="tableData" style="width: 100%">
        <el-table-column prop="id" label="编号" width="80" />
        <el-table-column prop="rfId" label="RFID" width="120" />
        <el-table-column prop="name" label="吊索具名称" width="120" />
        <el-table-column prop="slingTypeStr" label="型号" width="80" />
        <el-table-column prop="maxTonnageStr" label="最大吨位" width="100" />
        <el-table-column prop="useCount" label="使用次数" width="100" />
        <el-table-column prop="useStatusStr" label="使用状态" width="100" />
        <el-table-column prop="inspectStatusStr" label="点检状态" width="100" />
        <el-table-column prop="putTime" label="投用日期" width="180" />
        <el-table-column prop="location" label="存放位置" width="160" />
        <el-table-column prop="usePermission" label="领用权限" width="180" />
        <el-table-column fixed="right" label="操作" width="100">
          <template slot-scope="scope">
            <el-button type="text" size="small" @click="handleEditClick(scope.row)">编辑</el-button>
            <el-button type="text" size="small" @click="handleDeleteClick(scope.row)">删除</el-button>
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
    <el-dialog :title="slingNewTitle" :visible.sync="isAddSlingVisible" width="50%" top="10px">
      <SlingNew v-if="isAddSlingVisible" ref="slingNew" :sling="curSling" />
      <span slot="footer" class="dialog-footer">
        <el-button @click="isAddSlingVisible = false">取 消</el-button>
        <el-button type="primary" @click="onAddSlingOK">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import * as slingApi from '@/api/sling'
import { mapGetters, mapActions } from 'vuex'
import SlingNew from './components/SlingNew'
import * as _ from 'lodash'

export default {
  name: 'Sling',
  components: {
    SlingNew
  },
  data() {
    return {
      slingNewTitle: '添加吊索具',
      isAddSlingVisible: false,
      curSling: null,
      formData: {
        name: '',
        slingType: null,
        maxTonnage: null,
        useStatus: null,
        inspectStatus: null
      },
      tableData: [],
      dataTotal: 0,
      curPageIndex: 1,
      curPageSize: 10
    }
  },
  computed: {
    ...mapGetters('sling', {
      slingTypes: 'slingTypes',
      slingTons: 'slingTons',
      slingUseStatus: 'slingUseStatus',
      slingInspectStatus: 'slingInspectStatus'
    })
  },
  mounted() {
    this.initDict().then(d => {
      this.querySlings()
    })
  },
  methods: {
    ...mapActions('sling', {
      initDict: 'initDict'
    }),
    onSubmit() {
      this.querySlings()
    },
    onAddNew() {
      this.curSling = null
      this.slingNewTitle = '添加吊索具'
      this.isAddSlingVisible = true
    },
    onAddSlingOK() {
      const data = this.$refs.slingNew.getData()
      if (!data.id) {
        slingApi.addSling(data).then(d => {
          this.$message({
            message: '添加成功',
            type: 'success'
          })
          this.isAddSlingVisible = false
          this.querySlings()
        })
      } else {
        slingApi.updateSling(data).then(d => {
          this.$message({
            message: '修改成功',
            type: 'success'
          })
          this.isAddSlingVisible = false
          this.querySlings()
        })
      }
    },
    handleSizeChange(val) {
      this.curPageSize = val
      this.querySlings()
    },
    handleCurrentChange(val) {
      this.curPageIndex = val
      this.querySlings()
    },
    handleEditClick(data) {
      this.curSling = data
      this.slingNewTitle = '修改吊索具'
      this.isAddSlingVisible = true
    },
    handleDeleteClick(data) {
      this.$confirm('确认删除该吊索具？', '删除', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        slingApi.deleteSling(data.id).then(d => {
          if (d.code === 20000) {
            this.$message({
              message: '删除成功',
              type: 'success'
            })
            // 当前页只有一条数据时，查询前一页数据
            if (this.tableData.length === 1 && this.curPageIndex > 1) {
              this.curPageIndex -= 1
            }
            this.querySlings()
          }
        })
      })
    },
    querySlings() {
      slingApi.querySlings(this.formData, this.curPageSize, this.curPageIndex).then(d => {
        d.data.list.forEach(item => {
          let type = _.find(this.slingTypes, t => t.key === item.slingType)
          item.slingTypeStr = type ? type.name : '-'
          type = _.find(this.slingTons, t => t.key === item.maxTonnage)
          item.maxTonnageStr = type ? type.name : '-'
          type = _.find(this.slingUseStatus, t => t.key === item.useStatus)
          item.useStatusStr = type ? type.name : '-'
          type = _.find(this.slingInspectStatus, t => t.key === item.inspectStatus)
          item.inspectStatusStr = type ? type.name : '-'
          item.location = item.cabinetName ? item.cabinetName + ',' + item.gridNo + '号' : '-'
        })
        this.tableData = d.data.list
        this.dataTotal = d.data.total
      })
    }
  }
}
</script>
<style scoped>
.sling-conaitner {
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

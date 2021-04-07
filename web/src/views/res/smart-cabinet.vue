<template>
  <div class="cabinet-conaitner">
    <el-card class="search-box">
      <div style="position:relative;">
        <el-form :inline="true" :model="formData">
          <el-form-item label="智能柜名称">
            <el-input v-model="formData.name" clearable placeholder="智能柜名称" />
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
        <el-table-column prop="id" label="编号" width="80" />
        <el-table-column prop="name" label="智能柜名称" width="120" />
        <el-table-column prop="location" label="所在位置" width="250" />
        <el-table-column prop="gridCount" label="箱格数" width="80" />
        <el-table-column prop="usedCount" label="已使用" width="80" />
        <el-table-column prop="unUsedCount" label="未使用" width="80" />
        <el-table-column prop="usedRate" label="使用率" width="80" />
        <el-table-column prop="statusStr" label="状态" width="80">
          <template slot-scope="scope">
            <el-tag
              :type="scope.row.status === 0 ? 'success' : 'danger'"
              disable-transitions
            >{{ scope.row.statusStr }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" />
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
    <el-dialog :title="cabinetNewTitle" :visible.sync="isAddCabinetVisible" width="40%" top="10px">
      <SmartCabinetNew v-if="isAddCabinetVisible" ref="cabinetNew" :cabinet="curCabinet" />
      <span slot="footer" class="dialog-footer">
        <el-button @click="isAddCabinetVisible = false">取 消</el-button>
        <el-button type="primary" @click="onAddCabinetOK">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import * as cabinetApi from '@/api/cabinet'
import SmartCabinetNew from './components/SmartCabinetNew'

export default {
  name: 'SmartCabinet',
  components: {
    SmartCabinetNew
  },
  data() {
    return {
      cabinetNewTitle: '添加智能柜',
      isAddCabinetVisible: false,
      curCabinet: null,
      formData: {
        name: ''
      },
      tableData: [],
      dataTotal: 0,
      curPageIndex: 1,
      curPageSize: 10
    }
  },
  mounted() {
    this.queryCabinets()
  },
  methods: {
    onSubmit() {
      this.queryCabinets()
    },
    onAddNew() {
      this.curCabinet = null
      this.cabinetNewTitle = '添加智能柜'
      this.isAddCabinetVisible = true
    },
    onAddCabinetOK() {
      const data = this.$refs.cabinetNew.getData()
      if (!data.id) {
        cabinetApi.addCabinet(data).then(d => {
          this.$message({
            message: '添加成功',
            type: 'success'
          })
          this.isAddCabinetVisible = false
          this.queryCabinets()
        })
      } else {
        cabinetApi.updateCabinet(data).then(d => {
          this.$message({
            message: '修改成功',
            type: 'success'
          })
          this.isAddCabinetVisible = false
          this.queryCabinets()
        })
      }
    },
    handleSizeChange(val) {
      this.curPageSize = val
      this.queryCabinets()
    },
    handleCurrentChange(val) {
      this.curPageIndex = val
      this.queryCabinets()
    },
    handleEditClick(data) {
      this.curCabinet = data
      this.cabinetNewTitle = '修改智能柜'
      this.isAddCabinetVisible = true
    },
    handleDeleteClick(data) {
      this.$confirm('确认删除该智能柜？', '删除', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        cabinetApi.deleteCabinet(data.id).then(d => {
          if (d.code === 20000) {
            this.$message({
              message: '删除成功',
              type: 'success'
            })
            // 当前页只有一条数据时，查询前一页数据
            if (this.tableData.length === 1 && this.curPageIndex > 1) {
              this.curPageIndex -= 1
            }
            this.queryCabinets()
          }
        })
      })
    },
    queryCabinets() {
      cabinetApi.queryCabinets(this.formData.name, this.curPageSize, this.curPageIndex).then(d => {
        d.data.list.forEach(item => {
          item.statusStr = item.status === 0 ? '有效' : '无效'
          item.usedRate = Math.round(item.usedCount * 100.0 / item.gridCount, 2) + '%'
        })
        this.tableData = d.data.list
        this.dataTotal = d.data.total
      })
    }
  }
}
</script>
<style scoped>
.cabinet-conaitner {
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

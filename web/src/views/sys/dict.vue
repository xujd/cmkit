<template>
  <div class="dict-conaitner">
    <el-card class="cond-box">
      <div style="height: 400px">
        <el-input v-model="filterText" placeholder="输入关键字进行过滤" />

        <el-tree
          ref="tree"
          class="filter-tree"
          :data="treeData"
          :props="defaultProps"
          :highlight-current="true"
          default-expand-all
          :filter-node-method="filterNode"
          @node-click="onTreeNodeClick"
        />
      </div>
    </el-card>
    <el-card class="result-box">
      <div style="position: relative">
        <el-table height="400" :data="tableData" style="width: 100%">
          <el-table-column prop="id" label="ID" width="80" />
          <el-table-column prop="name" label="名称" width="150" />
          <el-table-column prop="key" label="键值" width="80" />
          <el-table-column prop="note" label="分类名称" />
          <el-table-column prop="type" label="分类标识" />
          <el-table-column prop="scene" label="归属大类" width="80" />
          <el-table-column fixed="right" label="操作" width="100">
            <template slot-scope="scope">
              <el-button
                type="text"
                size="small"
                @click="handleEditClick(scope.row)"
              >编辑</el-button>
              <el-button
                type="text"
                size="small"
                @click="handleDeleteClick(scope.row)"
              >删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <el-button
          v-show="curNode"
          style="position: absolute; right: 0; top: 0; z-index: 100"
          type="primary"
          icon="el-icon-plus"
          circle
          @click="handleAddClick"
        />
      </div>
    </el-card>
    <el-dialog
      :title="dictNewTitle"
      :visible.sync="isAddDictVisible"
      width="390px"
      top="50px"
    >
      <DictNew v-if="isAddDictVisible" ref="dictNew" :dict-data="curDcit" />
      <span slot="footer" class="dialog-footer">
        <el-button @click="isAddDictVisible = false">取 消</el-button>
        <el-button type="primary" @click="onAddDictOK">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import * as sysApi from '@/api/sys'
import * as _ from 'lodash'
import DictNew from './components/DictNew'

export default {
  components: {
    DictNew
  },

  data() {
    return {
      curDcit: null,
      dictNewTitle: '添加字典',
      isAddDictVisible: false,
      filterText: '',
      treeData: [],
      dictData: [],
      tableData: [],
      curNode: null,
      defaultProps: {
        children: 'children',
        label: 'label'
      }
    }
  },
  watch: {
    filterText(val) {
      this.$refs.tree.filter(val)
    }
  },
  mounted() {
    sysApi.queryDict().then((response) => {
      const { data } = response
      this.dictData = data
      const g1 = _.groupBy(data, 'scene')
      for (const k1 in g1) {
        const p = { id: k1, label: k1, children: [] }
        this.treeData.push(p)
        const g2 = _.groupBy(g1[k1], 'note')
        for (const k2 in g2) {
          const type = {
            id: k2,
            label: k2,
            parent: k1,
            key: g2[k2][0].type
            // children: _.map(g2[k2], (item) => {
            //   return { id: item.id, label: item.name, data: item };
            // }),
          }

          p.children.push(type)
        }
      }
    })
  },

  methods: {
    filterNode(value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    },
    onTreeNodeClick(data) {
      this.curNode = data
      this.tableData = _.orderBy(
        _.filter(this.dictData, {
          scene: data.parent,
          note: data.label
        }),
        ['id'],
        ['asc']
      )
    },
    handleEditClick(data) {
      this.curDcit = data
      this.dictNewTitle = '修改字典'
      this.isAddDictVisible = true
    },
    handleDeleteClick(data) {
      this.$confirm('确认删除该数据？', '删除', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        sysApi.deleteDict(data.id).then((d) => {
          if (d.code === 20000) {
            this.$message({
              message: '删除成功',
              type: 'success'
            })
            this.refreshData()
          }
        })
      })
    },
    handleAddClick() {
      this.curDcit = {
        id: this.dictData.length > 0 ? _.maxBy(this.dictData, 'id').id + 1 : 1,
        name: '',
        key: 0,
        note: this.curNode.label,
        type: this.curNode.key,
        scene: this.curNode.parent
      }
      this.dictNewTitle = '添加字典'
      this.isAddDictVisible = true
    },
    onAddDictOK() {
      const data = this.$refs.dictNew.getData()
      const oldDatas = _.filter(
        this.tableData,
        (item) =>
          item.id !== data.id &&
          (item.name === data.name || item.key === data.key)
      )
      if (oldDatas.length > 0) {
        this.$message({
          message: '名称或键值存在重复',
          type: 'error'
        })
        return
      }
      if (data.id > 0) {
        sysApi.updateDict(data).then((d) => {
          this.$message({
            message: '修改成功',
            type: 'success'
          })
          this.isAddDictVisible = false
          this.refreshData()
        })
      } else {
        sysApi.updateDict(data).then((d) => {
          this.$message({
            message: '添加成功',
            type: 'success'
          })
          this.isAddDictVisible = false
          this.refreshData()
        })
      }
    },
    refreshData() {
      sysApi.queryDict().then((response) => {
        const { data } = response
        this.dictData = data
        this.tableData = _.orderBy(
          _.filter(this.dictData, {
            scene: this.curNode.parent,
            note: this.curNode.label
          }),
          ['id'],
          ['asc']
        )
      })
    }
  }
}
</script>
<style scoped>
.dict-conaitner {
  width: 100%;
  height: 100%;
  padding: 10px;
}
.cond-box {
  width: 250px;
  height: calc(100% - 20px);
  float: left;
}
.result-box {
  width: calc(100% - 270px);
  margin-left: 10px;
  float: left;
}
</style>

<template>
  <el-form ref="form" :model="formData" :rules="rules" label-width="120px">
    <el-form-item label="吊索具名称">
      <el-input v-model="formData.resName" disabled autocomplete="off" />
    </el-form-item>
    <el-form-item label="归还人员" prop="staffId" required>
      <el-select v-model="formData.staffId" clearable placeholder="请选择">
        <el-option v-for="item in staffList" :key="item.id" :label="item.label" :value="item.id" />
      </el-select>
    </el-form-item>
    <el-form-item label="用途说明" prop="remark" required>
      <el-input v-model="formData.remark" type="textarea" :rows="2" placeholder="请输入" />
    </el-form-item>
  </el-form>
</template>
<script>
import { queryStaffs } from '@/api/staff'
import * as _ from 'lodash'
import dayjs from 'dayjs'
export default {
  name: 'ReturnNew',
  props: {
    useLog: {
      type: Object, default: null
    }
  },
  data() {
    return {
      formData: {
        name: '',
        staffId: null,
        remark: ''
      },
      rules: {
        staffId: [
          { required: true, message: '请输入选择归还人员', trigger: 'change' }
        ],
        remark: [
          { required: true, message: '请输入用途说明', trigger: 'blur' }
        ]
      },
      staffList: []
    }
  },
  watch: {
    useLog: {
      handler: function(newVal, oldVal) {
        if (newVal) {
          this.formData = {
            resId: newVal.resId,
            resName: newVal.resName,
            staffId: null,
            remark: newVal.remark
          }
        } else {
          this.formData = {
            resId: 0,
            name: '',
            staffId: null,
            remark: ''
          }
        }
      },
      immediate: true
    }
  },
  mounted() {
    queryStaffs({ name: '' }, 1000, 1).then(d => {
      d.data.list.forEach(item => {
        item.label = item.id + '|' + item.name
      })
      this.staffList = d.data.list.filter(item => item.status === 0)
    })
  },
  methods: {
    resetData() {
      this.formData = {
        resId: 0,
        name: '',
        staffId: null,
        remark: ''
      }
    },
    getData() {
      const tempData = _.cloneDeep(this.formData)
      tempData.returnTime = dayjs().format('YYYY-MM-DD HH:mm:ss')
      if (this.useLog && this.useLog.id > 0) {
        tempData.id = this.useLog.id
      }
      tempData.staffName = _.find(this.staffList, { id: tempData.staffId }).name
      return tempData
    }
  }
}
</script>
<style scoped>
</style>

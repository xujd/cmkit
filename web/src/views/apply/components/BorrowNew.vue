<template>
  <el-form ref="form" :model="formData" :rules="rules" label-width="120px">
    <el-form-item label="RFID">
      <el-input v-model="formData.rfId" disabled autocomplete="off" />
    </el-form-item>
    <el-form-item label="吊索具名称">
      <el-input v-model="formData.name" disabled autocomplete="off" />
    </el-form-item>
    <el-form-item label="借用人员" prop="staffId" required>
      <el-select v-model="formData.staffId" clearable placeholder="请选择">
        <el-option v-for="item in staffList" :key="item.id" :label="item.label" :value="item.id" />
      </el-select>
    </el-form-item>
    <el-form-item label="借用时间" prop="startTime" required>
      <el-date-picker
        v-model="formData.startTime"
        :value-format="'yyyy-MM-dd HH:mm:ss'"
        type="datetime"
        placeholder="选择时间"
      />
    </el-form-item>
    <el-form-item label="预计归还时间" prop="returnTime" required>
      <el-date-picker
        v-model="formData.returnTime"
        :value-format="'yyyy-MM-dd HH:mm:ss'"
        type="datetime"
        placeholder="选择时间"
      />
    </el-form-item>
    <el-form-item label="用途说明" prop="remark" required>
      <el-input v-model="formData.remark" type="textarea" :rows="2" placeholder="请输入" />
    </el-form-item>
  </el-form>
</template>
<script>
import { queryStaffs } from '@/api/staff'
import * as _ from 'lodash'
export default {
  name: 'BorrowNew',
  props: {
    sling: {
      type: Object, default: null
    }
  },
  data() {
    return {
      formData: {
        rfId: '',
        name: '',
        staffId: null,
        startTime: null,
        returnTime: null,
        remark: ''
      },
      rules: {
        staffId: [
          { required: true, message: '请输入选择借用人员', trigger: 'change' }
        ],
        startTime: [
          { required: true, message: '请输入选择借用时间', trigger: 'blur' }
        ],
        returnTime: [
          { required: true, message: '请输入选择归还时间', trigger: 'blur' }
        ],
        remark: [
          { required: true, message: '请输入用途说明', trigger: 'blur' }
        ]
      },
      staffList: []
    }
  },
  watch: {
    sling: {
      handler: function(newVal, oldVal) {
        if (newVal) {
          this.formData = {
            rfId: newVal.rfId,
            name: newVal.name,
            staffId: null,
            startTime: null,
            returnTime: null,
            remark: ''
          }
        } else {
          this.formData = {
            rfId: '',
            name: '',
            staffId: null,
            startTime: null,
            returnTime: null,
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
        rfId: '',
        name: '',
        staffId: null,
        startTime: null,
        returnTime: null,
        remark: ''
      }
    },
    getData() {
      const tempData = _.cloneDeep(this.formData)
      if (this.sling && this.sling.id) {
        tempData.resId = this.sling.id
      }
      return tempData
    }
  }
}
</script>
<style scoped>
</style>

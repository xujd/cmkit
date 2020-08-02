<template>
  <el-form :model="formData" :rules="rules" label-width="100px">
    <el-form-item label="员工姓名" prop="name" required>
      <el-input v-model="formData.name" autocomplete="off" />
    </el-form-item>
    <el-form-item label="所在公司">
      <el-select v-model="formData.companyId" clearable placeholder="请选择" @change="onCompanyChange">
        <el-option v-for="item in companys" :key="item.id" :label="item.name" :value="item.id" />
      </el-select>
    </el-form-item>
    <el-form-item label="所在部门">
      <el-select v-model="formData.departmentId" clearable placeholder="请选择">
        <el-option v-for="item in departments" :key="item.id" :label="item.name" :value="item.id" />
      </el-select>
    </el-form-item>
    <el-form-item label="出生日期">
      <el-date-picker
        v-model="formData.birthday"
        :value-format="'yyyy-MM-dd'"
        type="date"
        placeholder="选择日期"
      />
    </el-form-item>
    <el-form-item label="状态">
      <el-select v-model="formData.status" clearable placeholder="请选择">
        <el-option v-for="item in statusList" :key="item.id" :label="item.name" :value="item.id" />
      </el-select>
    </el-form-item>
    <el-form-item label="备注">
      <el-input v-model="formData.remark" type="textarea" :rows="2" placeholder="请输入" />
    </el-form-item>
  </el-form>
</template>
<script>
import * as sysApi from '@/api/sys'
import _ from 'lodash'
export default {
  name: 'StaffNew',
  props: {
    staff: {
      type: Object, default: null
    }
  },
  data() {
    return {
      companys: [],
      departments: [],
      formData: {
        birthday: null,
        companyId: null,
        departmentId: null,
        name: '',
        remark: '',
        status: 0
      },
      statusList: [
        { id: 0, name: '有效' },
        { id: 1, name: '无效' }
      ],
      rules: {
        name: [
          { required: true, message: '请输入员工姓名', trigger: 'blur' },
          { min: 2, max: 16, message: '长度在 2 到 16 个字符', trigger: 'blur' }
        ]
      }
    }
  },
  watch: {
    staff: {
      handler: function(newVal, oldVal) {
        if (newVal) {
          this.formData = {
            birthday: newVal.birthday,
            companyId: newVal.companyId,
            departmentId: newVal.departmentId,
            name: newVal.name,
            remark: newVal.remark,
            status: newVal.status
          }
          sysApi.queryDepartments('', this.formData.companyId).then(d => {
            this.departments = d.data.list
          })
        } else {
          this.formData = {
            birthday: null,
            companyId: null,
            departmentId: null,
            name: '',
            remark: '',
            status: 0
          }
          this.departments = []
        }
      },
      immediate: true
    }
  },
  mounted() {
    sysApi.queryCompanys('').then(d => {
      this.companys = d.data.list
    })
  },
  methods: {
    onCompanyChange() {
      this.formData.departmentId = null
      if (!this.formData.companyId) {
        this.departments = []
      } else {
        sysApi.queryDepartments('', this.formData.companyId).then(d => {
          this.departments = d.data.list
        })
      }
    },
    resetData() {
      this.formData = {
        birthday: null,
        companyId: null,
        departmentId: null,
        name: '',
        remark: '',
        status: 0
      }
      this.departments = []
    },
    getData() {
      const tempData = _.cloneDeep(this.formData)
      if (this.staff && this.staff.id) {
        tempData.id = this.staff.id
      }
      if (tempData.birthday && tempData.birthday.length < 12) {
        tempData.birthday += ' 00:00:00'
      }
      if (!tempData.companyId) {
        tempData.companyId = 1
      }
      if (!tempData.departmentId) {
        tempData.departmentId = 1
      }
      return tempData
    }
  }
}
</script>
<style scoped>
</style>

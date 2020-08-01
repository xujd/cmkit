<template>
  <el-form :model="formData" :rules="rules" label-width="100px">
    <el-form-item label="用户名" prop="name" required>
      <el-input v-model="formData.name" autocomplete="off" />
    </el-form-item>
    <el-form-item label="使用员工" prop="staffId" required>
      <el-select v-model="formData.staffId" clearable placeholder="请选择" filterable>
        <el-option v-for="item in staffList" :key="item.id" :label="item.label" :value="item.id" />
      </el-select>
    </el-form-item>
    <el-form-item label="生效开始时间">
      <el-date-picker
        v-model="formData.startTime"
        :value-format="'yyyy-MM-dd HH:mm:ss'"
        type="datetime"
        placeholder="选择时间"
      />
    </el-form-item>
    <el-form-item label="生效结束时间">
      <el-date-picker
        v-model="formData.endTime"
        :value-format="'yyyy-MM-dd HH:mm:ss'"
        type="datetime"
        placeholder="选择时间"
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
import * as staffApi from '@/api/staff'
export default {
  name: 'UserNew',
  props: {
    user: {
      type: Object, default: null
    }
  },
  data() {
    return {
      formData: {
        name: '',
        staffId: null,
        startTime: null,
        endTime: null,
        remark: '',
        status: 0
      },
      rules: {
        name: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 2, max: 16, message: '长度在 2 到 16 个字符', trigger: 'blur' }
        ],
        staffId: [
          { required: true, message: '请选择一个员工', trigger: 'change' }
        ]
      },
      staffList: [],
      statusList: [
        { id: 0, name: '有效' },
        { id: 1, name: '无效' }
      ]
    }
  },
  watch: {
    user: {
      handler: function(newVal, oldVal) {
        if (newVal) {
          this.formData = {
            name: newVal.name,
            staffId: newVal.staffId,
            startTime: newVal.startTime,
            endTime: newVal.endTime,
            remark: newVal.remark,
            status: newVal.status
          }
        } else {
          this.formData = {
            name: '',
            staffId: null,
            startTime: null,
            endTime: null,
            remark: '',
            status: 0
          }
        }
      },
      immediate: true
    }
  },
  mounted() {
    staffApi.queryStaffs({ name: '' }, 1000, 1).then(d => {
      d.data.list.forEach(item => {
        item.label = `${item.id}|${item.name}`
      })
      this.staffList = d.data.list
    })
  },
  methods: {
    resetData() {
      this.formData = {
        name: '',
        staffId: null,
        startTime: null,
        endTime: null,
        remark: '',
        status: 0
      }
    },
    getData() {
      const tempData = _.cloneDeep(this.formData)
      if (this.user && this.user.id) {
        tempData.id = this.user.id
      }
      return tempData
    }
  }
}
</script>
<style scoped>
</style>

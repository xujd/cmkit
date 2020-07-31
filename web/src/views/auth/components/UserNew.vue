<template>
  <el-form :model="formData" :rules="rules" label-width="100px">
    <el-form-item label="账号名称" prop="name" required>
      <el-input v-model="formData.name" autocomplete="off" />
    </el-form-item>
    <el-form-item label="使用员工">
      <el-select v-model="formData.departmentId" clearable placeholder="请选择">
        <el-option v-for="item in departments" :key="item.id" :label="item.name" :value="item.id" />
      </el-select>
    </el-form-item>
    <el-form-item label="有效日期">
      <el-date-picker
        v-model="formData.birthday"
        :value-format="'yyyy-MM-dd'"
        type="date"
        placeholder="选择日期"
      />
    </el-form-item>
    <el-form-item label="失效日期">
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
        remark: '',
        status: 0
      },
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
            remark: newVal.remark,
            status: newVal.status
          }
        } else {
          this.formData = {
            name: '',
            remark: '',
            status: 0
          }
        }
      },
      immediate: true
    }
  },
  mounted() {
  },
  methods: {
    resetData() {
      this.formData = {
        name: '',
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

<template>
  <el-form :model="formData" :rules="rules" label-width="100px">
    <el-form-item label="智能柜名称" prop="name" required>
      <el-input v-model="formData.name" style="width: 200px" autocomplete="off" />
    </el-form-item>
    <el-form-item label="箱格数" prop="gridCount" required>
      <el-input-number
        v-model="formData.gridCount"
        controls-position="right"
        :min="1"
        :max="100"
        :precision="0"
        style="width: 200px"
      />
    </el-form-item>
    <el-form-item label="状态">
      <el-select v-model="formData.status" style="width: 200px" clearable placeholder="请选择">
        <el-option v-for="item in statusList" :key="item.id" :label="item.name" :value="item.id" />
      </el-select>
    </el-form-item>
    <el-form-item label="所在位置">
      <el-input v-model="formData.location" autocomplete="off" />
    </el-form-item>
    <el-form-item label="备注">
      <el-input v-model="formData.remark" type="textarea" :rows="2" placeholder="请输入" />
    </el-form-item>
  </el-form>
</template>
<script>
export default {
  name: 'SmartCabinetNew',
  props: {
    cabinet: {
      type: Object, default: null
    }
  },
  data() {
    return {
      formData: {
        name: '',
        gridCount: 1,
        location: '',
        status: 0,
        remark: ''
      },
      rules: {
        name: [
          { required: true, message: '请输入智能柜名称', trigger: 'blur' },
          { min: 2, max: 16, message: '长度在 2 到 16 个字符', trigger: 'blur' }
        ],
        gridCount: [
          { required: true, message: '请输入箱格数', trigger: 'change' }
        ]
      },
      statusList: [
        { id: 0, name: '有效' },
        { id: 1, name: '无效' }
      ]
    }
  },
  watch: {
    cabinet: {
      handler: function(newVal, oldVal) {
        if (newVal) {
          this.formData = {
            name: newVal.name,
            gridCount: newVal.gridCount,
            location: newVal.location,
            remark: newVal.remark,
            status: newVal.status
          }
        } else {
          this.formData = {
            name: '',
            gridCount: 1,
            location: '',
            status: 0,
            remark: ''
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
        gridCount: 1,
        location: '',
        status: 0,
        remark: ''
      }
    },
    getData() {
      if (this.cabinet && this.cabinet.id) {
        this.formData.id = this.cabinet.id
      }
      return this.formData
    }
  }
}
</script>
<style scoped>
</style>

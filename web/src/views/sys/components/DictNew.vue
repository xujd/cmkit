<template>
  <el-form
    ref="form"
    status-icon
    :model="formData"
    :rules="rules"
    label-width="120px"
  >
    <el-form-item label="名称" prop="name" required>
      <el-input
        v-model="formData.name"
        :maxlength="64"
        style="width: 200px"
        autocomplete="off"
      />
    </el-form-item>
    <el-form-item label="键值" prop="key" required>
      <el-input v-model.number="formData.key" style="width: 200px" />
    </el-form-item>
    <el-form-item label="分类名称">
      <el-input
        v-model="formData.note"
        style="width: 200px"
        disabled
        autocomplete="off"
      />
    </el-form-item>
    <el-form-item label="归属大类">
      <el-input
        v-model="formData.scene"
        style="width: 200px"
        disabled
        autocomplete="off"
      />
    </el-form-item>
  </el-form>
</template>
<script>
export default {
  name: 'DictNew',
  props: {
    dictData: {
      type: Object,
      default: null
    }
  },
  data() {
    var checkKey = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('键值不能为空'))
      }
      setTimeout(() => {
        if (!Number.isInteger(value)) {
          callback(new Error('请输入数字值'))
        } else {
          if (value < 1) {
            callback(new Error('键值必须大于0'))
          } else {
            callback()
          }
        }
      }, 500)
    }
    return {
      formData: {
        id: 0,
        name: '',
        key: 0,
        note: '',
        type: '',
        scene: ''
      },
      rules: {
        name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
        key: [{ required: true, validator: checkKey, trigger: 'blur' }]
      }
    }
  },
  watch: {
    dictData: {
      handler: function(newVal, oldVal) {
        if (newVal) {
          this.formData = {
            id: newVal.id,
            name: newVal.name,
            key: newVal.key,
            note: newVal.note,
            type: newVal.type,
            scene: newVal.scene
          }
        } else {
          this.formData = {
            id: 0,
            name: '',
            key: 0,
            note: '',
            type: '',
            scene: ''
          }
        }
      },
      immediate: true
    }
  },
  mounted() {},
  methods: {
    resetData() {
      this.formData = {
        id: 0,
        name: '',
        key: 0,
        note: '',
        type: '',
        scene: ''
      }
    },
    getData() {
      return this.formData
    }
  }
}
</script>
<style scoped>
</style>

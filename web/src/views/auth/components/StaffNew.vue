<template>
  <div>
    <el-form :model="formData" :rules="rules" label-width="100px">
      <el-form-item label="员工姓名" prop="name" required>
        <el-input
          v-model="formData.name"
          autocomplete="off"
          style="width: 200px"
        />
      </el-form-item>
      <el-form-item label="所在公司">
        <el-select
          v-model="formData.companyId"
          style="width: 200px"
          clearable
          placeholder="请选择"
          @change="onCompanyChange"
        >
          <el-option
            v-for="item in companys"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="所在部门">
        <el-select
          v-model="formData.departmentId"
          style="width: 200px"
          clearable
          placeholder="请选择"
        >
          <el-option
            v-for="item in departments"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="出生日期">
        <el-date-picker
          v-model="formData.birthday"
          style="width: 200px"
          :value-format="'yyyy-MM-dd'"
          type="date"
          placeholder="选择日期"
        />
      </el-form-item>
      <el-form-item label="状态">
        <el-select
          v-model="formData.status"
          style="width: 200px"
          clearable
          placeholder="请选择"
        >
          <el-option
            v-for="item in statusList"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="备注">
        <el-input
          v-model="formData.remark"
          type="textarea"
          :rows="2"
          placeholder="请输入"
        />
      </el-form-item>
    </el-form>
    <img
      :src="imageUrl"
      class="person-img"
      @error="errorLoadImg"
      @click="imagecropperShow = true"
    >
    <image-cropper
      v-show="imagecropperShow"
      :key="imagecropperKey"
      field="uploadfile"
      :params="picParam"
      :width="300"
      :height="300"
      url="/file/upload"
      @close="close"
      @crop-upload-success="cropSuccess"
    />
  </div>
</template>
<script>
import * as sysApi from '@/api/sys'
import _ from 'lodash'
import ImageCropper from '@/components/ImageCropper'

export default {
  name: 'StaffNew',
  components: {
    ImageCropper
  },
  props: {
    staff: {
      type: Object, default: null
    }
  },
  data() {
    return {
      imagecropperShow: false,
      imagecropperKey: 0,
      imageUrl: 'assets/person.svg',
      defaultImageUrl: 'assets/person.svg',
      picParam: {
        fileName: ''
      },
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
          this.picParam.fileName = newVal.id.toString().padStart(6, '0') + '.png'
          this.imageUrl = 'assets/pictures/' + this.picParam.fileName
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
          this.picParam.fileName = ''
          this.imageUrl = this.defaultImageUrl
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
    errorLoadImg() {
      this.imageUrl = this.defaultImageUrl
    },
    cropSuccess(resData) {
      this.imagecropperShow = false
      this.imagecropperKey = this.imagecropperKey + 1
      this.imageUrl = 'assets/pictures/' + resData.data.fileName
    },
    close() {
      this.imagecropperShow = false
    },
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
.person-img {
  width: 180px;
  height: 180px;
  position: absolute;
  right: 24px;
  top: 84px;
  border: 1px solid lightgray;
  cursor: pointer;
}
</style>

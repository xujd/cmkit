<template>
  <el-form :model="formData" :rules="rules" label-width="100px">
    <el-form-item label="RFID" prop="rfId" required>
      <el-input v-model="formData.rfId" autocomplete="off" />
    </el-form-item>
    <el-form-item label="吊索具名称" prop="name" required>
      <el-input v-model="formData.name" autocomplete="off" />
    </el-form-item>
    <el-form-item label="型号">
      <el-select v-model="formData.slingType" clearable placeholder="请选择">
        <el-option v-for="item in slingTypes" :key="item.key" :label="item.name" :value="item.key" />
      </el-select>
    </el-form-item>
    <el-form-item label="最大吨位">
      <el-select v-model="formData.maxTonnage" clearable placeholder="请选择">
        <el-option v-for="item in slingTons" :key="item.key" :label="item.name" :value="item.key" />
      </el-select>
    </el-form-item>
    <el-form-item label="使用状态">
      <el-select v-model="formData.useStatus" clearable placeholder="请选择">
        <el-option
          v-for="item in slingUseStatus"
          :key="item.key"
          :label="item.name"
          :value="item.key"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="点检状态">
      <el-select v-model="formData.inspectStatus" clearable placeholder="请选择">
        <el-option
          v-for="item in slingInspectStatus"
          :key="item.key"
          :label="item.name"
          :value="item.key"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="投用日期">
      <el-date-picker
        v-model="formData.putTime"
        :value-format="'yyyy-MM-dd HH:mm:ss'"
        type="datetime"
        placeholder="选择日期"
      />
    </el-form-item>
    <el-form-item label="存放位置" prop="gridNo" required>
      <el-col :span="12">
        <el-select
          v-model="formData.cabinetId"
          @change="onCabinetChange"
          clearable
          placeholder="请选择"
        >
          <el-option
            v-for="item in cabinetList"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          />
        </el-select>柜
      </el-col>
      <el-col :span="12">
        <el-select v-model="formData.gridNo" clearable placeholder="请选择">
          <el-option
            v-for="item in gridList"
            :disabled="item.disabled"
            :key="item.gridNo"
            :label="item.gridNo"
            :value="item.gridNo"
          />
        </el-select>箱
      </el-col>
    </el-form-item>
    <el-form-item label="领用权限">
      <el-input v-model="formData.usePermission" type="textarea" :rows="2" placeholder="请输入" />
    </el-form-item>
  </el-form>
</template>
<script>
import { mapGetters } from 'vuex'
import { queryCabinets, queryGrids } from '@/api/cabinet'
import * as _ from 'lodash'
export default {
  name: 'SlingNew',
  props: {
    sling: {
      type: Object, default: null
    }
  },
  computed: {
    ...mapGetters('sling', {
      slingTypes: 'slingTypes',
      slingTons: 'slingTons',
      slingUseStatus: 'slingUseStatus',
      slingInspectStatus: 'slingInspectStatus',
    })
  },
  data() {
    return {
      formData: {
        rfId: '',
        name: '',
        slingType: null,
        maxTonnage: null,
        useStatus: null,
        inspectStatus: null,
        putTime: null,
        usePermission: '',
        cabinetId: null,
        gridNo: null
      },
      rules: {
        rfId: [
          { required: true, message: '请输入RFID', trigger: 'blur' }
        ],
        name: [
          { required: true, message: '请输入吊索具名称', trigger: 'blur' }
        ],
        gridNo: [
          { required: true, message: '请选择存放位置', trigger: 'change' }
        ]
      },
      cabinetList: [],
      gridList: []
    }
  },
  watch: {
    sling: {
      handler: function (newVal, oldVal) {
        if (newVal) {
          this.formData = {
            rfId: newVal.rfId,
            name: newVal.name,
            slingType: newVal.slingType === 0 ? null : newVal.slingType,
            maxTonnage: newVal.maxTonnage === 0 ? null : newVal.maxTonnage,
            useStatus: newVal.useStatus === 0 ? null : newVal.useStatus,
            inspectStatus: newVal.inspectStatus === 0 ? null : newVal.inspectStatus,
            putTime: newVal.putTime,
            usePermission: newVal.usePermission,
            cabinetId: newVal.cabinetId,
            gridNo: newVal.gridNo
          }
          queryGrids(this.formData.cabinetId).then(d => {
            d.data.list.forEach(item => {
              item.disabled = item.inResId !== 0
            })
            this.gridList = d.data.list
          })
        } else {
          this.formData = {
            rfId: '',
            name: '',
            slingType: null,
            maxTonnage: null,
            useStatus: null,
            inspectStatus: null,
            putTime: null,
            usePermission: '',
            cabinetId: null,
            gridNo: null
          }
          this.gridList = []
        }
      },
      immediate: true
    }
  },
  mounted() {
    queryCabinets('', 1000, 1).then(d => {
      this.cabinetList = d.data.list
    })
  },
  methods: {
    onCabinetChange() {
      this.formData.gridNo = null
      if (!this.formData.cabinetId) {
        this.gridList = []
      } else {
        queryGrids(this.formData.cabinetId).then(d => {
          d.data.list.forEach(item => {
            item.disabled = item.inResId !== 0
          })
          this.gridList = d.data.list
        })
      }
    },
    resetData() {
      this.formData = {
        rfId: '',
        name: '',
        slingType: null,
        maxTonnage: null,
        useStatus: null,
        inspectStatus: null,
        putTime: null,
        usePermission: '',
        cabinetId: null,
        gridNo: null
      }
    },
    getData() {
      const tempData = _.cloneDeep(this.formData)
      if (this.sling && this.sling.id) {
        tempData.id = this.sling.id
      }
      return tempData
    }
  }
}
</script>
<style scoped>
</style>

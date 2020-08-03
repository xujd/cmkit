<template>
  <el-card class="box-card-component" style="margin-left:8px;">
    <div slot="header" class="box-card-header">
      <span>吊索具状态统计</span>
    </div>
    <div style="position:relative;">
      <div v-for="item of statList" :key="item.name" class="progress-item">
        <span>{{ item.name }}：{{ item.count }}</span>
        <el-progress :percentage="item.percentage" />
      </div>
    </div>
  </el-card>
</template>

<script>
import { getSlingStatByStatus } from '@/api/home'
export default {
  data() {
    return {
      statList: []
    }
  },
  computed: {
  },
  mounted() {
    getSlingStatByStatus().then(d => {
      const totalItem = d.data.find(item => item.name === '总数')
      totalItem.count = totalItem.count || 1
      d.data.forEach(item => {
        item.percentage = parseFloat((item.count * 100 / totalItem.count).toFixed(1))
      })

      this.statList = d.data
    })
  }
}
</script>

<style lang="scss" >
.box-card-component {
  .el-card__header {
    padding: 0px !important;
  }
}
</style>
<style lang="scss" scoped>
.box-card-component {
  .box-card-header {
    position: relative;
    height: 51px;
    line-height: 51px;
    padding-left: 15px;
  }
  .mallki-text {
    position: absolute;
    top: 0px;
    right: 0px;
    font-size: 20px;
    font-weight: bold;
  }
  .panThumb {
    z-index: 100;
    height: 70px !important;
    width: 70px !important;
    position: absolute !important;
    top: -45px;
    left: 0px;
    border: 5px solid #ffffff;
    background-color: #fff;
    margin: auto;
    box-shadow: none !important;
    ::v-deep .pan-info {
      box-shadow: none !important;
    }
  }
  .progress-item {
    margin-bottom: 10px;
    font-size: 14px;
  }
  @media only screen and (max-width: 1510px) {
    .mallki-text {
      display: none;
    }
  }
}
</style>

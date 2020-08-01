<template>
  <div class="block">
    <el-timeline>
      <el-timeline-item v-for="(item,index) of timeline" :key="index" :timestamp="item.timestamp" placement="top">
        <el-card>
          <h4>{{ item.title }}</h4>
          <p>{{ item.content }}</p>
        </el-card>
      </el-timeline-item>
    </el-timeline>
  </div>
</template>

<script>
export default {
  data() {
    return {
      timeline: [
        {
          timestamp: '2019/4/20',
          title: '归还事件',
          content: '归还N2038号工具 2019/4/20 11:22'
        },
        {
          timestamp: '2019/4/21',
          title: '借出事件',
          content: '借出N2011号工具 2019/4/21 20:46'
        },
        {
          timestamp: '2019/4/22',
          title: '借出事件',
          content: '借出N1013号工具 2019/4/22 20:46'
        }
      ]
    }
  }
}
</script>

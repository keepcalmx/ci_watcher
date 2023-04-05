<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const activeIndex = ref('1')

const router = useRouter()
const toOverview = () => {
  router.push({ name: 'Overview' })
}
const toTestCase = () => {
  router.push({ name: 'TestCase' })
}
const toCaseLog = () => {
  router.push({ name: 'CaseLog' })
}

</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <el-menu :default-active="activeIndex" mode="horizontal" :ellipsis="false">
          <el-menu-item index="0">VISUAL PIPELINE DASHBORD</el-menu-item>
          <!-- <div class="flex-grow" /> -->
          <el-menu-item index="1" @click="toOverview">概览</el-menu-item>
          <el-menu-item index="2" @click="toTestCase">用例</el-menu-item>
          <el-menu-item index="3" @click="toCaseLog">日志</el-menu-item>
          <el-menu-item index="3" @click="toCaseLog">环境</el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <RouterView></RouterView>
      </el-main>
    </el-container>
  </div>
</template>

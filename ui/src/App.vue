<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const activeIndex = ref('Overview')
const handleSelect = (key: string, keyPath: string[]) => {
  router.push({ name: key })
}

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
        <el-menu :default-active="activeIndex" mode="horizontal" :ellipsis="false" @select="handleSelect">
          <el-menu-item index="log">VISUAL PIPELINE DASHBORD</el-menu-item>
          <el-menu-item index="Overview">概览</el-menu-item>
          <el-menu-item index="TestCase">用例</el-menu-item>
          <el-menu-item index="CaseLog">日志</el-menu-item>
          <el-menu-item index="EnvOps">环境</el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <RouterView></RouterView>
      </el-main>
    </el-container>
  </div>
</template>

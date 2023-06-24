<script setup lang="ts">
import { ref } from 'vue'
import { CaretRight, Edit } from '@element-plus/icons-vue'

const tableData = ref(
  [
    {
      id: 'f7e0355068434b67',
      name: '测试Max函数',
      owner_id: 'x30043367',
      owner_name: '向峰',
      env_type: 'SSH_CENT',
      work_dir: 'CENT/driver/go/test1',
      executor: 'go'
    },
    {
      id: '505d541701e547b8',
      name: '测试斐波那契数列',
      owner_id: 'x30043367',
      owner_name: '向峰',
      env_type: 'SSH_CENT',
      work_dir: 'CENT/driver/go/test_fib',
      executor: 'go'
    },
    {
      id: '93fe1ed716514b4a',
      name: '测试Roman函数',
      owner_id: 't30032423',
      owner_name: '唐佳乐',
      env_type: 'SSH_CENT',
      work_dir: 'CENT/driver/go/test2',
      executor: 'go'
    },
    {
      id: '93fe1ed716514b4a',
      name: '测试Roman函数',
      owner_id: 't30032423',
      owner_name: '唐佳乐',
      env_type: 'SSH_CENT',
      work_dir: 'CENT/driver/go/test2',
      executor: 'go'
    },
    {
      id: '93fe1ed716514b4a',
      name: '测试Roman函数',
      owner_id: 't30032423',
      owner_name: '唐佳乐',
      env_type: 'SSH_CENT',
      work_dir: 'CENT/driver/go/test2',
      executor: 'go'
    },
    {
      id: '93fe1ed716514b4a',
      name: '测试Roman函数',
      owner_id: 't30032423',
      owner_name: '唐佳乐',
      env_type: 'SSH_CENT',
      work_dir: 'CENT/driver/go/test2',
      executor: 'go'
    },
    {
      id: '93fe1ed716514b4a',
      name: '测试Roman函数',
      owner_id: 't30032423',
      owner_name: '唐佳乐',
      env_type: 'SSH_CENT',
      work_dir: 'CENT/driver/go/test2',
      executor: 'go'
    },
    {
      id: '93fe1ed716514b4a',
      name: '测试Roman函数',
      owner_id: 't30032423',
      owner_name: '唐佳乐',
      env_type: 'SSH_CENT',
      work_dir: 'CENT/driver/go/test2',
      executor: 'go'
    },
    {
      id: '93fe1ed716514b4a',
      name: '测试Roman函数',
      owner_id: 't30032423',
      owner_name: '唐佳乐',
      env_type: 'SSH_CENT',
      work_dir: 'CENT/driver/go/test2',
      executor: 'go'
    },
    {
      id: '93fe1ed716514b4a',
      name: '测试Roman函数',
      owner_id: 't30032423',
      owner_name: '唐佳乐',
      env_type: 'SSH_CENT',
      work_dir: 'CENT/driver/go/test2',
      executor: 'go'
    },
    {
      id: '93fe1ed716514b4a',
      name: '测试Roman函数',
      owner_id: 't30032423',
      owner_name: '唐佳乐',
      env_type: 'SSH_CENT',
      work_dir: 'CENT/driver/go/test2',
      executor: 'go'
    },
    {
      id: '93fe1ed716514b4a',
      name: '测试Roman函数',
      owner_id: 't30032423',
      owner_name: '唐佳乐',
      env_type: 'SSH_CENT',
      work_dir: 'CENT/driver/go/test2',
      executor: 'go'
    }
  ]
)

</script>
<template>
  <el-row justify="center">
    <el-col :span="16">
      <el-card>
        <template #header>
          <div class="card-header">
            <span>执行中的用例</span>
            <el-button type="danger" plain>终止</el-button>
          </div>
        </template>
        <el-descriptions border>
          <el-descriptions-item label="Owner">向峰</el-descriptions-item>
          <el-descriptions-item label="Name">测试Roman函数</el-descriptions-item>
          <el-descriptions-item label="组网方式">SSH_CENT</el-descriptions-item>
          <el-descriptions-item label="执行框架">
            <el-tag size="large">Go</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="已用时">32 分钟</el-descriptions-item>
          <el-descriptions-item label="总进度">80%</el-descriptions-item>
          <el-descriptions-item label="执行路径">CENT/driver/go/test2</el-descriptions-item>
        </el-descriptions>
      </el-card>
    </el-col>

    <el-col :span="16" style="margin-top: 20px;"></el-col>

    <el-col :span="16">
      <el-card>
        <template #header>
          <div class="card-header">
            <span>排队中的用例</span>
            <el-button type="danger" plain>清空</el-button>
          </div>
        </template>
        <div v-for="data in tableData">
          <el-descriptions border>
            <el-descriptions-item label="Owner">{{ data.owner_name }}</el-descriptions-item>
            <el-descriptions-item label="Name">{{ data.name }}</el-descriptions-item>
            <el-descriptions-item label="组网方式">{{ data.env_type }}</el-descriptions-item>
            <el-descriptions-item label="执行框架">
              <el-tag size="large">{{ data.executor }}</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="执行路径">{{ data.work_dir }}</el-descriptions-item>
          </el-descriptions>

          <el-divider />
        </div>

      </el-card>
    </el-col>

    <el-col :span="16" style="margin-top: 20px;"></el-col>

    <el-col :span="16">
      <el-card>
        <template #header>
          <div class="card-header">
            <span>全部用例</span>
          </div>
        </template>
        <el-table :data="tableData" stripe border style="width: 100%">
          <el-table-column type="selection" width="40" />
          <el-table-column prop="name" label="名称" width="240" />
          <el-table-column prop="owner_name" label="Owner" width="100" />
          <el-table-column prop="env_type" label="组网类型" width="100" />
          <el-table-column prop="executor" label="执行框架" width="100" />
          <el-table-column prop="work_dir" label="执行路径" />
          <el-table-column fixed="right" label="操作" width="110">
            <template #default>
              <el-button type="primary" size="small" :icon="CaretRight" />
              <el-button type="success" size="small" :icon="Edit" />
            </template>
          </el-table-column>
        </el-table>
      </el-card>
    </el-col>
  </el-row>

  <el-backtop :right="100" :bottom="100" />
</template>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}
</style>